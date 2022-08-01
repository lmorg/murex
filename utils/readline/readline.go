package readline

import (
	"bytes"
	"os"
	"regexp"
	"sync/atomic"
)

var rxMultiline = regexp.MustCompile(`[\r\n]+`)

// Readline displays the readline prompt.
// It will return a string (user entered data) or an error.
func (rl *Instance) Readline() (_ string, err error) {
	rl.mutex.Lock()
	rl.Active = true
	fd := int(os.Stdin.Fd())
	state, err := MakeRaw(fd)
	rl.mutex.Unlock()

	if err != nil {
		return "", err
	}
	defer func() {
		rl.mutex.Lock()
		rl.Active = false
		// return an error if Restore fails. However we don't want to return
		// `nil` if there is no error because there might be a CtrlC or EOF
		// that needs to be returned
		r := Restore(fd, state)
		if r != nil {
			err = r
		}
		rl.mutex.Unlock()
	}()

	x, _ := rl.getCursorPos()
	switch x {
	case -1:
		print(string(leftMost()))
	case 0:
		// do nothing
	default:
		print("\r\n")
	}
	print(rl.prompt)

	rl.line = []rune{}
	rl.lineChange = ""
	rl.viUndoHistory = []undoItem{{line: "", pos: 0}}
	rl.pos = 0
	rl.histPos = rl.History.Len()
	rl.modeViMode = vimInsert
	atomic.StoreInt32(&rl.delayedSyntaxCount, 0)
	rl.resetHintText()
	rl.resetTabCompletion()

	if len(rl.multiSplit) > 0 {
		r := []rune(rl.multiSplit[0])
		rl.readlineInput(r)
		rl.carriageReturn()
		if len(rl.multiSplit) > 1 {
			rl.multiSplit = rl.multiSplit[1:]
		} else {
			rl.multiSplit = []string{}
		}
		return string(rl.line), nil
	}

	rl.termWidth = GetTermWidth()
	rl.getHintText()
	rl.renderHelpers()

	for {
		if len(rl.line) == 0 {
			// clear the cache when the line is cleared
			rl.cacheHint.Init(rl)
			rl.cacheSyntax.Init(rl)
		} /*else if rl.line[0] == 16 { // fix bug editing pasted lines
			//rl.line = rl.line[1:]
		}*/

		go delayedSyntaxTimer(rl, atomic.LoadInt32(&rl.delayedSyntaxCount))
		rl.viUndoSkipAppend = false
		b := make([]byte, 1024*1024)
		var i int

		if !rl.skipStdinRead {
			i, err = read(b)
			if err != nil {
				return "", err
			}
			rl.termWidth = GetTermWidth()
		}
		atomic.AddInt32(&rl.delayedSyntaxCount, 1)

		rl.skipStdinRead = false
		r := []rune(string(b))

		if isMultiline(r[:i]) || len(rl.multiline) > 0 {
			rl.multiline = append(rl.multiline, b[:i]...)
			//if i == len(b) {
			//	continue
			//}

			if !rl.allowMultiline(rl.multiline) {
				rl.multiline = []byte{}
				continue
			}

			s := string(rl.multiline)
			rl.multiSplit = rxMultiline.Split(s, -1)

			r = []rune(rl.multiSplit[0])
			rl.modeViMode = vimInsert
			rl.readlineInput(r)
			rl.carriageReturn()
			rl.multiline = []byte{}
			if len(rl.multiSplit) > 1 {
				rl.multiSplit = rl.multiSplit[1:]
			} else {
				rl.multiSplit = []string{}
			}
			return string(rl.line), nil
		}

		s := string(r[:i])
		if rl.evtKeyPress[s] != nil {
			ret := rl.evtKeyPress[s](s, rl.line, rl.pos)

			rl.clearLine()
			rl.line = append(ret.NewLine, []rune{}...)
			rl.echo()
			rl.pos = ret.NewPos

			if ret.ClearHelpers {
				rl.resetHelpers()
			} else {
				rl.updateHelpers()
				rl.renderHelpers()
			}

			if len(ret.HintText) > 0 {
				rl.hintText = ret.HintText
				rl.clearHelpers()
				rl.renderHelpers()
			}
			if !ret.ForwardKey {
				continue
			}
			if ret.CloseReadline {
				rl.clearHelpers()
				return string(rl.line), nil
			}
		}

		i = removeNonPrintableChars(b)

		// Used for syntax completion
		rl.lineChange = string(b[:i])

		// Slow or invisible tab completions shouldn't lock up cursor movement
		if rl.modeTabCompletion && len(rl.tcSuggestions) == 0 {
			if rl.delayedTabContext.cancel != nil {
				rl.delayedTabContext.cancel()
			}
			rl.modeTabCompletion = false
			rl.updateHelpers()
		}

		switch b[0] {
		case charCtrlC:
			rl.clearHelpers()
			return "", CtrlC

		case charEOF:
			rl.clearHelpers()
			return "", EOF

		case charCtrlF:
			if !rl.modeTabCompletion {
				rl.modeAutoFind = true
				rl.getTabCompletion()
			}

			rl.modeTabFind = true
			rl.updateTabFind([]rune{})
			rl.viUndoSkipAppend = true

		case charCtrlR:
			rl.modeAutoFind = true
			rl.tcOffset = 0
			rl.modeTabCompletion = true
			rl.tcDisplayType = TabDisplayMap
			rl.tcSuggestions, rl.tcDescriptions = rl.autocompleteHistory()
			rl.initTabCompletion()

			rl.modeTabFind = true
			rl.updateTabFind([]rune{})
			rl.viUndoSkipAppend = true

		case charCtrlU:
			rl.clearLine()
			rl.resetHelpers()

		case charTab:
			if rl.modeTabCompletion {
				rl.moveTabCompletionHighlight(1, 0)
			} else {
				rl.getTabCompletion()
			}

			rl.renderHelpers()
			rl.viUndoSkipAppend = true

		case '\r':
			fallthrough
		case '\n':
			var suggestions []string
			if rl.modeTabFind {
				suggestions = rl.tfSuggestions
			} else {
				suggestions = rl.tcSuggestions
			}

			if rl.modeTabCompletion || len(rl.tfLine) != 0 /*&& len(suggestions) > 0*/ {
				tfLine := rl.tfLine
				cell := (rl.tcMaxX * (rl.tcPosY - 1)) + rl.tcOffset + rl.tcPosX - 1
				rl.clearHelpers()
				rl.resetTabCompletion()
				rl.renderHelpers()
				if len(suggestions) > 0 {
					rl.insert([]rune(suggestions[cell]))
				} else {
					rl.insert(tfLine)
				}
				continue
			}
			rl.carriageReturn()
			return string(rl.line), nil

		case charBackspace, charBackspace2:
			if rl.modeTabFind {
				rl.backspaceTabFind()
				rl.viUndoSkipAppend = true
			} else {
				rl.backspace()
				rl.renderHelpers()
			}

		case charEscape:
			rl.escapeSeq(r[:i])

		default:
			if rl.modeTabFind {
				rl.updateTabFind(r[:i])
				rl.viUndoSkipAppend = true
			} else {
				rl.readlineInput(r[:i])
				if len(rl.multiline) > 0 && rl.modeViMode == vimKeys {
					rl.skipStdinRead = true
				}
			}
		}

		//if !rl.viUndoSkipAppend {
		//	rl.viUndoHistory = append(rl.viUndoHistory, rl.line)
		//}
		rl.undoAppendHistory()
	}
}

func (rl *Instance) escapeSeq(r []rune) {
	switch string(r) {
	case string([]rune{charEscape}):
		switch {
		case rl.modeAutoFind:
			rl.resetTabFind()
			rl.clearHelpers()
			rl.resetTabCompletion()
			rl.renderHelpers()

		case rl.modeTabFind:
			rl.resetTabFind()

		case rl.modeTabCompletion:
			rl.clearHelpers()
			rl.resetTabCompletion()
			rl.renderHelpers()

		default:
			if rl.pos == len(rl.line) && len(rl.line) > 0 {
				rl.pos--
				moveCursorBackwards(1)
			}
			rl.modeViMode = vimKeys
			rl.viIteration = ""
			rl.viHintMessage()
		}
		rl.viUndoSkipAppend = true

	case seqDelete:
		if rl.modeTabFind {
			rl.backspaceTabFind()
		} else {
			rl.delete()
		}

	case seqUp:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(0, -1)
			rl.renderHelpers()
			return
		}

		// are we midway through a long line that wrap multiple terminal lines?
		_, posY := lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)
		if posY > 0 {
			rl.moveCursorByAdjust(-rl.termWidth + rl.promptLen)
			return
		}
		rl.walkHistory(-1)

	case seqDown:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(0, 1)
			rl.renderHelpers()
			return
		}

		// are we midway through a long line that wrap multiple terminal lines?
		_, posY := lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)
		_, lineY := lineWrapPos(rl.promptLen, len(rl.line), rl.termWidth)
		if posY < lineY {
			rl.moveCursorByAdjust(rl.termWidth - rl.promptLen)
			return
		}
		rl.walkHistory(1)

	case seqBackwards:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(-1, 0)
			rl.renderHelpers()
			return
		}

		if rl.pos > 0 {
			rl.moveCursorByAdjust(-1)
		}
		rl.viUndoSkipAppend = true

	case seqForwards:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(1, 0)
			rl.renderHelpers()
			return
		}

		if (rl.modeViMode == vimInsert && rl.pos < len(rl.line)) ||
			(rl.modeViMode != vimInsert && rl.pos < len(rl.line)-1) {
			rl.moveCursorByAdjust(1)
		}
		rl.viUndoSkipAppend = true

	case seqHome, seqHomeSc:
		if rl.modeTabCompletion {
			return
		}

		rl.moveCursorByAdjust(-rl.pos)
		rl.viUndoSkipAppend = true

	case seqEnd, seqEndSc:
		if rl.modeTabCompletion {
			return
		}

		rl.moveCursorByAdjust(len(rl.line) - rl.pos)
		rl.viUndoSkipAppend = true

	case seqShiftTab:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(-1, 0)
			rl.renderHelpers()
			return
		}

	default:
		if rl.modeTabFind {
			return
		}
		// alt+numeric append / delete
		if len(r) == 2 && '1' <= r[1] && r[1] <= '9' {
			if rl.modeViMode == vimDelete {
				rl.vimDelete(r)
				return
			}

			line, err := rl.History.GetLine(rl.History.Len() - 1)
			if err != nil {
				return
			}

			tokens, _, _ := tokeniseSplitSpaces([]rune(line), 0)
			pos := int(r[1]) - 48 // convert ASCII to integer
			if pos > len(tokens) {
				return
			}
			rl.insert([]rune(tokens[pos-1]))
		} else {
			rl.viUndoSkipAppend = true
		}
	}
}

// readlineInput is an unexported function used to determine what mode of text
// entry readline is currently configured for and then update the line entries
// accordingly.
func (rl *Instance) readlineInput(r []rune) {
	switch rl.modeViMode {
	case vimKeys:
		rl.vi(r[0])
		rl.viHintMessage()

	case vimDelete:
		rl.vimDelete(r)
		rl.viHintMessage()

	case vimReplaceOnce:
		rl.modeViMode = vimKeys
		rl.delete()
		rl.insert([]rune{r[0]})
		rl.viHintMessage()

	case vimReplaceMany:
		for _, char := range r {
			rl.delete()
			rl.insert([]rune{char})
		}
		rl.viHintMessage()

	default:
		rl.insert(r)
	}
}

// SetPrompt will define the readline prompt string.
// It also calculates the runes in the string as well as any non-printable
// escape codes.
func (rl *Instance) SetPrompt(s string) {
	rl.prompt = s
	rl.promptLen = strLen(s)
}

func (rl *Instance) carriageReturn() {
	rl.clearHelpers()
	print("\r\n")
	if rl.HistoryAutoWrite {
		var err error
		rl.histPos, err = rl.History.Write(string(rl.line))
		if err != nil {
			print(err.Error() + "\r\n")
		}
	}
}

func isMultiline(r []rune) bool {
	for i := range r {
		if (r[i] == '\r' || r[i] == '\n') && i != len(r)-1 {
			return true
		}
	}
	return false
}

func (rl *Instance) allowMultiline(data []byte) bool {
	rl.clearHelpers()
	printf("\r\nWARNING: %d bytes of multiline data was dumped into the shell!", len(data))
	for {
		print("\r\nDo you wish to proceed (yes|no|preview)? [y/n/p] ")

		b := make([]byte, 1024*1024)

		i, err := os.Stdin.Read(b)
		if err != nil {
			return false
		}

		if i > 1 {
			rl.multiline = append(rl.multiline, b[:i]...)
			moveCursorUp(2)
			return rl.allowMultiline(append(data, b[:i]...))
		}

		s := string(b[:i])
		print(s)

		switch s {
		case "y", "Y":
			print("\r\n" + rl.prompt)
			return true

		case "n", "N":
			print("\r\n" + rl.prompt)
			return false

		case "p", "P":
			preview := string(bytes.Replace(data, []byte{'\r'}, []byte{'\r', '\n'}, -1))
			if rl.SyntaxHighlighter != nil {
				preview = rl.SyntaxHighlighter([]rune(preview))
			}
			print("\r\n" + preview)

		default:
			print("\r\nInvalid response. Please answer `y` (yes), `n` (no) or `p` (preview)")
		}
	}
}

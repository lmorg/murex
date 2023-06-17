package readline

import (
	"bytes"
	"fmt"
	"regexp"
	"sync/atomic"

	"github.com/lmorg/murex/utils/readline/unicode"
)

var rxMultiline = regexp.MustCompile(`[\r\n]+`)

// Readline displays the readline prompt.
// It will return a string (user entered data) or an error.
func (rl *Instance) Readline() (_ string, err error) {
	rl.fdMutex.Lock()
	rl.Active = true

	state, err := MakeRaw(int(replica.Fd()))
	rl.sigwinch()

	rl.fdMutex.Unlock()

	if err != nil {
		return "", fmt.Errorf("unable to modify fd %d: %s", replica.Fd(), err.Error())
	}

	defer func() {
		rl.fdMutex.Lock()

		rl.closeSigwinch()

		rl.Active = false
		// return an error if Restore fails. However we don't want to return
		// `nil` if there is no error because there might be a CtrlC or EOF
		// that needs to be returned
		r := Restore(int(replica.Fd()), state)
		if r != nil {
			err = r
		}

		rl.fdMutex.Unlock()
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

	rl.line.Set([]rune{})
	rl.line.SetRunePos(0)
	rl.lineChange = ""
	rl.viUndoHistory = []*unicode.UnicodeT{rl.line.Duplicate()}
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
		return rl.line.String(), nil
	}

	rl.termWidth = GetTermWidth()
	rl.getHintText()
	rl.renderHelpers()

	for {
		if rl.line.RuneLen() == 0 {
			// clear the cache when the line is cleared
			rl.cacheHint.Init(rl)
			rl.cacheSyntax.Init(rl)
		}

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
			return rl.line.String(), nil
		}

		s := string(r[:i])
		if rl.evtKeyPress[s] != nil {
			ret := rl.evtKeyPress[s](s, rl.line.Runes(), rl.line.RunePos())

			rl.clearPrompt()
			rl.line.Set(append(ret.NewLine, []rune{}...))
			rl.echo()
			// TODO: should this be above echo?
			rl.line.SetRunePos(ret.NewPos)

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
				return rl.line.String(), nil
			}
		}

		i = removeNonPrintableChars(b[:i])

		// Used for syntax completion
		rl.lineChange = string(b[:i])

		// Slow or invisible tab completions shouldn't lock up cursor movement
		rl.tabMutex.Lock()
		lenTcS := len(rl.tcSuggestions)
		rl.tabMutex.Unlock()
		if rl.modeTabCompletion && lenTcS == 0 {
			if rl.delayedTabContext.cancel != nil {
				rl.delayedTabContext.cancel()
			}
			rl.modeTabCompletion = false
			rl.updateHelpers()
		}

		switch b[0] {
		case charCtrlA:
			HkFnMoveToStartOfLine(rl)

		case charCtrlC:
			rl.clearHelpers()
			return "", CtrlC

		case charEOF:
			rl.clearHelpers()
			return "", EOF

		case charCtrlE:
			HkFnMoveToEndOfLine(rl)

		case charCtrlF:
			HkFnFuzzyFind(rl)

		case charCtrlK:
			HkFnClearAfterCursor(rl)

		case charCtrlL:
			HkFnClearScreen(rl)

		case charCtrlR:
			HkFnSearchHistory(rl)

		case charCtrlU:
			HkFnClearLine(rl)

		case charTab:
			HkFnAutocomplete(rl)

		case '\r':
			fallthrough
		case '\n':
			var suggestions []string
			rl.tabMutex.Lock()
			if rl.modeTabFind {
				suggestions = rl.tfSuggestions
			} else {
				suggestions = rl.tcSuggestions
			}
			rl.tabMutex.Unlock()

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
			return rl.line.String(), nil

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
			if rl.line.RunePos() == rl.line.RuneLen() && rl.line.RuneLen() > 0 {
				rl.line.SetRunePos(rl.line.RunePos() - 1)
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
		posX, posY := rl.lineWrapCellPos()
		if posY > 0 {
			pos := rl.line.CellPos() - rl.termWidth + rl.promptLen
			rl.line.SetCellPos(pos)

			newX, _ := rl.lineWrapCellPos()
			offset := newX - posX
			switch {
			case offset > 0:
				moveCursorForwards(offset)
			case offset < 0:
				moveCursorBackwards(offset * -1)
			}

			moveCursorUp(1)
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
		posX, posY := rl.lineWrapCellPos()
		_, lineY := rl.lineWrapCellLen()
		if posY < lineY {
			pos := rl.line.CellPos() + rl.termWidth - rl.promptLen
			rl.line.SetCellPos(pos)

			newX, _ := rl.lineWrapCellPos()
			offset := newX - posX
			switch {
			case offset > 0:
				moveCursorForwards(offset)
			case offset < 0:
				moveCursorBackwards(offset * -1)
			}

			moveCursorDown(1)
			return
		}

		rl.walkHistory(1)

	case seqBackwards:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(-1, 0)
			rl.renderHelpers()
			return
		}

		rl.moveCursorByRuneAdjust(-1)
		rl.viUndoSkipAppend = true

	case seqForwards:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(1, 0)
			rl.renderHelpers()
			return
		}

		//if (rl.modeViMode == vimInsert && rl.line.RunePos() < rl.line.RuneLen()) ||
		//	(rl.modeViMode != vimInsert && rl.line.RunePos() < rl.line.RuneLen()-1) {
		rl.moveCursorByRuneAdjust(1)
		//}
		rl.viUndoSkipAppend = true

	case seqHome, seqHomeSc:
		if rl.modeTabCompletion {
			return
		}

		rl.moveCursorByRuneAdjust(-rl.line.RunePos())
		rl.viUndoSkipAppend = true

	case seqEnd, seqEndSc:
		if rl.modeTabCompletion {
			return
		}

		rl.moveCursorByRuneAdjust(rl.line.RuneLen() - rl.line.RunePos())
		rl.viUndoSkipAppend = true

	case seqShiftTab:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(-1, 0)
			rl.renderHelpers()
			return
		}

	case seqPageUp:
		rl.previewPageUp()
		return

	case seqPageDown:
		rl.previewPageDown()
		return

	case seqF1, seqF1VT100:
		rl.ShowPreviews = !rl.ShowPreviews
		//rl.screenRefresh()
		return

	case seqAltF:
		HkFnJumpForwards(rl)

	case seqAltB:
		HkFnJumpBackwards(rl)

	default:
		if rl.modeTabFind /*|| rl.modeAutoFind*/ {
			//rl.modeTabFind = false
			//rl.modeAutoFind = false
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
	if len(r) == 0 {
		return
	}

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
		rl.histPos, err = rl.History.Write(rl.line.String())
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

		i, err := read(b)
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

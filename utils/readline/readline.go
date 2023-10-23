package readline

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync/atomic"
)

var rxMultiline = regexp.MustCompile(`[\r\n]+`)

// Readline displays the readline prompt.
// It will return a string (user entered data) or an error.
func (rl *Instance) Readline() (_ string, err error) {
	rl.fdMutex.Lock()
	rl.Active = true

	state, err := MakeRaw(int(os.Stdin.Fd()))
	rl.sigwinch()

	rl.fdMutex.Unlock()

	if err != nil {
		return "", fmt.Errorf("unable to modify fd %d: %s", os.Stdout.Fd(), err.Error())
	}

	defer func() {
		print(rl.clearPreviewStr())

		rl.fdMutex.Lock()

		rl.closeSigwinch()

		rl.Active = false
		// return an error if Restore fails. However we don't want to return
		// `nil` if there is no error because there might be a CtrlC or EOF
		// that needs to be returned
		r := Restore(int(os.Stdin.Fd()), state)
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

	rl.line.Set(rl, []rune{})
	rl.line.SetRunePos(0)
	rl.lineChange = ""
	rl.viUndoHistory = []*UnicodeT{rl.line.Duplicate()}
	rl.histPos = rl.History.Len()
	rl.modeViMode = vimInsert
	atomic.StoreInt32(&rl.delayedSyntaxCount, 0)
	rl.resetHintText()
	rl.resetTabCompletion()

	if len(rl.multiSplit) > 0 {
		r := []rune(rl.multiSplit[0])
		print(rl.readlineInputStr(r))
		print(rl.carriageReturnStr())
		if len(rl.multiSplit) > 1 {
			rl.multiSplit = rl.multiSplit[1:]
		} else {
			rl.multiSplit = []string{}
		}
		return rl.line.String(), nil
	}

	rl.termWidth = GetTermWidth()
	rl.getHintText()
	print(rl.renderHelpersStr())

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
			print(rl.readlineInputStr(r))
			print(rl.carriageReturnStr())
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
			rl.line.Set(rl, append(ret.NewLine, []rune{}...))
			print(rl.echoStr())
			// TODO: should this be above echo?
			rl.line.SetRunePos(ret.NewPos)

			if ret.ClearHelpers {
				rl.resetHelpers()
			} else {
				output := rl.updateHelpersStr()
				output += rl.renderHelpersStr()
				print(output)
			}

			if len(ret.HintText) > 0 {
				rl.hintText = ret.HintText
				output := rl.clearHelpersStr()
				output += rl.renderHelpersStr()
				print(output)
			}

			if ret.DisplayPreview {
				if rl.previewMode == previewModeClosed {
					HkFnPreviewToggle(rl)
				}
			}

			//rl.previewItem

			if ret.Callback != nil {
				err = ret.Callback()
				if err != nil {
					rl.hintText = []rune(err.Error())
					output := rl.clearHelpersStr()
					output += rl.renderHelpersStr()
					print(output)
				}
			}

			if !ret.ForwardKey {
				continue
			}
			if ret.CloseReadline {
				print(rl.clearHelpersStr())
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
			print(rl.updateHelpersStr())
		}

		switch b[0] {
		case charCtrlA:
			HkFnMoveToStartOfLine(rl)

		case charCtrlC:
			output := rl.clearPreviewStr()
			output += rl.clearHelpersStr()
			print(output)
			return "", CtrlC

		case charEOF:
			if rl.line.RuneLen() == 0 {
				output := rl.clearPreviewStr()
				output += rl.clearHelpersStr()
				print(output)
				return "", EOF
			}

		case charCtrlE:
			HkFnMoveToEndOfLine(rl)

		case charCtrlF:
			HkFnFuzzyFind(rl)

		case charCtrlG:
			HkFnCancelAction(rl)

		case charCtrlK:
			HkFnClearAfterCursor(rl)

		case charCtrlL:
			HkFnClearScreen(rl)

		case charCtrlR:
			HkFnSearchHistory(rl)

		case charCtrlU:
			HkFnClearLine(rl)

		case charCtrlZ:
			HkFnUndo(rl)

		case charTab:
			HkFnAutocomplete(rl)

		case '\r':
			fallthrough
		case '\n':
			var output string
			rl.tabMutex.Lock()
			var suggestions *suggestionsT
			if rl.modeTabFind {
				suggestions = newSuggestionsT(rl, rl.tfSuggestions)
			} else {
				suggestions = newSuggestionsT(rl, rl.tcSuggestions)
			}
			rl.tabMutex.Unlock()

			switch {
			case rl.previewMode == previewModeOpen:
				output += rl.clearPreviewStr()
				output += rl.clearHelpersStr()
				print(output)
				continue
			case rl.previewMode == previewModeAutocomplete:
				rl.previewMode = previewModeOpen
				if !rl.modeTabCompletion {
					output += rl.clearPreviewStr()
					output += rl.clearHelpersStr()
					print(output)
					continue
				}
			}

			if rl.modeTabCompletion || len(rl.tfLine) != 0 /*&& len(suggestions) > 0*/ {
				tfLine := rl.tfLine
				cell := (rl.tcMaxX * (rl.tcPosY - 1)) + rl.tcOffset + rl.tcPosX - 1
				output += rl.clearHelpersStr()
				rl.resetTabCompletion()
				output += rl.renderHelpersStr()
				if suggestions.Len() > 0 {
					prefix, line := suggestions.ItemCompletionReturn(cell)
					if len(prefix) == 0 && len(rl.tcPrefix) > 0 {
						l := -len(rl.tcPrefix)
						if l == -1 && rl.line.RuneLen() > 0 && rl.line.RunePos() == rl.line.RuneLen() {
							rl.line.Set(rl, rl.line.Runes()[:rl.line.RuneLen()-1])
						} else {
							output += rl.viDeleteByAdjustStr(l)
						}
					}
					output += rl.insertStr([]rune(line))
				} else {
					output += rl.insertStr(tfLine)
				}
				print(output)
				continue
			}
			output += rl.carriageReturnStr()
			print(output)
			return rl.line.String(), nil

		case charBackspace, charBackspace2:
			if rl.modeTabFind {
				print(rl.backspaceTabFindStr())
				rl.viUndoSkipAppend = true
			} else {
				print(rl.backspaceStr())
			}

		case charEscape:
			print(rl.escapeSeq(r[:i]))

		default:
			if rl.modeTabFind {
				print(rl.updateTabFindStr(r[:i]))
				rl.viUndoSkipAppend = true
			} else {
				print(rl.readlineInputStr(r[:i]))
				if len(rl.multiline) > 0 && rl.modeViMode == vimKeys {
					rl.skipStdinRead = true
				}
			}
		}

		rl.undoAppendHistory()
	}
}

func (rl *Instance) escapeSeq(r []rune) string {
	var output string
	switch string(r) {
	case seqEscape:
		HkFnCancelAction(rl)

	case seqDelete:
		if rl.modeTabFind {
			output += rl.backspaceTabFindStr()
		} else {
			output += rl.deleteStr()
		}

	case seqUp:
		rl.viUndoSkipAppend = true

		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(0, -1)
			output += rl.renderHelpersStr()
			return output
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
				output += moveCursorForwardsStr(offset)
			case offset < 0:
				output += moveCursorBackwardsStr(offset * -1)
			}

			output += moveCursorUpStr(1)
			return output
		}

		rl.walkHistory(-1)

	case seqDown:
		rl.viUndoSkipAppend = true

		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(0, 1)
			output += rl.renderHelpersStr()
			return output
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
				output += moveCursorForwardsStr(offset)
			case offset < 0:
				output += moveCursorBackwardsStr(offset * -1)
			}

			output += moveCursorDownStr(1)
			return output
		}

		rl.walkHistory(1)

	case seqBackwards:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(-1, 0)
			output += rl.renderHelpersStr()
			return output
		}

		output += rl.moveCursorByRuneAdjustStr(-1)
		rl.viUndoSkipAppend = true

	case seqForwards:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(1, 0)
			output += rl.renderHelpersStr()
			return output
		}

		//if (rl.modeViMode == vimInsert && rl.line.RunePos() < rl.line.RuneLen()) ||
		//	(rl.modeViMode != vimInsert && rl.line.RunePos() < rl.line.RuneLen()-1) {
		output += rl.moveCursorByRuneAdjustStr(1)
		//}
		rl.viUndoSkipAppend = true

	case seqHome, seqHomeSc:
		if rl.modeTabCompletion {
			return output
		}

		output += rl.moveCursorByRuneAdjustStr(-rl.line.RunePos())
		rl.viUndoSkipAppend = true

	case seqEnd, seqEndSc:
		if rl.modeTabCompletion {
			return output
		}

		output += rl.moveCursorByRuneAdjustStr(rl.line.RuneLen() - rl.line.RunePos())
		rl.viUndoSkipAppend = true

	case seqShiftTab:
		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(-1, 0)
			output += rl.renderHelpersStr()
			return output
		}

	case seqPageUp, seqOptUp, seqCtrlUp:
		output += rl.previewPageUpStr()
		return output

	case seqPageDown, seqOptDown, seqCtrlDown:
		output += rl.previewPageDownStr()
		return output

	case seqF1, seqF1VT100:
		HkFnPreviewToggle(rl)
		return output

	case seqF9:
		HkFnPreviewLine(rl)
		return output

	case seqAltF, seqOptRight, seqCtrlRight:
		HkFnJumpForwards(rl)

	case seqAltB, seqOptLeft, seqCtrlLeft:
		HkFnJumpBackwards(rl)

		// TODO: test me
	case seqShiftF1:
		HkFnRecallWord1(rl)
	case seqShiftF2:
		HkFnRecallWord2(rl)
	case seqShiftF3:
		HkFnRecallWord3(rl)
	case seqShiftF4:
		HkFnRecallWord4(rl)
	case seqShiftF5:
		HkFnRecallWord5(rl)
	case seqShiftF6:
		HkFnRecallWord6(rl)
	case seqShiftF7:
		HkFnRecallWord7(rl)
	case seqShiftF8:
		HkFnRecallWord8(rl)
	case seqShiftF9:
		HkFnRecallWord9(rl)
	case seqShiftF10:
		HkFnRecallWord10(rl)
	case seqShiftF11:
		HkFnRecallWord11(rl)
	case seqShiftF12:
		HkFnRecallWord12(rl)

	default:
		if rl.modeTabFind /*|| rl.modeAutoFind*/ {
			//rl.modeTabFind = false
			//rl.modeAutoFind = false
			return output
		}
		// alt+numeric append / delete
		if len(r) == 2 && '1' <= r[1] && r[1] <= '9' {
			if rl.modeViMode == vimDelete {
				output += rl.vimDeleteStr(r)
				return output
			}

		} else {
			rl.viUndoSkipAppend = true
		}
	}

	return output
}

// readlineInput is an unexported function used to determine what mode of text
// entry readline is currently configured for and then update the line entries
// accordingly.
func (rl *Instance) readlineInputStr(r []rune) string {
	if len(r) == 0 {
		return ""
	}

	var output string

	switch rl.modeViMode {
	case vimKeys:
		rl.vi(r[0])
		output += rl.viHintMessageStr()

	case vimDelete:
		output += rl.vimDeleteStr(r)
		output += rl.viHintMessageStr()

	case vimReplaceOnce:
		rl.modeViMode = vimKeys
		output += rl.deleteStr()
		output += rl.insertStr([]rune{r[0]})
		output += rl.viHintMessageStr()

	case vimReplaceMany:
		for _, char := range r {
			output += rl.deleteStr()
			output += rl.insertStr([]rune{char})
		}
		output += rl.viHintMessageStr()

	default:
		output += rl.insertStr(r)
	}

	return output
}

// SetPrompt will define the readline prompt string.
// It also calculates the runes in the string as well as any non-printable
// escape codes.
func (rl *Instance) SetPrompt(s string) {
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\t", "    ")
	split := strings.Split(s, "\n")
	if len(split) > 1 {
		print(strings.Join(split[:len(split)-1], "\r\n") + "\r\n")
		s = split[len(split)-1]
	}
	rl.prompt = s
	rl.promptLen = strLen(s)
}

func (rl *Instance) carriageReturnStr() string {
	output := rl.clearHelpersStr()
	output += "\r\n"
	if rl.HistoryAutoWrite {
		var err error
		rl.histPos, err = rl.History.Write(rl.line.String())
		if err != nil {
			output += err.Error() + "\r\n"
		}
	}
	return output
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
	print(rl.clearHelpersStr())
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
			print(moveCursorUpStr(2))
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

package readline

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

var rxMultiline *regexp.Regexp = regexp.MustCompile(`[\r\n]+`)

// Readline displays the readline prompt.
// It will return a string (user entered data) or an error.
func (rl *Instance) Readline() (string, error) {
	fd := int(os.Stdin.Fd())
	state, err := MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer Restore(fd, state)

	fmt.Print(rl.prompt)

	rl.line = []rune{}
	rl.viUndoHistory = make([][]rune, 1)
	rl.pos = 0
	rl.histPos = rl.History.Len()
	rl.modeViMode = vimInsert

	if len(rl.multisplit) > 0 {
		b := []byte(rl.multisplit[0])
		rl.editorInput(b)
		rl.carridgeReturn()
		if len(rl.multisplit) > 1 {
			rl.multisplit = rl.multisplit[1:]
		} else {
			rl.multisplit = []string{}
		}
		return string(rl.line), nil
	}

	for {
		rl.viUndoSkipAppend = false
		b := make([]byte, 1024)

		i, err := os.Stdin.Read(b)
		if err != nil {
			return "", err
		}

		if isMultiline(b[:i]) || len(rl.multiline) > 0 {
			rl.multiline = append(rl.multiline, b[:i]...)
			if i == len(b) {
				continue
			}

			if !rl.allowMultiline(rl.multiline) {
				rl.multiline = []byte{}
				continue
			}

			s := string(rl.multiline)
			rl.multisplit = rxMultiline.Split(s, -1)

			b = []byte(rl.multisplit[0])
			rl.editorInput(b)
			rl.carridgeReturn()
			rl.multiline = []byte{}
			if len(rl.multisplit) > 1 {
				rl.multisplit = rl.multisplit[1:]
			} else {
				rl.multisplit = []string{}
			}
			return string(rl.line), nil
		}

		s := string(b[:i])

		if rl.evtKeyPress[s] != nil {
			ignoreKey, closeReadline, hintText := rl.evtKeyPress[s](s, rl.line, rl.pos)
			//getTermWidth()
			//rl.renderHintText()
			//rl.renderSuggestions()
			if len(hintText) > 0 {
				rl.pos--
				rl.writeHintText(hintText)
				rl.pos++
			}
			if ignoreKey {
				continue
			}
			if closeReadline {
				rl.clearHelpers()
				return string(rl.line), nil
			}
		}

		switch b[0] {
		case charCtrlC:
			rl.clearHelpers()
			return "", errors.New(ErrCtrlC)

		case charEOF:
			rl.clearHelpers()
			return "", errors.New(ErrEOF)

		case charTab:
			if rl.modeTabGrid {
				rl.moveTabHighlight(1, 0)
				continue
			}
			rl.tabCompletion()
			rl.viUndoSkipAppend = true

		case charCtrlU:
			//clearHintText()
			//clearLine()
			rl.clearTabSuggestions()
			moveCursorBackwards(rl.pos)
			fmt.Print(strings.Repeat(" ", len(rl.line)))
			//moveCursorBackwards(len(line))

			moveCursorBackwards(len(rl.line))
			rl.line = rl.line[rl.pos:]
			rl.pos = 0
			rl.echo()

			moveCursorBackwards(1)

		case '\r':
			fallthrough
		case '\n':
			if rl.modeTabGrid {
				cell := (rl.tcMaxX * (rl.tcPosY - 1)) + rl.tcPosX - 1
				rl.clearTabSuggestions()
				rl.insert([]byte(rl.tcSuggestions[cell]))
				continue
			}
			rl.carridgeReturn()
			return string(rl.line), nil

		case charBackspace, charBackspace2:
			rl.backspace()

		case charEscape:
			rl.escapeSeq(b[:i])

		default:
			rl.editorInput(b[:i])
		}

		if !rl.viUndoSkipAppend {
			rl.viUndoHistory = append(rl.viUndoHistory, rl.line)
		}
	}
}

func (rl *Instance) escapeSeq(b []byte) {
	switch string(b) {
	case string(charEscape):
		if rl.modeTabGrid {
			rl.clearTabSuggestions()
		} else {
			if rl.pos == len(rl.line) && len(rl.line) > 0 {
				rl.pos--
				moveCursorBackwards(1)
			}
			rl.modeViMode = vimKeys
			rl.viIteration = ""
		}
		rl.viUndoSkipAppend = true

	case seqDelete:
		rl.delete()

	case seqUp:
		if rl.modeTabGrid {
			rl.moveTabHighlight(0, -1)
			return
		}
		rl.walkHistory(-1)

	case seqDown:
		if rl.modeTabGrid {
			rl.moveTabHighlight(0, 1)
			return
		}
		rl.walkHistory(1)

	case seqBackwards:
		if rl.modeTabGrid {
			rl.moveTabHighlight(-1, 0)
			return
		}
		if rl.pos > 0 {
			moveCursorBackwards(1)
			rl.pos--
		}
		rl.viUndoSkipAppend = true

	case seqForwards:
		if rl.modeTabGrid {
			rl.moveTabHighlight(1, 0)
			return
		}
		if (rl.modeViMode == vimInsert && rl.pos < len(rl.line)) ||
			(rl.modeViMode != vimInsert && rl.pos < len(rl.line)-1) {
			//if pos < len(line) {
			moveCursorForwards(1)
			rl.pos++
		}
		rl.viUndoSkipAppend = true

	case seqHome, seqHomeSc:
		if rl.modeTabGrid {
			return
		}
		moveCursorBackwards(rl.pos)
		rl.pos = 0
		rl.viUndoSkipAppend = true

	case seqEnd, seqEndSc:
		if rl.modeTabGrid {
			return
		}
		moveCursorForwards(len(rl.line) - rl.pos)
		rl.pos = len(rl.line)
		rl.viUndoSkipAppend = true

	default:
		rl.viUndoSkipAppend = true
	}
}

// editorInput is an unexported function used to determine what mode
// of text entry readline is currently configured for and then update
// the line entries accordingly.
func (rl *Instance) editorInput(b []byte) {
	switch rl.modeViMode {
	case vimKeys:
		rl.vi(b[0])

	case vimReplaceOnce:
		rl.modeViMode = vimKeys
		rl.delete()
		r := []rune(string(b))
		rl.insert([]byte(string(r[0])))

	case vimReplaceMany:
		for _, r := range []rune(string(b)) {
			rl.delete()
			rl.insert([]byte(string(r)))
		}

	default:
		rl.insert(b)
	}

	rl.syntaxCompletion()
}

func (rl *Instance) echo() {
	moveCursorBackwards(rl.pos)

	switch {
	case rl.PasswordMask > 0:
		fmt.Print(strings.Repeat(string(rl.PasswordMask), len(rl.line)) + " ")

	case rl.SyntaxHighlighter == nil:
		fmt.Print(string(rl.line) + " ")

	default:
		fmt.Print(rl.SyntaxHighlighter(rl.line) + " ")
	}

	moveCursorBackwards(len(rl.line) - rl.pos)
	rl.renderHintText()
}

// SetPrompt will define the readline prompt string.
// It also calculates the runes in the string as well as any non-printable escape codes.
func (rl *Instance) SetPrompt(s string) {
	rl.prompt = s

	s = rxAnsiEscSeq.ReplaceAllString(s, "")
	rl.promptLen = utf8.RuneCountInString(s)
}

func (rl *Instance) carridgeReturn() {
	rl.clearHintText()
	fmt.Print("\r\n")
	if rl.HistoryAutoWrite {
		var err error
		rl.histPos, err = rl.History.Write(string(rl.line))
		if err != nil {
			fmt.Print(err.Error() + "\r\n")
		}
	}
}

func isMultiline(b []byte) bool {
	for i := range b {
		if (b[i] == '\r' || b[i] == '\n') && i != len(b)-1 {
			return true
		}
	}
	return false
}

func (rl *Instance) allowMultiline(data []byte) bool {
	fmt.Printf("\r\nWARNING: %d bytes of multiline data was dumped into the shell!", len(data))
	for {
		fmt.Print("\r\nDo you wish to proceed (yes|no|preview)? [y/n/p] ")

		b := make([]byte, 1024)

		i, err := os.Stdin.Read(b)
		if err != nil {
			return false
		}

		s := string(b[:i])
		fmt.Print(s)

		switch s {
		case "y", "Y":
			fmt.Print("\r\n" + rl.prompt)
			return true

		case "n", "N":
			fmt.Print("\r\n" + rl.prompt)
			return false

		case "p", "P":
			preview := bytes.Replace(data, []byte{'\r'}, []byte{'\r', '\n'}, -1)
			fmt.Print("\r\n" + string(preview))

		default:
			fmt.Print("\r\nInvalid response. Please answer `y` (yes), `n` (no) or `p` (preview)")
		}
	}
}

package readline

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/ssh/terminal"
)

// Readline displays the readline prompt.
// It will return a string (user entered data) or an error.
func (rl *Instance) Readline() (string, error) {
	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer terminal.Restore(fd, state)

	fmt.Print(rl.prompt)

	rl.line = []rune{}
	rl.pos = 0
	rl.histPos = rl.History.Len()
	rl.modeViMode = vimInsert

	for {
		b := make([]byte, 1024)
		i, err := os.Stdin.Read(b)
		if err != nil {
			return "", err
		}

		s := string(b[:i])
		if rl.evtKeyPress[s] != nil {
			ignoreKey, closeReadline := rl.evtKeyPress[s](s, rl.line, rl.pos)
			//getTermWidth()
			//rl.renderHintText()
			//rl.renderSuggestions()
			if ignoreKey {
				continue
			}
			if closeReadline {
				return string(rl.line), nil
			}
		}

		switch b[0] {
		case charCtrlC:
			if rl.modeTabGrid {
				rl.clearTabSuggestions()
			}
			return "", errors.New(ErrCtrlC)

		case charEOF:
			if rl.modeTabGrid {
				rl.clearTabSuggestions()
			}
			rl.clearHintText()
			return "", errors.New(ErrEOF)

		case charTab:
			if rl.modeTabGrid {
				rl.moveTabHighlight(1, 0)
				continue
			}
			rl.tabCompletion()

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
			rl.clearHintText()
			fmt.Print("\r\n")
			if rl.HistoryAutoWrite {
				rl.histPos, err = rl.History.Write(string(rl.line))
				if err != nil {
					fmt.Print(err.Error() + "\r\n")
				}
			}
			return string(rl.line), nil

		case charBackspace:
			rl.backspace()

		case charEscape:
			rl.escapeSeq(b[:i])

		default:
			rl.editorInput(b[:i])
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
		//renderHintText()

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
		//renderHintText()

	case seqHome, seqHomeSc:
		if rl.modeTabGrid {
			return
		}
		moveCursorBackwards(rl.pos)
		rl.pos = 0

	case seqEnd, seqEndSc:
		if rl.modeTabGrid {
			return
		}
		moveCursorForwards(len(rl.line) - rl.pos)
		rl.pos = len(rl.line)
	}
}

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

func (rl *Instance) SetPrompt(s string) {
	rl.prompt = s

	s = rxAnsiEscSeq.ReplaceAllString(s, "")
	rl.promptLen = utf8.RuneCountInString(s)
}

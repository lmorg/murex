package readline

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"unicode/utf8"
)

/*var (
	// PasswordMask is what character to hide password entry behind.
	// Once enabled, set to 0 (zero) to disable the mask again.
	PasswordMask rune

	// SyntaxHighlight is a helper function to provide syntax highlighting.
	// Once enabled, set to nil to disable again.
	SyntaxHighlighter func([]rune) string

	// History is an interface for querying the readline history.
	// This is exposed as an interface to allow you the flexibility to define how
	// you want your history managed (eg file on disk, database, cloud, or even no
	// history at all). By default it uses a dummy interface that only stores
	// historic items in memory.
	History LineHistory

	// HistoryAutoWrite defines whether items automatically get written to history.
	// Enabled by default. Set to false to disable.
	HistoryAutoWrite bool = true

	// TabCompleter is a simple function that offers completion suggestions.
	// It takes the readline line ([]rune) and cursor pos. Returns a prefix string
	// and an array of suggestions.
	TabCompleter func([]rune, int) (string, []string)

	// MaxTabCompletionRows is the maximum number of rows to display in the tab
	// completion grid.
	MaxTabCompleterRows int = 4

	// SyntaxCompletion is used to autocomplete code syntax (like braces and
	// quotation marks). If you want to complete words or phrases then you might
	// be better off using the TabCompletion function.
	// SyntaxCompletion takes the line ([]rune) and cursor position, and returns
	// the new line and cursor position.
	SyntaxCompleter func([]rune, int) ([]rune, int)

	// HintText is a helper function which displays hint text the prompt.
	// HintText takes the line input from the promt and the cursor position.
	// It returns the hint text to display.
	HintText func([]rune, int) []rune
)*/

// While it might normally seem bad practice to have global variables, you canot
// have two concurrent readline prompts anyway due to limitations in the way
// terminal emulators work. So storing these values as globals simplifies the
// API design immencely without sacricing functionality.
/*var (
	prompt      string = ">>> "
	promptLen   int    = 4
	line        []rune
	lineBuf     []rune
	pos         int
	histPos     int
	modeTabGrid bool
)*/

// Readline displays the readline prompt.
// It will return a string (user entered data) or an error.
func (rl *instance) Readline() (string, error) {
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

		//renderHintText()

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
			moveCursorBackwards(pos)
			fmt.Print(strings.Repeat(" ", len(rl.line)))
			//moveCursorBackwards(len(line))

			moveCursorBackwards(len(rl.line))
			rl.line = rl.line[pos:]
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
				rl.histPos, err = rl.History.Write(string(line))
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

func (rl *instance) escapeSeq(b []byte) {
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

	case seqHome:
		if rl.modeTabGrid {
			return
		}
		moveCursorBackwards(rl.pos)
		rl.pos = 0

	case seqEnd:
		if rl.modeTabGrid {
			return
		}
		moveCursorForwards(len(rl.line) - rl.pos)
		rl.pos = len(rl.line)
	}
}

func (rl *instance) editorInput(b []byte) {
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

func (rl *instance) echo() {
	moveCursorBackwards(pos)

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

func (rl *instance) SetPrompt(s string) {
	rl.prompt = s

	s = rxAnsiEscSeq.ReplaceAllString(s, "")
	rl.promptLen = utf8.RuneCountInString(s)
}

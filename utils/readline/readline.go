package readline

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
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

	HintText func([]rune) []rune
)

// While it might normally seem bad practice to have global variables, you canot
// have two concurrent readline prompts anyway due to limitations in the way
// terminal emulators work. So storing these values as globals simplifies the
// API design immencely without sacricing functionality.
var (
	prompt       string         = ">>> "
	promptLen    int            = 4
	rxAnsiEscSeq *regexp.Regexp = regexp.MustCompile("\x1b\\[[0-9]+[a-zA-Z]")
	line         []rune
	lineBuf      []rune
	pos          int
	histPos      int
	modeTabGrid  bool
	modeViKeys   bool
)

func init() {
	History = new(dummyLineHistory)
}

// Readline displays the readline prompt.
// It will return a string (user entered data) or an error.
func Readline() (string, error) {
	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer terminal.Restore(fd, state)

	fmt.Print(prompt)

	line = []rune{}
	pos = 0
	histPos = History.Len()

	for {
		b := make([]byte, 1024)
		i, err := os.Stdin.Read(b)
		if err != nil {
			return "", err
		}

		switch b[0] {
		case charCtrlC:
			if modeTabGrid {
				clearTabSuggestions()
			}
			return "", errors.New(ErrCtrlC)
		case charEOF:
			if modeTabGrid {
				clearTabSuggestions()
			}
			return "", errors.New(ErrEOF)
		case charTab:
			if modeTabGrid {
				moveTabHighlight(1, 0)
				continue
			}
			tabCompletion()
		case charCtrlU:
			clearLine()
		case '\r':
			fallthrough
		case '\n':
			if modeTabGrid {
				cell := (tcMaxX * (tcPosY - 1)) + tcPosX - 1
				clearTabSuggestions()
				insert([]byte(tcSuggestions[cell]))
				continue
			}
			fmt.Print("\r\n")
			moveCursorDown(hintY)
			if HistoryAutoWrite {
				histPos, err = History.Write(string(line))
				if err != nil {
					fmt.Print(err.Error() + "\r\n")
				}
			}
			hintY = 0
			modeViKeys = false
			return string(line), nil
		case charBackspace:
			backspace()
		case charEscape:
			escapeSeq(b[:i])
		default:
			if modeViKeys {
				vi(b[0])
			} else {
				insert(b[:i])
			}
			syntaxCompletion()
		}
	}
}

func escapeSeq(b []byte) {
	switch string(b) {
	case string(charEscape):
		if modeTabGrid {
			clearTabSuggestions()
		} else {
			fmt.Print("\r\nvi mode enabled\r\n")
			if pos == len(line) {
				pos--
			}
			fmt.Print(string(prompt))
			moveCursorForwards(pos)
			echo()
			moveCursorBackwards(1)
			modeViKeys = true
		}

	case seqDelete:
		delete()

	case seqUp:
		if modeTabGrid {
			moveTabHighlight(0, -1)
			return
		}
		walkHistory(-1)

	case seqDown:
		if modeTabGrid {
			moveTabHighlight(0, 1)
			return
		}
		walkHistory(1)

	case seqBackwards:
		if modeTabGrid {
			moveTabHighlight(-1, 0)
			return
		}
		if pos > 0 {
			moveCursorBackwards(1)
			pos--
		}

	case seqForwards:
		if modeTabGrid {
			moveTabHighlight(1, 0)
			return
		}
		if (!modeViKeys && pos < len(line)) ||
			(modeViKeys && pos < len(line)-1) {
			moveCursorForwards(1)
			pos++
		}

	case seqHome:
		if modeTabGrid {
			return
		}
		moveCursorBackwards(pos)
		pos = 0

	case seqEnd:
		if modeTabGrid {
			return
		}
		moveCursorForwards(len(line) - pos)
		pos = len(line)
	}
}

func echo() {
	moveCursorBackwards(pos)

	switch {
	case PasswordMask > 0:
		fmt.Print(strings.Repeat(string(PasswordMask), len(line)) + " ")

	case SyntaxHighlighter == nil:
		fmt.Print(string(line) + " ")

	default:
		fmt.Print(SyntaxHighlighter(line) + " ")
	}

	moveCursorBackwards(len(line) - pos)
	renderHintText()
}

func SetPrompt(s string) {
	prompt = s

	s = rxAnsiEscSeq.ReplaceAllString(s, "")
	promptLen = utf8.RuneCountInString(s)
}

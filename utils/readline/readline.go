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

	// HintText is a helper function which displays hint text the prompt.
	// HintText takes the line input from the promt and the cursor position.
	// It returns the hint text to display.
	HintText func([]rune, int) []rune
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
)

func init() {
	History = new(ExampleLineHistory)
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
	modeViMode = vimInsert

	for {
		b := make([]byte, 1024)
		i, err := os.Stdin.Read(b)
		if err != nil {
			return "", err
		}

		//renderHintText()

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
			clearHintText()
			return "", errors.New(ErrEOF)

		case charTab:
			if modeTabGrid {
				moveTabHighlight(1, 0)
				continue
			}
			tabCompletion()

		case charCtrlU:
			//clearHintText()
			//clearLine()
			moveCursorBackwards(pos)
			fmt.Print(strings.Repeat(" ", len(line)))
			//moveCursorBackwards(len(line))

			moveCursorBackwards(len(line))
			line = line[pos:]
			pos = 0
			echo()

			moveCursorBackwards(1)

		case '\r':
			fallthrough
		case '\n':
			if modeTabGrid {
				cell := (tcMaxX * (tcPosY - 1)) + tcPosX - 1
				clearTabSuggestions()
				insert([]byte(tcSuggestions[cell]))
				continue
			}
			clearHintText()
			fmt.Print("\r\n")
			if HistoryAutoWrite {
				histPos, err = History.Write(string(line))
				if err != nil {
					fmt.Print(err.Error() + "\r\n")
				}
			}
			return string(line), nil

		case charBackspace:
			backspace()

		case charEscape:
			escapeSeq(b[:i])

		default:
			editorInput(b[:i])
		}
	}
}

func escapeSeq(b []byte) {
	switch string(b) {
	case string(charEscape):
		if modeTabGrid {
			clearTabSuggestions()
		} else {
			if pos == len(line) && len(line) > 0 {
				pos--
				moveCursorBackwards(1)
			}
			modeViMode = vimKeys
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
		//renderHintText()

	case seqForwards:
		if modeTabGrid {
			moveTabHighlight(1, 0)
			return
		}
		if (modeViMode == vimInsert && pos < len(line)) ||
			(modeViMode != vimInsert && pos < len(line)-1) {
			//if pos < len(line) {
			moveCursorForwards(1)
			pos++
		}
		//renderHintText()

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

func editorInput(b []byte) {
	switch modeViMode {
	case vimKeys:
		vi(b[0])

	case vimReplaceOnce:
		modeViMode = vimKeys
		delete()
		r := []rune(string(b))
		insert([]byte(string(r[0])))

	case vimReplaceMany:
		for _, r := range []rune(string(b)) {
			delete()
			insert([]byte(string(r)))
		}

	default:
		insert(b)
	}

	syntaxCompletion()
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

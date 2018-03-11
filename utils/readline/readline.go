package readline

import (
	"errors"
	"fmt"
	"github.com/lmorg/murex/utils/ansi"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

var (
	// Prompt is the readline prompt
	Prompt string = ">>> "

	// PasswordMask is what character to hide password entry behind.
	// Once enabled, set to 0 (zero) to disable the mask again.
	PasswordMask rune

	// SyntaxHighlight is a helper function to provide syntax highlighting.
	// Once enabled, set to nil to disable again.
	SyntaxHighlight func([]rune) string

	// History is an interface for querying the readline history.
	// This is exposed as an interface to allow you the flexibility to define how
	// you want your history managed (eg file on disk, database, cloud, or even no
	// history at all). By default it uses a dummy interface that only stores
	// historic items in memory.
	History LineHistory
)

// While it might normally seem bad practice to have global variables, you canot
// have two concurrent readline prompts anyway due to limitations in the way
// terminal emulators work. So storing these values as globals simplifies the
// API design immencely without sacricing functionality.
var (
	line    []rune
	lineBuf []rune
	pos     int
	histPos int
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

	fmt.Print(Prompt)

	line = []rune{}
	pos = 0

	for {
		b := make([]byte, 1024)
		i, err := os.Stdin.Read(b)
		if err != nil {
			return "", err
		}

		//fmt.Print(b[:i])
		switch b[0] {
		case CtrlC:
			return "", errors.New(ErrCtrlC)
		case CtrlU:
			clearLine()
		case '\r':
			fallthrough
		case '\n':
			fmt.Print("\r\n")
			histPos, err = History.Append(string(line))
			if err != nil {
				fmt.Print(err.Error() + "\r\n")
			}
			return string(line), nil
		case Backspace:
			backspace()
		case Escape:
			escapeSequ(b[:i])
		default:
			insert(b[:i])
		}
	}
}

func echo() {
	if pos > 0 {
		fmt.Printf("\x1b[%dD", pos)
	}

	switch {
	case PasswordMask > 0:
		fmt.Print(strings.Repeat(string(PasswordMask), len(line)) + " ")

	case SyntaxHighlight == nil:
		fmt.Print(string(line) + " ")

	default:
		fmt.Print(SyntaxHighlight(line) + " ")
	}

	if pos < len(line) {
		fmt.Printf("\x1b[%dD", len(line)-pos)
	}
}

func escapeSequ(b []byte) {
	switch string(b) {
	case seqDelete:
		delete()
	case seqUp:
		walkHistory(-1)
	case seqDown:
		walkHistory(1)
	case seqBackwards:
		if pos > 0 {
			fmt.Print(ansi.Backwards)
			pos--
		}
	case seqForwards:
		if pos < len(line) {
			fmt.Print(ansi.Forwards)
			pos++
		}
	}
}

func backspace() {
	if len(line) == 0 || pos == 0 {
		return
	}

	fmt.Print(ansi.Backwards)
	pos--
	delete()
}

func insert(b []byte) {
	r := []rune(string(b))
	switch {
	case len(line) == 0:
		line = r
	case pos == 0:
		line = append(r, line...)
	case pos < len(line):
		r := append(r, line[pos:]...)
		line = append(line[:pos], r...)
	default:
		line = append(line, r...)
	}
	echo()
	pos++
}

func delete() {
	switch {
	case len(line) == 0:
		return
	case pos == 0:
		line = line[1:]
		echo()
		fmt.Print(ansi.Backwards)
	case pos > len(line):
		backspace()
	case pos == len(line):
		line = line[:pos]
		echo()
		fmt.Print(ansi.Backwards)
	default:
		line = append(line[:pos], line[pos+1:]...)
		echo()
		fmt.Print(ansi.Backwards)
	}
}

func clearLine() {
	if len(line) == 0 {
		return
	}

	if pos > 0 {
		fmt.Printf("\x1b[%dD", pos)
	}

	fmt.Print(strings.Repeat(" ", len(line)))
	fmt.Printf("\x1b[%dD", len(line))

	line = []rune{}
	pos = 0
}

func walkHistory(i int) {
	switch histPos + i {
	case -1, History.Len() + 1:
		return

	case History.Len():
		clearLine()
		histPos += i
		line = lineBuf

	default:
		s, err := History.GetLine(histPos + i)
		if err != nil {
			fmt.Print("\r\n" + err.Error() + "\r\n")
			fmt.Print(Prompt)
			return
		}

		if histPos == History.Len() {
			lineBuf = append(line, []rune{}...)
		}

		clearLine()
		histPos += i
		line = []rune(s)
	}

	echo()
	pos = len(line)
	if pos > 1 {
		fmt.Printf("\x1b[%dC", pos-1)
	} else if pos == 0 {
		fmt.Print("\x1b[1D")
	}
}

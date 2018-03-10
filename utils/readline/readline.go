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
	Prompt          string
	Echo            bool
	PasswordMask    string
	SyntaxHighlight func(string) string
)

var (
	line    string
	pos     int
	history []string
	histPos int
)

func Readline() (string, error) {
	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer terminal.Restore(fd, state)

	fmt.Print(Prompt)

	line = ""
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
			history = append(history, line)
			histPos = len(history)
			return line, nil
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

	if SyntaxHighlight == nil {
		fmt.Print(line + " ")
	} else {
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
		if histPos > 0 {
			histPos--
		}
		clearLine()
		line = history[histPos]
		echo()
		pos = len(line)
		fmt.Printf("\x1b[%dC", pos-1)
	case seqDown:
		if histPos < len(history)-1 {
			histPos++
		}
		clearLine()
		line = history[histPos]
		echo()
		pos = len(line)
		fmt.Printf("\x1b[%dC", pos-1)
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
	switch {
	case len(line) == 0:
		line = string(b)
		echo()
	case pos == 0:
		line = string(b) + line
		echo()
	case pos < len(line):
		line = line[:pos] + string(b) + line[pos:]
		echo()
	default:
		line += string(b)
		echo()
	}
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
		line = line[:pos] + line[pos+1:]
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
	line = ""
	pos = 0
}

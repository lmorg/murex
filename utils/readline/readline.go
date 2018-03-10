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
	Prompt       string
	Echo         bool
	PasswordMask string
)

var (
	line string
	pos  int
)

func Readline() (string, error) {
	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer terminal.Restore(fd, state)

	fmt.Print("\r" + Prompt)

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
			//fmt.Print(ansi.ClearLine)
			fmt.Printf("\x1b[%dD", len(line))
			fmt.Print(strings.Repeat(" ", len(line)))
			fmt.Printf("\x1b[%dD", len(line))
			line = ""
			pos = 0
		case '\r':
			fallthrough
			//fmt.Println("|")
		case '\n':
			fmt.Print("\r\n") // + ansi.ClearLine)
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

func insert(b []byte) {
	switch {
	case len(line) == 0:
		fmt.Print(string(b))
		line = string(b)
	case pos == 0:
		fmt.Print(string(b) + line)
		fmt.Printf("\x1b[%dD", len(line))
		line = string(b) + line
	case pos < len(line):
		fmt.Print(string(b) + line[pos:])
		fmt.Printf("\x1b[%dD", len(line[pos:]))
		line = line[:pos] + string(b) + line[pos:]
	default:
		fmt.Print(string(b))
		line += string(b)
	}
	pos++
}

func backspace() {
	if len(line) == 0 || pos == 0 {
		return
	}

	fmt.Print(ansi.Backwards)
	pos--
	delete()
}

func delete() {
	switch {
	/*if len(line) > 0 {
		fmt.Print(ansi.Backwards + " " + ansi.Backwards)
		line = line[:len(line)-1]
	}*/
	case len(line) == 0:
		return
	case pos == 0:
		line = line[1:]
		fmt.Print(line + " ")
		fmt.Printf("\x1b[%dD", len(line)+1)
	case pos > len(line):
		backspace()
	case pos == len(line):
		line = line[:pos]
		fmt.Print(" " + ansi.Backwards)
	default:
		fmt.Print(line[pos+1:] + " ")
		fmt.Printf("\x1b[%dD", len(line[pos:]))
		line = line[:pos] + line[pos+1:]
	}
}

func escapeSequ(b []byte) {
	switch string(b) {
	case up:
	case down:
	case backwards:
		if pos > 0 {
			fmt.Print(ansi.Backwards)
			pos--
		}
	case forwards:
		if pos < len(line) {
			fmt.Print(ansi.Forwards)
			pos++
		}
	}
}

package readline

import (
	"fmt"
	"strings"
)

func moveCursorUp(i int) {
	if i < 1 {
		return
	}

	fmt.Printf("\x1b[%dA", i)
}

func moveCursorForwards(i int) {
	if i < 1 {
		return
	}

	fmt.Printf("\x1b[%dC", i)
}

func moveCursorBackwards(i int) {
	if i < 1 {
		return
	}

	fmt.Printf("\x1b[%dD", i)
}

func backspace() {
	if len(line) == 0 || pos == 0 {
		return
	}

	moveCursorBackwards(1)
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

	moveCursorForwards(len(r) - 1)
	pos += len(r)

	if mode == modeTabCompletion {
		clearTabSuggestions()
		tabCompletion()
	}
}

func delete() {
	switch {
	case len(line) == 0:
		return
	case pos == 0:
		line = line[1:]
		echo()
		moveCursorBackwards(1)
	case pos > len(line):
		backspace()
	case pos == len(line):
		line = line[:pos]
		echo()
		moveCursorBackwards(1)
	default:
		line = append(line[:pos], line[pos+1:]...)
		echo()
		moveCursorBackwards(1)
	}

	if mode == modeTabCompletion {
		clearTabSuggestions()
		tabCompletion()
	}
}

func clearLine() {
	if len(line) == 0 {
		return
	}

	moveCursorBackwards(pos)
	fmt.Print(strings.Repeat(" ", len(line)))
	moveCursorBackwards(len(line))

	line = []rune{}
	pos = 0

	if mode == modeTabCompletion {
		clearTabSuggestions()
		tabCompletion()
	}
}

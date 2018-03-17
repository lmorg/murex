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

func moveCursorDown(i int) {
	if i < 1 {
		return
	}

	fmt.Printf("\x1b[%dB", i)
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

func (rl *instance) backspace() {
	if len(rl.line) == 0 || rl.pos == 0 {
		return
	}

	moveCursorBackwards(1)
	rl.pos--
	rl.delete()
}

func (rl *instance) insert(b []byte) {
	r := []rune(string(b))
	switch {
	case len(rl.line) == 0:
		rl.line = r
	case rl.pos == 0:
		line = append(r, rl.line...)
	case rl.pos < len(rl.line):
		r := append(r, rl.line[rl.pos:]...)
		rl.line = append(rl.line[:rl.pos], r...)
	default:
		rl.line = append(rl.line, r...)
	}

	rl.echo()

	moveCursorForwards(len(r) - 1)
	rl.pos += len(r)

	if rl.modeTabGrid {
		rl.clearTabSuggestions()
		rl.tabCompletion()
	}
}

func (rl *instance) delete() {
	switch {
	case len(rl.line) == 0:
		return
	case rl.pos == 0:
		rl.line = rl.line[1:]
		rl.echo()
		moveCursorBackwards(1)
	case rl.pos > len(rl.line):
		backspace()
	case rl.pos == len(rl.line):
		rl.line = rl.line[:rl.pos]
		rl.echo()
		moveCursorBackwards(1)
	default:
		rl.line = append(rl.line[:rl.pos], rl.line[rl.pos+1:]...)
		rl.echo()
		moveCursorBackwards(1)
	}

	if rl.modeTabGrid {
		rl.clearTabSuggestions()
		rl.tabCompletion()
	}
}

func rl.clearLine() {
	if len(rl.line) == 0 {
		return
	}

	moveCursorBackwards(rl.pos)
	fmt.Print(strings.Repeat(" ", len(rl.line)))
	moveCursorBackwards(len(rl.line))

	rl.line = []rune{}
	rl.pos = 0

	if rl.modeTabGrid {
		rl.clearTabSuggestions()
		rl.tabCompletion()
	}
}

package readline

import (
	"fmt"
)

// LineHistory is an interface to allow you to write your own history logging
// tools. eg sqlite backend instead of a file system.
// By default readline will just use the dummyLineHistory interface which only
// logs the history to memory ([]string to be precise).
type LineHistory interface {
	// Append takes the line and returns an updated number of lines or an error
	Write(string) (int, error)

	// GetLine takes the historic line number and returns the line or an error
	GetLine(int) (string, error)

	// Len returns the number of history lines
	Len() int

	// Dump returns everything in readline. The return is an interface because{} not all
	// LineHistory implementations might want to structure the history the same
	Dump() interface{}
}

// An example of a LineHistory interface:

type ExampleLineHistory struct {
	items []string
}

func (h *ExampleLineHistory) Write(s string) (int, error) {
	h.items = append(h.items, s)
	return len(h.items), nil
}

func (h *ExampleLineHistory) GetLine(i int) (string, error) {
	return h.items[i], nil
}

func (h *ExampleLineHistory) Len() int {
	return len(h.items)
}

func (h *ExampleLineHistory) Dump() interface{} {
	return h.items
}

// Browse historic lines:

func (rl *instance) walkHistory(i int) {
	switch rl.histPos + i {
	case -1, rl.History.Len() + 1:
		return

	case rl.History.Len():
		rl.clearLine()
		rl.histPos += i
		rl.line = rl.lineBuf

	default:
		s, err := rl.History.GetLine(rl.histPos + i)
		if err != nil {
			fmt.Print("\r\n" + err.Error() + "\r\n")
			fmt.Print(rl.prompt)
			return
		}

		if rl.histPos == rl.History.Len() {
			rl.lineBuf = append(rl.line, []rune{}...)
		}

		rl.clearLine()
		rl.histPos += i
		rl.line = []rune(s)
	}

	rl.echo()
	rl.pos = len(rl.line)
	if rl.pos > 1 {
		moveCursorForwards(rl.pos - 1)
	} else if rl.pos == 0 {
		moveCursorBackwards(1)
	}
}

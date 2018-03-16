package readline

import (
	"fmt"
	"strings"
)

type viMode int

const (
	vimInsert viMode = iota
	vimReplaceOnce
	vimReplaceMany
	vimKeys
)

var modeViMode viMode = vimInsert

func vi(b byte) {
	switch b {
	case 'a':
		moveCursorForwards(1)
		pos++
		modeViMode = vimInsert
	case 'A':
		moveCursorForwards(len(line) - pos)
		pos = len(line)
		modeViMode = vimInsert
	case 'D':
		moveCursorBackwards(pos)
		fmt.Print(strings.Repeat(" ", len(line)))
		//moveCursorBackwards(len(line))

		moveCursorBackwards(len(line) - pos)
		line = line[:pos]
		echo()

		moveCursorBackwards(2)
		pos--
	case 'i':
		modeViMode = vimInsert
	case 'r':
		modeViMode = vimReplaceOnce
	case 'R':
		modeViMode = vimReplaceMany
	case 'x':
		delete()
	}
}

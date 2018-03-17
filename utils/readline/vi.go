package readline

import (
	"fmt"
	"strconv"
	"strings"
)

type viMode int

const (
	vimInsert viMode = iota
	vimReplaceOnce
	vimReplaceMany
	vimKeys
)

func (rl *instance) vi(b byte) {
	switch b {
	case 'a':
		moveCursorForwards(1)
		rl.pos++
		rl.modeViMode = vimInsert
		rl.viIteration = ""
	case 'A':
		moveCursorForwards(len(rl.line) - rl.pos)
		rl.pos = len(line)
		rl.modeViMode = vimInsert
		rl.viIteration = ""
	case 'D':
		moveCursorBackwards(pos)
		fmt.Print(strings.Repeat(" ", len(rl.line)))
		//moveCursorBackwards(len(line))

		moveCursorBackwards(len(rl.line) - rl.pos)
		rl.line = rl.line[:rl.pos]
		rl.echo()

		moveCursorBackwards(2)
		rl.pos--
		rl.viIteration = ""
	case 'i':
		rl.modeViMode = vimInsert
		rl.viIteration = ""
	case 'r':
		rl.modeViMode = vimReplaceOnce
		rl.viIteration = ""
	case 'R':
		rl.modeViMode = vimReplaceMany
		rl.viIteration = ""
	case 'x':
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.delete()
		}
	default:
		if b <= '9' && '0' <= b {
			rl.viIteration += string(b)
		}
	}
}

func (rl *instance) getViIterations() int {
	i, _ := strconv.Atoi(rl.viIteration)
	if i < 1 {
		i = 1
	}
	rl.viIteration = ""
	return i
}

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

var modeViMode viMode = vimInsert
var viIteration string

func vi(b byte) {
	switch b {
	case 'a':
		moveCursorForwards(1)
		pos++
		modeViMode = vimInsert
		viIteration = ""
	case 'A':
		moveCursorForwards(len(line) - pos)
		pos = len(line)
		modeViMode = vimInsert
		viIteration = ""
	case 'D':
		moveCursorBackwards(pos)
		fmt.Print(strings.Repeat(" ", len(line)))
		//moveCursorBackwards(len(line))

		moveCursorBackwards(len(line) - pos)
		line = line[:pos]
		echo()

		moveCursorBackwards(2)
		pos--
		viIteration = ""
	case 'i':
		modeViMode = vimInsert
		viIteration = ""
	case 'r':
		modeViMode = vimReplaceOnce
		viIteration = ""
	case 'R':
		modeViMode = vimReplaceMany
		viIteration = ""
	case 'x':
		vii := getViIterations()
		for i := 1; i <= vii; i++ {
			delete()
		}
	default:
		if b <= '9' && '0' <= b {
			viIteration += string(b)
		}
	}
}

func getViIterations() int {
	i, _ := strconv.Atoi(viIteration)
	if i < 1 {
		i = 1
	}
	viIteration = ""
	return i
}

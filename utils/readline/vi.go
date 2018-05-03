package readline

import (
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

func (rl *Instance) vi(b byte) {
	switch b {
	case 'a':
		moveCursorForwards(1)
		rl.pos++
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'A':
		moveCursorForwards(len(rl.line) - rl.pos)
		rl.pos = len(rl.line)
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'D':
		moveCursorBackwards(rl.pos)
		print(strings.Repeat(" ", len(rl.line)))
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
		rl.viUndoSkipAppend = true

	case 'r':
		rl.modeViMode = vimReplaceOnce
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'R':
		rl.modeViMode = vimReplaceMany
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'u':
		if len(rl.viUndoHistory) > 0 {
			newline := append(rl.viUndoHistory[len(rl.viUndoHistory)-1], []rune{}...)
			rl.viUndoHistory = rl.viUndoHistory[:len(rl.viUndoHistory)-1]
			rl.clearHelpers()
			print("\r\n" + rl.prompt)
			rl.line = newline
			rl.pos = -1
			rl.echo()
		}
		rl.viUndoSkipAppend = true

	case 'v':
		rl.clearHelpers()
		var multiline []rune
		if rl.GetMultiLine == nil {
			multiline = rl.line
		} else {
			multiline = append(rl.GetMultiLine(), rl.line...)
		}

		new, err := rl.launchEditor(multiline)
		if err != nil || len(new) == 0 || string(new) == string(multiline) {
			rl.viUndoSkipAppend = true
			return
		}
		rl.clearLine()
		rl.multiline = []byte(string(new))

	case 'x':
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.delete()
		}

	default:
		if b <= '9' && '0' <= b {
			rl.viIteration += string(b)
		}
		rl.viUndoSkipAppend = true

	}
}

func (rl *Instance) getViIterations() int {
	i, _ := strconv.Atoi(rl.viIteration)
	if i < 1 {
		i = 1
	}
	rl.viIteration = ""
	return i
}

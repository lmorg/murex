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
	vimDelete
	vimKeys
)

func (rl *Instance) vi(r rune) {
	switch r {
	case 'a':
		moveCursorForwards(1)
		rl.pos++
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true
		rl.viHintInsert()

	case 'A':
		moveCursorForwards(len(rl.line) - rl.pos)
		rl.pos = len(rl.line)
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true
		rl.viHintInsert()

	case 'd':
		rl.modeViMode = vimDelete
		rl.viUndoSkipAppend = true
		rl.viHintDelete()

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
		rl.viHintInsert()

	case 'r':
		rl.modeViMode = vimReplaceOnce
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'R':
		rl.modeViMode = vimReplaceMany
		rl.viIteration = ""
		rl.viUndoSkipAppend = true
		rl.viHintReplace()

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
		if r <= '9' && '0' <= r {
			rl.viIteration += string(r)
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

func (rl *Instance) vimDelete(r rune) {
	switch r {
	case 'd':
		rl.clearLine()
		rl.resetHelpers()
		rl.getHintText()

	case 'w':
		split, index, pos := rl.tokeniseLine()
		//fmt.Println("|", split, index, pos, "|")
		var before, partial, after string
		if index > 0 {
			before = strings.Join(split[:index], "")
		}

		partial = split[index][:pos]

		if index < len(split)-1 {
			after = strings.Join(split[index+1:], "")
		}

		moveCursorBackwards(rl.pos)
		print(strings.Repeat(" ", len(rl.line)))
		moveCursorBackwards(len(rl.line) - rl.pos)

		rl.line = []rune(before + partial + after)
		rl.echo()
		rl.getHintText()
	}

	rl.modeViMode = vimKeys
	rl.renderHelpers()
}

func (rl *Instance) viHintVimKeys() {
	rl.viHintMessage([]rune("-- VIM KEYS -- (press `i` to return to normal editing mode)"))
}
func (rl *Instance) viHintInsert()  { rl.viHintMessage([]rune("-- INSERT --")) }
func (rl *Instance) viHintReplace() { rl.viHintMessage([]rune("-- REPLACE --")) }
func (rl *Instance) viHintDelete()  { rl.viHintMessage([]rune("-- DELETE --")) }

func (rl *Instance) viHintMessage(message []rune) {
	rl.hintText = message
	rl.clearHelpers()
	rl.renderHelpers()
}

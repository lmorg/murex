package readline

import "strings"

func (rl *Instance) vimDelete(r rune) {
	switch r {
	case 'b':
		rl.viDeleteByAdjust(rl.viJumpB(tokeniseLine))

	case 'B':
		rl.viDeleteByAdjust(rl.viJumpB(tokeniseSplitSpaces))

	case 'd':
		rl.clearLine()
		rl.resetHelpers()
		rl.getHintText()

	case 'e':
		rl.viDeleteByAdjust(rl.viJumpE(tokeniseLine))

	case 'E':
		rl.viDeleteByAdjust(rl.viJumpE(tokeniseSplitSpaces))

	case 'w':
		rl.viDeleteByAdjust(rl.viJumpW(tokeniseLine))

	case 'W':
		rl.viDeleteByAdjust(rl.viJumpW(tokeniseSplitSpaces))

	default:
		rl.viUndoSkipAppend = true
	}

	rl.modeViMode = vimKeys
}

func (rl *Instance) viDeleteByAdjust(adjust int) {
	var newLine []rune
	switch {
	case adjust == 0:
		rl.viUndoSkipAppend = true
		return
	case rl.pos+adjust == len(rl.line)-1:
		newLine = rl.line[:rl.pos]
	case rl.pos+adjust == 0:
		newLine = rl.line[rl.pos:]
	case adjust < 0:
		newLine = append(rl.line[:rl.pos+adjust], rl.line[rl.pos+1:]...)
	default:
		newLine = append(rl.line[:rl.pos], rl.line[rl.pos+adjust+1:]...)
	}

	moveCursorBackwards(rl.pos)
	print(strings.Repeat(" ", len(rl.line)))
	moveCursorBackwards(len(rl.line) - rl.pos)

	rl.line = newLine

	rl.echo()

	if adjust < 0 {
		rl.moveCursorByAdjust(adjust)
	}

}

func (rl *Instance) viDeleteW(tokeniser func(*Instance) ([]string, int, int)) {
	split, index, pos := tokeniser(rl)
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

	if rl.pos == len(rl.line)-1 {
		moveCursorBackwards(1)
	}
}

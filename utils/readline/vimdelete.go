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

	case '%':
		rl.viDeleteByAdjust(rl.viJumpBracket())

	default:
		rl.viUndoSkipAppend = true
	}

	rl.modeViMode = vimKeys
}

func (rl *Instance) viDeleteByAdjust(adjust int) {
	var (
		newLine []rune
		backOne bool
	)

	switch {
	case adjust == 0:
		rl.viUndoSkipAppend = true
		return
	case rl.pos+adjust == len(rl.line)-1:
		newLine = rl.line[:rl.pos]
		backOne = true
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

	if backOne {
		moveCursorBackwards(1)
		rl.pos--
	}
}

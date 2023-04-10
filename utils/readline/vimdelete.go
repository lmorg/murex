package readline

import "strings"

func (rl *Instance) vimDelete(r []rune) {
	defer func() { rl.modeViMode = vimKeys }()

	switch r[0] {
	case 'b':
		rl.viDeleteByAdjust(rl.viJumpB(tokeniseLine))

	case 'B':
		rl.viDeleteByAdjust(rl.viJumpB(tokeniseSplitSpaces))

	case 'd':
		rl.clearPrompt()
		rl.resetHelpers()
		rl.getHintText()

	case 'e':
		rl.viDeleteByAdjust(rl.viJumpE(tokeniseLine) + 1)

	case 'E':
		rl.viDeleteByAdjust(rl.viJumpE(tokeniseSplitSpaces) + 1)

	case 'w':
		rl.viDeleteByAdjust(rl.viJumpW(tokeniseLine))

	case 'W':
		rl.viDeleteByAdjust(rl.viJumpW(tokeniseSplitSpaces))

	case '%':
		rl.viDeleteByAdjust(rl.viJumpBracket())

	case 27:
		if len(r) > 1 && '1' <= r[1] && r[1] <= '9' {
			if rl.vimDeleteToken(r[1]) {
				return
			}
		}
		fallthrough

	default:
		rl.viUndoSkipAppend = true
	}
}

func (rl *Instance) viDeleteByAdjust(adjust int) {
	if adjust == 0 {
		rl.viUndoSkipAppend = true
		return
	}

	// Separate out the cursor movement from the logic so we can run tests on
	// the logic
	newLine, backOne := rl.viDeleteByAdjustLogic(&adjust)

	rl.line.Set(newLine)

	rl.echo()

	if adjust < 0 {
		rl.moveCursorByRuneAdjust(adjust)
	}

	if backOne {
		moveCursorBackwards(1)
		rl.line.SetRunePos(rl.line.RunePos() - 1)
	}
}

func (rl *Instance) viDeleteByAdjustLogic(adjust *int) (newLine []rune, backOne bool) {
	switch {
	case rl.line.RuneLen() == 0:
		*adjust = 0
		newLine = []rune{}

	case rl.line.RunePos()+*adjust > rl.line.RuneLen()-1:
		*adjust -= (rl.line.RunePos() + *adjust) - (rl.line.RuneLen() - 1)
		fallthrough

	case rl.line.RunePos()+*adjust == rl.line.RuneLen()-1:
		newLine = rl.line.Runes()[:rl.line.RunePos()]
		backOne = true

	case rl.line.RunePos()+*adjust < 0:
		*adjust = rl.line.RunePos()
		fallthrough

	case rl.line.RunePos()+*adjust == 0:
		newLine = rl.line.Runes()[rl.line.RunePos():]

	case *adjust < 0:
		newLine = append(rl.line.Runes()[:rl.line.RunePos()+*adjust], rl.line.Runes()[rl.line.RunePos():]...)

	default:
		newLine = append(rl.line.Runes()[:rl.line.RunePos()], rl.line.Runes()[rl.line.RunePos()+*adjust:]...)
	}

	return
}

func (rl *Instance) vimDeleteToken(r rune) bool {
	tokens, _, _ := tokeniseSplitSpaces(rl.line.Runes(), 0)
	pos := int(r) - 48 // convert ASCII to integer
	if pos > len(tokens) {
		return false
	}

	s := rl.line.String()
	newLine := strings.Replace(s, tokens[pos-1], "", -1)
	if newLine == s {
		return false
	}

	moveCursorBackwards(rl.line.CellPos())
	print(strings.Repeat(" ", rl.line.CellLen()))
	moveCursorBackwards(rl.line.CellLen() - rl.line.CellPos())

	rl.line.Set([]rune(newLine))

	rl.echo()

	if rl.line.RunePos() > rl.line.RuneLen() {
		moveCursorBackwards(GetTermWidth())
		moveCursorForwards(rl.promptLen + rl.line.CellLen() - 1)
		// ^ this is lazy
		rl.line.SetRunePos(rl.line.RuneLen() - 1)
	}

	return true
}

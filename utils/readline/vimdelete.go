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
		rl.clearLine()
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

	moveCursorBackwards(rl.pos)
	print(strings.Repeat(" ", rl.line.Len()))
	moveCursorBackwards(rl.line.Len() - rl.pos)

	rl.line.Value = newLine

	rl.echo()

	if adjust < 0 {
		rl.moveCursorByAdjust(adjust)
	}

	if backOne {
		moveCursorBackwards(1)
		rl.pos--
	}
}

func (rl *Instance) viDeleteByAdjustLogic(adjust *int) (newLine []rune, backOne bool) {
	switch {
	case rl.line.Len() == 0:
		*adjust = 0
		newLine = []rune{}
	case rl.pos+*adjust > rl.line.Len()-1:
		*adjust -= (rl.pos + *adjust) - (rl.line.Len() - 1)
		fallthrough
	case rl.pos+*adjust == rl.line.Len()-1:
		newLine = rl.line.Value[:rl.pos]
		backOne = true
	case rl.pos+*adjust < 0:
		*adjust = rl.pos
		fallthrough
	case rl.pos+*adjust == 0:
		newLine = rl.line.Value[rl.pos:]
	case *adjust < 0:
		newLine = append(rl.line.Value[:rl.pos+*adjust], rl.line.Value[rl.pos:]...)
	default:
		newLine = append(rl.line.Value[:rl.pos], rl.line.Value[rl.pos+*adjust:]...)
	}

	return
}

func (rl *Instance) vimDeleteToken(r rune) bool {
	/*line, err := rl.History.GetLine(rl.History.Len() - 1)
	if err != nil {
		return false
	}*/

	//tokens, _, _ := tokeniseSplitSpaces([]rune(line), 0)
	tokens, _, _ := tokeniseSplitSpaces(rl.line.Value, 0)
	pos := int(r) - 48 // convert ASCII to integer
	if pos > len(tokens) {
		return false
	}

	s := rl.line.String()
	newLine := strings.Replace(s, tokens[pos-1], "", -1)
	if newLine == s {
		return false
	}

	moveCursorBackwards(rl.pos)
	print(strings.Repeat(" ", rl.line.Len()))
	moveCursorBackwards(rl.line.Len() - rl.pos)

	rl.line.Value = []rune(newLine)

	rl.echo()

	if rl.pos > rl.line.Len() {
		moveCursorBackwards(GetTermWidth())
		moveCursorForwards(rl.promptLen + rl.line.Len() - 1)
		// ^ this is lazy
		rl.pos = rl.line.Len() - 1
	}

	return true
}

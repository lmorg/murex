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

func (rl *Instance) viDeleteByAdjustLogic(adjust *int) (newLine []rune, backOne bool) {
	switch {
	case len(rl.line) == 0:
		*adjust = 0
		newLine = []rune{}
	case rl.pos+*adjust > len(rl.line)-1:
		*adjust -= (rl.pos + *adjust) - (len(rl.line) - 1)
		fallthrough
	case rl.pos+*adjust == len(rl.line)-1:
		newLine = rl.line[:rl.pos]
		backOne = true
	case rl.pos+*adjust < 0:
		*adjust = rl.pos
		fallthrough
	case rl.pos+*adjust == 0:
		newLine = rl.line[rl.pos:]
	case *adjust < 0:
		newLine = append(rl.line[:rl.pos+*adjust], rl.line[rl.pos:]...)
	default:
		newLine = append(rl.line[:rl.pos], rl.line[rl.pos+*adjust:]...)
	}

	return
}

func (rl *Instance) vimDeleteToken(r rune) bool {
	/*line, err := rl.History.GetLine(rl.History.Len() - 1)
	if err != nil {
		return false
	}*/

	//tokens, _, _ := tokeniseSplitSpaces([]rune(line), 0)
	tokens, _, _ := tokeniseSplitSpaces(rl.line, 0)
	pos := int(r) - 48 // convert ASCII to integer
	if pos > len(tokens) {
		return false
	}

	s := string(rl.line)
	newLine := strings.Replace(s, tokens[pos-1], "", -1)
	if newLine == s {
		return false
	}

	moveCursorBackwards(rl.pos)
	print(strings.Repeat(" ", len(rl.line)))
	moveCursorBackwards(len(rl.line) - rl.pos)

	rl.line = []rune(newLine)

	rl.echo()

	if rl.pos > len(rl.line) {
		moveCursorBackwards(GetTermWidth())
		moveCursorForwards(rl.promptLen + len(rl.line) - 1)
		// ^ this is lazy
		rl.pos = len(rl.line) - 1
	}

	return true
}

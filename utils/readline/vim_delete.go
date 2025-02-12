package readline

import "strings"

func (rl *Instance) vimDeleteStr(r []rune) string {
	// TODO: test me
	//defer func() { rl.modeViMode = vimKeys }()

	var output string
	switch r[0] {
	case 'b':
		output = rl.viDeleteByAdjustStr(rl.viJumpB(tokeniseLine))

	case 'B':
		output = rl.viDeleteByAdjustStr(rl.viJumpB(tokeniseSplitSpaces))

	case 'd':
		rl.clearPrompt()
		rl.resetHelpers()
		rl.getHintText()

	case 'e':
		output = rl.viDeleteByAdjustStr(rl.viJumpE(tokeniseLine) + 1)

	case 'E':
		output = rl.viDeleteByAdjustStr(rl.viJumpE(tokeniseSplitSpaces) + 1)

	case 'w':
		output = rl.viDeleteByAdjustStr(rl.viJumpW(tokeniseLine))

	case 'W':
		output = rl.viDeleteByAdjustStr(rl.viJumpW(tokeniseSplitSpaces))

	case '%':
		output = rl.viDeleteByAdjustStr(rl.viJumpBracket())

	case 27:
		if len(r) > 1 && '1' <= r[1] && r[1] <= '9' {
			output = rl.vimDeleteTokenStr(r[1])
			if output != "" {
				rl.modeViMode = vimKeys
				return output
			}
		}
		fallthrough

	default:
		rl.viUndoSkipAppend = true
	}

	rl.modeViMode = vimKeys

	return output
}

func (rl *Instance) viDeleteByAdjustStr(adjust int) string {
	if adjust == 0 {
		rl.viUndoSkipAppend = true
		return ""
	}

	// Separate out the cursor movement from the logic so we can run tests on
	// the logic
	newLine, backOne := rl.viDeleteByAdjustLogic(&adjust)

	rl.line.Set(rl, newLine)

	output := rl.echoStr()

	if adjust < 0 {
		output += rl.moveCursorByRuneAdjustStr(adjust)
	}

	if backOne {
		output += moveCursorBackwardsStr(1)
		rl.line.SetRunePos(rl.line.RunePos() - 1)
	}

	return output
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

func (rl *Instance) vimDeleteTokenStr(r rune) string {
	tokens, _, _ := tokeniseSplitSpaces(rl.line.Runes(), 0)
	pos := int(r) - 48 // convert ASCII to integer
	if pos > len(tokens) {
		return ""
	}

	s := rl.line.String()
	newLine := strings.Replace(s, tokens[pos-1], "", -1)
	if newLine == s {
		return ""
	}

	output := moveCursorBackwardsStr(rl.line.CellPos())
	output += strings.Repeat(" ", rl.line.CellLen())
	output += moveCursorBackwardsStr(rl.line.CellLen() - rl.line.CellPos())

	rl.line.Set(rl, []rune(newLine))

	output += rl.echoStr()

	if rl.line.RunePos() > rl.line.RuneLen() {
		output += "\r" //moveCursorBackwardsStr(GetTermWidth())
		output += moveCursorForwardsStr(rl.promptLen + rl.line.CellLen() - 1)
		// ^ this is lazy
		rl.line.SetRunePos(rl.line.RuneLen() - 1)
	}

	return output
}

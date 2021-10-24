package readline

import "strings"

func (rl *Instance) getHintText() {
	if rl.HintText == nil {
		rl.resetHintText()
		return
	}

	rl.hintText = rl.HintText(rl.line, rl.pos)
}

func (rl *Instance) writeHintText(resetCursorPos bool) {
	if len(rl.hintText) == 0 {
		rl.hintY = 0
		return
	}

	hintText := string(rl.hintText)

	if rl.modeTabCompletion && rl.tcDisplayType == TabDisplayGrid &&
		!rl.modeTabFind && len(rl.tcSuggestions) > 0 {
		cell := (rl.tcMaxX * (rl.tcPosY - 1)) + rl.tcOffset + rl.tcPosX - 1
		description := rl.tcDescriptions[rl.tcSuggestions[cell]]
		if description != "" {
			hintText = description
		}
	}

	// fix bug https://github.com/lmorg/murex/issues/376
	if rl.termWidth == 0 {
		rl.termWidth = GetTermWidth()
	}

	// Determine how many lines hintText spans over
	// (Currently there is no support for carridge returns / new lines)
	hintLength := strLen(hintText)
	n := float64(hintLength) / float64(rl.termWidth)
	if float64(int(n)) != n {
		n++
	}
	rl.hintY = int(n)

	if rl.hintY > 3 {
		rl.hintY = 3
		//hintText = hintText[:(rl.termWidth*3)-4] + "..."
		hintText = hintText[:(rl.termWidth*3)-2] + "…"
	} else {
		padding := (rl.hintY * rl.termWidth) - len(hintText)
		if padding < 0 {
			padding = 0
		}
		hintText += strings.Repeat(" ", padding)
	}

	if resetCursorPos {
		_, lineY := lineWrapPos(rl.promptLen, len(rl.line), rl.termWidth)
		posX, posY := lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)
		y := lineY - posY
		moveCursorDown(y)

		print("\r\n" + rl.HintFormatting + hintText + seqReset)

		moveCursorUp(rl.hintY + lineY - posY)
		moveCursorBackwards(rl.termWidth)
		moveCursorForwards(posX)
	}
}

func (rl *Instance) resetHintText() {
	rl.hintY = 0
	rl.hintText = []rune{}
}

// SetHintText is a nasty function for force writing a new hint text. Use sparingly!
func (rl *Instance) SetHintText(s string) {
	rl.hintText = []rune(s)
	rl.writeHintText(true)
}

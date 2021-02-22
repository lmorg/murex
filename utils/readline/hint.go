package readline

import "strings"

func (rl *Instance) getHintText() {
	if rl.HintText == nil {
		rl.resetHintText()
		return
	}

	rl.hintText = rl.HintText(rl.line, rl.pos)
}

func (rl *Instance) writeHintText() {
	if len(rl.hintText) == 0 {
		rl.hintY = 0
		return
	}

	// Determine how many lines hintText spans over
	// (Currently there is no support for carridge returns / new lines)
	hintLength := strLen(string(rl.hintText))
	n := float64(hintLength) / float64(rl.termWidth)
	if float64(int(n)) != n {
		n++
	}
	rl.hintY = int(n)

	if rl.hintY > 3 {
		rl.hintY = 3
		rl.hintText = rl.hintText[:(rl.termWidth*3)-4]
		rl.hintText = append(rl.hintText, '.', '.', '.')
	}

	hintText := rl.hintText

	if rl.modeTabCompletion && rl.tcDisplayType == TabDisplayGrid &&
		!rl.modeTabFind && len(rl.tcSuggestions) > 0 {
		cell := (rl.tcMaxX * (rl.tcPosY - 1)) + rl.tcOffset + rl.tcPosX - 1
		description := rl.tcDescriptions[rl.tcSuggestions[cell]]
		if description != "" {
			hintText = []rune(description)
		}
	}

	_, lineY := lineWrapPos(rl.promptLen, len(rl.line), rl.termWidth)
	_, posY := lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)
	y := lineY - posY + 1
	print(strings.Repeat("\r\n", y) + rl.HintFormatting + string(hintText) + seqReset)
}

func (rl *Instance) resetHintText() {
	rl.hintY = 0
	rl.hintText = []rune{}
}

// SetHintText is a nasty function for force writing a new hint text. Use sparingly!
func (rl *Instance) SetHintText(s string) {
	rl.hintText = []rune(s)
	//rl.writeHintText()
}

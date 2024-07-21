package readline

import "strings"

func (rl *Instance) getHintText() {
	if rl.HintText == nil {
		rl.resetHintText()
		return
	}

	hint := rl.cacheHint.Get(rl.line.Runes())
	if len(hint) > 0 {
		rl.hintText = hint
		return
	}

	rl.hintText = rl.HintText(rl.line.Runes(), rl.line.RunePos())
	rl.cacheHint.Append(rl.line.Runes(), rl.hintText)
}

func (rl *Instance) writeHintTextStr() string {
	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	if rl.HintText == nil {
		rl.hintY = 0
		return ""
	}

	if len(rl.hintText) == 0 {
		rl.hintText = []rune{' '}
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

	s := rl._writeHintTextStr(hintText)
	return s
}

// ForceHintTextUpdate is a nasty function for force writing a new hint text. Use sparingly!
func (rl *Instance) ForceHintTextUpdate(s string) {
	rl.hintText = []rune(s)

	rl.tabMutex.Lock()
	print(rl._writeHintTextStr(s))
	rl.tabMutex.Unlock()
}

// _writeHintTextStr doesn't contain any mutex locks. This will have to be
// handled from the calling function. eg writeHintTextStr()
func (rl *Instance) _writeHintTextStr(hintText string) string {
	// fix bug https://github.com/lmorg/murex/issues/376
	if rl.termWidth == 0 {
		rl.termWidth = GetTermWidth()
	}

	// Determine how many lines hintText spans over
	// (Currently there is no support for carriage returns / new lines)
	hintLength := strLen(hintText)
	n := float64(hintLength) / float64(rl.termWidth)
	if float64(int(n)) != n {
		n++
	}
	rl.hintY = int(n)

	if rl.hintY > 3 {
		rl.hintY = 3
		hintText = hintText[:(rl.termWidth*3)-2] + "…"
	} else {
		padding := (rl.hintY * rl.termWidth) - len(hintText)
		if padding < 0 {
			padding = 0
		}
		hintText += strings.Repeat(" ", padding)
	}

	_, lineY := rl.lineWrapCellLen()
	posX, posY := rl.lineWrapCellPos()
	y := lineY - posY
	write := moveCursorDownStr(y)

	write += "\r\n" + rl.HintFormatting + hintText + seqReset

	write += moveCursorUpStr(rl.hintY + lineY - posY)
	write += moveCursorBackwardsStr(rl.termWidth)
	write += moveCursorForwardsStr(posX)

	return write
}

func (rl *Instance) resetHintText() {
	rl.hintY = 0
	rl.hintText = []rune{}
}

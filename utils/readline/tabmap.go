package readline

import "fmt"

func (rl *Instance) initTabMap() {
	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	var suggestions *suggestionsT
	if rl.modeTabFind {
		suggestions = newSuggestionsT(rl, rl.tfSuggestions)
	} else {
		suggestions = newSuggestionsT(rl, rl.tcSuggestions)
	}

	rl.tcMaxLength = 1
	//for i := range suggestions {
	for i := 0; i < suggestions.Len(); i++ {
		if rl.tcDisplayType == TabDisplayList {
			if suggestions.ItemLen(i) > rl.tcMaxLength {
				rl.tcMaxLength = suggestions.ItemLen(i)
			}

		} else {
			if len(rl.tcDescriptions[suggestions.ItemLookupValue(i)]) > rl.tcMaxLength {
				rl.tcMaxLength = len(rl.tcDescriptions[suggestions.ItemLookupValue(i)])
			}
		}
	}

	rl.tcPosX = 1
	rl.tcPosY = 1
	rl.tcOffset = 0
	rl.tcMaxX = 1

	if suggestions.Len() > rl.MaxTabCompleterRows {
		rl.tcMaxY = rl.MaxTabCompleterRows
	} else {
		rl.tcMaxY = suggestions.Len()
	}
}

func (rl *Instance) moveTabMapHighlight(x, y int) {
	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	var suggestions *suggestionsT
	if rl.modeTabFind {
		suggestions = newSuggestionsT(rl, rl.tfSuggestions)
	} else {
		suggestions = newSuggestionsT(rl, rl.tcSuggestions)
	}

	rl.tcPosY += x
	rl.tcPosY += y

	if rl.tcPosY < 1 {
		rl.tcPosY = 1
		rl.tcOffset--
	}

	if rl.tcPosY > rl.tcMaxY {
		rl.tcPosY--
		rl.tcOffset++
	}

	if rl.tcOffset+rl.tcPosY < 1 && suggestions.Len() > 0 {
		rl.tcPosY = rl.tcMaxY
		rl.tcOffset = suggestions.Len() - rl.tcMaxY
	}

	if rl.tcOffset < 0 {
		rl.tcOffset = 0
	}

	if rl.tcOffset+rl.tcPosY > suggestions.Len() {
		rl.tcPosY = 1
		rl.tcOffset = 0
	}
}

func (rl *Instance) _writeTabMap() string {
	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	var suggestions *suggestionsT
	if rl.modeTabFind {
		suggestions = newSuggestionsT(rl, rl.tfSuggestions)
	} else {
		suggestions = newSuggestionsT(rl, rl.tcSuggestions)
	}

	if rl.termWidth < 10 {
		// terminal too small. Probably better we do nothing instead of crash
		return ""
	}

	maxLength := rl.tcMaxLength
	if maxLength > rl.termWidth-9 {
		maxLength = rl.termWidth - 9
	}
	maxDescWidth := rl.termWidth - maxLength - 4

	y := 0
	rl.previewItem = ""

	// bit of a kludge. Really should find where the code is "\n"ing
	output := moveCursorUpStr(1)

	isTabDisplayList := rl.tcDisplayType == TabDisplayList

	var item, description string
	for i := rl.tcOffset; i < suggestions.Len(); i++ {
		y++
		if y > rl.tcMaxY {
			break
		}

		if isTabDisplayList {
			item = runeWidthTruncate(suggestions.ItemValue(i), maxLength)
			description = runeWidthTruncate(rl.tcDescriptions[suggestions.ItemLookupValue(i)], maxDescWidth)

		} else {
			item = runeWidthTruncate(suggestions.ItemValue(i), maxDescWidth)
			description = runeWidthTruncate(rl.tcDescriptions[suggestions.ItemLookupValue(i)], maxLength)
		}

		if isTabDisplayList {
			output += fmt.Sprintf("\r\n%s %s %s %s",
				highlight(rl, y), runeWidthFillRight(item, maxLength),
				seqReset, description)

		} else {
			output += fmt.Sprintf("\r\n %s %s %s %s",
				runeWidthFillRight(description, maxLength), highlight(rl, y),
				runeWidthFillRight(item, maxDescWidth), seqReset)

		}

		if y == rl.tcPosY {
			rl.previewItem = suggestions.ItemValue(i)
		}
	}

	//print(output)

	if suggestions.Len() < rl.tcMaxX {
		rl.tcUsedY = suggestions.Len()
	} else {
		rl.tcUsedY = rl.tcMaxY
	}

	return output
}

func highlight(rl *Instance, y int) string {
	if y == rl.tcPosY {
		return seqBgWhite + seqFgBlack
	}
	return ""
}

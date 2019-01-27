package readline

import (
	"strconv"
)

func (rl *Instance) initTabMap() {
	//width := GetTermWidth()

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	rl.tcMaxLength = 1
	for i := range suggestions {
		if rl.tcDisplayType == TabDisplayList {
			if len(rl.tcPrefix+suggestions[i]) > rl.tcMaxLength {
				rl.tcMaxLength = len([]rune(rl.tcPrefix + suggestions[i]))
			}

		} else {
			if len(rl.tcDescriptions[suggestions[i]]) > rl.tcMaxLength {
				rl.tcMaxLength = len(rl.tcDescriptions[suggestions[i]])
			}
		}
	}

	rl.modeTabCompletion = true
	rl.tcPosX = 1
	rl.tcPosY = 1
	rl.tcOffset = 0
	rl.tcMaxX = 1

	if len(suggestions) > rl.MaxTabCompleterRows {
		rl.tcMaxY = rl.MaxTabCompleterRows
	} else {
		rl.tcMaxY = len(suggestions)
	}
}

func (rl *Instance) moveTabMapHighlight(x, y int) {
	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
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

	if rl.tcOffset+rl.tcPosY < 1 && len(suggestions) > 0 {
		rl.tcPosY = rl.tcMaxY
		rl.tcOffset = len(suggestions) - rl.tcMaxY
	}

	if rl.tcOffset < 0 {
		rl.tcOffset = 0
	}

	if rl.tcOffset+rl.tcPosY > len(suggestions) {
		rl.tcPosY = 1
		rl.tcOffset = 0
	}
}

func (rl *Instance) writeTabMap() {
	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	termWidth := GetTermWidth()
	if termWidth < 10 {
		// terminal too small. Probably better we do nothing instead of crash
		return
	}

	maxLength := rl.tcMaxLength
	if maxLength > termWidth-9 {
		maxLength = termWidth - 9
	}
	maxDescWidth := termWidth - maxLength - 4

	cellWidth := strconv.Itoa(maxLength)
	itemWidth := strconv.Itoa(maxDescWidth)

	y := 0
	print(seqClearScreenBelow)

	highlight := func(y int) string {
		if y == rl.tcPosY {
			return seqBgWhite + seqFgBlack
		}
		return ""
	}

	var item, description string
	for i := rl.tcOffset; i < len(suggestions); i++ {
		y++
		if y > rl.tcMaxY {
			break
		}

		if rl.tcDisplayType == TabDisplayList {
			item = rl.tcPrefix + suggestions[i]
			if len(item) > maxLength {
				item = item[:maxLength-3] + "..."
			}

			description = rl.tcDescriptions[suggestions[i]]
			if len(description) > maxDescWidth {
				description = description[:maxDescWidth-3] + "..."
			}

		} else {
			item = rl.tcPrefix + suggestions[i]
			if len(item) > maxDescWidth {
				item = item[:maxDescWidth-3] + "..."
			}

			description = rl.tcDescriptions[suggestions[i]]
			if len(description) > maxLength {
				description = description[:maxLength-3] + "..."
			}
		}

		if rl.tcDisplayType == TabDisplayList {
			printf("\r\n%s %-"+cellWidth+"s %s %s",
				highlight(y), item, seqReset, description)
		} else {
			printf("\r\n %-"+cellWidth+"s %s %-"+itemWidth+"s %s",
				description, highlight(y), item, seqReset)
		}
	}

	if len(suggestions) < rl.tcMaxX {
		rl.tcUsedY = len(suggestions)
	} else {
		rl.tcUsedY = rl.tcMaxY
	}
}

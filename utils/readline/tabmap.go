package readline

import (
	"strconv"
)

func (rl *Instance) initTabMap() {
	//width := getTermWidth()

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	rl.tcMaxLength = 1
	for i := range suggestions {
		if len(rl.tcPrefix+suggestions[i]) > rl.tcMaxLength {
			rl.tcMaxLength = len([]rune(rl.tcPrefix + suggestions[i]))
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

	maxDescWidth := (getTermWidth() - rl.tcMaxLength) - 4
	cellWidth := strconv.Itoa(rl.tcMaxLength)
	y := 0

	print(seqClearScreenBelow)

	for i := rl.tcOffset; i < len(suggestions); i++ {
		y++
		if y > rl.tcMaxY {
			break
		}

		if y == rl.tcPosY {
			print(seqBgWhite + seqFgBlack)
		}

		description := rl.tcDescriptions[suggestions[i]]
		if len(description) > maxDescWidth {
			description = description[:maxDescWidth-3] + "..."
		}

		printf("\r\n %-"+cellWidth+"s %s %s",
			rl.tcPrefix+suggestions[i], seqReset, description)
	}

	if len(suggestions) < rl.tcMaxX {
		rl.tcUsedY = len(suggestions)
	} else {
		rl.tcUsedY = rl.tcMaxY
	}
}

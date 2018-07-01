package readline

import (
	"strconv"
	"strings"
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

	rl.tcMaxX = 1

	if len(suggestions) > rl.MaxTabCompleterRows {
		rl.tcMaxY = rl.MaxTabCompleterRows
	} else {
		rl.tcMaxY = len(suggestions) - 1
	}
}

func (rl *Instance) moveTabMapHighlight(x, y int) {
	/*var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}*/

	rl.tcPosY += x
	rl.tcPosY += y

	if rl.tcPosY < 1 {
		rl.tcPosY = rl.tcMaxY
	}

	if rl.tcPosY > rl.tcMaxY {
		rl.tcPosY = 1
	}
}

func (rl *Instance) writeTabMap() {
	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	maxDefinitionWidth := (getTermWidth() - rl.tcMaxLength) - 4
	cellWidth := strconv.Itoa(rl.tcMaxLength)
	y := 0

	for _, item := range suggestions {
		y++
		if y > rl.tcMaxY {
			break
		}

		if y == rl.tcPosY {
			print(seqBgWhite + seqFgBlack)
		}

		definition := strings.TrimSpace(rl.tcDefinitions[item])
		if len(definition) > maxDefinitionWidth {
			definition = definition[:maxDefinitionWidth-3] + "..."
		}

		printf("\r\n %-"+cellWidth+"s %s %s", rl.tcPrefix+item, seqReset, definition)
	}

	rl.tcUsedY = y - 1
	if rl.tcUsedY < 0 {
		rl.tcUsedY = 0
	}
}

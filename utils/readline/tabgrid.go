package readline

import (
	"strconv"
)

func (rl *Instance) initTabGrid() {
	width := GetTermWidth()

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	tcMaxLength := 1
	for i := range suggestions {
		if len(rl.tcPrefix+suggestions[i]) > tcMaxLength {
			tcMaxLength = len([]rune(rl.tcPrefix + suggestions[i]))
		}
	}

	rl.modeTabCompletion = true
	rl.tcPosX = 1
	rl.tcPosY = 1
	rl.tcMaxX = width / (tcMaxLength + 2)
	rl.tcOffset = 0

	// avoid a divide by zero error
	if rl.tcMaxX < 1 {
		rl.tcMaxX = 1
	}

	rl.tcMaxY = rl.MaxTabCompleterRows
}

func (rl *Instance) moveTabGridHighlight(x, y int) {
	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	rl.tcPosX += x
	rl.tcPosY += y

	if rl.tcPosX < 1 {
		rl.tcPosX = rl.tcMaxX
		rl.tcPosY--
	}

	if rl.tcPosX > rl.tcMaxX {
		rl.tcPosX = 1
		rl.tcPosY++
	}

	if rl.tcPosY < 1 {
		rl.tcPosY = rl.tcUsedY
	}

	if rl.tcPosY > rl.tcUsedY {
		rl.tcPosY = 1
	}

	if rl.tcPosY == rl.tcUsedY && (rl.tcMaxX*(rl.tcPosY-1))+rl.tcPosX > len(suggestions) {
		if x < 0 {
			rl.tcPosX = len(suggestions) - (rl.tcMaxX * (rl.tcPosY - 1))
		}

		if x > 0 {
			rl.tcPosX = 1
			rl.tcPosY = 1
		}

		if y < 0 {
			rl.tcPosY--
		}

		if y > 0 {
			rl.tcPosY = 1
		}
	}

	/*if !rl.modeTabFind && len(suggestions) > 0 {
		cell := (rl.tcMaxX * (rl.tcPosY - 1)) + rl.tcOffset + rl.tcPosX - 1
		description := rl.tcDescriptions[suggestions[cell]]
		if description != "" {
			rl.hintText = []rune(description)
		} else {
			rl.getHintText()
		}
	}*/
}

func (rl *Instance) writeTabGrid() {
	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	print(seqClearScreenBelow + "\r\n")

	cellWidth := strconv.Itoa((GetTermWidth() / rl.tcMaxX) - 2)
	x := 0
	y := 1

	for i := range suggestions {
		x++
		if x > rl.tcMaxX {
			x = 1
			y++
			if y > rl.tcMaxY {
				y--
				break
			} else {
				print("\r\n")
			}
		}

		if x == rl.tcPosX && y == rl.tcPosY {
			print(seqBgWhite + seqFgBlack)
		}
		printf(" %-"+cellWidth+"s %s", rl.tcPrefix+suggestions[i], seqReset)
	}

	rl.tcUsedY = y
}

package readline

import (
	"sort"
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

	sort.Strings(suggestions) // I don't like doing this here

	rl.tcMaxLength = rl.MinTabItemLength
	for i := range suggestions {
		if len(rl.tcPrefix+suggestions[i]) > rl.tcMaxLength {
			rl.tcMaxLength = len([]rune(rl.tcPrefix + suggestions[i]))
		}
	}
	if rl.tcMaxLength > rl.MaxTabItemLength && rl.MaxTabItemLength > 0 && rl.MaxTabItemLength > rl.MinTabItemLength {
		rl.tcMaxLength = rl.MaxTabItemLength
	}
	if rl.tcMaxLength == 0 {
		rl.tcMaxLength = 1
	}

	rl.modeTabCompletion = true
	rl.tcPosX = 1
	rl.tcPosY = 1
	rl.tcMaxX = width / (rl.tcMaxLength + 2)
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

	iCellWidth := (GetTermWidth() / rl.tcMaxX) - 2
	cellWidth := strconv.Itoa(iCellWidth)

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

		caption := cropCaption(rl.tcPrefix+suggestions[i], rl.tcMaxLength, iCellWidth)

		printf(" %-"+cellWidth+"s %s", caption, seqReset)
	}

	rl.tcUsedY = y
}

func cropCaption(caption string, tcMaxLength int, iCellWidth int) string {
	switch {
	case iCellWidth == 0:
		// this condition shouldn't ever happen but lets cover it just in case
		return ""
	case len(caption) < tcMaxLength:
		return caption
	case len(caption) < 5:
		return caption
	case len(caption) <= iCellWidth:
		return caption
	case len(caption)-iCellWidth+6 < 1:
		return caption[:iCellWidth-1] + "…"
	case len(caption) > 5+len(caption)-iCellWidth+6:
		return caption[:4] + "…" + caption[len(caption)-iCellWidth+6:]
	default:
		return caption[:iCellWidth-1] + "…"
	}
}

package readline

import (
	"strconv"
)

func (rl *Instance) initTabGrid() {
	rl.tabMutex.Lock()

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

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
		rl.tcMaxLength = 20
	}

	rl.tcPosX = 1
	rl.tcPosY = 1
	rl.tcMaxX = rl.termWidth / (rl.tcMaxLength + 2)
	rl.tcOffset = 0

	// avoid a divide by zero error
	if rl.tcMaxX < 1 {
		rl.tcMaxX = 1
	}

	rl.tcMaxY = rl.MaxTabCompleterRows

	// pre-cache
	max := rl.tcMaxX * rl.tcMaxY
	if max > len(rl.tcSuggestions) {
		max = len(rl.tcSuggestions)
	}
	subset := rl.tcSuggestions[:max]

	rl.tabMutex.Unlock()

	if rl.tcr.HintCache == nil {
		return
	}

	go rl.tabHintCache(subset)
}

func (rl *Instance) tabHintCache(subset []string) {
	hints := rl.tcr.HintCache(rl.tcPrefix, subset)
	if len(hints) != len(subset) {
		return
	}

	rl.tabMutex.Lock()
	for i := range subset {
		rl.tcDescriptions[subset[i]] = hints[i]
	}
	rl.tabMutex.Unlock()
}

func (rl *Instance) moveTabGridHighlight(x, y int) {
	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

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
}

func (rl *Instance) writeTabGrid() {
	rl.tabMutex.Lock()

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	iCellWidth := (rl.termWidth / rl.tcMaxX) - 2
	cellWidth := strconv.Itoa(iCellWidth)

	x := 0
	y := 1
	var item string

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
			item = rl.tcPrefix + suggestions[i]
		}

		value := rl.tcPrefix + suggestions[i]
		caption := cropCaption(value, rl.tcMaxLength, iCellWidth)
		if caption != value {
			rl.tcDescriptions[suggestions[i]] = value
		}

		printf(" %-"+cellWidth+"s %s", caption, seqReset)
	}

	rl.tcUsedY = y
	rl.tabMutex.Unlock()

	rl.writePreview(item)
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

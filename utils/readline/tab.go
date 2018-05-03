package readline

import (
	"strconv"
)

func (rl *Instance) getTabCompletion() {
	if rl.TabCompleter == nil {
		return
	}

	rl.tcPrefix, rl.tcSuggestions = rl.TabCompleter(rl.line, rl.pos)
	if len(rl.tcSuggestions) == 0 {
		return
	}

	/*if len(rl.tcSuggestions) == 1 && !rl.modeTabGrid {
		if len(rl.tcSuggestions[0]) == 0 || rl.tcSuggestions[0] == " " || rl.tcSuggestions[0] == "\t" {
			return
		}
		rl.insert([]byte(rl.tcSuggestions[0]))
		return
	}*/

	rl.initTabGrid()
}

func (rl *Instance) initTabGrid() {
	width := getTermWidth()

	tcMaxLength := 1
	for i := range rl.tcSuggestions {
		if len(rl.tcPrefix+rl.tcSuggestions[i]) > tcMaxLength {
			tcMaxLength = len([]rune(rl.tcPrefix + rl.tcSuggestions[i]))
		}
	}

	rl.modeTabGrid = true
	rl.tcPosX = 1
	rl.tcPosY = 1
	rl.tcMaxX = width / (tcMaxLength + 2)

	// avoid a divide by zero error
	if rl.tcMaxX < 1 {
		rl.tcMaxX = 1
	}

	rl.tcMaxY = rl.MaxTabCompleterRows
	//if rl.tcMaxY < 1 {
	//	rl.tcMaxY = 1
	//}
}

func (rl *Instance) moveTabHighlight(x, y int) {
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

	if rl.tcPosY == rl.tcUsedY && (rl.tcMaxX*(rl.tcPosY-1))+rl.tcPosX > len(rl.tcSuggestions) {
		if x < 0 {
			rl.tcPosX = len(rl.tcSuggestions) - (rl.tcMaxX * (rl.tcPosY - 1))
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
	if !rl.modeTabGrid {
		return
	}

	print("\r\n")

	cellWidth := strconv.Itoa((getTermWidth() / rl.tcMaxX) - 2)
	x := 0
	y := 1

	for i := range rl.tcSuggestions {
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
		printf(" %-"+cellWidth+"s %s", rl.tcPrefix+rl.tcSuggestions[i], seqReset)
	}

	rl.tcUsedY = y
}

func (rl *Instance) resetTabCompletion() {
	rl.modeTabGrid = false
	rl.tcUsedY = 0
}

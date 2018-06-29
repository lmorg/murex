package readline

import (
	"regexp"
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
	if !rl.modeTabGrid {
		return
	}

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	print("\r\n")

	cellWidth := strconv.Itoa((getTermWidth() / rl.tcMaxX) - 2)
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

func (rl *Instance) resetTabCompletion() {
	rl.modeTabGrid = false
	rl.tcUsedY = 0
	rl.modeTabFind = false
	rl.tfLine = []rune{}
}

func (rl *Instance) backspaceTabFind() {
	if len(rl.tfLine) > 0 {
		rl.tfLine = rl.tfLine[:len(rl.tfLine)-1]
	}
	rl.updateTabFind([]rune{})
}

func (rl *Instance) updateTabFind(r []rune) {
	rl.tfLine = append(rl.tfLine, r...)
	rl.hintText = append([]rune("regexp find: "), rl.tfLine...)

	defer func() {
		rl.clearHelpers()
		rl.initTabGrid()
		rl.renderHelpers()
	}()

	if len(rl.tfLine) == 0 {
		rl.tfSuggestions = append(rl.tcSuggestions, []string{}...)
		return
	}

	rx, err := regexp.Compile(string(rl.tfLine))
	if err != nil {
		rl.tfSuggestions = []string{err.Error()}
		return
	}

	rl.tfSuggestions = make([]string, 0)
	for i := range rl.tcSuggestions {
		if rx.MatchString(rl.tcSuggestions[i]) {
			rl.tfSuggestions = append(rl.tfSuggestions, rl.tcSuggestions[i])
		}
	}
}

func (rl *Instance) resetTabFind() {
	rl.modeTabFind = false
	rl.tfLine = []rune{}
	rl.hintText = []rune("Cancelled regexp suggestion find.")

	rl.clearHelpers()
	rl.initTabGrid()
	rl.renderHelpers()
}

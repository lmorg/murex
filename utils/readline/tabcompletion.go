package readline

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
)

const (
	modeNormal = iota
	modeTabCompletion
)

var (
	tcPrefix      string
	tcSuggestions []string
	tcPosX        int
	tcPosY        int
	tcMaxX        int
	tcMaxY        int
	tcUsedY       int
	tcMaxLength   int
	termWidth     int
)

func tabCompletion() {
	if TabCompleter == nil {
		return
	}

	tcPrefix, tcSuggestions = TabCompleter(line, pos)
	if len(tcSuggestions) == 0 {
		return
	}

	if len(tcSuggestions) == 1 {
		if len(tcSuggestions[0]) == 0 || tcSuggestions[0] == " " || tcSuggestions[0] == "\t" {
			return
		}
		insert([]byte(tcSuggestions[0]))
		return
	}

	initTabGrid()
	renderSuggestions()
}

func initTabGrid() {
	var err error
	fd := int(os.Stdout.Fd())
	termWidth, _, err = terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	tcMaxLength := 1
	for i := range tcSuggestions {
		if len(tcPrefix+tcSuggestions[i]) > tcMaxLength {
			tcMaxLength = len([]rune(tcPrefix + tcSuggestions[i]))
		}
	}

	mode = modeTabCompletion
	tcPosX = 1
	tcPosY = 1
	tcMaxX = termWidth / (tcMaxLength + 2)
	tcMaxY = 4
}

func moveTabHighlight(x, y int) {
	tcPosX += x
	tcPosY += y

	if tcPosX < 1 {
		tcPosX = tcMaxX
		tcPosY--
	}

	if tcPosX > tcMaxX {
		tcPosX = 1
		tcPosY++
	}

	if tcPosY < 1 {
		tcPosY = tcUsedY
	}

	if tcPosY > tcUsedY {
		tcPosY = 1
	}

	if tcPosY == tcUsedY && (tcMaxX*(tcPosY-1))+tcPosX > len(tcSuggestions) {
		if x < 0 {
			tcPosX = len(tcSuggestions) - (tcMaxX * (tcPosY - 1))
		}

		if x > 0 {
			tcPosX = 1
			tcPosY = 1
		}

		if y < 0 {
			tcPosY--
		}

		if y > 0 {
			tcPosY = 1
		}
	}

	renderSuggestions()
}

func renderSuggestions() {
	fmt.Print("\r\n")

	cellWidth := strconv.Itoa((termWidth / tcMaxX) - 2)
	x := 0
	y := 1

	for i := range tcSuggestions {
		x++
		if x > tcMaxX {
			x = 1
			y++
			if y > tcMaxY {
				y--
				break
			} else {
				fmt.Print("\r\n")
			}
		}

		if x == tcPosX && y == tcPosY {
			fmt.Print(seqBgWhite + seqFgBlack)
		}
		fmt.Printf(" %-"+cellWidth+"s %s", tcPrefix+tcSuggestions[i], seqReset)
	}

	tcUsedY = y
	moveCursorUp(y)
	moveCursorBackwards(termWidth)
	moveCursorForwards(promptLen + pos)
}

func clearTabSuggestions() {
	move := termWidth * tcUsedY
	blank := strings.Repeat(" ", move)

	fmt.Print(seqPosSave + "\r\n" + blank + seqPosRestore)
	//fmt.Print("\r\n" + blank)
	//moveCursorBackwards(termWidth)
	//moveCursorUp(tcMaxY)
	//moveCursorForwards(len(Prompt) + len(line))
	mode = modeNormal
}

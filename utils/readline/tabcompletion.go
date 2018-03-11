package readline

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
)

const modeTabCompletion = 1

var (
	tcSuggestions []string
	tcPosX        int
	tcPosY        int
	tcMaxX        int
	tcMaxY        int
	tcMaxLength   int
	termWidth     int
)

func tabCompletion() {
	if TabCompleter == nil {
		return
	}

	tcSuggestions = TabCompleter(line, pos)
	if len(tcSuggestions) == 0 {
		return
	}

	if len(tcSuggestions) == 1 {
		if len(tcSuggestions[0]) == 0 {
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

	fmt.Print("\r\n")

	s := string(line)
	tcMaxLength := 1
	for i := range tcSuggestions {
		tcSuggestions[i] = s + tcSuggestions[i]
		if len(tcSuggestions[i]) > tcMaxLength {
			tcMaxLength = len(tcSuggestions[i])
		}
	}

	mode = modeTabCompletion
	tcPosX = 1
	tcPosY = 1
	tcMaxX = termWidth / (tcMaxLength + 2)
	tcMaxY = 4
}

func moveHighlight(x, y int) {
	//switch x {
	//case -1:
	//}
	tcPosX += x
	tcPosY += y
}

func renderSuggestions() {
	cellWidth := strconv.Itoa((termWidth / tcMaxX) - 2)
	x := 0
	y := 1
	for i := range tcSuggestions {
		x++
		if x > tcMaxX {
			x = 1
			y++
			if y > tcMaxY {
				break
			} else {
				fmt.Print("\r\n")
			}
		}

		if x == tcPosX && y == tcPosY {
			fmt.Print(seqBgWhite + seqFgBlack)
		}
		fmt.Printf(" %-"+cellWidth+"s %s", tcSuggestions[i], seqReset)
	}

	//fmt.Print(seqPosRestore)
	moveCursorUp(y)
	moveCursorBackwards(termWidth)
	moveCursorForwards(len(Prompt) + pos)
}

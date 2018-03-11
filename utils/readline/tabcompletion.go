package readline

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
)

func tabCompletion() {
	if TabCompleter == nil {
		return
	}

	suggestions := TabCompleter(line, pos)
	if len(suggestions) == 0 {
		return
	}

	if len(suggestions) == 1 {
		if len(suggestions[0]) == 0 {
			return
		}
		insert([]byte(suggestions[0]))
		return
	}

	//fmt.Print("\r\n")
	//fmt.Print(suggestions)
	//fmt.Print("\r\n" + Prompt + string(line))

	renderSuggestions(suggestions)
}

func renderSuggestions(suggestions []string) {
	fd := int(os.Stdout.Fd())
	width, _, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	//fmt.Print(seqPosSave + "\r\n")
	fmt.Print("\r\n")

	s := string(line)
	maxLength := 1
	for i := range suggestions {
		suggestions[i] = s + suggestions[i]
		if len(suggestions[i]) > maxLength {
			maxLength = len(suggestions[i])
		}
	}

	xCells := int(width / (maxLength + 2))
	cellWidth := strconv.Itoa((width / xCells) - 2)

	yCells := 4
	x := 0
	y := 1
	for i := range suggestions {
		x++
		if x > xCells {
			x = 1
			y++
			if y > yCells {
				break
			} else {
				fmt.Print("\r\n")
			}
		}

		fmt.Printf(" %-"+cellWidth+"s ", suggestions[i])
	}

	//fmt.Print(seqPosRestore)
	moveCursorUp(y)
	moveCursorBackwards(width)
	moveCursorForwards(len(Prompt) + pos)
}

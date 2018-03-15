package readline

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

var (
	termWidth int
	hintY     int = 0
)

func init() {
	getTermWidth()
}

func getTermWidth() {
	var err error
	fd := int(os.Stdout.Fd())
	termWidth, _, err = terminal.GetSize(fd)
	if err != nil {
		termWidth = 100
	}
}

func syntaxCompletion() {
	if SyntaxCompleter == nil {
		return
	}

	x := pos
	if !modeViKeys && pos > 0 {
		x--
	}
	newLine, newPos := SyntaxCompleter(line, x)
	if string(newLine) == string(line) {
		return
	}

	newPos++

	line = newLine
	echo()
	moveCursorForwards(newPos - pos - 1)
	moveCursorBackwards(pos - newPos + 1)
	pos = newPos
}

func renderHintText() {
	if HintText == nil {
		hintY = 0
		return
	}

	r := HintText(line, pos)
	if len(r) == 0 && hintY == 0 {
		return
	}

	moveY := hintY
	if moveY == 0 {
		moveY = 1
	}

	blankChars := (termWidth * moveY) - len(r)
	if blankChars < 0 {
		blankChars = 1
	}

	blank := strings.Repeat(" ", blankChars)

	if len(r) > 0 {
		hintY = (len(r) / termWidth) + 1
		if hintY > moveY {
			moveY = hintY
		}
	} else {
		hintY = 0
	}

	fmt.Print("\r\n" + seqFgBlue + string(r) + seqReset + blank)
	moveCursorBackwards(termWidth)
	moveCursorUp(moveY)
	moveCursorForwards(promptLen + pos + 1)
}

func clearHintText() {
	if hintY == 0 {
		return
	}

	move := termWidth * hintY
	blank := strings.Repeat(" ", move)

	fmt.Print("\r\n" + blank)
	moveCursorUp(hintY)
	moveCursorBackwards(termWidth)
	moveCursorForwards(promptLen + pos)

	hintY = 0
}

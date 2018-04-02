package readline

import (
	"fmt"
	"os"
	"strings"
)

func getTermWidth() {
	var err error
	fd := int(os.Stdout.Fd())
	termWidth, _, err = GetSize(fd)
	if err != nil {
		termWidth = 100
	}
}

func (rl *Instance) syntaxCompletion() {
	if rl.SyntaxCompleter == nil {
		return
	}

	//x := pos
	//if modeViMode == vimInsert && pos > 0 {
	//	x--
	//}
	newLine, newPos := rl.SyntaxCompleter(rl.line, rl.pos-1)
	if string(newLine) == string(rl.line) {
		return
	}

	newPos++

	rl.line = newLine
	rl.echo()
	moveCursorForwards(newPos - rl.pos - 1)
	moveCursorBackwards(rl.pos - newPos + 1)
	rl.pos = newPos
}

func (rl *Instance) renderHintText() {
	if rl.HintText == nil && rl.hintY == 0 {
		//rl.hintY = 0
		return
	}

	r := rl.HintText(rl.line, rl.pos)
	if len(r) == 0 && rl.hintY == 0 {
		return
	}

	rl.writeHintText(r)
}

func (rl *Instance) writeHintText(r []rune) {
	moveY := rl.hintY
	if moveY == 0 {
		moveY = 1
	}

	blankChars := (termWidth * moveY) - len(r)
	if blankChars < 0 {
		blankChars = 1
	}

	blank := strings.Repeat(" ", blankChars)

	if len(r) > 0 {
		rl.hintY = (len(r) / termWidth) + 1
		if rl.hintY > moveY {
			moveY = rl.hintY
		}
	} else {
		rl.hintY = 0
	}

	fmt.Print("\r\n" + seqFgBlue + string(r) + seqReset + blank)
	moveCursorBackwards(termWidth)
	moveCursorUp(moveY)
	moveCursorForwards(rl.promptLen + rl.pos + 1)
}

func (rl *Instance) clearHintText() {
	if rl.hintY == 0 {
		return
	}

	move := termWidth * rl.hintY
	blank := strings.Repeat(" ", move)

	fmt.Print("\r\n" + blank)
	moveCursorUp(rl.hintY)
	moveCursorBackwards(termWidth)
	moveCursorForwards(rl.promptLen + rl.pos)

	rl.hintY = 0
}

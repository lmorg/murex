package readline

import (
	"fmt"
	"strings"
)

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
	getTermWidth()

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
		rl.hintY += strings.Count(string(r), "\n")
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

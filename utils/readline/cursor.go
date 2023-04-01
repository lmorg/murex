package readline

import (
	"regexp"
	"strconv"
)

func leftMost() []byte {
	fd := int(primary.Fd())
	w, _, err := GetSize(fd)
	if err != nil {
		return []byte{'\r', '\n'}
	}

	b := make([]byte, w+1)
	for i := 0; i < w; i++ {
		b[i] = ' '
	}
	b[w] = '\r'

	return b
}

var rxRcvCursorPos = regexp.MustCompile("^\x1b([0-9]+);([0-9]+)R$")

func (rl *Instance) getCursorPos() (x int, y int) {
	if !ForceCrLf {
		return 0, 0
	}

	if !rl.EnableGetCursorPos {
		return -1, -1
	}

	disable := func() (int, int) {
		printErr("\r\ngetCursorPos() not supported by terminal emulator, disabling....\r\n")
		rl.EnableGetCursorPos = false
		return -1, -1
	}

	print(seqGetCursorPos)
	b := make([]byte, 64)
	i, err := read(b)
	if err != nil {
		return disable()
	}

	if !rxRcvCursorPos.Match(b[:i]) {
		return disable()
	}

	match := rxRcvCursorPos.FindAllStringSubmatch(string(b[:i]), 1)
	y, err = strconv.Atoi(match[0][1])
	if err != nil {
		return disable()
	}

	x, err = strconv.Atoi(match[0][2])
	if err != nil {
		return disable()
	}

	return x, y
}

func moveCursorUp(i int) {
	if i < 1 {
		return
	}

	printf("\x1b[%dA", i)
}

func moveCursorDown(i int) {
	if i < 1 {
		return
	}

	printf("\x1b[%dB", i)
}

func moveCursorForwards(i int) {
	if i < 1 {
		return
	}

	printf("\x1b[%dC", i)
}

func moveCursorBackwards(i int) {
	if i < 1 {
		return
	}

	printf("\x1b[%dD", i)
}

func (rl *Instance) moveCursorToStart() {
	posX, posY := lineWrapCellPos(rl.promptLen, rl.line.CellPos(), rl.termWidth)

	moveCursorBackwards(posX - rl.promptLen)
	moveCursorUp(posY)
}

func (rl *Instance) moveCursorFromStartToLinePos() {
	posX, posY := lineWrapCellPos(rl.promptLen, rl.line.CellPos(), rl.termWidth)
	moveCursorForwards(posX)
	moveCursorDown(posY)
}

func (rl *Instance) moveCursorFromEndToLinePos() {
	lineX, lineY := lineWrapCellPos(rl.promptLen, rl.line.CellLen(), rl.termWidth)
	posX, posY := lineWrapCellPos(rl.promptLen, rl.line.CellPos(), rl.termWidth)
	moveCursorBackwards(lineX - posX)
	moveCursorUp(lineY - posY)
}

func (rl *Instance) moveCursorByRuneAdjust(rAdjust int) {
	oldX, oldY := lineWrapCellPos(rl.promptLen, rl.line.CellPos(), rl.termWidth)

	rl.line.SetRunePos(rl.line.RunePos() + rAdjust)

	if rl.line.RunePos() < 0 {
		rl.line.SetRunePos(0)
	}
	if rl.line.RunePos() > rl.line.RuneLen() {
		rl.line.SetRunePos(rl.line.RuneLen())
	}

	if rl.modeViMode != vimInsert && rl.line.RunePos() == rl.line.RuneLen() {
		rl.line.SetRunePos(rl.line.RunePos() - 1)
	}

	newX, newY := lineWrapCellPos(rl.promptLen, rl.line.CellPos(), rl.termWidth)

	y := newY - oldY
	switch {
	case y < 0:
		moveCursorUp(-y)
	case y > 0:
		moveCursorDown(y)
	}

	x := newX - oldX
	switch {
	case x < 0:
		moveCursorBackwards(-x)
	case x > 0:
		moveCursorForwards(x)
	}
}

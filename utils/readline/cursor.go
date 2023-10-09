package readline

import (
	"fmt"
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

const (
	cursorUpf   = "\x1b[%dA"
	cursorDownf = "\x1b[%dB"
	cursorForwf = "\x1b[%dC"
	cursorBackf = "\x1b[%dD"
)

func moveCursorUpStr(i int) string {
	if i < 1 {
		return ""
	}

	return fmt.Sprintf(cursorUpf, i)
}

func moveCursorDownStr(i int) string {
	if i < 1 {
		return ""
	}

	return fmt.Sprintf(cursorDownf, i)
}

func moveCursorForwardsStr(i int) string {
	if i < 1 {
		return ""
	}

	return fmt.Sprintf(cursorForwf, i)
}

func moveCursorBackwardsStr(i int) string {
	if i < 1 {
		return ""
	}

	return fmt.Sprintf(cursorBackf, i)
}

func (rl *Instance) moveCursorToStartStr() string {
	posX, posY := rl.lineWrapCellPos()
	return moveCursorBackwardsStr(posX-rl.promptLen) + moveCursorUpStr(posY)
}

func (rl *Instance) moveCursorFromStartToLinePosStr() string {
	posX, posY := rl.lineWrapCellPos()
	output := moveCursorForwardsStr(posX)
	output += moveCursorDownStr(posY)
	return output
}

func (rl *Instance) moveCursorFromEndToLinePosStr() string {
	lineX, lineY := rl.lineWrapCellLen()
	posX, posY := rl.lineWrapCellPos()
	output := moveCursorBackwardsStr(lineX - posX)
	output += moveCursorUpStr(lineY - posY)
	return output
}

func (rl *Instance) moveCursorByRuneAdjustStr(rAdjust int) string {
	oldX, oldY := rl.lineWrapCellPos()

	rl.line.SetRunePos(rl.line.RunePos() + rAdjust)

	newX, newY := rl.lineWrapCellPos()

	y := newY - oldY

	var output string

	switch {
	case y < 0:
		output += moveCursorUpStr(-y)
	case y > 0:
		output += moveCursorDownStr(y)
	}

	x := newX - oldX
	switch {
	case x < 0:
		output += moveCursorBackwardsStr(-x)
	case x > 0:
		output += moveCursorForwardsStr(x)
	}

	return output
}

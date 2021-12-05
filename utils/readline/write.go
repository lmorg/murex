package readline

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode/utf8"
)

func printf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	print(s)
}

//var rxAnsiSgr = regexp.MustCompile("\x1b\\[[:;0-9]+m")
var rxAnsiSgr = regexp.MustCompile(`\x1b\[([0-9]{1,2}(;[0-9]{1,2})*)?[m|K]`)

// Gets the number of runes in a string and
func strLen(s string) int {
	s = rxAnsiSgr.ReplaceAllString(s, "")
	return utf8.RuneCountInString(s)
}

func (rl *Instance) echo() {
	if len(rl.multisplit) == 0 {
		rl.syntaxCompletion()
	}

	lineX, lineY := lineWrapPos(rl.promptLen, len(rl.line), rl.termWidth)
	posX, posY := lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)

	moveCursorBackwards(posX)
	moveCursorUp(posY)
	if rl.promptLen < rl.termWidth {
		print(rl.prompt)
	}

	switch {
	case rl.PasswordMask != 0:
		print(strings.Repeat(string(rl.PasswordMask), len(rl.line)) + " \r\n")

	case len(rl.line)+rl.promptLen > rl.termWidth:
		fallthrough

	case rl.SyntaxHighlighter == nil:
		wrap := lineWrap(rl, rl.termWidth)
		for i := range wrap {
			print(wrap[i] + "\r\n")
		}

	default:
		print(rl.SyntaxHighlighter(rl.line) + " \r\n")
	}

	//lineX, lineY := lineWrapPos(rl.promptLen, strLen(string(rl.line)), rl.termWidth)
	//posX, posY = lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)

	moveCursorUp(lineY + 1)
	moveCursorDown(posY)
	moveCursorBackwards(lineX - posX + 1)
}

func lineWrap(rl *Instance, termWidth int) []string {
	var promptLen int
	if rl.promptLen < termWidth {
		promptLen = rl.promptLen
	}

	n := float64(len(rl.line)+1) / (float64(termWidth) - float64(promptLen))
	if n < 0 {
		return []string{" "}
	}

	var (
		ceil = int(math.Ceil(n))
		wrap = make([]string, ceil)
		l    = termWidth - promptLen
		line = string(rl.line) + " "
	)

	for i := 0; i < ceil; i++ {
		if i > 0 {
			wrap[i] = strings.Repeat(" ", promptLen)
		}
		if i == ceil-1 {
			wrap[i] += line[l*i:]
			break
		}
		wrap[i] += line[l*i : l*(i+1)]
	}

	return wrap
}

func lineWrapPos(promptLen, lineLength, termWidth int) (x, y int) {
	if promptLen >= termWidth {
		promptLen = 0
	}

	y = lineLength / (termWidth - promptLen)
	if y < 0 {
		return 0, 0
	}

	l := termWidth - promptLen
	x = lineLength - (l * y)
	x += promptLen

	return
}

func (rl *Instance) clearLine() {
	if len(rl.line) == 0 {
		return
	}

	rl.moveCursorToStart()

	if rl.termWidth > rl.promptLen {
		print(strings.Repeat(" ", rl.termWidth-rl.promptLen))
	}
	print(seqClearScreenBelow)

	moveCursorBackwards(rl.termWidth)
	print(rl.prompt)

	rl.line = []rune{}
	rl.pos = 0
}

func (rl *Instance) resetHelpers() {
	rl.modeAutoFind = false
	rl.clearHelpers()
	rl.resetHintText()
	rl.resetTabCompletion()
}

func (rl *Instance) clearHelpers() {
	posX, posY := lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)
	_, lineY := lineWrapPos(rl.promptLen, len(rl.line), rl.termWidth)
	y := lineY - posY

	moveCursorDown(y)
	print("\r\n" + seqClearScreenBelow)

	moveCursorUp(y + 1)
	moveCursorForwards(posX)
}

func (rl *Instance) renderHelpers() {
	rl.writeHintText(true)
	rl.writeTabCompletion(true)
}

func (rl *Instance) updateHelpers() {
	rl.tcOffset = 0
	rl.getHintText()
	if rl.modeTabCompletion {
		rl.getTabCompletion()
	}
	rl.clearHelpers()
	rl.renderHelpers()
}

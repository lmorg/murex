package readline

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
)

func printf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	print(s)
}

// var rxAnsiSgr = regexp.MustCompile("\x1b\\[[:;0-9]+m")
var rxAnsiSgr = regexp.MustCompile(`\x1b\[([0-9]{1,2}(;[0-9]{1,2})*)?[m|K]`)

// Gets the number of runes in a string and
func strLen(s string) int {
	s = rxAnsiSgr.ReplaceAllString(s, "")
	//return utf8.RuneCountInString(s)
	return runewidth.StringWidth(s)
}

func (rl *Instance) echo() {
	if len(rl.multiSplit) == 0 {
		rl.syntaxCompletion()
	}

	lineX, lineY := lineWrapPos(rl.promptLen, rl.line.Len(), rl.termWidth)
	posX, posY := lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)

	moveCursorBackwards(posX)
	moveCursorUp(posY)
	if rl.promptLen < rl.termWidth {
		print(rl.prompt)
	}

	switch {
	case rl.PasswordMask != 0:
		print(strings.Repeat(string(rl.PasswordMask), rl.line.Len()) + " \r\n")

	case rl.line.Len()+rl.promptLen > rl.termWidth:
		fallthrough

	case rl.SyntaxHighlighter == nil:
		wrap := lineWrap(rl, rl.termWidth)
		for i := range wrap {
			print(wrap[i] + "\r\n")
		}

	default:
		syntax := rl.cacheSyntax.Get(rl.line.Value)
		if len(syntax) > 0 {
			print(syntax + " \r\n")

		} else {
			syntax = rl.SyntaxHighlighter(rl.line.Value)
			print(syntax + " \r\n")

			if rl.DelayedSyntaxWorker == nil {
				rl.cacheSyntax.Append(rl.line.Value, syntax)
			}
		}
		//print(string(rl.line) + " \r\n")
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

	n := float64(rl.line.Len()+1) / (float64(termWidth) - float64(promptLen))
	ceil := int(math.Ceil(n))
	if ceil < 1 || ceil > 2000000000 {
		return []string{" "}
	}

	var (
		wrap = make([]string, ceil)
		l    = termWidth - promptLen
		line = rl.line.String() + " "
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

func lineWrapCellPos(promptLen, lineLength, termWidth int) (x, y int) {
	if promptLen >= termWidth {
		promptLen = 0
	}

	// avoid divide by zero error
	if termWidth-promptLen == 0 {
		return 0, 0
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
	if rl.line.Len() == 0 {
		return
	}

	rl.moveCursorToStart()

	if rl.termWidth > rl.promptLen {
		print(strings.Repeat(" ", rl.termWidth-rl.promptLen))
	}
	print(seqClearScreenBelow)

	moveCursorBackwards(rl.termWidth)
	print(rl.prompt)

	rl.line.Value = []rune{}
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
	_, lineY := lineWrapPos(rl.promptLen, rl.line.Len(), rl.termWidth)
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

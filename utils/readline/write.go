package readline

import (
	"fmt"
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
	return runewidth.StringWidth(s)
}

func (rl *Instance) echo() {
	if len(rl.multiSplit) == 0 {
		rl.syntaxCompletion()
	}

	lineX, lineY := rl.lineWrapCellLen()
	posX, posY := rl.lineWrapCellPos()

	moveCursorBackwards(posX)
	moveCursorUp(posY)
	if rl.promptLen < rl.termWidth {
		print(rl.prompt)
	}

	switch {
	case rl.PasswordMask != 0:
		print(strings.Repeat(string(rl.PasswordMask), rl.line.CellLen()) + " \r\n")

	case rl.line.CellLen()+rl.promptLen > rl.termWidth:
		fallthrough

	case rl.SyntaxHighlighter == nil:
		wrap := lineWrap(rl, rl.termWidth)
		for i := range wrap {
			print(wrap[i] + "\r\n")
		}

	default:
		syntax := rl.cacheSyntax.Get(rl.line.Runes())
		if len(syntax) > 0 {
			print(syntax + " \r\n")

		} else {
			syntax = rl.SyntaxHighlighter(rl.line.Runes())
			print(syntax + " \r\n")

			if rl.DelayedSyntaxWorker == nil {
				rl.cacheSyntax.Append(rl.line.Runes(), syntax)
			}
		}
	}

	moveCursorUp(lineY + 1)
	moveCursorDown(posY)
	moveCursorBackwards(lineX - posX + 1)
}

func lineWrap(rl *Instance, termWidth int) []string {
	var promptLen int
	if rl.promptLen < termWidth {
		promptLen = rl.promptLen
	}

	var (
		wrap       []string
		wrapRunes  [][]rune
		bufCellLen int
		length     = termWidth - promptLen
		line       = append(rl.line.Runes(), []rune{' '}...) // double space to work around wide characters
		lPos       int
	)

	wrapRunes = append(wrapRunes, []rune{})

	for r := range line {
		w := runewidth.RuneWidth(line[r])
		if bufCellLen+w > length {
			wrapRunes = append(wrapRunes, []rune(strings.Repeat(" ", promptLen)))
			lPos++
			bufCellLen = 0
		}
		bufCellLen += w
		wrapRunes[lPos] = append(wrapRunes[lPos], line[r])
	}

	wrap = make([]string, lPos+1)
	for i := range wrap {
		wrap[i] = string(wrapRunes[i])
	}

	// shouldn't get this far either, but just in case
	return wrap
}

func (rl *Instance) lineWrapCellLen() (x, y int) {
	return lineWrapCell(rl.promptLen, rl.line.Runes(), rl.termWidth)
}

func (rl *Instance) lineWrapCellPos() (x, y int) {
	return lineWrapCell(rl.promptLen, rl.line.Runes()[:rl.line.RunePos()], rl.termWidth)
}

func lineWrapCell(promptLen int, line []rune, termWidth int) (x, y int) {
	if promptLen >= termWidth {
		promptLen = 0
	}

	// avoid divide by zero error
	if termWidth-promptLen == 0 {
		return 0, 0
	}

	x = promptLen
	for i := range line {
		w := runewidth.RuneWidth(line[i])
		if x+w > termWidth {
			x = promptLen
			y++
		}
		x += w
	}

	return
}

func (rl *Instance) clearLine() {
	if rl.line.RuneLen() == 0 {
		return
	}

	rl.moveCursorToStart()

	if rl.termWidth > rl.promptLen {
		print(strings.Repeat(" ", rl.termWidth-rl.promptLen))
	}
	print(seqClearScreenBelow)

	moveCursorBackwards(rl.termWidth)
	print(rl.prompt)

	rl.line.Set([]rune{})
	rl.line.SetRunePos(0)
}

func (rl *Instance) resetHelpers() {
	rl.modeAutoFind = false
	rl.clearHelpers()
	rl.resetHintText()
	rl.resetTabCompletion()
}

func (rl *Instance) clearHelpers() {
	posX, posY := rl.lineWrapCellPos()
	_, lineY := rl.lineWrapCellLen()
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

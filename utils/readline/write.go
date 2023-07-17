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

	// reset cursor to start
	line := "\r"
	if posY > 0 {
		line += fmt.Sprintf(cursorUpf, posY)
	}

	// clear the line
	line += strings.Repeat("\x1b[2K\n", lineY+1) // clear line + move cursor down 1
	line += fmt.Sprintf(cursorUpf, lineY+1)
	//line += seqClearScreenBelow

	promptLen := rl.promptLen
	if promptLen < rl.termWidth {
		line += rl.prompt
	} else {
		promptLen = 0
	}

	switch {
	case rl.PasswordMask != 0:
		line += strings.Repeat(string(rl.PasswordMask), rl.line.CellLen())

	case rl.line.CellLen()+promptLen > rl.termWidth:
		fallthrough

	case rl.SyntaxHighlighter == nil:
		line += strings.Join(lineWrap(rl, rl.termWidth), "\r\n")

	default:
		syntax := rl.cacheSyntax.Get(rl.line.Runes())
		if len(syntax) == 0 {
			syntax = rl.SyntaxHighlighter(rl.line.Runes())

			if rl.DelayedSyntaxWorker == nil {
				rl.cacheSyntax.Append(rl.line.Runes(), syntax)
			}
		}
		line += syntax
	}

	y := lineY - posY
	if y > 0 {
		line += fmt.Sprintf(cursorUpf, y)
	}
	x := lineX - posX + 1
	if x > 0 {
		line += fmt.Sprintf(cursorBackf, x)
	}
	print(line + moveCursorForwardsStr(1))
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
		line       = rl.line.Runes() //append(rl.line.Runes(), []rune{' ', ' '}...) // double space to work around wide characters
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

func (rl *Instance) clearPrompt() {
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

	output := moveCursorDownStr(y)
	output += "\r\n" + seqClearScreenBelow

	output += moveCursorUpStr(y + 1)
	output += moveCursorForwardsStr(posX)

	print(output)
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

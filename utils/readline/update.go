package readline

import (
	"strings"
)

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

func moveCursorToLinePos(rl *Instance) {
	moveCursorForwards(rl.promptLen + rl.pos)
}

func (rl *Instance) moveCursorByAdjust(adjust int) {
	switch {
	case adjust > 0:
		moveCursorForwards(adjust)
		rl.pos += adjust
	case adjust < 0:
		moveCursorBackwards(adjust * -1)
		rl.pos += adjust
	}

	if rl.modeViMode != vimInsert && rl.pos == len(rl.line) {
		moveCursorBackwards(1)
		rl.pos--
	}
}

func (rl *Instance) insert(r []rune) {
	for {
		// I don't really understand why `0` is creaping in at the end of the
		// array but it only happens with unicode characters.
		if len(r) > 1 && r[len(r)-1] == 0 {
			r = r[:len(r)-1]
			continue
		}
		break
	}

	switch {
	case len(rl.line) == 0:
		rl.line = r
	case rl.pos == 0:
		rl.line = append(r, rl.line...)
	case rl.pos < len(rl.line):
		r := append(r, rl.line[rl.pos:]...)
		rl.line = append(rl.line[:rl.pos], r...)
	default:
		rl.line = append(rl.line, r...)
	}

	rl.echo()

	rl.pos += len(r)
	moveCursorForwards(len(r) - 1)

	if rl.modeViMode == vimInsert {
		rl.updateHelpers()
	}
}

func (rl *Instance) backspace() {
	if len(rl.line) == 0 || rl.pos == 0 {
		return
	}

	moveCursorBackwards(1)
	rl.pos--
	rl.delete()
}

func (rl *Instance) delete() {
	switch {
	case len(rl.line) == 0:
		return
	case rl.pos == 0:
		rl.line = rl.line[1:]
		rl.echo()
		moveCursorBackwards(1)
	case rl.pos > len(rl.line):
		rl.backspace()
	case rl.pos == len(rl.line):
		rl.line = rl.line[:rl.pos]
		rl.echo()
		moveCursorBackwards(1)
	default:
		rl.line = append(rl.line[:rl.pos], rl.line[rl.pos+1:]...)
		rl.echo()
		moveCursorBackwards(1)
	}

	rl.updateHelpers()
}

func (rl *Instance) echo() {
	moveCursorBackwards(rl.pos)

	switch {
	case rl.PasswordMask > 0:
		print(strings.Repeat(string(rl.PasswordMask), len(rl.line)) + " ")

	case rl.SyntaxHighlighter == nil:
		print(string(rl.line) + " ")

	default:
		print(rl.SyntaxHighlighter(rl.line) + " ")
	}

	moveCursorBackwards(len(rl.line) - rl.pos)
}

func (rl *Instance) clearLine() {
	if len(rl.line) == 0 {
		return
	}

	moveCursorBackwards(rl.pos)
	print(strings.Repeat(" ", len(rl.line)))
	moveCursorBackwards(len(rl.line))

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
	print("\r\n" + seqClearScreenBelow)
	moveCursorUp(1)
	moveCursorToLinePos(rl)
}

func (rl *Instance) renderHelpers() {
	rl.writeHintText()
	rl.writeTabCompletion()

	moveCursorUp(rl.hintY + rl.tcUsedY)
	moveCursorBackwards(GetTermWidth())
	moveCursorToLinePos(rl)
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

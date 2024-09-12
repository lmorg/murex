package readline

import (
	"strconv"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/mattn/go-runewidth"
)

type viMode int

const (
	vimInsert viMode = iota
	vimReplaceOnce
	vimReplaceMany
	vimDelete
	vimKeys
)

func (rl *Instance) vi(r rune) string {
	if !debug.Enabled {
		// This would normally be a massive anti-pattern. But in this instance
		// any edge case exceptions are better off ignored and the interactive
		// prompt kept alive. The worst case scenario is the cursor might
		// become misaligned but that would resolve itself very quickly under
		// normal operation.
		defer recover()
	}

	var output string
	switch r {
	case 'a':
		if rl.line.CellLen() > 0 {
			output = moveCursorForwardsStr(1)
			rl.line.SetRunePos(rl.line.RunePos())
		}
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'A':
		if rl.line.RuneLen() > 0 {
			output = moveCursorForwardsStr(rl.line.CellLen() - rl.line.CellPos())
			rl.line.SetRunePos(rl.line.RuneLen())
		}
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'b':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.moveCursorByRuneAdjustStr(rl.viJumpB(tokeniseLine))
		}

	case 'B':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.moveCursorByRuneAdjustStr(rl.viJumpB(tokeniseSplitSpaces))
		}

	case 'd':
		rl.modeViMode = vimDelete
		rl.viUndoSkipAppend = true

	case 'D':
		output = moveCursorBackwardsStr(rl.line.CellPos())
		output += strings.Repeat(" ", rl.line.CellLen())

		output += moveCursorBackwardsStr(rl.line.CellLen() - rl.line.CellPos())
		rl.line.Set(rl, rl.line.Runes()[:rl.line.RunePos()])
		output += rl.echoStr()

		r := rl.line.Runes()[rl.line.RuneLen()-1]
		output += moveCursorBackwardsStr(1 + runewidth.RuneWidth(r))
		rl.line.SetRunePos(rl.line.RunePos() - 1)
		rl.viIteration = ""

	case 'e':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.moveCursorByRuneAdjustStr(rl.viJumpE(tokeniseLine))
		}

	case 'E':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.moveCursorByRuneAdjustStr(rl.viJumpE(tokeniseSplitSpaces))
		}

	case 'h':
		if rl.line.RunePos() > 0 {
			r := rl.line.Runes()[rl.line.RunePos()-1]
			output += moveCursorBackwardsStr(runewidth.RuneWidth(r))
			rl.line.SetRunePos(rl.line.RunePos() - 1)
		}
		rl.viUndoSkipAppend = true

	case 'i':
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'I':
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true
		output += moveCursorBackwardsStr(rl.line.CellPos())
		rl.line.SetRunePos(0)

	case 'l':
		// TODO: test me
		if (rl.modeViMode == vimInsert && rl.line.RunePos() < rl.line.RuneLen()) ||
			(rl.modeViMode != vimInsert && rl.line.RunePos() < rl.line.RuneLen()-1) {
			r := rl.line.Runes()[rl.line.RunePos()+1]
			output += moveCursorForwardsStr(runewidth.RuneWidth(r))
			rl.line.SetRunePos(rl.line.RunePos() + 1)
		}
		rl.viUndoSkipAppend = true

	case 'p':
		// paste after
		if len(rl.line.Runes()) == 0 {
			return ""
		}

		rl.viUndoSkipAppend = true
		w := runewidth.RuneWidth(rl.line.Runes()[rl.line.RunePos()])

		rl.line.SetRunePos(rl.line.RunePos() + 1)
		output += moveCursorForwardsStr(w)

		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.insertStr([]rune(rl.viYankBuffer))
		}

		rl.line.SetRunePos(rl.line.RunePos() - 1)
		output += moveCursorBackwardsStr(w)

	case 'P':
		// paste before
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.insertStr([]rune(rl.viYankBuffer))
		}

	case 'r':
		rl.modeViMode = vimReplaceOnce
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'R':
		rl.modeViMode = vimReplaceMany
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'u':
		output = rl.undoLastStr()
		rl.viUndoSkipAppend = true

	case 'v':
		output = rl.clearHelpersStr()
		var multiline []rune
		if rl.GetMultiLine == nil {
			multiline = rl.line.Runes()
		} else {
			multiline = rl.GetMultiLine(rl.line.Runes())
		}

		new, err := rl.launchEditor(multiline)
		if err != nil || len(new) == 0 || string(new) == string(multiline) {
			rl.viUndoSkipAppend = true
			return ""
		}
		rl.clearPrompt()
		rl.multiline = []byte(string(new))

	case 'w':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.moveCursorByRuneAdjustStr(rl.viJumpW(tokeniseLine))
		}

	case 'W':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.moveCursorByRuneAdjustStr(rl.viJumpW(tokeniseSplitSpaces))
		}

	case 'x':
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			output += rl.deleteStr()
		}
		if rl.line.RunePos() == rl.line.RuneLen() && rl.line.RuneLen() > 0 {
			///// TODO !!!!!!!!!!
			r := rl.line.Runes()[rl.line.RunePos()-1]
			output += moveCursorBackwardsStr(runewidth.RuneWidth(r))
			rl.line.SetRunePos(rl.line.RunePos() - 1)
		}

	case 'y', 'Y':
		rl.viYankBuffer = rl.line.String()
		rl.viUndoSkipAppend = true
		//rl.hintText = []rune("-- LINE YANKED --")
		//rl.renderHelpers()

	case '[':
		rl.viUndoSkipAppend = true
		output = rl.moveCursorByRuneAdjustStr(rl.viJumpPreviousBrace())

	case ']':
		rl.viUndoSkipAppend = true
		output = rl.moveCursorByRuneAdjustStr(rl.viJumpNextBrace())

	case '$':
		output = moveCursorForwardsStr(rl.line.CellLen() - rl.line.CellPos())
		rl.line.SetRunePos(rl.line.RuneLen())
		rl.viUndoSkipAppend = true

	case '%':
		rl.viUndoSkipAppend = true
		output = rl.moveCursorByRuneAdjustStr(rl.viJumpBracket())

	default:
		if r <= '9' && '0' <= r {
			rl.viIteration += string(r)
		}
		rl.viUndoSkipAppend = true

	}

	return output
}

func (rl *Instance) getViIterations() int {
	i, _ := strconv.Atoi(rl.viIteration)
	if i < 1 {
		i = 1
	}
	rl.viIteration = ""
	return i
}

func (rl *Instance) viHintMessageStr() string {
	switch rl.modeViMode {
	case vimKeys:
		rl.hintText = []rune("-- VIM KEYS -- (press `i` to return to normal editing mode)")
	case vimInsert:
		rl.hintText = []rune("-- INSERT --")
	case vimReplaceOnce:
		rl.hintText = []rune("-- REPLACE CHARACTER --")
	case vimReplaceMany:
		rl.hintText = []rune("-- REPLACE --")
	case vimDelete:
		rl.hintText = []rune("-- DELETE --")
	default:
		rl.getHintText()
	}

	output := rl.clearHelpersStr()
	output += rl.renderHelpersStr()
	return output
}

func (rl *Instance) viJumpB(tokeniser func([]rune, int) ([]string, int, int)) (adjust int) {
	split, index, pos := tokeniser(rl.line.Runes(), rl.line.RunePos())
	switch {
	case len(split) == 0:
		return
	case index == 0 && pos == 0:
		return
	case pos == 0:
		adjust = len(split[index-1])
	default:
		adjust = pos
	}
	return adjust * -1
}

func (rl *Instance) viJumpE(tokeniser func([]rune, int) ([]string, int, int)) (adjust int) {
	split, index, pos := tokeniser(rl.line.Runes(), rl.line.RunePos())
	if len(split) == 0 {
		return
	}

	word := rTrimWhiteSpace(split[index])

	switch {
	case len(split) == 0:
		return
	case index == len(split)-1 && pos >= len(word)-1:
		return
	case pos >= len(word)-1:
		word = rTrimWhiteSpace(split[index+1])
		adjust = len(split[index]) - pos
		adjust += len(word) - 1
	default:
		adjust = len(word) - pos - 1
	}
	return
}

func (rl *Instance) viJumpW(tokeniser func([]rune, int) ([]string, int, int)) (adjust int) {
	split, index, pos := tokeniser(rl.line.Runes(), rl.line.RunePos())
	switch {
	case len(split) == 0:
		return
	case index+1 == len(split):
		adjust = rl.line.RuneLen() - 1 - rl.line.RunePos()
	default:
		adjust = len(split[index]) - pos
	}
	return
}

func (rl *Instance) viJumpPreviousBrace() (adjust int) {
	if rl.line.RunePos() == 0 {
		return 0
	}

	for i := rl.line.RunePos() - 1; i != 0; i-- {
		if rl.line.Runes()[i] == '{' {
			return i - rl.line.RunePos()
		}
	}

	return 0
}

func (rl *Instance) viJumpNextBrace() (adjust int) {
	if rl.line.RunePos() >= rl.line.RuneLen()-1 {
		return 0
	}

	for i := rl.line.RunePos() + 1; i < rl.line.RuneLen(); i++ {
		if rl.line.Runes()[i] == '{' {
			return i - rl.line.RunePos()
		}
	}

	return 0
}

func (rl *Instance) viJumpBracket() (adjust int) {
	split, index, pos := tokeniseBrackets(rl.line.Runes(), rl.line.RunePos())
	switch {
	case len(split) == 0:
		return
	case pos == 0:
		adjust = len(split[index])
	default:
		adjust = pos * -1
	}
	return
}

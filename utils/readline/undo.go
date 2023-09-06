package readline

import (
	"strings"

	"github.com/lmorg/murex/utils/readline/unicode"
)

func (rl *Instance) undoAppendHistory() {
	if rl.viUndoSkipAppend {
		rl.viUndoSkipAppend = false
		return
	}

	rl.viUndoHistory = append(rl.viUndoHistory, rl.line.Duplicate())
}

func (rl *Instance) undoLast() {
	var undo *unicode.UnicodeT
	for {
		if len(rl.viUndoHistory) == 0 {
			return
		}
		undo = rl.viUndoHistory[len(rl.viUndoHistory)-1]
		rl.viUndoHistory = rl.viUndoHistory[:len(rl.viUndoHistory)-1]
		if undo.String() != rl.line.String() {
			break
		}
	}

	rl.clearHelpers()

	moveCursorBackwards(rl.line.CellPos())
	print(strings.Repeat(" ", rl.line.CellLen()))
	moveCursorBackwards(rl.line.CellLen())
	moveCursorForwards(undo.CellPos())

	rl.line = undo.Duplicate()

	rl.echo()

	if rl.modeViMode != vimInsert && rl.line.RuneLen() > 0 && rl.line.RunePos() == rl.line.RuneLen() {
		rl.line.SetRunePos(rl.line.RuneLen() - 1)
		moveCursorBackwards(1)
	}
}

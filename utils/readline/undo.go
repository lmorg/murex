package readline

import (
	"strings"
)

func (rl *Instance) undoAppendHistory() {
	if rl.viUndoSkipAppend {
		rl.viUndoSkipAppend = false
		return
	}

	rl.viUndoHistory = append(rl.viUndoHistory, rl.line.Duplicate())
}

func (rl *Instance) undoLastStr() string {
	var undo *UnicodeT
	for {
		if len(rl.viUndoHistory) == 0 {
			return ""
		}
		undo = rl.viUndoHistory[len(rl.viUndoHistory)-1]
		rl.viUndoHistory = rl.viUndoHistory[:len(rl.viUndoHistory)-1]
		if undo.String() != rl.line.String() {
			break
		}
	}

	output := rl.clearHelpersStr()

	output += moveCursorBackwardsStr(rl.line.CellPos())
	output += strings.Repeat(" ", rl.line.CellLen())
	output += moveCursorBackwardsStr(rl.line.CellLen())
	output += moveCursorForwardsStr(undo.CellPos())

	rl.line = undo.Duplicate()

	output += rl.echoStr()

	// TODO: check me
	if rl.modeViMode != vimInsert && rl.line.RuneLen() > 0 && rl.line.RunePos() == rl.line.RuneLen() {
		rl.line.SetRunePos(rl.line.RuneLen() - 1)
		output += moveCursorBackwardsStr(1)
	}

	return output
}

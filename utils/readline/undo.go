package readline

import "strings"

type undoItem struct {
	line string
	pos  int
}

func (rl *Instance) undoAppendHistory() {
	defer func() { rl.viUndoSkipAppend = false }()

	if rl.viUndoSkipAppend {
		return
	}

	rl.viUndoHistory = append(rl.viUndoHistory, undoItem{
		line: rl.line.String(),
		pos:  rl.pos,
	})
}

func (rl *Instance) undoLast() {
	var undo undoItem
	for {
		if len(rl.viUndoHistory) == 0 {
			return
		}
		undo = rl.viUndoHistory[len(rl.viUndoHistory)-1]
		rl.viUndoHistory = rl.viUndoHistory[:len(rl.viUndoHistory)-1]
		if undo.line != rl.line.String() {
			break
		}
	}

	rl.clearHelpers()

	moveCursorBackwards(rl.pos)
	print(strings.Repeat(" ", rl.line.Len()))
	moveCursorBackwards(rl.line.Len())
	moveCursorForwards(undo.pos)

	rl.line.Value = []rune(undo.line)
	rl.pos = undo.pos

	rl.echo()

	if rl.modeViMode != vimInsert && rl.line.Len() > 0 && rl.pos == rl.line.Len() {
		rl.pos--
		moveCursorBackwards(1)
	}

}

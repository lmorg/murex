package readline

func (rl *Instance) syntaxCompletion() {
	if rl.SyntaxCompleter == nil {
		return
	}

	newLine, newPos := rl.SyntaxCompleter(rl.line.Value, rl.lineChange, rl.pos-1)
	if string(newLine) == rl.line.String() {
		return
	}

	newPos++

	rl.line.Value = newLine
	rl.pos = newPos
}

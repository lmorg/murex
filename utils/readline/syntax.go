package readline

func (rl *Instance) syntaxCompletion() {
	if rl.SyntaxCompleter == nil {
		return
	}

	newLine, newPos := rl.SyntaxCompleter(rl.line.Runes(), rl.lineChange, rl.line.RunePos()-1)
	if string(newLine) == rl.line.String() {
		return
	}

	newPos++

	rl.line.Set(rl, newLine)
	rl.line.SetRunePos(newPos)
}

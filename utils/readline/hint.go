package readline

func (rl *Instance) getHintText() {
	if rl.HintText == nil {
		rl.resetHintText()
		return
	}

	rl.hintText = rl.HintText(rl.line, rl.pos)
}

func (rl *Instance) writeHintText() {
	if len(rl.hintText) == 0 {
		rl.hintY = 0
		return
	}

	width := getTermWidth()

	// Determine how many lines hintText spans over
	// (Currently there is no support for carridge returns / new lines)
	hintLength := strLen(string(rl.hintText))
	n := float64(hintLength) / float64(width)
	if float64(int(n)) != n {
		n++
	}
	rl.hintY = int(n)

	print("\r\n" + seqFgBlue + string(rl.hintText) + seqReset)
}

func (rl *Instance) resetHintText() {
	rl.hintY = 0
	rl.hintText = []rune{}
}

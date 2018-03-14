package readline

func vi(b byte) {
	switch b {
	case 'a':
		moveCursorForwards(len(line) - pos)
		pos = len(line)
		modeViKeys = false
	case 'A':
		moveCursorForwards(len(line) - pos)
		pos = len(line)
		modeViKeys = false
	case 'i':
		modeViKeys = false
	case 'x':
		delete()
	}
}

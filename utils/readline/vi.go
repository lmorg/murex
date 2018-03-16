package readline

type viMode int

const (
	vimInsert viMode = iota
	vimReplaceOnce
	vimReplaceMany
	vimKeys
)

var modeViMode viMode = vimInsert

func vi(b byte) {
	switch b {
	case 'a':
		moveCursorForwards(1)
		pos++
		modeViMode = vimInsert
	case 'A':
		moveCursorForwards(len(line) - pos)
		pos = len(line)
		modeViMode = vimInsert
	case 'i':
		//moveCursorForwards(1)
		//pos++
		modeViMode = vimInsert
	case 'r':
		modeViMode = vimReplaceOnce
	case 'R':
		modeViMode = vimReplaceMany
	case 'x':
		delete()
	}
}

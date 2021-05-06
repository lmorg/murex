package termemu

const (
	charEsc           = 27
	charBackspaceIso  = 8
	charBackspaceAnsi = 127
)

func Parse(grid *Grid, text []rune) {
	//grid := NewGrid(120, 40)

	// parser state
	var (
		escape bool
	)

	for _, r := range text {
		switch r {
		case charBackspaceIso, charBackspaceAnsi:
			panic("TODO: backspace")
		case charEsc:
			escape = true

		case '[':
			if escape {
				// separate parser
				continue
			}
			grid.writeCell(r)

		case '\r':
			//if grid.LfIncCr {
			//	continue
			//}
			grid.curPos.X = 0

		case '\n':
			if grid.LfIncCr {
				grid.curPos.X = 0
			}
			if grid.moveCursorDownwards(1) > 0 {
				grid.moveGridUp()
				grid.moveCursorDownwards(1)
			}

		default:
			grid.writeCell(r)
		}
	}
}

func SgrToHtml(s string) string {
	return s
}

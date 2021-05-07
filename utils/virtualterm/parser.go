package virtualterm

import "github.com/lmorg/murex/utils/ansi"

const (
	charEsc           = 27
	charBackspaceIso  = 8
	charBackspaceAnsi = 127
)

func Parse(term *Term, text []rune) {
	//grid := NewGrid(120, 40)

	// parser state
	var (
		escape bool
	)

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case charBackspaceIso, charBackspaceAnsi:
			panic("TODO: backspace")

		case charEsc:
			escape = true

		case '[':
			if escape {
				i += parseSgr(term, text[i:])
				continue
			}
			term.writeCell(text[i])

		case '\r':
			//if grid.LfIncCr {
			//	continue
			//}
			term.curPos.X = 0

		case '\n':
			if term.LfIncCr {
				term.curPos.X = 0
			}
			if term.moveCursorDownwards(1) > 0 {
				term.moveContentsUp()
				term.moveCursorDownwards(1)
			}

		default:
			term.writeCell(text[i])
		}
	}
}

func parseSgr(term *Term, text []rune) int {
	for i, r := range text {
		switch {
		case r >= '0' && '9' >= r:
			// do nothing
		case r == 'm':
			lookupSgr(term, text[:i])
			return i
		default:
			return 0
		}
	}
	return 0
}

func lookupSgr(term *Term, text []rune) {
	sgr := string(text)
	switch sgr {
	case ansi.Reset:
		term.sgrReset()
	case ansi.Bold:
		term.sgrEffect(sgrBold)
	}
}

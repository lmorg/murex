package virtualterm

const (
	charEsc           = 27
	charBackspaceIso  = 8
	charBackspaceAnsi = 127
)

func (term *Term) Write(text []rune) {
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
				i += parseSgr(term, text[i-1:])
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
	i := 2

	for ; i < len(text); i++ {
		switch {
		case text[i] >= '0' && '9' >= text[i]:
			// do nothing
		case text[i] == 'm':
			lookupSgr(term, text[:i+1])
			return i - 1
		default:
			return i - 1
		}
	}
	return i
}

func parseCsi(term *Term, text []rune) int {
	i := 2

	/*for ; i < len(text); i++ {
		switch {
		case text[i] >= '0' && '9' >= text[i]:
			// do nothing
		case text[i] == 'm':
			lookupSgr(term, text[:i+1])
			return i - 1
		default:
			return i - 1
		}
	}*/
	return i
}

func lookupSgr(term *Term, text []rune) {
	sgr := string(text)
	switch sgr {
	case "\x1b[m", "\x1b[0m":
		// reset / normal
		term.sgrReset()

	case "\x1b[1m":
		// bold
		term.sgrEffect(sgrBold)

	case "\x1b[4m":
		// underscore
		term.sgrEffect(sgrUnderscore)

	case "\x1b[5m":
		// blink
		term.sgrEffect(sgrBlink)

		//
		// 4bit foreground colour:
		//

	case "\x1b[30m":
		// fg black
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Black

	case "\x1b[31m":
		// fg red
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Red

	case "\x1b[32m":
		// fg green
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Green

	case "\x1b[33m":
		// fg yellow
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Yellow

	case "\x1b[34m":
		// fg blue
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Blue

	case "\x1b[35m":
		// fg magenta
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Magenta

	case "\x1b[36m":
		// fg cyan
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Cyan

	case "\x1b[37m":
		// fg white
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4White
	}
}

package virtualterm

const (
	charEsc           = 27
	charBackspaceIso  = 8
	charBackspaceAnsi = 127
)

func (term *Term) writeCell(r rune) {
	term.cell().char = r
	term.cell().sgr = term.sgr
	term.wrapCursorForwards()
}

// Write multiple characters to the virtual terminal
func (term *Term) Write(text []rune) {
	var (
		escape bool
	)

	term.mutex.Lock()

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
			if term.state.LfIncCr {
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

	term.mutex.Unlock()
}

func parseSgr(term *Term, text []rune) int {
	i := 2

	var n rune

	for ; i < len(text); i++ {
		switch {
		case text[i] >= '0' && '9' >= text[i]:
			n = (n * 10) + (text[i] - 48)

		case text[i] == 'm': // SGR
			lookupSgr(term, n)
			return i - 1

		case text[i] == 'A': // moveCursorUp
			term.moveCursorUpwards(int(n))
			return i - 1

		case text[i] == 'B': // moveCursorDown
			term.moveCursorDownwards(int(n))
			return i - 1

		case text[i] == 'C': // moveCursorForwards
			term.moveCursorForwards(int(n))
			return i - 1

		case text[i] == 'D': // moveCursorBackwards
			term.moveCursorBackwards(int(n))
			return i - 1

		case text[i] == 'J': // eraseDisplay...
			if n == 0 {
				term.eraseDisplayAfter()
			}

		default:
			return i - 1
		}
	}
	return i
}

func lookupSgr(term *Term, n rune) {
	switch n {
	case 0: // reset / normal
		term.sgrReset()

	case 1: // bold
		term.sgrEffect(sgrBold)

	case 4: // underscore
		term.sgrEffect(sgrUnderscore)

	case 5: // blink
		term.sgrEffect(sgrBlink)

		//
		// 4bit foreground colour:
		//

	case 30: // fg black
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Black

	case 31: // fg red
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Red

	case 32: // fg green
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Green

	case 33: // fg yellow
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Yellow

	case 34: // fg blue
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Blue

	case 35: // fg magenta
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Magenta

	case 36: // fg cyan
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4Cyan

	case 37: // fg white
		term.sgrEffect(sgrFgColour4)
		term.sgr.fg.Red = sgrColour4White
	}
}

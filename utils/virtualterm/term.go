package virtualterm

import (
	"html"
	"strings"
)

type Term struct {
	cells   [][]cell
	size    Xy
	curPos  Xy
	sgr     sgr
	LfIncCr bool // if false, \n acts as a \r\n
}

type cell struct {
	char rune
	sgr  sgr
}

type sgr struct {
	bitwise sgrFlag
	fg      Rgb
	bg      Rgb
}

func (sgr *sgr) Differs(old *sgr) bool {
	if sgr.bitwise != old.bitwise {
		return true
	}

	if sgr.fg.Red != old.fg.Red || sgr.fg.Green != old.fg.Green || sgr.fg.Blue != old.fg.Blue {
		return true
	}

	if sgr.bg.Red != old.bg.Red || sgr.bg.Green != old.bg.Green || sgr.bg.Blue != old.bg.Blue {
		return true
	}

	return false
}

func (sgr *sgr) CheckFlag(flag sgrFlag) bool {
	return sgr.bitwise&flag != 0
}

type Xy struct {
	X int
	Y int
}

type Rgb struct {
	Red, Green, Blue byte
}

func NewTerminal(x, y int) *Term {
	cells := make([][]cell, y, y)
	for i := range cells {
		cells[i] = make([]cell, x, x)
	}

	return &Term{
		cells: cells,
		size:  Xy{x, y},
	}
}

// format

func (t *Term) sgrReset() {
	t.sgr.bitwise = 0
	t.sgr.fg = Rgb{}
	t.sgr.bg = Rgb{}
}

func (t *Term) sgrEffect(flag sgrFlag) {
	t.sgr.bitwise |= flag
}

// moveCursor functions DON'T effect other contents in the grid

func (t *Term) moveCursorBackwards(i int) (overflow int) {
	t.curPos.X -= i
	if t.curPos.X < 0 {
		overflow = t.curPos.X * -1
		t.curPos.X = 0
	}

	return
}

func (t *Term) moveCursorForwards(i int) (overflow int) {
	t.curPos.X += i
	if t.curPos.X >= t.size.X {
		overflow = t.curPos.X - (t.size.X - 1)
		t.curPos.X = t.size.X - 1
	}

	return
}

func (t *Term) moveCursorUpwards(i int) (overflow int) {
	t.curPos.Y -= i
	if t.curPos.Y < 0 {
		overflow = t.curPos.Y * -1
		t.curPos.Y = 0
	}

	return
}

func (t *Term) moveCursorDownwards(i int) (overflow int) {
	t.curPos.Y += i
	if t.curPos.Y >= t.size.Y {
		overflow = t.curPos.Y - (t.size.Y - 1)
		t.curPos.Y = t.size.Y - 1
	}

	return
}

func (t *Term) Cell() *cell { return &t.cells[t.curPos.Y][t.curPos.X] }

// moveGridPos functions DO effect other contents in the grid

func (t *Term) moveContentsUp() {
	var i int
	for ; i < t.size.Y-1; i++ {
		t.cells[i] = t.cells[i+1]
	}
	t.cells[i] = make([]cell, t.size.X, t.size.X)
}

func (t *Term) wrapCursorForwards() {
	t.curPos.X += 1

	if t.curPos.X >= t.size.X {
		overflow := t.curPos.X - (t.size.X - 1)
		t.curPos.X = 0

		if overflow > 0 && t.moveCursorDownwards(1) > 0 {
			t.moveContentsUp()
			t.moveCursorDownwards(1)
		}
	}

	return
}

func (t *Term) writeCell(r rune) {
	t.Cell().char = r
	t.Cell().sgr = t.sgr
	t.wrapCursorForwards()
}

// export contents

func (t *Term) Export() string {
	gridLen := (t.size.X + 1) * t.size.Y
	r := make([]rune, gridLen, gridLen)
	var i int
	for y := range t.cells {
		for x := range t.cells[y] {
			if t.cells[y][x].char != 0 { // if cell contains no data then lets assume it's a space character
				r[i] = t.cells[y][x].char
			} else {
				r[i] = ' '
			}
			i++
		}
		r[i] = '\n'
		i++
	}

	return string(r)
}

func (t *Term) ExportHtml() string {
	s := `<span class="">`

	last := &sgr{}

	for y := range t.cells {
		for x := range t.cells[y] {

			if t.Cell().sgr.Differs(last) {
				s += `</span><span class="` + sgrHtmlClassLookup(&t.Cell().sgr) + `">`
			}

			if t.cells[y][x].char != 0 { // if cell contains no data then lets assume it's a space character
				s += html.EscapeString(string(t.cells[y][x].char))
			} else {
				s += " "
			}

			last = &t.Cell().sgr
		}
		s += "\n"
	}

	return s + "</span>"
}

func sgrHtmlClassLookup(sgr *sgr) string {
	classes := make([]string, 0)

	for bit, class := range sgrHtmlClassNames {
		if sgr.CheckFlag(bit) {
			classes = append(classes, class)
		}
	}

	return strings.Join(classes, " ")
}

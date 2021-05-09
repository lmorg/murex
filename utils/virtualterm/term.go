package virtualterm

type Term struct {
	cells   [][]cell
	size    xy
	curPos  xy
	sgr     sgr
	LfIncCr bool // if false, \n acts as a \r\n
}

type cell struct {
	char rune
	sgr  sgr
}

type sgr struct {
	bitwise sgrFlag
	fg      rgb
	bg      rgb
}

func (c *cell) differs(oldChar rune, oldSgr *sgr) bool {
	if c.sgr.bitwise != oldSgr.bitwise {
		return true
	}

	if c.char == 0 && oldChar != 0 {
		return true
	}

	if c.sgr.fg.Red != oldSgr.fg.Red ||
		c.sgr.fg.Green != oldSgr.fg.Green ||
		c.sgr.fg.Blue != oldSgr.fg.Blue {
		return true
	}

	if c.sgr.bg.Red != oldSgr.bg.Red ||
		c.sgr.bg.Green != oldSgr.bg.Green ||
		c.sgr.bg.Blue != oldSgr.bg.Blue {
		return true
	}

	return false
}

func (sgr *sgr) checkFlag(flag sgrFlag) bool {
	return sgr.bitwise&flag != 0
}

type xy struct {
	X int
	Y int
}

type rgb struct {
	Red, Green, Blue byte
}

func NewTerminal(x, y int) *Term {
	cells := make([][]cell, y, y)
	for i := range cells {
		cells[i] = make([]cell, x, x)
	}

	return &Term{
		cells: cells,
		size:  xy{x, y},
	}
}

// format

func (t *Term) sgrReset() {
	t.sgr.bitwise = 0
	t.sgr.fg = rgb{}
	t.sgr.bg = rgb{}
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

func (t *Term) cell() *cell { return &t.cells[t.curPos.Y][t.curPos.X] }

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
	t.cell().char = r
	t.cell().sgr = t.sgr
	t.wrapCursorForwards()
}

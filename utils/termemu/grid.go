package termemu

type Grid struct {
	cells   [][]cell
	size    Xy
	curPos  Xy
	LfIncCr bool // if false, \n acts as a \r\n
}

type cell struct {
	char   rune
	format uint32
	colour Rgb
}

type Xy struct {
	X int
	Y int
}

type Rgb struct {
	Red, Green, Blue byte
}

func NewGrid(x, y int) *Grid {
	cells := make([][]cell, y, y)
	for i := range cells {
		cells[i] = make([]cell, x, x)
	}

	return &Grid{
		cells: cells,
		size:  Xy{x, y},
	}
}

// moveCursor functions DON'T effect other contents in the grid

func (g *Grid) moveCursorBackwards(i int) (overflow int) {
	g.curPos.X -= i
	if g.curPos.X < 0 {
		overflow = g.curPos.X * -1
		g.curPos.X = 0
	}

	return
}

func (g *Grid) moveCursorForwards(i int) (overflow int) {
	g.curPos.X += i
	if g.curPos.X >= g.size.X {
		overflow = g.curPos.X - (g.size.X - 1)
		g.curPos.X = g.size.X - 1
	}

	return
}

func (g *Grid) moveCursorUpwards(i int) (overflow int) {
	g.curPos.Y -= i
	if g.curPos.Y < 0 {
		overflow = g.curPos.Y * -1
		g.curPos.Y = 0
	}

	return
}

func (g *Grid) moveCursorDownwards(i int) (overflow int) {
	g.curPos.Y += i
	if g.curPos.Y >= g.size.Y {
		overflow = g.curPos.Y - (g.size.Y - 1)
		g.curPos.Y = g.size.Y - 1
	}

	return
}

func (g *Grid) Cell() *cell { return &g.cells[g.curPos.Y][g.curPos.X] }

// moveGridPos functions DO effect other contents in the grid

func (g *Grid) moveGridUp() {
	var i int
	for ; i < g.size.Y-1; i++ {
		g.cells[i] = g.cells[i+1]
	}
	g.cells[i] = make([]cell, g.size.X, g.size.X)
}

func (g *Grid) wrapCursorForwards() {
	g.curPos.X += 1

	if g.curPos.X >= g.size.X {
		overflow := g.curPos.X - (g.size.X - 1)
		g.curPos.X = 0

		if overflow > 0 && g.moveCursorDownwards(1) > 0 {
			g.moveGridUp()
			g.moveCursorDownwards(1)
		}
	}

	return
}

func (g *Grid) writeCell(r rune) {
	g.Cell().char = r
	g.wrapCursorForwards()
}

// export contents

func (g *Grid) Export() string {
	gridLen := (g.size.X + 1) * g.size.Y
	r := make([]rune, gridLen, gridLen)
	var i int
	for y := range g.cells {
		for x := range g.cells[y] {
			if g.cells[y][x].char != 0 { // if cell contains no data then lets assume it's a space character
				r[i] = g.cells[y][x].char
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

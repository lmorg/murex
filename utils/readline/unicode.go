package readline

import "github.com/mattn/go-runewidth"

type UnicodeT struct {
	rl    *Instance
	value []rune
	rPos  int
	cPos  int
}

func (u *UnicodeT) Set(rl *Instance, r []rune) {
	u.rl = rl
	u.value = r
	u.cPos = u.cellPos()
}

func (u *UnicodeT) Runes() []rune {
	return u.value
}

func (u *UnicodeT) String() string {
	return string(u.value)
}

func (u *UnicodeT) RuneLen() int {
	return len(u.value)
}

func (u *UnicodeT) RunePos() int {
	return u.rPos
}

func (u *UnicodeT) _offByOne(i int) int {
	if len(u.value) == 0 {
		return 0
	}
	if u.rl != nil && u.rl.modeViMode != vimInsert && i == len(u.value) {
		i = len(u.value) - 1
	}
	return i
}

func (u *UnicodeT) SetRunePos(i int) {
	if i < 0 {
		i = 0
	}
	if i > len(u.value) {
		i = len(u.value)
	}

	u.rPos = u._offByOne(i)
	u.cPos = u.cellPos()
}

func (u *UnicodeT) Duplicate() *UnicodeT {
	dup := new(UnicodeT)
	dup.value = make([]rune, len(u.value))
	copy(dup.value, u.value)
	dup.rPos = u.rPos
	dup.cPos = u.cPos
	return dup
}

func (u *UnicodeT) CellLen() int {
	return runewidth.StringWidth(u.String())
}

func (u *UnicodeT) cellPos() int {
	var cPos, i, last int
	for ; i < len(u.value) && i < u.rPos; i++ {
		w := runewidth.RuneWidth(u.value[i])
		cPos += w
		last = w
	}
	if last == 2 {
		cPos--
	}

	return cPos
}

func (u *UnicodeT) CellPos() int {
	return u.cPos
}

func (u *UnicodeT) SetCellPos(cPos int) {
	u._setCellPos(cPos)
	i := u._offByOne(u.rPos)
	if i != u.rPos {
		u.rPos--
		u.cPos -= runewidth.RuneWidth(u.value[u.rPos])
	}
}

func (u *UnicodeT) _setCellPos(cPos int) {
	if len(u.value) == 0 {
		return
	}

	u.cPos = 0
	var last int
	for u.rPos = 0; u.rPos < len(u.value); u.rPos++ {
		if u.cPos >= cPos {
			if last == 2 {
				u.cPos--
			}
			return
		}
		w := runewidth.RuneWidth(u.value[u.rPos])
		u.cPos += w
		last = w
	}

	if last == 2 {
		u.cPos--
	}
	u.rPos = len(u.value)
	if u.rPos < 0 {
		u.rPos = 0
	}
}

package virtualterm

import (
	"html"
	"strings"
)

// Export returns a character map of the virtual terminal
func (t *Term) Export() string {
	t.mutex.Lock()

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

	t.mutex.Unlock()

	return string(r)
}

// ExportHTML returns a HTML reder of the virtual terminal
func (t *Term) ExportHtml() string {
	s := `<span class="">`

	lastSgr := &sgr{}
	var lastChar rune = 0

	t.mutex.Lock()

	for y := range t.cells {
		for x := range t.cells[y] {
			sgr := &t.cells[y][x].sgr
			char := t.cells[y][x].char

			if t.cells[y][x].differs(lastChar, lastSgr) {
				s += `</span><span class="` + sgrHtmlClassLookup(sgr) + `">`
			}

			if char != 0 { // if cell contains no data then lets assume it's a space character
				s += html.EscapeString(string(char))
			} else {
				s += " "
			}

			lastSgr = sgr
			lastChar = char
		}
		s += "\n"
	}

	t.mutex.Unlock()

	return s + "</span>"
}

func sgrHtmlClassLookup(sgr *sgr) string {
	classes := make([]string, 0)

	for bit, class := range sgrHtmlClassNames {
		if sgr.checkFlag(bit) {
			classes = append(classes, class)
		}
	}

	if sgr.checkFlag(sgrFgColour4) {
		classes = append(classes, sgrColourHtmlClassNames[sgr.fg.Red])
	}

	return strings.Join(classes, " ")
}

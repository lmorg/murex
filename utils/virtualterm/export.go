package virtualterm

import (
	"html"
	"sort"
	"strings"
)

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

	lastSgr := &sgr{}
	var lastChar rune = 0

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

	// It doesn't actually matter if the classes are sorted. However it does
	// make it harder to write tests when classes can be returned in any order.
	// So while this adds a bit of overhead to the function, it does ensure
	// tests pass consistently. The prod version of this code might see this
	// this removed with regexp matches in the tests rather than string == and
	// have this sort() function removed.
	sort.Strings(classes)

	return strings.Join(classes, " ")
}

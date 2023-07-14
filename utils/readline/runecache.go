package readline

import (
	"github.com/mattn/go-runewidth"
)

//var _runeWidthTruncate = make(map[string]string)

func runeWidthTruncate(s string, maxLength int) string {
	/*key := fmt.Sprintf("%s:%d", s, maxLength)

	r, ok := _runeWidthTruncate[key]
	if ok {
		return r
	}

	r = runewidth.Truncate(s, maxLength, "…")
	_runeWidthTruncate[key] = r
	return r*/

	return runewidth.Truncate(s, maxLength, "…")
}

//var _runeWidthFillRight = make(map[string]string)

func runeWidthFillRight(s string, maxLength int) string {
	/*key := fmt.Sprintf("%s:%d", s, maxLength)

	r, ok := _runeWidthFillRight[key]
	if ok {
		return r
	}

	r = runewidth.FillRight(s, maxLength)
	_runeWidthFillRight[key] = r
	return r*/

	return runewidth.FillRight(s, maxLength)
}

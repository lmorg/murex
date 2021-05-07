package virtualterm

type sgrFlag uint32

const (
	sgrReset sgrFlag = 0

	sgrBold sgrFlag = 1 << iota
	sgrItalic
	sgrUnderscore
	sgrBlink

	// colour bit pallets
	sgrFgColour4
	sgrFgColour8
	sgrFgColour24

	sgrBgColour4
	sgrBgColour8
	sgrBgColour24
)

var sgrHtmlClassNames = map[sgrFlag]string{
	sgrBold:       "sgr-bold",
	sgrItalic:     "sgr-italic",
	sgrUnderscore: "sgr-underscore",
	sgrBlink:      "sgr-blink",
}

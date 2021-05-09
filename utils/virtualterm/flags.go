package virtualterm

type sgrFlag uint32

// Flags
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

const (
	sgrColour4Black = 0
	sgrColour4Red   = iota
	sgrColour4Green
	sgrColour4Yellow
	sgrColour4Blue
	sgrColour4Magenta
	sgrColour4Cyan
	sgrColour4White

	sgrColour4BlackBright
	sgrColour4RedBright
	sgrColour4GreenBright
	sgrColour4YellowBright
	sgrColour4BlueBright
	sgrColour4MagentaBright
	sgrColour4CyanBright
	sgrColour4WhiteBright
)

var sgrColourHtmlClassNames = []string{
	"sgr-black",
	"sgr-red",
	"sgr-green",
	"sgr-yellow",
	"sgr-blue",
	"sgr-magenta",
	"sgr-cyan",
	"sgr-white",

	"sgr-black-bright",
	"sgr-red-bright",
	"sgr-green-bright",
	"sgr-yellow-bright",
	"sgr-blue-bright",
	"sgr-magenta-bright",
	"sgr-cyan-bright",
	"sgr-white-bright",
}

package codes

// Text effects
const (
	Reset      = "\x1b[0m"
	Bold       = "\x1b[1m"
	Underscore = "\x1b[4m"
	Blink      = "\x1b[5m"
	Invert     = "\x1b[7m"

	NormalIntensity = "\x1b[22m"
	NoUnderscore    = "\x1b[24m"
	NoBlink         = "\x1b[25m"
	NoInvert        = "\x1b[27m"
)

// Text colours
const (
	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"

	FgBlackBright   = "\x1b[90m"
	FgRedBright     = "\x1b[91m"
	FgGreenBright   = "\x1b[92m"
	FgYellowBright  = "\x1b[93m"
	FgBlueBright    = "\x1b[94m"
	FgMagentaBright = "\x1b[95m"
	FgCyanBright    = "\x1b[96m"
	FgWhiteBright   = "\x1b[97m"
)

// Background colours
const (
	BgBlack   = "\x1b[40m"
	BgRed     = "\x1b[41m"
	BgGreen   = "\x1b[42m"
	BgYellow  = "\x1b[43m"
	BgBlue    = "\x1b[44m"
	BgMagenta = "\x1b[45m"
	BgCyan    = "\x1b[46m"
	BgWhite   = "\x1b[47m"

	BgBlackBright   = "\x1b[100m"
	BgRedBright     = "\x1b[101m"
	BgGreenBright   = "\x1b[102m"
	BgYellowBright  = "\x1b[103m"
	BgBlueBright    = "\x1b[104m"
	BgMagentaBright = "\x1b[105m"
	BgCyanBright    = "\x1b[106m"
	BgWhiteBright   = "\x1b[107m"
)

// Clear terminal
const (
	ClearScreen      = "\x1b[2J"
	ClearLine        = "\x1b[K"
	ClearScreenBelow = "\x1b[J"
)

// Cursor Movement
const (
	Home      = "\x1b[H"
	Up        = "\x1b[1A"
	Down      = "\x1b[1B"
	Forwards  = "\x1b[1C"
	Backwards = "\x1b[1D"
)

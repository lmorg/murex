package readline

// Character codes
const (
	charCtrlA = iota + 1
	charCtrlB
	charCtrlC
	charEOF
	charCtrlE
	charCtrlF
	charCtrlG
	charBackspace // ISO 646
	charTab
	charCtrlJ
	charCtrlK
	charCtrlL
	charCtrlM
	charCtrlN
	charCtrlO
	charCtrlP
	charCtrlQ
	charCtrlR
	charCtrlS
	charCtrlT
	charCtrlU
	charCtrlV
	charCtrlW
	charCtrlX
	charCtrlY
	charCtrlZ
	charEscape
	charCtrlSlash             // ^\
	charCtrlCloseSquare       // ^]
	charCtrlHat               // ^^
	charCtrlUnderscore        // ^_
	charBackspace2      = 127 // ASCII 1963
)

// Escape sequences
var (
	seqUp        = string([]byte{27, 91, 65})
	seqDown      = string([]byte{27, 91, 66})
	seqForwards  = string([]byte{27, 91, 67})
	seqBackwards = string([]byte{27, 91, 68})
	seqHome      = string([]byte{27, 91, 72})
	seqHomeSc    = string([]byte{27, 91, 49, 126})
	seqEnd       = string([]byte{27, 91, 70})
	seqEndSc     = string([]byte{27, 91, 52, 126})
	seqDelete    = string([]byte{27, 91, 51, 126})
	seqShiftTab  = string([]byte{27, 91, 90})
	seqPageUp    = string([]byte{27, 91, 53, 126})
	seqPageDown  = string([]byte{27, 91, 54, 126})

	seqF1VT100 = string([]byte{27, 79, 80})
	//seqF2VT100 = string([]byte{27, 79, 81})
	//seqF3VT100 = string([]byte{27, 79, 82})
	//seqF4VT100 = string([]byte{27, 79, 83})
	seqF1 = string([]byte{27, 91, 49, 49, 126})
	//seqF2      = string([]byte{27, 91, 49, 50, 126})
	//seqF3      = string([]byte{27, 91, 49, 51, 126})
	//seqF4      = string([]byte{27, 91, 49, 52, 126})
	//seqF5      = string([]byte{27, 91, 49, 53, 126})
	//seqF6      = string([]byte{27, 91, 49, 55, 126})
	//seqF7      = string([]byte{27, 91, 49, 56, 126})
	//seqF8      = string([]byte{27, 91, 49, 57, 126})
	//seqF9      = string([]byte{27, 91, 50, 48, 126})
	//seqF10     = string([]byte{27, 91, 50, 49, 126})
	//seqF11     = string([]byte{27, 91, 50, 51, 126})
	//seqF12     = string([]byte{27, 91, 50, 52, 126})
)

const (
	//seqPosSave    = "\x1b[s"
	//seqPosRestore = "\x1b[u"

	seqClearLineAfter   = "\x1b[0k"
	seqClearLineBefore  = "\x1b[1k"
	seqClearLine        = "\x1b[2k"
	seqClearScreenBelow = "\x1b[J"
	seqClearScreen      = "\x1b[2J"

	seqGetCursorPos = "\x1b6n" // response: "\x1b{Line};{Column}R"

	seqSetCursorPosTopLeft = "\x1b[1;1H"
)

// Text effects
const (
	seqReset      = "\x1b[0m"
	seqBold       = "\x1b[1m"
	seqUnderscore = "\x1b[4m"
	seqBlink      = "\x1b[5m"
)

// Text colours
const (
	seqFgBlack   = "\x1b[30m"
	seqFgRed     = "\x1b[31m"
	seqFgGreen   = "\x1b[32m"
	seqFgYellow  = "\x1b[33m"
	seqFgBlue    = "\x1b[34m"
	seqFgMagenta = "\x1b[35m"
	seqFgCyan    = "\x1b[36m"
	seqFgWhite   = "\x1b[37m"

	/*seqFgBlackBright   = "\x1b[1;30m"
	seqFgRedBright     = "\x1b[1;31m"
	seqFgGreenBright   = "\x1b[1;32m"
	seqFgYellowBright  = "\x1b[1;33m"
	seqFgBlueBright    = "\x1b[1;34m"
	seqFgMagentaBright = "\x1b[1;35m"
	seqFgCyanBright    = "\x1b[1;36m"
	seqFgWhiteBright   = "\x1b[1;37m"*/
)

// Background colours
const (
	seqBgBlack   = "\x1b[40m"
	seqBgRed     = "\x1b[41m"
	seqBgGreen   = "\x1b[42m"
	seqBgYellow  = "\x1b[43m"
	seqBgBlue    = "\x1b[44m"
	seqBgMagenta = "\x1b[45m"
	seqBgCyan    = "\x1b[46m"
	seqBgWhite   = "\x1b[47m"

	/*seqBgBlackBright   = "\x1b[1;40m"
	seqBgRedBright     = "\x1b[1;41m"
	seqBgGreenBright   = "\x1b[1;42m"
	seqBgYellowBright  = "\x1b[1;43m"
	seqBgBlueBright    = "\x1b[1;44m"
	seqBgMagentaBright = "\x1b[1;45m"
	seqBgCyanBright    = "\x1b[1;46m"
	seqBgWhiteBright   = "\x1b[1;47m"*/
)

const (
	seqEscape = "\x1b"

	// generated using:
	// a [a..z] -> foreach c { -> tr [:lower:] [:upper:] -> set C; out (seqAlt$C = "\x1b$c") }
	seqAltA = "\x1ba"
	seqAltB = "\x1bb"
	seqAltC = "\x1bc"
	seqAltD = "\x1bd"
	seqAltE = "\x1be"
	seqAltF = "\x1bf"
	seqAltG = "\x1bg"
	seqAltH = "\x1bh"
	seqAltI = "\x1bi"
	seqAltJ = "\x1bj"
	seqAltK = "\x1bk"
	seqAltL = "\x1bl"
	seqAltM = "\x1bm"
	seqAltN = "\x1bn"
	seqAltO = "\x1bo"
	seqAltP = "\x1bp"
	seqAltQ = "\x1bq"
	seqAltR = "\x1br"
	seqAltS = "\x1bs"
	seqAltT = "\x1bt"
	seqAltU = "\x1bu"
	seqAltV = "\x1bv"
	seqAltW = "\x1bw"
	seqAltX = "\x1bx"
	seqAltY = "\x1by"
	seqAltZ = "\x1bz"

	seqAltShiftA = "\x1bA"
	seqAltShiftB = "\x1bB"
	seqAltShiftC = "\x1bC"
	seqAltShiftD = "\x1bD"
	seqAltShiftE = "\x1bE"
	seqAltShiftF = "\x1bF"
	seqAltShiftG = "\x1bG"
	seqAltShiftH = "\x1bH"
	seqAltShiftI = "\x1bI"
	seqAltShiftJ = "\x1bJ"
	seqAltShiftK = "\x1bK"
	seqAltShiftL = "\x1bL"
	seqAltShiftM = "\x1bM"
	seqAltShiftN = "\x1bN"
	seqAltShiftO = "\x1bO"
	seqAltShiftP = "\x1bP"
	seqAltShiftQ = "\x1bQ"
	seqAltShiftR = "\x1bR"
	seqAltShiftS = "\x1bS"
	seqAltShiftT = "\x1bT"
	seqAltShiftU = "\x1bU"
	seqAltShiftV = "\x1bV"
	seqAltShiftW = "\x1bW"
	seqAltShiftX = "\x1bX"
	seqAltShiftY = "\x1bY"
	seqAltShiftZ = "\x1bZ"
)

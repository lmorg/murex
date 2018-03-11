package readline

// Character codes
const (
	charCtrlC     = 3
	charCtrlU     = 21
	charEOF       = 4
	charBackspace = 127
	charEscape    = 27
)

// Escape sequences
var (
	seqUp        = string([]byte{27, 91, 65})
	seqDown      = string([]byte{27, 91, 66})
	seqForwards  = string([]byte{27, 91, 67})
	seqBackwards = string([]byte{27, 91, 68})
	seqDelete    = string([]byte{27, 91, 51, 126})
)

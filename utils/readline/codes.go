package readline

const (
	CtrlC     = 3
	CtrlU     = 21
	EOF       = 4
	Backspace = 127
	Escape    = 27
)

var (
	up        = string([]byte{27, 91, 65})
	down      = string([]byte{27, 91, 66})
	forwards  = string([]byte{27, 91, 67})
	backwards = string([]byte{27, 91, 68})
)

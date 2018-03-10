package readline

const (
	CtrlC     = 3
	CtrlU     = 21
	EOF       = 4
	Backspace = 127
	Escape    = 27
)

var (
	seqUp        = string([]byte{27, 91, 65})
	seqDown      = string([]byte{27, 91, 66})
	seqForwards  = string([]byte{27, 91, 67})
	seqBackwards = string([]byte{27, 91, 68})
	seqDelete    = string([]byte{27, 91, 51, 126})
)

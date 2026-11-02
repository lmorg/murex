package terminal

import (
	"os"
)

type TermT int

const (
	TERM_TTYPHOON = 1
	TERM_VSCODE   = 2
	TERM_OTHER    = 0
)

func IsTtyphoon() bool {
	return os.Getenv("MXTTY") == "true"
}

func IsVSCode() bool {
	return os.Getenv("TERM_PROGRAM") == "vscode"
}

func Application() TermT {
	switch {
	case IsTtyphoon():
		return TERM_TTYPHOON
	case IsVSCode():
		return TERM_VSCODE
	default:
		return 0
	}
}

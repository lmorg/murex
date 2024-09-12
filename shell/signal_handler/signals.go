package signalhandler

import (
	"os"
)

var signalChan chan os.Signal = make(chan os.Signal, 1)

const (
	// PromptEOF defines the string to write when ctrl+d is pressed
	PromptEOF = "^D"

	// PromptSIGTSTP defines the string to write when ctrl+z is pressed
	PromptSIGTSTP = "^Z"

	// PromptSIGINT defines the string to write when ctrl+c is pressed
	PromptSIGINT = "^C"

	// PromptSIGQUIT defines the string to write when ctrl+\ is pressed
	PromptSIGQUIT = "^\\"
)

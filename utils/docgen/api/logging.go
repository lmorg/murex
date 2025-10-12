package docgen

import (
	"fmt"
	"io"
	golog "log"
)

var (
	// Verbose enables verbose logging ([LOG] and [WARNING] messages)
	Verbose bool

	// Warning enables [WARNING] messages
	Warning bool

	// ExitStatus is what the recommended exit status should be
	ExitStatus int
)

// SetLogger allows output messages to be redirected
func SetLogger(w io.Writer) {
	golog.SetOutput(w)
}

func log(v ...any) {
	if Verbose {
		golog.Println(append([]any{"[LOG]"}, v...)...)
	}
}

func warning(file string, v ...any) {
	if Warning || Verbose {
		warning := fmt.Sprintf("[WARNING] %s:", file)
		golog.Println(append([]any{warning}, v...)...)
	}
}

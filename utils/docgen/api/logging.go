package docgen

import (
	"io"
	golog "log"
)

var (
	// Verbose enables verbose logging ([LOG] and [WARNING] messages)
	Verbose bool

	// Warning enables [WARNING] messages
	Warning bool
)

// SetLogger allows output messages to be redirected
func SetLogger(w io.Writer) {
	golog.SetOutput(w)
}

func log(v ...interface{}) {
	if Verbose {
		golog.Println(append([]interface{}{"[LOG]"}, v...)...)
	}
}

func warning(v ...interface{}) {
	if Warning || Verbose {
		golog.Println(append([]interface{}{"[WARNING]"}, v...)...)
	}
}

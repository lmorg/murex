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

func log(v ...interface{}) {
	if Verbose {
		golog.Println(append([]interface{}{"[LOG]"}, v...)...)
	}
}

func warning(file string, v ...interface{}) {
	if Warning || Verbose {
		/*ExitStatus++
		if ExitStatus > 254 {
			ExitStatus = 254
		}*/

		warning := fmt.Sprintf("[WARNING] %s:", file)
		golog.Println(append([]interface{}{warning}, v...)...)
	}
}

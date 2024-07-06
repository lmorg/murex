package crash

import (
	"fmt"
	"os"
	"runtime"

	"github.com/lmorg/murex/debug"
)

var disable_handler bool

func Handler() bool {
	if debug.Enabled || disable_handler {
		return false
	}

	r := recover()
	if r == nil {
		return false
	}

	_, _ = os.Stderr.WriteString(fmt.Sprintf("Error: %v\n", r))
	_, _ = os.Stderr.WriteString(_crashStack())
	_, _ = os.Stderr.WriteString(_crashHostReport())
	_, _ = os.Stderr.WriteString(crashMessage)

	return true
}

var crashMessage = `
---

!!! Murex has crashed !!!

Above is a crash report that can be shared in https://github.com/lmorg/murex/issues

Your session state, including stored variables will be retained. However you may need to ctrl+\ to continue
`

func _crashStack() string {
	pc := make([]uintptr, 5)
	l := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:l])

	var (
		frame runtime.Frame
		more  = true
		msg   = "Stack:\n"
	)

	for more {
		frame, more = frames.Next()
		msg += fmt.Sprintf(crashStackMessage, frame.Function, frame.File, frame.Line)
	}

	return msg
}

var crashStackMessage = `
  - function: %s(...)
    filename: %s:%d
`

func _crashHostReport() string {
	build := fmt.Sprintf("\n  OS:  %s\n  CPU: %s", runtime.GOOS, runtime.GOARCH)
	compiler := runtime.Version()

	return fmt.Sprintf("Build: %s\nCompiler: %s", build, compiler)
}

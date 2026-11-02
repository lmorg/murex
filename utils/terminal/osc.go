package terminal

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lmorg/murex/debug"
)

func emitOSC(code int, fields ...string) {
	writeOSC(os.Stdout, code, fields...)
}

func writeOSC(writer io.Writer, code int, fields ...string) {
	if _, err := fmt.Fprintf(writer, "\x1b]%d;%s\x1b\\", code, strings.Join(fields, ";")); err != nil {
		debug.Log("term: writing OSC:", err)
	}
}

// +build js

package term

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Error - Coloured Red

// ErrRed is the Stderr interface for term - with output coloured red
type ErrRed struct {
	term
}

const (
	fgRed = "\x1b[31m"
	reset = "\x1b[0m"
)

// Write is the io.Writer() interface for term
func (t *ErrRed) Write(b []byte) (int, error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	t.mutex.Unlock()

	new := fgRed + string(b) + reset

	vtermWrite([]rune(new))

	return len(b), nil
}

// Writeln writes an OS-specific terminated line to the stderr
func (t *ErrRed) Writeln(b []byte) (int, error) {
	//line := append(b, utils.NewLineByte...)
	//return t.Write(line)
	return t.Write(appendBytes(b, utils.NewLineByte...))
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *ErrRed) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(t, dataType)
}

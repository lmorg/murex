// +build js

package term

import (
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Out

// Out is the Stdout interface for term
type Out struct {
	term
}

// Write is the io.Writer() interface for term
func (t *Out) Write(b []byte) (int, error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	t.mutex.Unlock()

	vtermWrite([]rune(string(b)))

	return len(b), nil
}

// Writeln writes an OS-specific terminated line to the stdout
func (t *Out) Writeln(b []byte) (int, error) {
	return t.Write(appendBytes(b, utils.NewLineByte...))
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *Out) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(t, dataType)
}

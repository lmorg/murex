//go:build !js
// +build !js

package term

import (
	"os"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi/codes"
)

// Terminal: Standard Error - Coloured Red

// ErrRed is the Stderr interface for term - with output coloured red
type ErrRed struct {
	term
}

// SetDataType is a null method because the term interface is write-only
func (t *ErrRed) SetDataType(string) {}

// Write is the io.Writer() interface for term
func (t *ErrRed) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	t.mutex.Unlock()

	i, err = os.Stderr.WriteString(codes.FgRed + string(b) + codes.Reset)
	if err != nil {
		os.Stdout.WriteString(codes.FgRed + err.Error() + codes.Reset)
	}

	return i - 9, err
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

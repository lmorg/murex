package term

import (
	"os"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Error - Coloured Red

// ErrRed is the Stderr interface for term - with output coloured red
type ErrRed struct {
	term
}

const (
	reset = "\x1b[0m"
	fgRed = "\x1b[31m"
)

// Write is the io.Writer() interface for term
func (t *ErrRed) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	i, err = os.Stderr.WriteString(fgRed + string(b) + reset)
	if err != nil {
		os.Stdout.WriteString(fgRed + err.Error() + reset)
	}
	t.mutex.Unlock()
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

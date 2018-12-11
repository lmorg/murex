package term

import (
	"os"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Error

// Err is the Stderr interface for term
type Err struct {
	term
}

// Write is the io.Writer() interface for term
func (t *Err) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	i, err = os.Stderr.Write(b)
	if err != nil {
		os.Stdout.WriteString(err.Error())
	} else if len(b) > 0 {
		CrLf.set(b[len(b)-1])
	}
	t.mutex.Unlock()
	return
}

// Writeln writes an OS-specific terminated line to the stderr
func (t *Err) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	return t.Write(line)
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *Err) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(t, dataType)
}

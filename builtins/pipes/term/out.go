package term

import (
	"os"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Out

// Out is the Stdout interface for term
type Out struct {
	term
}

// Write is the io.Writer() interface for term
func (t *Out) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	t.mutex.Unlock()

	i, err = os.Stdout.Write(b)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	} else if len(b) > 0 {
		CrLf.set(b[len(b)-1])
	}

	return
}

// Writeln writes an OS-specific terminated line to the stdout
func (t *Out) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	return t.Write(line)
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *Out) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(t, dataType)
}

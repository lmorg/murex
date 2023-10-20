package term

import (
	"errors"
	"os"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/tty"
)

// Terminal: Standard In

// In is the Stdin interface for term
type In struct {
	streams.Stdin
}

func NewIn(dataType string) stdio.Io {
	t := new(In)
	t.Stdin = *streams.NewStdin()
	t.SetDataType(dataType)
	go backgroundRead(t)
	return t
}

func backgroundRead(t *In) {
	_, err := t.ReadFrom(tty.Stdin)
	if err != nil {
		tty.Stderr.WriteString("Error reading from STDIN: " + err.Error())
	}
}

func (t *In) File() *os.File {
	return tty.Stdin
}

// Write is the io.Writer() interface for term
func (t *In) Write(_ []byte) (int, error) {
	return 0, errors.New("attempting to write to a readonly STDIN interface: Write()")
}

// Writeln writes an OS-specific terminated line to the stdout
func (t *In) Writeln(_ []byte) (int, error) {
	return 0, errors.New("attempting to write to a readonly STDIN interface: Writeln()")
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *In) WriteArray(_ string) (stdio.ArrayWriter, error) {
	return nil, errors.New("attempting to write to a readonly STDIN interface: WriteArray()")
}

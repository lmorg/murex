//go:build !js
// +build !js

package term

import (
	"os"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Out

// Out is the Stdout interface for term
type Out struct {
	term
}

func OutSetDataTypeFd3() {
	s, _ := os.LookupEnv("MUREX_EXEC")
	if s != "yes" {
		return
	}

	outSetDataTypeFd3 = true
}

var outSetDataTypeFd3 bool

// SetDataType writes the data type to a special pipe when run under murex
func (t *Out) SetDataType(dt string) {
	if !outSetDataTypeFd3 || len(dt) == 0 || dt == types.Null {
		return
	}

	f := os.NewFile(3, "dt")
	_, err := f.WriteString(dt + "\n")
	if err != nil {
		os.Stderr.WriteString("Error writing data type: " + err.Error() + "\n")
	}

	//f.Close()
}

// Write is the io.Writer() interface for term
func (t *Out) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	t.mutex.Unlock()

	i, err = os.Stdout.Write(b)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	return
}

// Writeln writes an OS-specific terminated line to the stdout
func (t *Out) Writeln(b []byte) (int, error) {
	return t.Write(appendBytes(b, utils.NewLineByte...))
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *Out) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(t, dataType)
}

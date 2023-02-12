//go:build !js
// +build !js

package term

import (
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/tty"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Out

// Out is the Stdout interface for term
type Out struct {
	term
}

func OutSetDataTypeIPC() {
	/*murexPid, exists := os.LookupEnv(consts.EnvMurexPid)

	if !exists {
		return
	}

	if strconv.Itoa(os.Getppid()) != murexPid {
		return
	}

	outSetDataTypeFd3 = true*/
}

//var OutSetDataTypeIPC bool

// SetDataType writes the data type to a special pipe when run under murex
func (t *Out) SetDataType(dt string) {
	/*if !OutSetDataTypeIPC || len(dt) == 0 || dt == types.Null {
		return
	}

	f := os.NewFile(3, "dt")
	_, err := f.WriteString(dt + "\n")
	if err != nil && debug.Enabled {
		tty.Stderr.WriteString("Error writing data type: " + err.Error() + "\n")
	}

	OutSetDataTypeIPC = false
	f.Close()*/
}

// Write is the io.Writer() interface for term
func (t *Out) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	t.mutex.Unlock()

	i, err = tty.Stdout.Write(b)
	if err != nil {
		tty.Stderr.WriteString(err.Error())
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

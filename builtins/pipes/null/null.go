package null

import (
	"io"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

// Since I don't want you to create null pipes, lets not register it
/*func init() {
	stdio.RegisterPipe("null", func(string) (stdio.Io, error) {
		return nil, errors.New("null pipes cannot be created. Use `null` if you require a null pipe")
	})
}*/

// Null is null interface for named pipes
type Null struct{}

// Read - null interface
func (t *Null) Read([]byte) (int, error) { return 0, nil }

// ReadLine - null interface
func (t *Null) ReadLine(func([]byte)) error { return nil }

// ReadArray - null interface
func (t *Null) ReadArray(func([]byte)) error { return nil }

// ReadArrayWithType - null interface
func (t *Null) ReadArrayWithType(func([]byte, string)) error { return nil }

// ReadMap - null interface
func (t *Null) ReadMap(*config.Config, func(string, string, bool)) error { return nil }

// ReadAll - null interface
func (t *Null) ReadAll() ([]byte, error) { return []byte{}, nil }

// WriteTo - null interface
func (t *Null) WriteTo(io.Writer) (int64, error) { return 0, nil }

// Write - null interface
func (t *Null) Write(b []byte) (int, error) { return len(b), nil }

// Writeln - null interface
func (t *Null) Writeln(b []byte) (int, error) { return len(b), nil }

// WriteArray - null interface
func (t *Null) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(t, dataType)
}

// Stats - null interface
func (t *Null) Stats() (uint64, uint64) { return 0, 0 }

// GetDataType - null interface
func (t *Null) GetDataType() string { return types.Null }

// SetDataType - null interface
func (t *Null) SetDataType(string) {}

// DefaultDataType - null interface
func (t *Null) DefaultDataType(bool) {}

// IsTTY - null interface
func (t *Null) IsTTY() bool { return false }

// Open - null interface
func (t *Null) Open() {}

// Close - null interface
func (t *Null) Close() {}

// ForceClose - null interface
func (t *Null) ForceClose() {}

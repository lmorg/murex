package streams

import (
	"io"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
)

type Tee struct {
	primary   stdio.Io
	secondary Stdin
}

func NewTee(primary stdio.Io) (tee *Tee, secondary *Stdin) {
	tee = new(Tee)
	tee.primary = primary
	tee.secondary.max = 0
	//tee.secondary.Open()
	secondary = &tee.secondary
	return
}

func (tee *Tee) IsTTY() bool                           { return tee.primary.IsTTY() }
func (tee *Tee) MakePipe()                             { tee.primary.MakePipe() }
func (tee *Tee) Stats() (uint64, uint64)               { return tee.primary.Stats() }
func (tee *Tee) Read(p []byte) (int, error)            { return tee.primary.Read(p) }
func (tee *Tee) ReadLine(callback func([]byte)) error  { return tee.primary.ReadLine(callback) }
func (tee *Tee) ReadArray(callback func([]byte)) error { return tee.primary.ReadArray(callback) }
func (tee *Tee) ReadAll() ([]byte, error)              { return tee.primary.ReadAll() }

func (tee *Tee) ReadMap(config *config.Config, callback func(string, string, bool)) error {
	return tee.primary.ReadMap(config, callback)
}

// Write is the standard Writer interface Write() method.
func (tee *Tee) Write(p []byte) (int, error) {
	tee.secondary.Write(p)
	return tee.primary.Write(p)
}

// Writeln just calls Write() but with an appended, OS specific, new line.
func (tee *Tee) Writeln(p []byte) (int, error) {
	tee.secondary.Writeln(p)
	return tee.primary.Writeln(p)
}

// Open the stream.Io interface for another dependant
func (tee *Tee) Open() {
	tee.primary.Open()
}

// Close the stream.Io interface
func (tee *Tee) Close() {
	tee.primary.Close()
}

// WriteTo reads from the stream.Io interface and writes to a destination
// io.Writer interface
func (tee *Tee) WriteTo(w io.Writer) (n int64, err error) {
	return tee.primary.WriteTo(w)
}

// GetDataType returns the murex data type for the stream.Io interface
func (tee *Tee) GetDataType() (dt string) {
	return tee.primary.GetDataType()
}

// SetDataType defines the murex data type for the stream.Io interface
func (tee *Tee) SetDataType(dt string) {
	tee.secondary.SetDataType(dt)
	tee.primary.SetDataType(dt)
}

// DefaultDataType defines the murex data type for the stream.Io interface if
// it's not already set
func (tee *Tee) DefaultDataType(err bool) {
	tee.secondary.DefaultDataType(err)
	tee.primary.DefaultDataType(err)
}

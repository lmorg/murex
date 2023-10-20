package psuedotty

import (
	"os"
	"sync/atomic"
)

// GetDataType returns the murex data type for the stream.Io interface
func (p *PTY) GetDataType() (dt string) { return p.out.GetDataType() }

// SetDataType defines the murex data type for the stream.Io interface
func (p *PTY) SetDataType(dt string) { p.out.SetDataType(dt) }

// Stats provides real time stream stats. Useful for progress bars etc.
func (p *PTY) Stats() (uint64, uint64) { return p.out.Stats() }

// IsTTY returns true because the PTY stream is a pseudo-TTY
func (p *PTY) IsTTY() bool { return true }

// File returns the os.File struct for the stream.Io interface if a TTY
func (p *PTY) File() *os.File { return p.in }

// Open the stream.Io interface for another dependant
func (p *PTY) Open() {
	atomic.AddInt32(&p.dependents, 1)
	//p.out.Open()
}

// Close the stream.Io interface
func (p *PTY) Close() {
	i := atomic.AddInt32(&p.dependents, -1)
	if i < 0 {
		panic("More closed dependents than open")
	}
}

// ForceClose forces the stream.Io interface to close. This should only be called by a STDIN reader
func (p *PTY) ForceClose() {
	p.in.Close()
	p.out.ForceClose()
}

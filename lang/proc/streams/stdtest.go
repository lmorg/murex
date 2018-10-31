package streams

import "context"

// Stdtest is a testing Io interface user to workaround writing to the terminal
type Stdtest struct {
	Stdin
}

// NewStdtest creates a new stream.Io interface for piping data between test
// processes that need to spoof an PTY
func NewStdtest() (stdtest *Stdtest) {
	stdtest = new(Stdtest)
	stdtest.max = DefaultMaxBufferSize
	stdtest.ctx, stdtest.forceClose = context.WithCancel(context.Background())
	return
}

// IsTTY returns true because the Stdtest stream is a pseudo-TTY mockup
func (stdtest *Stdtest) IsTTY() bool { return true }

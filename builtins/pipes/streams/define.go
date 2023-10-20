package streams

import (
	"context"
	"os"
	"sync"
	"sync/atomic"

	"github.com/lmorg/murex/lang/stdio"
)

func init() {
	stdio.RegisterPipe("std", newStream)
}

func newStream(_ string) (io stdio.Io, err error) {
	io = NewStdin()
	return
}

// Stdin is the default stdio.Io interface.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
type Stdin struct {
	mutex      sync.Mutex
	ctx        context.Context
	forceClose func()
	buffer     []byte
	bRead      uint64
	bWritten   uint64
	dependents int32
	dataType   string
	max        int
}

// DefaultMaxBufferSize is the maximum size of buffer for stdin
// var DefaultMaxBufferSize = 1024 * 1024 * 1000 // 10 meg
var DefaultMaxBufferSize = 1024 * 1024 * 1 // 1 meg

// NewStdin creates a new stream.Io interface for piping data between processes.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
	stdin.max = DefaultMaxBufferSize
	//stdin.buffer = make([]byte, 0, 1024*1024)
	stdin.ctx, stdin.forceClose = context.WithCancel(context.Background())
	return
}

// NewStdinWithContext creates a new stream.Io interface for piping data between processes.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
// This function is also useful as a context aware version of ioutil.ReadAll
func NewStdinWithContext(ctx context.Context, forceClose context.CancelFunc) (stdin *Stdin) {
	stdin = new(Stdin)
	stdin.max = DefaultMaxBufferSize
	stdin.ctx = ctx
	stdin.forceClose = forceClose
	return
}

func (stdin *Stdin) File() *os.File {
	return nil
}

// Open the stream.Io interface for another dependant
func (stdin *Stdin) Open() {
	stdin.mutex.Lock()
	atomic.AddInt32(&stdin.dependents, 1)
	stdin.mutex.Unlock()
}

// Close the stream.Io interface
func (stdin *Stdin) Close() {
	stdin.mutex.Lock()

	i := atomic.AddInt32(&stdin.dependents, -1)
	if i < 0 {
		panic("More closed dependents than open")
	}

	stdin.mutex.Unlock()
}

// ForceClose forces the stream.Io interface to close. This should only be called by a STDIN reader
func (stdin *Stdin) ForceClose() {
	if stdin.forceClose != nil {
		stdin.forceClose()
	}
}

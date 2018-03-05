// +build ingnore

package streams

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"

	"bufio"
	"bytes"
	"io"
	"sync"
)

// Stdin is the default stdio.Io interface.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
type Stdin struct {
	mutex       sync.Mutex
	buffer      *bytes.Buffer
	b           []byte
	closed      bool
	bRead       uint64
	bWritten    uint64
	isParent    bool
	isNamedPipe bool
	dataType    string
	dtLock      sync.Mutex
	max         int
}

// DefaultMaxBufferSize is the maximum size of buffer for stdin. Hwever this will
// get automatically overriden by ReadAll
var DefaultMaxBufferSize int = 1024 * 1024 * 10 // 10 meg

// NewStdin creates a new stream.Io interface for piping data between processes.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
	stdin.max = DefaultMaxBufferSize
	//stdin.b = make([]byte, stdin.max)
	stdin.buffer = bytes.NewBuffer(stdin.b)
	return
}

// IsTTY returns false because the Stdin stream is not a pseudo-TTY
func (stdin *Stdin) IsTTY() bool { return false }

// MakeParent is used for subshells so they don't accidentally close the parent stream.
func (stdin *Stdin) MakeParent() {
	stdin.mutex.Lock()
	stdin.isParent = true
	stdin.mutex.Unlock()
}

// UnmakeParent is used when the subshell has terminated, now we allow the stream to be closable again.
func (stdin *Stdin) UnmakeParent() {
	stdin.mutex.Lock()
	if !stdin.isParent {
		//debug.Log("Trying to unmake parent on a non-parent pipe.")
	}
	stdin.isParent = false

	stdin.mutex.Unlock()
}

// MakePipe is used for named pipes. Basically just used to relax the exception handling since we can make fewer
// guarantees about the state of named pipes.
func (stdin *Stdin) MakePipe() {
	stdin.mutex.Lock()
	stdin.isParent = true
	stdin.isNamedPipe = true
	stdin.mutex.Unlock()
}

// Stats provides real time stream stats. Useful for progress bars etc.
func (stdin *Stdin) Stats() (bytesWritten, bytesRead uint64) {
	stdin.mutex.Lock()
	bytesWritten = stdin.bWritten
	bytesRead = stdin.bRead
	stdin.mutex.Unlock()
	return
}

// Read is the standard Reader interface Read() method.
func (stdin *Stdin) Read(p []byte) (int, error) {
	stdin.mutex.Lock()
	defer stdin.mutex.Unlock()

	/*if stdin.buffer.Len() == 0 && stdin.closed {
		return 0, io.EOF
	}*/

	i, err := stdin.buffer.Read(p)
	stdin.bRead += uint64(i)
	return i, err
}

// ReadLine returns each line in the stream as a callback function
func (stdin *Stdin) ReadLine(callback func([]byte)) error {
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		callback(append(scanner.Bytes(), utils.NewLineByte...))
	}

	return scanner.Err()
}

// ReadAll reads everything and dump it into one byte slice.
func (stdin *Stdin) ReadAll() ([]byte, error) {
	stdin.mutex.Lock()
	stdin.max = int((^uint(0)) >> 1)
	stdin.mutex.Unlock()

	for {
		stdin.mutex.Lock()
		closed := stdin.closed
		stdin.mutex.Unlock()

		if closed {
			break
		}
	}

	stdin.mutex.Lock()
	defer stdin.mutex.Unlock()

	stdin.bRead = uint64(stdin.buffer.Len())
	return stdin.buffer.Bytes(), nil
}

// ReadArray returns a data type-specific array returned via a callback function
func (stdin *Stdin) ReadArray(callback func([]byte)) error {
	return readArray(stdin, callback)
}

// ReadMap returns a data type-specific key/values returned via a callback function
func (stdin *Stdin) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	return readMap(stdin, config, callback)
}

// Write is the standard Writer interface Write() method.
func (stdin *Stdin) Write(p []byte) (int, error) {
	/*stdin.mutex.Lock()
	isClosed := stdin.closed
	stdin.mutex.Unlock()

	if isClosed {
		return 0, io.ErrClosedPipe
	}*/

	/*for {
		stdin.mutex.Lock()
		buffSize := stdin.buffer.Len()
		maxBufferSize := stdin.max

		stdin.mutex.Unlock()

		if buffSize < maxBufferSize {
			break
		}
	}*/

	stdin.mutex.Lock()
	i, err := stdin.buffer.Write(p)
	stdin.bWritten += uint64(i)
	stdin.mutex.Unlock()

	return i, err
}

// Writeln just calls Write() but with an appended, OS specific, new line.
func (stdin *Stdin) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	stdin.Write(line)
	return len(b), nil
}

// Close the stream.Io interface
func (stdin *Stdin) Close() {
	stdin.mutex.Lock()
	defer stdin.mutex.Unlock()

	if stdin.isParent {
		// This will legitimately happen a lot since the reason we mark a stream as parent is to prevent
		// accidental closing. However it's worth pushing a message out in debug mode during this alpha build.
		//debug.Log("Cannot Close() stdin marked as parent. We don't want to EOT parent streams multiple times")
		return
	}

	if stdin.closed {
		if stdin.isNamedPipe {
			//debug.Log("Error with murex named pipes: Trying to close an already closed named pipe.")
		} else {
			//debug.Log("Trying to close an already closed stdin.")
		}
		return
	}

	stdin.closed = true
}

// WriteTo reads from the stream.Io interface and writes to a destination io.Writer interface
func (stdin *Stdin) WriteTo(dst io.Writer) (int64, error) {
	stdin.mutex.Lock()
	defer stdin.mutex.Unlock()

	i, err := stdin.buffer.WriteTo(dst)
	stdin.bRead += uint64(i)

	return i, err
}

// GetDataType returns the murex data type for the stream.Io interface
func (stdin *Stdin) GetDataType() (dt string) {
	for {
		stdin.dtLock.Lock()
		dt = stdin.dataType
		stdin.dtLock.Unlock()
		if dt != "" {
			return
		}
	}
}

// SetDataType defines the murex data type for the stream.Io interface
func (stdin *Stdin) SetDataType(dt string) {
	stdin.dtLock.Lock()
	stdin.dataType = dt
	stdin.dtLock.Unlock()
	return
}

// DefaultDataType defines the murex data type for the stream.Io interface if it's not already set
func (stdin *Stdin) DefaultDataType(err bool) {
	stdin.dtLock.Lock()
	dt := stdin.dataType
	stdin.dtLock.Unlock()

	if dt == "" {
		if err {
			stdin.dtLock.Lock()
			stdin.dataType = types.Null
			stdin.dtLock.Unlock()
		} else {
			stdin.dtLock.Lock()
			stdin.dataType = types.Generic
			stdin.dtLock.Unlock()
		}
	}
}

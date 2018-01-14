package streams

import (
	"bufio"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"sync"
)

// Stdin is the default stdio.Io interface.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
type Stdin struct {
	mutex       sync.Mutex
	buffer      []byte
	closed      bool
	bRead       uint64
	bWritten    uint64
	isParent    bool
	isNamedPipe bool
	dataType    string
	dtLock      sync.Mutex
}

// MaxBufferSize is the maximum size of buffer for stdin
var MaxBufferSize int = 1024 * 1024 * 10 // 10 meg

// NewStdin creates a new stream.Io interface for piping data between processes.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
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
func (stdin *Stdin) Read(p []byte) (i int, err error) {
	defer stdin.mutex.Unlock()
	for {
		stdin.mutex.Lock()

		if len(stdin.buffer) == 0 && stdin.closed {
			return 0, io.EOF
		}

		if len(stdin.buffer) == 0 && !stdin.closed {
			stdin.mutex.Unlock()
			continue
		}

		break
	}

	if len(p) >= len(stdin.buffer) {
		i = len(stdin.buffer)
		copy(p, stdin.buffer)
		stdin.buffer = make([]byte, 0)

	} else {
		i = len(p)
		copy(p[:i], stdin.buffer[:i])
		stdin.buffer = stdin.buffer[i+1:]
	}

	stdin.bRead += uint64(i)

	return i, err
}

// readerFunc is a callback function for reading raw data.
func (stdin *Stdin) readerFunc(callback func([]byte)) {
	for {
		stdin.mutex.Lock()
		if len(stdin.buffer) == 0 {
			if stdin.closed {
				stdin.mutex.Unlock()
				return
			}
			stdin.mutex.Unlock()
			continue
		}

		b := stdin.buffer
		stdin.buffer = make([]byte, 0)

		stdin.bRead += uint64(len(b))

		stdin.mutex.Unlock()

		callback(b)
	}
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
	for {
		stdin.mutex.Lock()
		closed := stdin.closed
		stdin.mutex.Unlock()

		if closed {
			break
		}
	}

	stdin.mutex.Lock()
	b := stdin.buffer
	stdin.buffer = make([]byte, 0)
	stdin.bRead = uint64(len(b))
	stdin.mutex.Unlock()
	return b, nil
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
func (stdin *Stdin) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	stdin.mutex.Lock()
	isClosed := stdin.closed
	stdin.mutex.Unlock()

	if isClosed {
		//return 0, errors.New("Writing to closed pipe.")
		return 0, io.ErrClosedPipe
	}

	for {
		stdin.mutex.Lock()
		buffSize := len(stdin.buffer)
		stdin.mutex.Unlock()

		if buffSize < MaxBufferSize {
			break
		}
	}

	stdin.mutex.Lock()
	stdin.buffer = append(stdin.buffer, b...)
	stdin.bWritten += uint64(len(b))
	stdin.mutex.Unlock()

	return len(b), nil
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
func (stdin *Stdin) WriteTo(dst io.Writer) (n int64, err error) {
	var i int
	stdin.readerFunc(func(b []byte) {
		i, err = dst.Write(b)
		n += int64(i)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
	})
	return
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

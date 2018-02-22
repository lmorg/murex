package streams

import (
	"bufio"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"sync"
)

// Stdtest is a testing Io interface user to workaround writing to the terminal
type Stdtest struct {
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

// NewStdtest creates a new stream.Io interface for piping data between processes.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
func NewStdtest() (stdtest *Stdtest) {
	stdtest = new(Stdtest)
	return
}

// IsTTY returns true because the Stdtest stream is a pseudo-TTY mockup
func (stdtest *Stdtest) IsTTY() bool { return true }

// MakeParent is used for subshells so they don't accidentally close the parent stream.
func (stdtest *Stdtest) MakeParent() {
	stdtest.mutex.Lock()
	stdtest.isParent = true
	stdtest.mutex.Unlock()
}

// UnmakeParent is used when the subshell has terminated, now we allow the stream to be closable again.
func (stdtest *Stdtest) UnmakeParent() {
	stdtest.mutex.Lock()
	if !stdtest.isParent {
		//debug.Log("Trying to unmake parent on a non-parent pipe.")
	}
	stdtest.isParent = false

	stdtest.mutex.Unlock()
}

// MakePipe is used for named pipes. Basically just used to relax the exception handling since we can make fewer
// guarantees about the state of named pipes.
func (stdtest *Stdtest) MakePipe() {
	stdtest.mutex.Lock()
	stdtest.isParent = true
	stdtest.isNamedPipe = true
	stdtest.mutex.Unlock()
}

// Stats provides real time stream stats. Useful for progress bars etc.
func (stdtest *Stdtest) Stats() (bytesWritten, bytesRead uint64) {
	stdtest.mutex.Lock()
	bytesWritten = stdtest.bWritten
	bytesRead = stdtest.bRead
	stdtest.mutex.Unlock()
	return
}

// Read is the standard Reader interface Read() method.
func (stdtest *Stdtest) Read(p []byte) (i int, err error) {
	defer stdtest.mutex.Unlock()
	for {
		stdtest.mutex.Lock()

		if len(stdtest.buffer) == 0 && stdtest.closed {
			return 0, io.EOF
		}

		if len(stdtest.buffer) == 0 && !stdtest.closed {
			stdtest.mutex.Unlock()
			continue
		}

		break
	}

	if len(p) >= len(stdtest.buffer) {
		i = len(stdtest.buffer)
		copy(p, stdtest.buffer)
		stdtest.buffer = make([]byte, 0)

	} else {
		i = len(p)
		copy(p[:i], stdtest.buffer[:i])
		stdtest.buffer = stdtest.buffer[i+1:]
	}

	stdtest.bRead += uint64(i)

	return i, err
}

// readerFunc is a callback function for reading raw data.
func (stdtest *Stdtest) readerFunc(callback func([]byte)) {
	for {
		stdtest.mutex.Lock()
		if len(stdtest.buffer) == 0 {
			if stdtest.closed {
				stdtest.mutex.Unlock()
				return
			}
			stdtest.mutex.Unlock()
			continue
		}

		b := stdtest.buffer
		stdtest.buffer = make([]byte, 0)

		stdtest.bRead += uint64(len(b))

		stdtest.mutex.Unlock()

		callback(b)
	}
}

// ReadLine returns each line in the stream as a callback function
func (stdtest *Stdtest) ReadLine(callback func([]byte)) error {
	scanner := bufio.NewScanner(stdtest)
	for scanner.Scan() {
		callback(append(scanner.Bytes(), utils.NewLineByte...))
	}

	return scanner.Err()
}

// ReadAll reads everything and dump it into one byte slice.
func (stdtest *Stdtest) ReadAll() ([]byte, error) {
	for {
		stdtest.mutex.Lock()
		closed := stdtest.closed
		stdtest.mutex.Unlock()

		if closed {
			break
		}
	}

	stdtest.mutex.Lock()
	b := stdtest.buffer
	stdtest.buffer = make([]byte, 0)
	stdtest.bRead = uint64(len(b))
	stdtest.mutex.Unlock()
	return b, nil
}

// ReadArray returns a data type-specific array returned via a callback function
func (stdtest *Stdtest) ReadArray(callback func([]byte)) error {
	return readArray(stdtest, callback)
}

// ReadMap returns a data type-specific key/values returned via a callback function
func (stdtest *Stdtest) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	return readMap(stdtest, config, callback)
}

// Write is the standard Writer interface Write() method.
func (stdtest *Stdtest) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	stdtest.mutex.Lock()
	isClosed := stdtest.closed
	stdtest.mutex.Unlock()

	if isClosed {
		//return 0, errors.New("Writing to closed pipe.")
		return 0, io.ErrClosedPipe
	}

	for {
		stdtest.mutex.Lock()
		buffSize := len(stdtest.buffer)
		stdtest.mutex.Unlock()

		if buffSize < MaxBufferSize {
			break
		}
	}

	stdtest.mutex.Lock()
	stdtest.buffer = append(stdtest.buffer, b...)
	stdtest.bWritten += uint64(len(b))
	stdtest.mutex.Unlock()

	return len(b), nil
}

// Writeln just calls Write() but with an appended, OS specific, new line.
func (stdtest *Stdtest) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	stdtest.Write(line)
	return len(b), nil
}

// Close the stream.Io interface
func (stdtest *Stdtest) Close() {
	stdtest.mutex.Lock()
	defer stdtest.mutex.Unlock()

	if stdtest.isParent {
		// This will legitimately happen a lot since the reason we mark a stream as parent is to prevent
		// accidental closing. However it's worth pushing a message out in debug mode during this alpha build.
		//debug.Log("Cannot Close() stdtest marked as parent. We don't want to EOT parent streams multiple times")
		return
	}

	if stdtest.closed {
		if stdtest.isNamedPipe {
			//debug.Log("Error with murex named pipes: Trying to close an already closed named pipe.")
		} else {
			//debug.Log("Trying to close an already closed stdtest.")
		}
		return
	}

	stdtest.closed = true
}

// WriteTo reads from the stream.Io interface and writes to a destination io.Writer interface
func (stdtest *Stdtest) WriteTo(dst io.Writer) (n int64, err error) {
	var i int
	stdtest.readerFunc(func(b []byte) {
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
func (stdtest *Stdtest) GetDataType() (dt string) {
	for {
		stdtest.dtLock.Lock()
		dt = stdtest.dataType
		stdtest.dtLock.Unlock()
		if dt != "" {
			return
		}
	}
}

// SetDataType defines the murex data type for the stream.Io interface
func (stdtest *Stdtest) SetDataType(dt string) {
	stdtest.dtLock.Lock()
	stdtest.dataType = dt
	stdtest.dtLock.Unlock()
	return
}

// DefaultDataType defines the murex data type for the stream.Io interface if it's not already set
func (stdtest *Stdtest) DefaultDataType(err bool) {
	stdtest.dtLock.Lock()
	dt := stdtest.dataType
	stdtest.dtLock.Unlock()

	if dt == "" {
		if err {
			stdtest.dtLock.Lock()
			stdtest.dataType = types.Null
			stdtest.dtLock.Unlock()
		} else {
			stdtest.dtLock.Lock()
			stdtest.dataType = types.Generic
			stdtest.dtLock.Unlock()
		}
	}
}

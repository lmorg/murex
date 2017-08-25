package streams

import (
	"bufio"
	"errors"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"sync"
)

// Stdin is the default stream.Io interface.
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

// NewStdin creates a new stream.Io interface for piping data between processes.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
	return
}

// IsTTY returns false because the Stdin stream is not a pseudo-TTY
func (in *Stdin) IsTTY() bool { return false }

// MakeParent is used for subshells so they don't accidentally close the parent stream.
func (rw *Stdin) MakeParent() {
	rw.mutex.Lock()
	rw.isParent = true
	rw.mutex.Unlock()
}

// UnmakeParent is used when the subshell has terminated, now we allow the stream to be closable again.
func (rw *Stdin) UnmakeParent() {
	rw.mutex.Lock()
	if !rw.isParent {
		debug.Log("Trying to unmake parent on a non-parent pipe.")
	}
	rw.isParent = false

	rw.mutex.Unlock()
}

// MakePipe is used for named pipes. Basically just used to relax the exception handling since we can make fewer
// guarantees about the state of named pipes.
func (rw *Stdin) MakePipe() {
	rw.mutex.Lock()
	rw.isParent = true
	rw.isNamedPipe = true
	rw.mutex.Unlock()
}

// Stats provides real time stream stats. Useful for progress bars etc.
func (rw *Stdin) Stats() (bytesWritten, bytesRead uint64) {
	rw.mutex.Lock()
	bytesWritten = rw.bWritten
	bytesRead = rw.bRead
	rw.mutex.Unlock()
	return
}

// Read is the standard Reader interface Read() method.
func (read *Stdin) Read(p []byte) (i int, err error) {
	defer read.mutex.Unlock()
	for {
		read.mutex.Lock()

		if len(read.buffer) == 0 && read.closed {
			return 0, io.EOF
		}

		if len(read.buffer) == 0 && !read.closed {
			read.mutex.Unlock()
			continue
		}

		break
	}

	if len(p) >= len(read.buffer) {
		i = len(read.buffer)
		copy(p, read.buffer)
		read.buffer = make([]byte, 0)

	} else {
		i = len(p)
		copy(p[:i], read.buffer[:i])
		read.buffer = read.buffer[i+1:]
	}

	read.bRead += uint64(i)

	return i, err
}

// ReaderFunc is a callback function for reading raw data.
func (read *Stdin) ReaderFunc(callback func([]byte)) {
	for {
		read.mutex.Lock()
		if len(read.buffer) == 0 {
			if read.closed {
				read.mutex.Unlock()
				return
			}
			read.mutex.Unlock()
			continue
		}

		b := read.buffer
		read.buffer = make([]byte, 0)

		read.bRead += uint64(len(b))

		read.mutex.Unlock()

		callback(b)
	}
}

// ReadLine returns each line in the stream as a callback function
func (read *Stdin) ReadLine(callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(append(scanner.Bytes(), utils.NewLineByte...))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// ReadAll reads everything and dump it into one byte slice.
func (read *Stdin) ReadAll() ([]byte, error) {
	for {
		read.mutex.Lock()
		closed := read.closed
		read.mutex.Unlock()

		if closed {
			break
		}
	}

	read.mutex.Lock()
	b := read.buffer
	read.buffer = make([]byte, 0)
	read.bRead = uint64(len(b))
	read.mutex.Unlock()
	return b, nil
}

// ReadArray returns a data type-specific array returned via a callback function
func (read *Stdin) ReadArray(callback func([]byte)) error {
	return readArray(read, callback)
}

// ReadMap returns a data type-specific key/values returned via a callback function
func (read *Stdin) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	return readMap(read, config, callback)
}

// Write is the standard Writer interface Write() method.
func (write *Stdin) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	write.mutex.Lock()
	defer write.mutex.Unlock()

	if write.closed {
		return 0, errors.New("Writing to closed pipe.")
	}

	write.buffer = append(write.buffer, b...)
	write.bWritten += uint64(len(b))

	return len(b), nil
}

// Writeln just calls Write() but with an appended, OS specific, new line.
func (write *Stdin) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	write.Write(line)
	return len(b), nil
}

// Close the stream.Io interface
func (write *Stdin) Close() {
	write.mutex.Lock()
	defer write.mutex.Unlock()

	if write.isParent {
		// This will legitimately happen a lot since the reason we mark a stream as parent is to prevent
		// accidental closing. However it's worth pushing a message out in debug mode during this alpha build.
		debug.Log("Cannot Close() stdin marked as parent. We don't want to EOT parent streams multiple times")
		return
	}

	if write.closed {
		if write.isNamedPipe {
			debug.Log("Error with murex named pipes: Trying to close an already closed named pipe.")
			return
		} else {
			debug.Log("Trying to close an already closed stdin.")
			return
		}
	}

	write.closed = true
}

// WriteTo reads from the stream.Io interface and writes to a destination io.Writer interface
func (rw *Stdin) WriteTo(dst io.Writer) (n int64, err error) {
	var i int
	rw.ReaderFunc(func(b []byte) {
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
func (in *Stdin) GetDataType() (dt string) {
	for {
		in.dtLock.Lock()
		dt = in.dataType
		in.dtLock.Unlock()
		if dt != "" {
			return
		}
	}
}

// SetDataType defines the murex data type for the stream.Io interface
func (in *Stdin) SetDataType(dt string) {
	in.dtLock.Lock()
	in.dataType = dt
	in.dtLock.Unlock()
	return
}

// DefaultDataType defines the murex data type for the stream.Io interface if it's not already set
func (in *Stdin) DefaultDataType(err bool) {
	in.dtLock.Lock()
	dt := in.dataType
	in.dtLock.Unlock()

	if dt == "" {
		if err {
			in.dtLock.Lock()
			in.dataType = types.Null
			in.dtLock.Unlock()
		} else {
			in.dtLock.Lock()
			in.dataType = types.Generic
			in.dtLock.Unlock()
		}
	}
}

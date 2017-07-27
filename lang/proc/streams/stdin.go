package streams

import (
	"bufio"
	"errors"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
	"sync"
)

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

func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
	return
}

func (in *Stdin) IsTTY() bool { return false }

// This is used for subshells so they don't accidentally close the parent stream.
func (rw *Stdin) MakeParent() {
	rw.mutex.Lock()
	rw.isParent = true
	rw.mutex.Unlock()
}

// Subshell terminated, now we allow the stream to be closable again.
func (rw *Stdin) UnmakeParent() {
	rw.mutex.Lock()
	if !rw.isParent {
		if rw.isNamedPipe {
			os.Stderr.WriteString("Error with murex named pipes: Trying to unmake parent on a non-parent pipe." + utils.NewLineString)
		} else {
			// Should be fine to panic because this runtime error is generated from block compilation.
			//panic("Cannot call UnmakeParent() on stdin not marked as Parent.")
		}
	}
	rw.isParent = false

	rw.mutex.Unlock()
}

// This is used for named pipes. Basically just used to relax the exception handling since we can make fewer guarantees
// about the state of named pipes.
func (rw *Stdin) MakePipe() {
	rw.mutex.Lock()
	rw.isParent = true
	rw.isNamedPipe = true
	rw.mutex.Unlock()
}

// Real time stream stats. Useful for progress bars etc.
func (rw *Stdin) Stats() (bytesWritten, bytesRead uint64) {
	rw.mutex.Lock()
	bytesWritten = rw.bWritten
	bytesRead = rw.bRead
	rw.mutex.Unlock()
	return
}

// Standard Reader interface Read() method.
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

// A callback function for reading raw data.
func (read *Stdin) ReaderFunc(callback func([]byte)) {
	//defer read.mutex.Unlock()
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

// Should be more performant than ReadLine() because it's uses callback functions (ie does not need to be stateless) so
// we don't need to keep pushing stuff back to the interfaces buffer.
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

// Read everything and dump it into one byte slice.
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

func (read *Stdin) ReadArray(callback func([]byte)) error {
	return readArray(read, callback)
}

func (read *Stdin) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	return readMap(read, config, callback)
}

// Standard Writer interface Write() method.
func (write *Stdin) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	write.mutex.Lock()
	defer write.mutex.Unlock()

	if write.closed {
		if write.isNamedPipe {
			return 0, errors.New("Error with murex named pipes: Trying to write to a closed pipe.")
		} else {
			// This shouldn't happen because it then means we have lost track of the state of the streams.
			// So we'll throw a panic to highlight our error early on and force better code.
			panic("Writing to closed pipe.")
		}
	}

	write.buffer = append(write.buffer, b...)
	write.bWritten += uint64(len(b))

	return len(b), nil
}

// Just calls Write() but with an appended, OS specific, new line.
func (write *Stdin) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	write.Write(line)
	return len(b), nil
}

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
			os.Stderr.WriteString("Error with murex named pipes: Trying to close an already closed named pipe." + utils.NewLineString)
			return
		} else {
			// This shouldn't happen because it then means we have lost track of the state of the streams.
			// So we'll throw a panic to highlight our error early on and force better code.
			//panic("Trying to close an already closed stdin.")
			os.Stderr.WriteString("Trying to close an already closed stdin." + utils.NewLineString)
		}
	}

	write.closed = true
}

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

func (in *Stdin) SetDataType(dt string) {
	in.dtLock.Lock()
	in.dataType = dt
	in.dtLock.Unlock()
	return
}

func (in *Stdin) DefaultDataType(err bool) {
	return
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

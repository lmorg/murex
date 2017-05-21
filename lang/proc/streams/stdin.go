package streams

import (
	"bufio"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/utils"
	"io"
	"sync"
)

type Stdin struct {
	sync.Mutex
	buffer   []byte
	closed   bool
	bRead    uint64
	bWritten uint64
	isParent bool
}

func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
	//stdin.buffer = make([]byte, 1)
	return
}

// This is used for subshells so they don't accidentally close the parent stream.
func (rw *Stdin) MakeParent() {
	rw.Lock()
	rw.isParent = true
	rw.Unlock()
}

// Subshell terminated, now we allow the stream to be closable again.
func (rw *Stdin) UnmakeParent() {
	rw.Lock()
	if !rw.isParent {
		// Should be fine to panic because this runtime error is generated from block compilation.
		panic("Cannot call UnmakeParent() on stdin not marked as Parent.")
	}
	rw.isParent = false

	rw.Unlock()
}

// Real time stream stats. Useful for progress bars etc.
func (rw *Stdin) Stats() (bytesWritten, bytesRead uint64) {
	rw.Lock()
	bytesWritten = rw.bWritten
	bytesRead = rw.bRead
	rw.Unlock()
	return
}

// Standard Reader interface Read() method.
func (read *Stdin) Read(p []byte) (i int, err error) {
	for {
		read.Lock()
		//defer read.Unlock()

		if len(read.buffer) == 0 && read.closed {
			read.Unlock()
			return 0, io.EOF
		}

		if len(read.buffer) == 0 && !read.closed {
			read.Unlock()
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

	read.Unlock()
	return i, err
}

// A callback function for reading raw data.
func (read *Stdin) ReaderFunc(callback func([]byte)) {
	for {
		read.Lock()
		if len(read.buffer) == 0 {
			if read.closed {
				read.Unlock()
				return
			}
			read.Unlock()
			continue
		}

		b := read.buffer
		read.buffer = make([]byte, 0)

		read.bRead += uint64(len(b))

		read.Unlock()

		callback(b)
	}
}

// Should be more performant than ReadLine() because it's uses callback functions (ie does not need to be stateless) so
// we don't need to keep pushing stuff back to the interfaces buffer.
func (read *Stdin) ReadLineFunc(callback func([]byte)) {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(append(scanner.Bytes(), utils.NewLineByte...))
	}

	if err := scanner.Err(); err != nil {
		panic("ReadLine: " + err.Error())
	}

	return
}

// Read everything and dump it into one byte slice.
func (read *Stdin) ReadAll() (b []byte) {
	for !read.closed {
		// Wait for interface to close.
	}

	read.Lock()
	b = read.buffer
	read.buffer = make([]byte, 0)
	read.bRead = uint64(len(b))
	read.Unlock()
	return
}

// Standard Writer interface Write() method.
func (write *Stdin) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	write.Lock()

	if write.closed {
		// This shouldn't happen because it then means we have lost track of the state of the streams.
		// So we'll throw a panic to highlight our error early on and force better code.
		panic("Writing to closed pipe.")
	}

	write.buffer = append(write.buffer, b...)
	write.bWritten += uint64(len(b))

	write.Unlock()

	return len(b), nil
}

// Just calls Write() but with an appended, OS specific, new line.
func (write *Stdin) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	write.Write(line)
	return len(b), nil
}

func (write *Stdin) Close() {
	write.Lock()
	defer write.Unlock()

	if write.isParent {
		// This will legitimately happen a lot since the reason we mark a stream as parent is to prevent
		// accidental closing. However it's worth pushing a message out in debug mode during this alpha build.
		debug.Log("Cannot Close() stdin marked as parent. We don't want to EOT parent streams multiple times")
		return
	}

	if write.closed {
		// This shouldn't happen because it then means we have lost track of the state of the streams.
		// So we'll throw a panic to highlight our error early on and force better code.
		panic("Trying to close an already closed stdin.")
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

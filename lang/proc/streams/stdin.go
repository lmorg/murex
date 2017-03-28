package streams

import (
	"bytes"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/utils"
	"io"
	"sync"
)

type Stdin struct {
	sync.Mutex
	buffer   [][]byte
	closed   bool
	bRead    uint64
	bWritten uint64
	isParent bool
}

func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
	stdin.buffer = make([][]byte, 1)
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
		panic("Cannot call CloseParent() on stdin not marked as Parent.")
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
	read.Lock()
	defer read.Unlock()

	if len(read.buffer) == 0 && read.closed {
		return 0, io.EOF
	}

	if len(read.buffer) == 0 && !read.closed {
		return 0, nil
	}

	if len(p) >= len(read.buffer[0]) {
		i = len(read.buffer[0])
		copy(p, read.buffer[0])
		read.buffer = read.buffer[1:]

	} else {
		i = len(p)
		copy(p[:i], read.buffer[0][:i])
		read.buffer[0] = read.buffer[0][i+1:]
	}

	read.bRead += uint64(i)

	return i, err
}

// Reads a line at a time. This method will be slower than the other Read* methods because ReadLine() needs to be
// stateless. So ReadLine() will write back to the interface on occasions where \n appears mid-buffer.
func (read *Stdin) ReadLine(line *string) (more bool) {
	defer func() {
		read.bRead += uint64(len(*line))
		read.Unlock()
	}()

	more = true

start:
	for {
		read.Lock()
		if len(read.buffer) == 0 {
			if read.closed {
				return false
			}
			read.Unlock()
			continue
		}
	}

	b := read.buffer[0]
	read.buffer = read.buffer[1:]

	lines := bytes.SplitAfter(b, []byte{'\n'})

	if len(lines) > 1 {
		*line = string(lines[0])
		read.buffer = append(lines[1:], read.buffer...)
		return
	}

	if len(lines[0]) > 0 && lines[0][len(lines[0])] == '\n' {
		*line = string(lines[0])
		return
	}

	if len(read.buffer) > 0 {
		read.buffer[0] = append(lines[0], read.buffer[0]...)

	} else {
		read.buffer = lines
	}

	goto start
}

// Faster than ReadLine but doesn't chunk the data based on new lines.
func (read *Stdin) ReadData() (b []byte, more bool) {
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

		b = read.buffer[0]
		read.buffer = read.buffer[1:]
		read.bRead += uint64(len(b))

		read.Unlock()
		more = true

		if len(b) == 0 {
			continue
		}
		return
	}
	return
}

// A callback function for reading raw data. I'm thinking this is largely unnecessary so might delete it.
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

		b := read.buffer[0]
		read.buffer = read.buffer[1:]
		read.bRead += uint64(len(b))

		read.Unlock()

		callback(b)
	}
}

// Should be more performant than ReadLine() because it's uses callback functions (ie does not need to be stateless).
func (read *Stdin) ReadLineFunc(callback func([]byte)) {
	var remainder []byte
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

		b := read.buffer[0]
		read.buffer = read.buffer[1:]

		read.Unlock()

		lines := bytes.SplitAfter(b, []byte{'\n'})
		lines[0] = append(remainder, lines[0]...)
		remainder = []byte{}

		for len(lines) > 1 {
			if len(lines[0]) > 0 {
				read.Lock()
				read.bRead += uint64(len(lines[0]))
				read.Unlock()
				callback(lines[0])
			}
			lines = lines[1:]
		}

		if len(lines[0]) > 0 && lines[0][len(lines[0])-1] == '\n' {
			read.Lock()
			read.bRead += uint64(len(lines[0]))
			read.Unlock()
			callback(lines[0])

			lines = lines[1:]

		} else {
			if read.closed {
				lines[0] = append(lines[0], utils.NewLineByte...)
				read.Lock()
				read.bRead += uint64(len(lines[0]))
				read.Unlock()
				callback(lines[0])
			} else {
				remainder = lines[0]
			}
		}

	}

	return
}

// Read everything and dump it into one byte slice.
func (read *Stdin) ReadAll() (b []byte) {
	read.ReaderFunc(func(line []byte) {
		b = append(b, line...)
	})

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

	write.buffer = append(write.buffer, b)
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

/*func (rw *Stdin) ReadFrom(src io.Reader) (n int64, err error) {
	b := make([]byte, 1024)
	for {
		i, err := src.Read(b)
		debug.Log("#############readfrom#####################", i, string(b))
		rw.Write(b[:i])

		n += int64(i)

		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return n, err
		}
	}
	return n, nil
}

func (rw *Stdin) WriteTo(dst io.Writer) (n int64, err error) {
	var i int
	rw.ReaderFunc(func(b []byte) {
		i, err = dst.Write(b)
		debug.Log("#############writeto#####################", i, string(b))
		n += int64(i)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
	})
	return
}*/

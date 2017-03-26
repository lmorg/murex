package streams

import (
	"bytes"
	"io"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/utils"
	"sync"
)

// See comments on Close() for rational behind this buffer structure.
type buffer struct {
	data []byte
	eot  bool
}

type Stdin struct {
	sync.Mutex
	data     []buffer
	closed   bool
	bRead    uint64
	bWritten uint64
	isParent bool
}

func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
	stdin.data = make([]buffer, 0)
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
	for {
		read.Lock()
		if len(read.data) == 0 {
			read.Unlock()
			continue
		}
		break
	}

	var eot bool
	copy(p, make([]byte, len(p)))
	if len(p) >= len(read.data[0].data) {
		i = len(read.data[0].data)
		copy(p, read.data[0].data)
		eot = read.data[0].eot
		read.data = read.data[1:]
		debug.Log("read: [2]:", string(p), len(p), len(read.data))
	} else {
		i = len(p)
		copy(p, read.data[0].data[:i])
		read.data[0].data = read.data[0].data[i+1:]
	}

	read.bRead += uint64(i)
	read.Unlock()

	if eot {
		err = io.EOF
	}

	return
}

// Reads a line at a time. This method will be slower than the other Read* methods because ReadLine() needs to be
// stateless. So ReadLine() will write back to the interface on occasions where \n appears mid-buffer.
func (read *Stdin) ReadLine(line *string) (more bool) {
	var remainder []byte
	more = true

scan:
	for {
		read.Lock()
		if len(read.data) == 0 {
			read.Unlock()
			continue
		}
		break
	}

	in := read.data[0]
	read.data = read.data[1:]
	read.Unlock()

	lines := bytes.SplitAfter(in.data, []byte{'\n'})

	if len(lines) > 1 {
		read.Lock()
		for i := len(lines) - 1; i > 1; i-- {
			if len(lines[i]) != 0 {
				read.data = append([]buffer{{data: lines[i]}}, read.data...)
			}
		}
		if in.eot {
			read.data[len(read.data)-1].eot = true
		}
		read.Unlock()
		*line = string(lines[0])

	} else if len(lines[0]) > 0 && lines[0][len(lines[0])-1] == '\n' {
		*line = string(lines[0])
		more = !in.eot

	} else {
		if in.eot {
			*line = string(lines[0]) + utils.NewLineString
			more = false

		} else {
			*line = ""
			remainder = lines[0]
		}
	}

	if len(remainder) != 0 {
		read.Lock()
		if len(read.data) > 0 {
			read.data[0].data = append(remainder, read.data[0].data...)
		} else {
			read.data = append(read.data, buffer{
				data: remainder,
				eot:  false,
			})
		}
		read.Unlock()
		remainder = []byte{}
		goto scan // I know goto's are "ugly", but it makes some structural sense in the context of this function.
	}

	read.Lock()
	read.bRead += uint64(len(*line))
	read.Unlock()
	return
}

// Faster than ReadLine but doesn't chunk the data based on new lines.
func (read *Stdin) ReadData() (b []byte, more bool) {
	for {
		read.Lock()
		if len(read.data) == 0 {
			read.Unlock()
			continue
		}

		in := read.data[0]
		read.data = read.data[1:]
		read.bRead += uint64(len(in.data))

		read.Unlock()

		if len(in.data) == 0 && !in.eot {
			continue
		}

		b = in.data
		more = !in.eot

		break
	}
	return
}

// A callback function for reading raw data. I'm thinking this is largely unnecessary so might delete it.
func (read *Stdin) ReaderFunc(callback func([]byte)) {
	for {
		read.Lock()
		if len(read.data) == 0 {
			read.Unlock()
			continue
		}

		in := read.data[0]
		read.data = read.data[1:]
		read.bRead += uint64(len(in.data))

		read.Unlock()

		line := bytes.SplitAfter(in.data, utils.NewLineByte)
		for i := range line {
			if len(line[i]) > 0 {
				callback(line[i])
			}
		}

		if in.eot {
			break
		}
	}
}

// Should be more performant than ReadLine() because it's uses callback functions (ie does not need to be stateless).
func (read *Stdin) ReadLineFunc(callback func([]byte)) {
	var remainder []byte
	for {
		read.Lock()
		if len(read.data) == 0 {
			read.Unlock()
			continue
		}

		in := read.data[0]
		read.data = read.data[1:]

		read.Unlock()

		lines := bytes.SplitAfter(in.data, []byte{'\n'})
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
			if in.eot {
				lines[0] = append(lines[0], utils.NewLineByte...)
				read.Lock()
				read.bRead += uint64(len(lines[0]))
				read.Unlock()
				callback(lines[0])
			} else {
				remainder = lines[0]
			}
		}

		if in.eot {
			break
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
		return len(b), nil
	}
	write.Lock()

	if write.closed {
		// This shouldn't happen because it then means we have lost track of the state of the streams.
		// So we'll throw a panic to highlight our error early on and force better code.
		panic("Writing to closed pipe.")
	}

	write.data = append(write.data, buffer{data: b})
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

// Close the interface and mark the stream with an EOT.
// Inside this shell we don't use an EOT byte(4) because I want to support binary streams that may also contain byte(4).
// So instead we have a struct which marks whether that element is the EOT. This structure is more memory expensive
// but makes sense in terms of clean code design and rapid prototyping. Eventually I may replace the buffer structure
// with a straight []byte and an slice size counter to mark the remainder in the slice if an EOT is expected.
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
	write.data = append(write.data, buffer{
		eot: true,
	})
}

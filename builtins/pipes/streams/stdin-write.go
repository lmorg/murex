package streams

import (
	"io"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils"
)

// Write is the standard Writer interface Write() method.
func (stdin *Stdin) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	for {
		select {
		case <-stdin.ctx.Done():
			stdin.mutex.Lock()
			stdin.buffer = []byte{}
			stdin.mutex.Unlock()
			return 0, io.ErrClosedPipe
		default:
		}

		stdin.mutex.Lock()
		buffSize := len(stdin.buffer)
		maxBufferSize := stdin.max
		stdin.mutex.Unlock()

		if buffSize < maxBufferSize || maxBufferSize == 0 {
			break
		}
	}

	stdin.mutex.Lock()
	stdin.buffer = appendBytes(stdin.buffer, p...)
	stdin.bWritten += uint64(len(p))
	stdin.mutex.Unlock()

	return len(p), nil
}

// Writeln just calls Write() but with an appended, OS specific, new line.
func (stdin *Stdin) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	stdin.Write(line)
	return len(b), nil
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (stdin *Stdin) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(stdin, dataType)
}

// ReadFrom reads data from r until EOF and appends it to the buffer.
func (stdin *Stdin) ReadFrom(r io.Reader) (int64, error) {
	var total int64

	stdin.mutex.Lock()
	stdin.max = 0
	stdin.mutex.Unlock()

	for {
		select {
		case <-stdin.ctx.Done():
			return total, io.ErrClosedPipe

		default:
			p := make([]byte, 1024)
			i, err := r.Read(p)

			if err == io.EOF {
				return total, nil
			}

			if err != nil {
				return total, err
			}

			i, err = stdin.Write(p[:i])
			if err != nil {
				return total, err
			}

			total += int64(i)
		}
	}
}

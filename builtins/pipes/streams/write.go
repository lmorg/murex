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

		//stdin.mutex.RLock()
		stdin.mutex.Lock()
		buffSize := len(stdin.buffer)
		maxBufferSize := stdin.max
		//stdin.mutex.RUnlock()
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
	return stdin.Write(appendBytes(b, utils.NewLineByte...))
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

	var rErr, wErr error
	i := 0

	for {
		select {
		case <-stdin.ctx.Done():
			return total, io.ErrClosedPipe

		default:
			p := make([]byte, 1024)
			i, rErr = r.Read(p)

			if rErr != nil && rErr != io.EOF {
				return total, rErr
			}

			i, wErr = stdin.Write(p[:i])
			if wErr != nil {
				return total, wErr
			}

			total += int64(i)
			if rErr == io.EOF {
				return total, nil
			}
		}
	}
}

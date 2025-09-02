package streams

import (
    "io"

    "github.com/lmorg/murex/lang/stdio"
    "github.com/lmorg/murex/utils"
)

// Write is the standard Writer interface Write() method.
func (stdin *Stdin) Write(p []byte) (int, error) {
    if len(p) == 0 {
        return 0, nil
    }

    stdin.mutex.Lock()
    defer stdin.mutex.Unlock()

    for {
        if stdin.ctx.Err() != nil {
            stdin.buffer = stdin.buffer[:0]
            return 0, io.ErrClosedPipe
        }
        // If max==0 then unbounded; otherwise wait while at or above max
        if stdin.max == 0 || len(stdin.buffer) < stdin.max {
            break
        }
        stdin.cond.Wait()
    }

    stdin.buffer = appendBytes(stdin.buffer, p...)
    stdin.bWritten += uint64(len(p))
    // Wake potential readers waiting for data
    stdin.cond.Signal()
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

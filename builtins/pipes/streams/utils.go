package streams

import (
	"github.com/lmorg/murex/lang/types"
)

// Shamelessly stolen from https://blog.golang.org/go-slices-usage-and-internals
// (it works well so why reinvent the wheel?)
func appendBytes(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// IsTTY returns false because the Stdin stream is not a pseudo-TTY
func (stdin *Stdin) IsTTY() bool { return false }

// Stats provides real time stream stats. Useful for progress bars etc.
func (stdin *Stdin) Stats() (bytesWritten, bytesRead uint64) {
	//stdin.mutex.RLock()
	stdin.mutex.Lock()
	bytesWritten = stdin.bWritten
	bytesRead = stdin.bRead
	//stdin.mutex.RUnlock()
	stdin.mutex.Unlock()
	return
}

// GetDataType returns the murex data type for the stream.Io interface
func (stdin *Stdin) GetDataType() (dt string) {
	for {
		select {
		case <-stdin.ctx.Done():
			// This should probably be locked to avoid a data race, but I'm also
			// quite scared locking it might also cause deadlocks given processes
			// can be terminated at random points by users. Thus I'm think those
			// edge cases of a data race will have more desirable side effects
			// than those edge case of deadlocks.
			//stdin.dtLock.Lock()
			dt = stdin.dataType
			//stdin.dtLock.Unlock()
			if dt != "" {
				return dt
			}
			return types.Generic
		default:
		}

		stdin.mutex.Lock()
		//stdin.dtLock.Lock()
		dt = stdin.dataType
		//stdin.dtLock.Unlock()

		if dt != "" {
			stdin.mutex.Unlock()
			return
		}

		fin := stdin.dependants < 1
		stdin.mutex.Unlock()

		if fin {
			return types.Generic
		}
	}
}

// SetDataType defines the murex data type for the stream.Io interface
func (stdin *Stdin) SetDataType(dt string) {
	//stdin.dtLock.Lock()
	stdin.mutex.Lock()
	if stdin.dataType == "" {
		stdin.dataType = dt
	}
	//stdin.dtLock.Unlock()
	stdin.mutex.Unlock()
}

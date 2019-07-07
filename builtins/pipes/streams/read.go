package streams

import (
	"bufio"
	"io"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils"
)

// Read is the standard Reader interface Read() method.
func (stdin *Stdin) Read(p []byte) (i int, err error) {
	for {
		select {
		case <-stdin.ctx.Done():
			return 0, io.EOF
		default:
		}

		//stdin.mutex.RLock()
		stdin.mutex.Lock()
		l := len(stdin.buffer)
		deps := stdin.dependants
		//stdin.mutex.RUnlock()
		stdin.mutex.Unlock()

		if l == 0 {
			if deps < 1 {
				return 0, io.EOF
			}

			continue
		}

		break
	}

	stdin.mutex.Lock()
	//stdin.mutex.RLock()

	if len(p) >= len(stdin.buffer) {
		i = len(stdin.buffer)
		copy(p, stdin.buffer)
		//stdin.mutex.RUnlock()
		//stdin.mutex.Lock()
		stdin.buffer = make([]byte, 0)

	} else {
		i = len(p)
		copy(p, stdin.buffer[:i])
		//stdin.mutex.RUnlock()
		//stdin.mutex.Lock()
		stdin.buffer = stdin.buffer[i:]
	}

	stdin.bRead += uint64(i)
	stdin.mutex.Unlock()

	return i, err
}

// ReadLine returns each line in the stream as a callback function
func (stdin *Stdin) ReadLine(callback func([]byte)) error {
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		b := scanner.Bytes()
		// surely this is covered by Read() ...?
		//stdin.mutex.Lock()
		//stdin.bRead += uint64(len(b))
		//stdin.mutex.Unlock()
		callback(append(b, utils.NewLineByte...))
	}

	return scanner.Err()
}

// ReadAll reads everything and dump it into one byte slice.
func (stdin *Stdin) ReadAll() ([]byte, error) {
	stdin.mutex.Lock()
	stdin.max = 0
	stdin.mutex.Unlock()

	for {
		select {
		case <-stdin.ctx.Done():
			break
		default:
		}

		//stdin.mutex.RLock()
		stdin.mutex.Lock()
		closed := stdin.dependants < 1
		//stdin.mutex.RUnlock()
		stdin.mutex.Unlock()

		if closed {
			break
		}
	}

	stdin.mutex.Lock()
	defer stdin.mutex.Unlock()
	stdin.bRead = uint64(len(stdin.buffer))
	return stdin.buffer, nil
}

// ReadArray returns a data type-specific array returned via a callback function
func (stdin *Stdin) ReadArray(callback func([]byte)) error {
	return stdio.ReadArray(stdin, callback)
}

// ReadMap returns a data type-specific key/values returned via a callback function
func (stdin *Stdin) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	return stdio.ReadMap(stdin, config, callback)
}

// WriteTo reads from the stream.Io interface and writes to a destination
// io.Writer interface
func (stdin *Stdin) WriteTo(w io.Writer) (int64, error) {
	return stdio.WriteTo(stdin, w)
}

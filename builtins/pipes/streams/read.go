package streams

import (
	"bufio"
	"context"
	"io"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
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

		stdin.mutex.Lock()
		l := len(stdin.buffer)
		deps := stdin.dependents

		stdin.mutex.Unlock()

		if l == 0 {
			if deps < 1 {
				return 0, io.EOF
			}
			//time.Sleep(3 * time.Millisecond)
			continue
		}

		break
	}

	stdin.mutex.Lock()

	if len(p) >= len(stdin.buffer) {
		i = len(stdin.buffer)
		copy(p, stdin.buffer)

		stdin.buffer = make([]byte, 0)

	} else {
		i = len(p)
		copy(p, stdin.buffer[:i])

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
			goto read
		default:
		}

		stdin.mutex.Lock()
		closed := stdin.dependents < 1

		stdin.mutex.Unlock()

		if closed {
			break
		}
	}

read:
	stdin.mutex.Lock()
	stdin.bRead = uint64(len(stdin.buffer))
	b := stdin.buffer
	stdin.mutex.Unlock()
	return b, nil
}

// ReadArray returns a data type-specific array returned via a callback function
func (stdin *Stdin) ReadArray(ctx context.Context, callback func([]byte)) error {
	return stdio.ReadArray(ctx, stdin, callback)
}

// ReadArrayWithType returns an array like "ReadArray" plus data type via a callback function
func (stdin *Stdin) ReadArrayWithType(ctx context.Context, callback func(interface{}, string)) error {
	return stdio.ReadArrayWithType(ctx, stdin, callback)
}

// ReadMap returns a data type-specific key/values returned via a callback function
func (stdin *Stdin) ReadMap(config *config.Config, callback func(*stdio.Map)) error {
	return stdio.ReadMap(stdin, config, callback)
}

// WriteTo reads from the stream.Io interface and writes to a destination
// io.Writer interface
func (stdin *Stdin) WriteTo(w io.Writer) (int64, error) {
	return stdio.WriteTo(stdin, w)
}

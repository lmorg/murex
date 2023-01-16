package string

import (
	"bufio"
	"bytes"
	"context"

	"github.com/lmorg/murex/lang/stdio"
)

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(bytes.TrimSpace(scanner.Bytes()))
		}
	}

	return scanner.Err()
}

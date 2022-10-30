package string

import (
	"bufio"
	"bytes"
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArrayWithType(ctx context.Context, read stdio.Io, callback func([]byte, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(bytes.TrimSpace(scanner.Bytes()), types.String)
		}
	}

	return scanner.Err()
}

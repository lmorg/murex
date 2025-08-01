package apachelogs

import (
	"bufio"
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
			callback(scanner.Bytes())
		}
	}

	return scanner.Err()
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(scanner.Bytes(), typeAccess)
		}
	}

	return scanner.Err()
}

package generic

import (
	"bufio"
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
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

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(interface{}, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(scanner.Text(), types.Generic)
		}
	}

	return scanner.Err()
}

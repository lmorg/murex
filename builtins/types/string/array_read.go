package string

import (
	"bufio"
	"bytes"
	"context"
	"fmt"

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
			callback(bytes.TrimSpace(scanner.Bytes()))
		}
	}

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("error while reading a %s array: %s", types.String, err.Error())
	}

	return nil
}

package generic

import (
	"bufio"
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
			callback(scanner.Bytes())
		}
	}

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("error while reading a %s array: %s", types.Generic, err.Error())
	}
	return nil
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(scanner.Text(), types.Generic)
		}
	}

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("error while reading a %s array: %s", types.Generic, err.Error())
	}
	return nil
}

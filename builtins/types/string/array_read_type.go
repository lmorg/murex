package string

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(strings.TrimSpace(scanner.Text()), types.String)
		}
	}

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("error while reading a %s array: %s", types.String, err.Error())
	}
	return nil
}

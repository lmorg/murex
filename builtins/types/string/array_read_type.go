package string

import (
	"bufio"
	"context"
	"strings"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(interface{}, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(strings.TrimSpace(scanner.Text()), types.String)
		}
	}

	return scanner.Err()
}

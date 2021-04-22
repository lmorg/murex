package string

import (
	"bufio"
	"bytes"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArrayWithType(read stdio.Io, callback func([]byte, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(bytes.TrimSpace(scanner.Bytes()), types.String)
	}

	return scanner.Err()
}

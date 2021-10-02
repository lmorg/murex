package string

import (
	"bufio"
	"bytes"

	"github.com/lmorg/murex/lang/stdio"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(bytes.TrimSpace(scanner.Bytes()))
	}

	return scanner.Err()
}

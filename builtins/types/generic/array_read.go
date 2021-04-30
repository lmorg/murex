package generic

import (
	"bufio"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(scanner.Bytes())
	}

	return scanner.Err()
}

func readArrayWithType(read stdio.Io, callback func([]byte, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(scanner.Bytes(), types.Generic)
	}

	return scanner.Err()
}

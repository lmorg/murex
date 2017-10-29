package apachelogs

import (
	"bufio"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(scanner.Bytes())
	}

	return scanner.Err()
}

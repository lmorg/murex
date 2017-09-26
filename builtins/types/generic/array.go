package generic

import (
	"bufio"
	"bytes"
	"github.com/lmorg/murex/lang/proc/streams"
)

func readArray(read streams.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(bytes.TrimSpace(scanner.Bytes()))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

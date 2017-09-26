package generic

import (
	"bufio"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams"
	"strconv"
)

func readMap(read streams.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		recs := rxWhitespace.Split(scanner.Text(), -1)
		for i := range recs {
			callback(strconv.Itoa(i), string(recs[i]), i+1 == len(recs))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

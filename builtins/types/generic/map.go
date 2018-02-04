package generic

import (
	"bufio"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"strconv"
)

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	/*scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		recs := rxWhitespace.Split(scanner.Text(), -1)
		for i := range recs {
			callback(strconv.Itoa(i), string(recs[i]), i+1 == len(recs))
		}
	}

	err := scanner.Err()
	return err*/

	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		row := rxWhitespace.Split(scanner.Text(), -1)

		for i := range row {
			callback(strconv.Itoa(i), row[i], i+1 == len(row))
		}

	}

	err := scanner.Err()
	return err

}

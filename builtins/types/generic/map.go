package generic

import (
	"bufio"
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		row := rxWhitespace.Split(scanner.Text(), -1)

		for i := range row {
			callback(&stdio.Map{
				Key:      strconv.Itoa(i),
				Value:    row[i],
				DataType: types.String,
				Last:     i+1 == len(row),
			})
		}
	}

	err := scanner.Err()
	return err

}

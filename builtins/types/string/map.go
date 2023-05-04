package string

import (
	"bufio"
	"bytes"
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
	scanner := bufio.NewScanner(read)
	i := -1
	for scanner.Scan() {
		i++

		callback(&stdio.Map{
			Key:      strconv.Itoa(i),
			Value:    string(bytes.TrimSpace(scanner.Bytes())),
			DataType: types.String,
			Last:     false,
		})
	}

	return scanner.Err()
}

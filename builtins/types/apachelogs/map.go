// +build ignore

package apachelogs

import (
	"bufio"
	"github.com/lmorg/apachelogs"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
)

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		access, err, _ := apachelogs.ParseAccessLine(scanner.Text())
		if err != nil {
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

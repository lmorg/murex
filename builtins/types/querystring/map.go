package string

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/stdio"
	"net/url"
)

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	if len(b) == 0 {
		return nil
	}

	if b[0] == '?' {
		if len(b) == 1 {
			return nil
		}
		b = b[1:]
	}

	values, err := url.ParseQuery(string(b))
	if err != nil {
		return err
	}

	for k := range values {
		for i := range values[k] {
			callback(k, values[k][i], true)
		}
	}

	return nil
}

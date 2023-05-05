package string

import (
	"net/url"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
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
			callback(&stdio.Map{
				Key:      k,
				Value:    values[k][i],
				DataType: types.String,
				Last:     true,
			})
		}
	}

	return nil
}

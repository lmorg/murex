package sexp

import (
	"strconv"

	"github.com/abesto/sexp"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
)

func readMapC(read stdio.Io, config *config.Config, callback func(*stdio.Map)) error {
	return readMap(read, config, callback, csexp)
}

func readMapS(read stdio.Io, config *config.Config, callback func(*stdio.Map)) error {
	return readMap(read, config, callback, sexpr)
}

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map), dataType string) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	se, err := sexp.Unmarshal(b)
	if err == nil {

		for i := range se {
			j, err := sexp.Marshal(se[i], dataType == csexp)
			if err != nil {
				return err
			}

			callback(&stdio.Map{
				Key:      strconv.Itoa(i),
				Value:    string(j),
				DataType: dataType,
				Last:     i != len(se)-1,
			})
		}

		return nil
	}
	return err
}

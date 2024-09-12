package lists

import (
	"bytes"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("mjoin", cmdMjoin, types.ReadArray, types.String)
}

func cmdMjoin(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	separator, err := p.Parameters.Byte(0)
	if err != nil {
		return err
	}

	if p.IsMethod {
		return mjoinMethod(p, separator)
	}

	return mjoinFunction(p, separator)
}

func mjoinMethod(p *lang.Process, separator []byte) error {
	var slice [][]byte

	err := p.Stdin.ReadArray(p.Context, func(b []byte) {
		slice = append(slice, b)
	})

	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(bytes.Join(slice, separator))
	return err
}

func mjoinFunction(p *lang.Process, separator []byte) error {
	slice := p.Parameters.StringArray()[1:]

	s := strings.Join(slice, string(separator))
	_, err := p.Stdout.Write([]byte(s))
	return err
}

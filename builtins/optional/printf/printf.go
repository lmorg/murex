package printf

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("printf", cmdPrintf, types.String)
}

func cmdPrintf(p *lang.Process) error {
	str, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	s := fmt.Sprintf(str, convSlice(p.Parameters.StringArray()[1:])...)

	_, err = p.Stdout.Write([]byte(s))
	return err
}

func convSlice(s []string) []any {
	r := make([]any, len(s))
	for i := range s {
		r[i] = s[i]
	}
	return r
}

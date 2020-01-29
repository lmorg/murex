// +build plan9

package processes

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func mkbg(p *lang.Process) error {
	return errors.New("Invalid parameters. `bg` only supports a code block on Plan 9 because backgrounding a stopped process is not currently supported on Plan 9.")
}

func cmdForeground(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	return errors.New("This function is currently not supported on Plan 9.")
}

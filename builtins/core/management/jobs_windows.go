// +build windows

package management

import (
	"errors"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func mkbg(p *proc.Process) error {
	return errors.New("Invalid parameters. `bg` only supports a code block on Windows because backgrounding a suspended process is not currently supported on Windows.")
}

func cmdForeground(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	return errors.New("This function is currently not supported on Windows.")
}

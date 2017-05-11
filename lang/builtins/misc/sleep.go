package misc

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"time"
)

func init() {
	proc.GoFunctions["sleep"] = proc.GoFunction{Func: cmdSleep, TypeIn: types.Null, TypeOut: types.Null}
}

func cmdSleep(p *proc.Process) error {
	i, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(int64(i)) * time.Second)

	return nil
}

package misc

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"strconv"
	"time"
)

func init() {
	proc.GoFunctions["sleep"] = proc.GoFunction{Func: cmdSleep, TypeIn: types.Null, TypeOut: types.Null}
}

func cmdSleep(p *proc.Process) (err error) {
	var i int64
	i, err = strconv.ParseInt(string(p.Parameters[0]), 10, 0)
	if err != nil {
		p.Stderr.Writeln([]byte(err.Error()))
	}

	time.Sleep(time.Duration(i) * time.Second)

	return
}

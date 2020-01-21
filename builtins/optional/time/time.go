package time

import (
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["sleep"] = cmdSleep
}

func cmdSleep(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	i, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	select {
	case <-p.Context.Done():
		return nil
	case <-time.After(time.Duration(int64(i)) * time.Second):
		return nil
	}
}

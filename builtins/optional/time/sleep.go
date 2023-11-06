package time

import (
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("sleep", cmdSleep, types.Null)
}

func cmdSleep(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	i, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	sleep := time.After(time.Duration(int64(i)) * time.Second)

	for {
		select {
		case <-p.Context.Done():
			return nil

		case <-sleep:
			return nil

		case <-p.HasStopped:
			<-p.WaitForStopped
		}
	}
}

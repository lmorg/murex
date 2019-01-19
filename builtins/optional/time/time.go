package time

import (
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["sleep"] = cmdSleep
	lang.GoFunctions["time"] = cmdTime
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

func cmdTime(p *lang.Process) error {
	p.Stdout.SetDataType(types.Integer)
	block := p.Parameters.ByteAll()

	if types.IsBlock(block) {
		block, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		start := time.Now()

		p.ExitNum, err = lang.RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stdout, p)
		if err != nil {
			return err
		}

		s := types.FloatToString(time.Now().Sub(start).Seconds())
		_, err = p.Stderr.Write([]byte(s))
		return err
	}

	p.Parameters.Params = append([]string{"time"}, p.Parameters.Params...)
	err := lang.External(p)
	return err
}

package misc

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"time"
)

func init() {
	proc.GoFunctions["sleep"] = proc.GoFunction{Func: cmdSleep, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["time"] = proc.GoFunction{Func: cmdTime, TypeIn: types.Null, TypeOut: types.Integer}
}

func cmdSleep(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)
	i, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(int64(i)) * time.Second)

	return nil
}

func cmdTime(p *proc.Process) error {
	p.Stdout.SetDataType(types.Integer)
	block := p.Parameters.ByteAll()

	if types.IsBlock(block) {
		block, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		start := time.Now()

		p.ExitNum, err = lang.ProcessNewBlock(block, p.Stdin, p.Stdout, p.Stdout, "time")
		if err != nil {
			return err
		}

		s := types.FloatToString(time.Now().Sub(start).Seconds())
		_, err = p.Stderr.Write([]byte(s))
		return err
	}

	p.Parameters.Params = append([]string{"time"}, p.Parameters.Params...)
	err := proc.External(p)
	return err
}

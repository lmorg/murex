package time

import (
	"errors"
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

func cmdTime(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Integer)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing parameters")
	}

	block := p.Parameters.StringAll()

	start := time.Now()

	//p.ExitNum, err = lang.RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stdout, p)
	p.ExitNum, err = p.Fork(lang.F_DEFAULTS).Execute([]rune(block))
	if err != nil {
		return
	}

	s := types.FloatToString(time.Now().Sub(start).Seconds())
	_, err = p.Stderr.Write([]byte(s))
	return
}

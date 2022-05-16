package time

import (
	"errors"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("time", cmdTime, types.Any)
}

func cmdTime(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Integer)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing parameters")
	}

	block := p.Parameters.StringAll()
	start := time.Now()

	p.ExitNum, err = p.Fork(lang.F_DEFAULTS).Execute([]rune(block))
	if err != nil {
		return
	}

	s := types.FloatToString(time.Now().Sub(start).Seconds())
	_, err = p.Stderr.Write([]byte(s))
	return
}

package encoders

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"strconv"
)

func init() {
	proc.GoFunctions["escape"] = proc.GoFunction{Func: cmdEscape, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["!escape"] = proc.GoFunction{Func: cmdEscape, TypeIn: types.String, TypeOut: types.String}
}

func cmdEscape(p *proc.Process) (err error) {
	var str string
	if p.Parameters.Len() == 0 {
		str = string(p.Stdin.ReadAll())
		//debug.Log("[cmdEscape] [p.Stdin.ReadAll]", str)
	} else {
		str = p.Parameters.AllString()
		//debug.Log("[cmdEscape] [len(p.Parameters)]", len(p.Parameters), "'"+p.Parameters.String()+"'")
	}

	if p.Not {
		str, err = strconv.Unquote(str)
		if err != nil {
			return
		}
	} else {
		str = strconv.Quote(str)
	}

	p.Stdout.Write([]byte(str))

	return nil
}

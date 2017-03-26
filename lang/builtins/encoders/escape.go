package encoders

import (
	"html"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["escape"] = proc.GoFunction{Func: cmdEscape, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["escape!"] = proc.GoFunction{Func: cmdEscape, TypeIn: types.String, TypeOut: types.String}
}

func cmdEscape(p *proc.Process) error {
	var str string
	if p.Parameters.Len() == 0 {
		str = string(p.Stdin.ReadAll())
		//debug.Log("[cmdEscape] [p.Stdin.ReadAll]", str)
	} else {
		str = p.Parameters.AllString()
		//debug.Log("[cmdEscape] [len(p.Parameters)]", len(p.Parameters), "'"+p.Parameters.String()+"'")
	}

	if p.Not {
		str = html.UnescapeString(str)
	} else {
		str = html.EscapeString(str)
	}

	p.Stdout.Write([]byte(str))

	return nil
}

package encoders

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"html"
	"strconv"
)

func init() {
	proc.GoFunctions["escape"] = cmdEscape
	proc.GoFunctions["!escape"] = cmdEscape
}

func cmdEscape(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)
	var str string
	if p.Parameters.Len() == 0 {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		str = string(b)

	} else {
		str = p.Parameters.StringAll()

	}

	if p.IsNot {
		unescape, err := strconv.Unquote(str)
		if err != nil {
			unescape = html.UnescapeString(str)
		}
		str = unescape

	} else {
		str = strconv.Quote(str)
	}

	p.Stdout.Write([]byte(str))

	return nil
}

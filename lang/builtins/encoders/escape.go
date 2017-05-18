package encoders

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"html"
	"strconv"
)

func init() {
	lang.GoFunctions["escape"] = lang.GoFunction{Func: cmdEscape, TypeIn: types.String, TypeOut: types.String}
	lang.GoFunctions["!escape"] = lang.GoFunction{Func: cmdEscape, TypeIn: types.String, TypeOut: types.String}
}

func cmdEscape(p *lang.Process) error {
	var str string
	if p.Parameters.Len() == 0 {
		str = string(p.Stdin.ReadAll())

	} else {
		str = p.Parameters.StringAll()

	}

	if p.Not {
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

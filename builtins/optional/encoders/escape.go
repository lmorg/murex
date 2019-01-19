package encoders

import (
	"html"
	"net/url"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["escape"] = cmdEscape
	lang.GoFunctions["!escape"] = cmdEscape
	lang.GoFunctions["htmlesc"] = cmdHtmlEscape
	lang.GoFunctions["!htmlesc"] = cmdHtmlEscape
	lang.GoFunctions["urlesc"] = cmdUrlEscape
	lang.GoFunctions["!urlesc"] = cmdUrlEscape
}

func cmdEscape(p *lang.Process) error {
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

	_, err := p.Stdout.Write([]byte(str))
	return err
}

func cmdHtmlEscape(p *lang.Process) error {
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
		str = html.UnescapeString(str)

	} else {
		str = html.EscapeString(str)
	}

	_, err := p.Stdout.Write([]byte(str))
	return err
}

func cmdUrlEscape(p *lang.Process) (err error) {
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
		str, err = url.PathUnescape(str)
		if err != nil {
			return err
		}

	} else {
		str = url.PathEscape(str)
	}

	_, err = p.Stdout.Write([]byte(str))
	return err
}

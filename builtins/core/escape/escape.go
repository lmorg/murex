package escape

import (
	"html"
	"net/url"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/escape"
)

func init() {
	lang.GoFunctions["escape"] = cmdEscape
	lang.GoFunctions["!escape"] = cmdEscape
	lang.GoFunctions["eschtml"] = cmdHtml
	lang.GoFunctions["!eschtml"] = cmdHtml
	lang.GoFunctions["escurl"] = cmdUrl
	lang.GoFunctions["!escurl"] = cmdUrl
	lang.GoFunctions["esccli"] = cmdEscapeCli
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

func cmdHtml(p *lang.Process) error {
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

func cmdUrl(p *lang.Process) (err error) {
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

func cmdEscapeCli(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	var s []string

	if p.IsMethod {
		err := p.Stdin.ReadArray(func(b []byte) {
			s = append(s, string(b))
		})
		if err != nil {
			return err
		}
	} else {
		s = p.Parameters.StringArray()
	}

	escape.CommandLine(s)

	_, err := p.Stdout.Writeln([]byte(strings.Join(s, " ")))
	return err
}

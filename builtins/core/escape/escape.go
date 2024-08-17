package escape

import (
	"html"
	"net/url"
	"strconv"
	"strings"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/escape"
)

func init() {
	lang.DefineMethod("escape", cmdEscape, types.Text, types.String)
	lang.DefineMethod("!escape", cmdEscape, types.Text, types.String)
	lang.DefineMethod("eschtml", cmdHtml, types.Text, types.String)
	lang.DefineMethod("!eschtml", cmdHtml, types.Text, types.String)
	lang.DefineMethod("escurl", cmdUrl, types.Text, types.String)
	lang.DefineMethod("!escurl", cmdUrl, types.Text, types.String)
	lang.DefineMethod("esccli", cmdEscapeCli, types.Text, types.String)

	defaults.AppendProfile(`
		alias  escape.quote =  escape
		alias !escape.quote = !escape
		alias  escape.html  =  eschtml
		alias !escape.html  = !eschtml
		alias  escape.url   =  escurl
		alias !escape.url   = !escurl
		alias  escape.cli   =  esccli
	`)
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
		err := p.Stdin.ReadArray(p.Context, func(b []byte) {
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

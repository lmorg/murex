package lists

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineMethod("list.case", cmdCase, types.ReadArray, types.WriteArray)

	defaults.AppendProfile(`
		autocomplete set list.case %[
			{
				DynamicDesc: '{ list.case help }'
			}
		]
	`)
}

const (
	_CASE_UPPER      = "upper"
	_CASE_LOWER      = "lower"
	_CASE_TITLE      = "title"
	_CASE_TITLE_CAPS = "title+caps"
	_CASE_HELP       = "help"
)

var caseActions = map[string]string{
	_CASE_UPPER:      "Convert all elements in list to uppercase",
	_CASE_LOWER:      "Convert all elements in list to lowercase",
	_CASE_TITLE:      "Capitalize the first character of each word",
	_CASE_TITLE_CAPS: "Capitalize the first character of each word, do not lowercase characters already in uppercase",
}

func caseTitle(s string) string {
	return cases.Title(language.English).String(s)
}

func caseTitlePlusCaps(s string) string {
	return cases.Title(language.English, cases.NoLower).String(s)
}

func cmdCase(p *lang.Process) error {
	action, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	var fn func(string) string

	switch action {
	case _CASE_UPPER:
		fn = strings.ToUpper

	case _CASE_LOWER:
		fn = strings.ToLower

	case _CASE_TITLE:
		fn = caseTitle

	case _CASE_TITLE_CAPS:
		fn = caseTitlePlusCaps

	case _CASE_HELP:
		return caseHelp(p)

	default:
		return fmt.Errorf("invalid action `%s` -- run `%s %s` for usage",
			action, p.Name.String(), _CASE_HELP)
	}

	if p.IsMethod {
		return caseMethod(p, fn)
	}

	return caseFunction(p, fn)
}

func caseMethod(p *lang.Process, fn func(string) string) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)
	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	err = p.Stdin.ReadArray(p.Context, func(b []byte) {
		s := fn(string(b))
		err := aw.WriteString(s)
		if err != nil {
			p.Done()
			_, err = p.Stderr.Writeln([]byte(fmt.Sprintf("error: %s", err.Error())))
			p.ExitNum = 2
			if err != nil {
				panic(err)
			}
		}
	})
	if err != nil {
		return err
	}

	return aw.Close()
}

func caseFunction(p *lang.Process, fn func(string) string) error {
	p.Stdout.SetDataType(types.Json)

	slice := p.Parameters.StringArray()[1:]

	for i := range slice {
		slice[i] = fn(slice[i])
	}

	b, err := json.Marshal(slice, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func caseHelp(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := json.Marshal(caseActions, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

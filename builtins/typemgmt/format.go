package typemgmt

import (
	"bufio"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	proc.GoFunctions["table"] = proc.GoFunction{Func: cmdTable, TypeIn: types.Generic, TypeOut: types.Csv}
	proc.GoFunctions["format"] = proc.GoFunction{Func: cmdFormat, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdTable(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Csv)

	separator, err := p.Parameters.String(0)
	if err != nil {
		return
	}

	var (
		a []string
		s string
	)

	join := func(b []byte) {
		a = append(a, string(b))
	}

	if p.IsMethod {
		p.Stdin.ReadArray(join)
		s = strings.Join(a, separator)
	} else {
		s = strings.Join(p.Parameters.StringArray()[1:], string(separator))
	}

	_, err = p.Stdout.Writeln([]byte(s))
	return
}

func cmdFormat(p *proc.Process) (err error) {
	format, err := p.Parameters.String(0)
	if err != nil {
		return
	}

	p.Stdout.SetDataType(format)

	dt := p.Stdin.GetDataType()

	switch dt {
	case types.String, types.Generic:
		return fStringGeneric(p, format)
	case types.Csv:
		return fCsv(p, format)
	}

	return errors.New("I don't know how to convert this data")
}

func fStringGeneric(p *proc.Process, format string) (err error) {
	inSep, _ := p.Parameters.String(1)
	outSep, _ := p.Parameters.String(2)

	if inSep == "" {
		inSep = `[\s][\s]+`
	}

	if outSep == "" {
		iface, _ := proc.GlobalConf.Get("shell", "Csv-Separator", types.String)
		outSep = iface.(string)
	}

	rxWhiteSpaceSplit, err := regexp.Compile(inSep)
	if err != nil {
		return err
	}

	switch format {
	case types.Json:
		var (
			headings []string
			jObj     []map[string]string
		)

		scanner := bufio.NewScanner(p.Stdin)
		for scanner.Scan() {
			s := scanner.Text()
			split := rxWhiteSpaceSplit.Split(s, -1)
			if len(headings) == 0 {
				headings = split
			} else {
				m := make(map[string]string)
				for i := range split {
					if i >= len(headings) {
						m[strconv.Itoa(i)] = split[i]
					} else {
						m[headings[i]] = split[i]
					}
				}
				jObj = append(jObj, m)
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}

		b, err := utils.JsonMarshal(jObj)
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err

	case types.Csv:
		scanner := bufio.NewScanner(p.Stdin)
		for scanner.Scan() {
			s := scanner.Text()
			s = strings.Replace(s, `\`, `\\`, -1)
			s = strings.Replace(s, `"`, `\"`, -1)
			s = `"` + rxWhiteSpaceSplit.ReplaceAllString(s, `"`+outSep+`"`) + `"`
			p.Stdout.Writeln([]byte(s))
		}
		if err := scanner.Err(); err != nil {
			return err
		}
		return
	}

	return errors.New("I don't know how to convert this data")
}

func fCsv(p *proc.Process, format string) (err error) {
	outSep, _ := p.Parameters.String(1)

	if outSep == "" {
		iface, _ := proc.GlobalConf.Get("shell", "Csv-Separator", types.String)
		outSep = iface.(string)
	}

	switch format {
	case types.Json:
		var (
			a []string
			s string
		)

		join := func(b []byte) {
			a = append(a, string(b))
		}

		p.Stdin.ReadArray(join)
		s = strings.Join(a, outSep)

		_, err = p.Stdout.Writeln([]byte(s))
		return
	}

	return errors.New("I don't know how to convert this data")
}

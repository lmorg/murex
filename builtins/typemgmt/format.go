package typemgmt

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/csv"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	proc.GoFunctions["table"] = proc.GoFunction{Func: cmdTable, TypeIn: types.Generic, TypeOut: types.Csv}
	proc.GoFunctions["format"] = proc.GoFunction{Func: cmdFormat, TypeIn: types.Generic, TypeOut: types.Generic}
}

const iDontKnow = "I don't know how to convert %s into %s."

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
		return fStringGeneric(p, dt, format)
	case types.Csv:
		return fCsv(p, dt, format)
	case types.Json:
		return fJson(p, dt, format)
	}

	return errors.New(fmt.Sprintf(iDontKnow, dt, format))
}

func fStringGeneric(p *proc.Process, dt, format string) error {
	inSep, _ := p.Parameters.String(1)
	outSep, _ := p.Parameters.String(2)

	if inSep == "" {
		//inSep = `[\s][\s]+`
		inSep = `\s+`
	}

	if outSep == "" {
		iface, _ := proc.GlobalConf.Get("shell", "Csv-Separator", types.String)
		outSep = iface.(string)
	}

	rxSplit, err := regexp.Compile(inSep)
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
			split := rxSplit.Split(s, -1)
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
		parser, err := csv.NewParser(nil, &proc.GlobalConf)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(p.Stdin)
		for scanner.Scan() {
			s := scanner.Text()
			split := rxSplit.Split(s, -1)
			b := parser.ArrayToCsv(split)
			p.Stdout.Writeln(b)
		}
		if err := scanner.Err(); err != nil {
			return err
		}
		return nil
	}

	return errors.New(fmt.Sprintf(iDontKnow, dt, format))
}

func fJson(p *proc.Process, dt, format string) error {
	switch format {
	case types.Csv:
		var a []string

		p.Stdin.ReadArray(func(b []byte) {
			a = append(a, string(b))
		})

		csvParser, err := csv.NewParser(nil, &proc.GlobalConf)
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(csvParser.ArrayToCsv(a))
		return err

	case types.String:
		separator, _ := p.Parameters.String(1)

		var (
			a []string
			s string
		)

		p.Stdin.ReadArray(func(b []byte) {
			a = append(a, string(b))
		})

		s = strings.Join(a, separator)

		_, err := p.Stdout.Writeln([]byte(s))
		return err
	}

	return errors.New(fmt.Sprintf(iDontKnow, dt, format))
}

func fCsv(p *proc.Process, dt, format string) error {
	switch format {
	case types.Json:
		csvParser, err := csv.NewParser(p.Stdin, &proc.GlobalConf)
		if err != nil {
			return err
		}

		var jObj []map[string]string
		csvParser.ReadLine(func(records []string, headings []string) {
			obj := make(map[string]string)
			for i := range records {
				obj[headings[i]] = records[i]
			}
			jObj = append(jObj, obj)
		})

		b, err := utils.JsonMarshal(jObj)
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err
	}

	return errors.New(fmt.Sprintf(iDontKnow, dt, format))
}

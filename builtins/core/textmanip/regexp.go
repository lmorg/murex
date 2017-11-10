package textmanip

import (
	"bytes"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
	"regexp"
	"strings"
)

func init() {
	proc.GoFunctions["match"] = cmdMatch
	proc.GoFunctions["!match"] = cmdMatch
	proc.GoFunctions["regexp"] = cmdRegexp
	proc.GoFunctions["!regexp"] = cmdRegexp
}

func cmdMatch(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.StringAll() == "" {
		return errors.New("No parameters supplied.")
	}

	var output []string

	p.Stdin.ReadArray(func(b []byte) {
		matched := bytes.Contains(b, p.Parameters.ByteAll())
		if (matched && !p.IsNot) || (!matched && p.IsNot) {
			output = append(output, string(b))
		}
	})

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdRegexp(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.StringAll() == "" {
		return errors.New("No parameters supplied.")
	}

	var sRegex []string
	if p.Parameters.Len() == 1 {
		sRegex = splitRegexParams(p.Parameters.StringAll())
	} else {
		sRegex = p.Parameters.StringArray()
	}

	var rx *regexp.Regexp
	if len(sRegex) < 2 || len(sRegex) > 4 {
		return errors.New("Invalid regexp.")
	}
	if rx, err = regexp.Compile(sRegex[1]); err != nil {
		return
	}

	switch sRegex[0][0] {
	case 'm': // match
		var output []string

		p.Stdin.ReadArray(func(b []byte) {
			matched := rx.Match(b)
			if (matched && !p.IsNot) || (!matched && p.IsNot) {
				output = append(output, string(b))
			}
		})

		b, err := define.MarshalData(p, dt, output)
		if err != nil {
			return err
		}

		_, err = p.Stdout.Write(b)
		return err

	case 's': // substitute
		var output []string

		p.Stdin.ReadArray(func(b []byte) {
			output = append(output, rx.ReplaceAllString(string(b), sRegex[2]))
		})

		b, err := define.MarshalData(p, dt, output)
		if err != nil {
			return err
		}

		_, err = p.Stdout.Write(b)
		return err

	case 'f': // find
		var output [][]string

		p.Stdin.ReadArray(func(b []byte) {
			found := rx.FindStringSubmatch(string(b))
			if len(found) > 0 {
				output = append(output, found)
			}
		})

		b, err := define.MarshalData(p, dt, output)
		if err != nil {
			return err
		}

		_, err = p.Stdout.Write(b)
		return err

	default:
		return errors.New("Invalid regexp. Please use either match (m), substitute (s) or find (f).")
	}
}

func splitRegexParams(s string) (regex []string) {
	if len(s) < 2 {
		return
	}
	switch s[1] {
	case '/':
		regex = strings.Split(s, "/")
	case '#':
		regex = strings.Split(s, "#")
	case ',':
		regex = strings.Split(s, ",")

		//case '{':
		//	b = append([]byte{'}'}, b...)
		//	b = append(b, '{')
		//	b = bytes.Split(b, []byte{'}', '{'})
	}
	return
}

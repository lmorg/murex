package textmanip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"regexp"
	"strings"
)

func init() {
	proc.GoFunctions["match"] = proc.GoFunction{Func: cmdMatch, TypeIn: types.Generic, TypeOut: types.String}
	proc.GoFunctions["!match"] = proc.GoFunction{Func: cmdMatch, TypeIn: types.Generic, TypeOut: types.String}
	proc.GoFunctions["regex"] = proc.GoFunction{Func: cmdRegexp, TypeIn: types.Generic, TypeOut: types.String}
	proc.GoFunctions["!regex"] = proc.GoFunction{Func: cmdRegexp, TypeIn: types.Generic, TypeOut: types.String}
	proc.GoFunctions["left"] = proc.GoFunction{Func: cmdLeft, TypeIn: types.Generic, TypeOut: types.String}
	proc.GoFunctions["right"] = proc.GoFunction{Func: cmdRight, TypeIn: types.Generic, TypeOut: types.String}
	proc.GoFunctions["append"] = proc.GoFunction{Func: cmdAppend, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["prepend"] = proc.GoFunction{Func: cmdPrepend, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["pretty"] = proc.GoFunction{Func: cmdPretty, TypeIn: types.Json, TypeOut: types.String}
	proc.GoFunctions["sprintf"] = proc.GoFunction{Func: cmdSprintf, TypeIn: types.Generic, TypeOut: types.String}
}

func cmdMatch(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)

	if p.Parameters.StringAll() == "" {
		return errors.New("No parameters supplied.")
	}

	p.Stdin.ReadArray(func(b []byte) {
		matched := bytes.Contains(b, p.Parameters.ByteAll())
		if (matched && !p.IsNot) || (!matched && p.IsNot) {
			p.Stdout.Writeln(b)
		}
	})

	return nil
}

func cmdRegexp(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)

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
		//if len(sRegex) != 2 {
		//	return errors.New("Invalid regexp.")
		//}
		p.Stdin.ReadArray(func(b []byte) {
			matched := rx.Match(b)
			if (matched && !p.IsNot) || (!matched && p.IsNot) {
				p.Stdout.Writeln(b)
			}
		})

	case 's': // substitute
		//if len(sRegex) != 3 {
		//	return errors.New("Invalid regexp.")
		//}
		p.Stdin.ReadArray(func(b []byte) {
			p.Stdout.Writeln(rx.ReplaceAll(b, []byte(sRegex[2])))
		})

	case 'f': // match
		//if len(sRegex) != 2 {
		//	return errors.New("Invalid regexp.")
		//}
		p.Stdin.ReadArray(func(b []byte) {
			found := rx.Find(b)
			if len(found) != 0 {
				p.Stdout.Writeln(found)
			}
			//debug.Log("[cmdRegexp] [line]", string(b), string(found))
			//if (matched && !p.Not) || (!matched && p.Not) {
			//}
		})

	default:
		return errors.New("Invalid regexp. Please use either match (m), substitute (s) or find (f).")
	}

	return
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

func cmdLeft(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)

	left, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(func(b []byte) {
		if len(b) < left {
			_, err = p.Stdout.Writeln(b)
		} else {
			_, err = p.Stdout.Writeln(b[:left])
		}
	})

	return err
}

func cmdRight(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)

	right, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(func(b []byte) {
		if len(b) < right {
			_, err = p.Stdout.Writeln(b)
		} else {
			_, err = p.Stdout.Writeln(b[len(b)-right:])
		}
	})

	return err
}

func cmdPrepend(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)

	prepend := p.Parameters.ByteAll()
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(append(prepend, b...))

	return
}

func cmdAppend(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)

	b := p.Parameters.ByteAll()
	text, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	if len(text) > 1 && text[len(text)-1] == '\n' {
		text = text[:len(text)-1]
	}
	if len(text) > 1 && text[len(text)-1] == '\r' {
		text = text[:len(text)-1]
	}
	_, err = p.Stdout.Write(append(text, b...))

	return
}

func cmdPretty(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, b, "", "\t")
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(prettyJSON.Bytes())
	return err
}

func cmdSprintf(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)

	if !p.IsMethod {
		return errors.New("I must be called as a method.")
	}

	if p.Parameters.Len() == 0 {
		return errors.New("Parameters missing.")
	}

	s := p.Parameters.StringAll()
	var a []interface{}

	err := p.Stdin.ReadArray(func(b []byte) {
		a = append(a, string(b))
	})

	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(fmt.Sprintf(s, a...)))
	return err
}

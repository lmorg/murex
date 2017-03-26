package textmanip

import (
	"bytes"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"regexp"
	"strings"
)

func init() {
	proc.GoFunctions["match"] = proc.GoFunction{Func: cmdMatch, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["!match"] = proc.GoFunction{Func: cmdMatch, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["regex"] = proc.GoFunction{Func: cmdRegexp, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["!regex"] = proc.GoFunction{Func: cmdRegexp, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["left"] = proc.GoFunction{Func: cmdLeft, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["right"] = proc.GoFunction{Func: cmdRight, TypeIn: types.String, TypeOut: types.String}
	proc.GoFunctions["prepend"] = proc.GoFunction{Func: cmdPrepend, TypeIn: types.String, TypeOut: types.String}
}

func cmdMatch(p *proc.Process) error {
	if p.Parameters.AllString() == "" {
		return errors.New("No parameters supplied.")
	}

	p.Stdin.ReadLineFunc(func(b []byte) {
		matched := bytes.Contains(b, p.Parameters.AllByte())
		if (matched && !p.Not) || (!matched && p.Not) {
			p.Stdout.Write(b)
		}
	})

	return nil
}

func cmdRegexp(p *proc.Process) (err error) {
	if p.Parameters.AllString() == "" {
		return errors.New("No parameters supplied.")
	}

	var sRegex []string
	if len(p.Parameters) == 1 {
		sRegex = splitRegexParams(p.Parameters.AllString())
	} else {
		sRegex = p.Parameters
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
		p.Stdin.ReadLineFunc(func(b []byte) {
			matched := rx.Match(b)
			if (matched && !p.Not) || (!matched && p.Not) {
				p.Stdout.Write(b)
			}
		})

	case 's': // substitute
		//if len(sRegex) != 3 {
		//	return errors.New("Invalid regexp.")
		//}
		p.Stdin.ReadLineFunc(func(b []byte) {
			p.Stdout.Write(rx.ReplaceAll(b, []byte(sRegex[2])))
		})

	case 'f': // match
		//if len(sRegex) != 2 {
		//	return errors.New("Invalid regexp.")
		//}
		p.Stdin.ReadLineFunc(func(b []byte) {
			found := rx.Find(b)
			if len(found) != 0 {
				p.Stdout.Writeln(found)
			}
			//debug.Log("[cmdRegexp] [line]", string(b), string(found))
			//if (matched && !p.Not) || (!matched && p.Not) {
			//}
		})

	default:
		return errors.New("Invalid regexp. Please use either match (m), substitude (s) or find (f).")
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
	left, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	p.Stdin.ReadLineFunc(func(b []byte) {
		if len(b) < left {
			_, err = p.Stdout.Write(b)
		} else {
			_, err = p.Stdout.Write(b[:left])
		}
	})

	return err
}

func cmdRight(p *proc.Process) error {
	right, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	p.Stdin.ReadLineFunc(func(b []byte) {
		if len(b) < right {
			_, err = p.Stdout.Write(b)
		} else {
			_, err = p.Stdout.Write(b[len(b)-right:])
		}
	})

	return err
}

func cmdPrepend(p *proc.Process) (err error) {
	prepend := p.Parameters.AllByte()
	_, err = p.Stdout.Write(append(prepend, p.Stdin.ReadAll()...))

	return
}

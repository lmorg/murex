package lists

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("match", cmdMatch, types.ReadArray, types.WriteArray)
	lang.DefineMethod("!match", cmdMatch, types.ReadArray, types.WriteArray)
	lang.DefineMethod("regexp", cmdRegexp, types.ReadArray, types.WriteArray)
	lang.DefineMethod("!regexp", cmdRegexp, types.ReadArray, types.WriteArray)
}

func cmdMatch(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.StringAll() == "" {
		return errors.New("no parameters supplied")
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(p.Context, func(b []byte) {
		matched := bytes.Contains(b, p.Parameters.ByteAll())
		if (matched && !p.IsNot) || (!matched && p.IsNot) {
			err = aw.Write(b)
			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}
		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}

func cmdRegexp(p *lang.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.StringAll() == "" {
		return errors.New("no parameters supplied")
	}

	var sRegex []string
	if p.Parameters.Len() == 1 {
		sRegex, err = splitRegexParams(p.Parameters.ByteAll())
		if err != nil {
			return err
		}

	} else {
		// No need to get clever with the regex parser because the parameters are already split by murex's parser
		sRegex = p.Parameters.StringArray()
	}

	if len(sRegex) < 2 {
		return fmt.Errorf("invalid regexp (too few parameters) in: `%s`", p.Parameters.StringAll())
	}
	if len(sRegex) > 4 {
		return fmt.Errorf("invalid regexp (too many parameters) in: `%s`", p.Parameters.StringAll())
	}

	var rx *regexp.Regexp
	if rx, err = regexp.Compile(sRegex[1]); err != nil {
		return
	}

	switch sRegex[0][0] {
	case 'm':
		return regexMatch(p, rx, dt)

	case 's':
		if p.IsNot {
			return fmt.Errorf("cannot use `%s` with `%s` flag in `%s`", p.Name.String(), string(sRegex[0][0]), p.Parameters.StringAll())
		}
		return regexSubstitute(p, rx, sRegex, dt)

	case 'f':
		if p.IsNot {
			return fmt.Errorf("cannot use `%s` with `%s` flag in `%s`", p.Name.String(), string(sRegex[0][0]), p.Parameters.StringAll())
		}
		return regexFind(p, rx, dt)

	default:
		return errors.New("invalid regexp. Please use either match (m), substitute (s) or find (f)")
	}
}

func splitRegexParams(regex []byte) ([]string, error) {
	if len(regex) < 2 {
		return nil, fmt.Errorf("invalid regexp (too few characters) in: `%s`", string(regex))
	}

	switch regex[1] {
	default:
		return splitRegexDefault(regex)

	case '{':
		return nil, fmt.Errorf("the `{` character is not supported for separating regex parameters in: `%s`", string(regex))
		//return splitRegexBraces(regex)

	case '\\':
		return nil, fmt.Errorf("the `\\` character is not valid for separating regex parameters in: `%s`", string(regex))
	}
}

func splitRegexDefault(regex []byte) (s []string, _ error) {
	var (
		param   []byte
		escaped bool
		token   = regex[1]
	)

	for _, c := range regex {
		switch c {
		default:
			if escaped {
				param = append(param, '\\', c)
				escaped = false
				continue
			}
			param = append(param, c)

		case '\\':
			if escaped {
				param = append(param, '\\', c)
				escaped = false
				continue
			}
			escaped = true

		case token:
			if escaped {
				escaped = false
				param = append(param, c)
				continue
			}

			s = append(s, string(param))
			param = []byte{}
		}
	}
	s = append(s, string(param))

	return
}

var rxCurlyBraceSplit = regexp.MustCompile(`\{(.*?)\}`)

func splitRegexBraces(regex []byte) ([]string, error) {
	s := rxCurlyBraceSplit.FindAllString(string(regex), -1)
	s = append([]string{string(regex[0])}, s...)
	debug.Json("s", s)
	return s, nil
}

// -------- regex functions --------

func regexMatch(p *lang.Process, rx *regexp.Regexp, dt string) error {
	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(p.Context, func(b []byte) {
		matched := rx.Match(b)
		if (matched && !p.IsNot) || (!matched && p.IsNot) {

			err = aw.Write(b)
			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}

		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}

func regexSubstitute(p *lang.Process, rx *regexp.Regexp, sRegex []string, dt string) error {
	if len(sRegex) < 3 {
		return fmt.Errorf("invalid regex: too few parameters\nexpecting s/find/substitute/ in: `%s`", p.Parameters.StringAll())
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	sub := []byte(sRegex[2])

	p.Stdin.ReadArray(p.Context, func(b []byte) {
		err = aw.Write(rx.ReplaceAll(b, sub))
		if err != nil {
			p.Stdin.ForceClose()
			p.Done()
		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}

func regexFind(p *lang.Process, rx *regexp.Regexp, dt string) error {
	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(p.Context, func(b []byte) {
		match := rx.FindAllStringSubmatch(string(b), -1)
		for _, found := range match {
			if len(found) > 1 {

				for i := 1; i < len(found); i++ {
					err = aw.WriteString(found[i])
					if err != nil {
						p.Stdin.ForceClose()
						p.Done()
					}

				}

			}
		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}

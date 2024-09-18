package lists

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("regexp", cmdRegexp, types.ReadArray, types.WriteArray)
	lang.DefineMethod("!regexp", cmdRegexp, types.ReadArray, types.WriteArray)
}

const (
	_WITH_HEADING    = true
	_WITHOUT_HEADING = false
)

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
		sRegex, err = splitRegexParams(p.Parameters.RuneAll())
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
		return regexMatch(p, rx, dt, _WITHOUT_HEADING)

	case 'M':
		return regexMatch(p, rx, dt, _WITH_HEADING)

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

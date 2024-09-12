package typemgmt

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

var (
	rxLet    = regexp.MustCompile(`^([_a-zA-Z0-9]+)\s*=(.*)$`)
	rxItself = regexp.MustCompile(`^([_a-zA-Z0-9]+)\s*([\-\*\+/])\s*=\s*(.*)$`)
	rxMinus  = regexp.MustCompile(`^([_a-zA-Z0-9]+)--$`)
	rxPlus   = regexp.MustCompile(`^([_a-zA-Z0-9]+)\+\+$`)
)

func init() {
	lang.DefineMethod("let", cmdLet, types.Math, types.Null)
}

func cmdLet(p *lang.Process) (err error) {
	lang.FeatureDeprecatedBuiltin(p)

	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				p.ExitNum = 2
				err = errors.New(fmt.Sprint("Panic caught: ", r))
			}
		}()
	}

	params := p.Parameters.StringAll()

	variable, expression, err := letBuilder(params)
	if err != nil {
		return err
	}

	value, dt, err := evaluate(p, expression)
	if err != nil {
		return err
	}

	err = p.Variables.Set(p, variable, value, dt)
	return err
}

func letBuilder(params string) (variable, expression string, err error) {
	switch {
	case rxLet.MatchString(params):
		match := rxLet.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = match[0][2]

	case rxItself.MatchString(params):
		match := rxItself.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = variable + match[0][2] + match[0][3]

	case rxPlus.MatchString(params):
		match := rxPlus.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = variable + "+1"

	case rxMinus.MatchString(params):
		match := rxMinus.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = variable + "-1"

	default:
		err = errors.New("invalid syntax for `let`. Should be `let variable-name = expression`")
	}

	return
}

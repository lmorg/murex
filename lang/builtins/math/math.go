package math

import (
	"errors"
	"fmt"
	gov "github.com/Knetic/govaluate"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"regexp"
	"strings"
)

func init() {
	proc.GoFunctions["eval"] = proc.GoFunction{Func: cmdEval, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["let"] = proc.GoFunction{Func: cmdLet, TypeIn: types.Null, TypeOut: types.Boolean}
}

var rxLet *regexp.Regexp = regexp.MustCompile(`^([-_a-zA-Z0-9]+)\s*=`)

func cmdLet(p *proc.Process) error {
	params := p.Parameters.AllString()

	if !rxLet.MatchString(params) {
		return errors.New("Invalid syntax for `let`. Should be `let variable-name = expression`.")
	}

	params = rxLet.ReplaceAllString(params, "")

	return cmdEval(p)
}

func cmdEval(p *proc.Process) (err error) {
	defer func() {
		if r := recover(); r != nil {
			p.ExitNum = 2
			err = errors.New(fmt.Sprint("Panic caught: ", r))
			//msg := fmt.Sprint("Panic caught: ", r)
			//p.Stderr.Writeln([]byte(msg))
		}
	}()

	expression, err := gov.NewEvaluableExpression(p.Parameters.AllString())
	if err != nil {
		return err
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		return err
	}

	cgt, err := types.ConvertGoType(result, types.String)
	if err != nil {
		return err
	}

	s := strings.Replace(cgt.(string), ".000000", "", 1) // TODO: this is hacky!!!
	_, err = p.Stdout.Write([]byte(s))
	return err
}

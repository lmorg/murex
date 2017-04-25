package math

import (
	"fmt"
	gov "github.com/Knetic/govaluate"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"strings"
)

func init() {
	proc.GoFunctions["eval"] = proc.GoFunction{Func: cmdEval, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdEval(p *proc.Process) error {
	defer func() {
		if r := recover(); r != nil {
			p.ExitNum = 2
			msg := fmt.Sprint("Panic caught: ", r)
			p.Stderr.Writeln([]byte(msg))
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

	v, err := types.ConvertGoType(result, types.String)
	if err != nil {
		return err
	}

	s := strings.Replace(v.(string), ".000000", "", 1) // TODO: this is hacky!!!
	_, err = p.Stdout.Write([]byte(s))
	return err
}

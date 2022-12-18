package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction(lang.ExpressionFunctionName, cmdExpressions, types.Any)
}

func cmdExpressions(p *lang.Process) error {
	expression := []rune(p.Parameters.StringAll())

	result, err := expressions.ExecuteExpr(p, expression)
	if err != nil {
		return err
	}

	dt := result.DataType()
	p.Stdout.SetDataType(dt)

	b, err := lang.MarshalData(p, dt, result.Value)
	if err != nil {
		return err
	}

	if dt == types.Boolean && !types.IsTrue(b, 0) {
		p.ExitNum = 1
	}

	_, err = p.Stdout.Write(b)
	return err
}

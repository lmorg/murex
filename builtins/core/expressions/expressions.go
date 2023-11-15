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

	val, err := result.GetValue()
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(val.DataType)

	if val.Value == nil && val.DataType == types.Json {
		_, err = p.Stdout.Write([]byte{'n', 'u', 'l', 'l'})
		return err
	}

	var b []byte

	switch val.Value.(type) {
	case string:
		b = []byte(val.Value.(string))
	default:
		b, err = lang.MarshalData(p, val.DataType, val.Value)
		if err != nil {
			return err
		}
	}

	if val.DataType == types.Boolean {
		if !types.IsTrue(b, 0) {
			p.ExitNum = 1
		}
		if p.Next.OperatorLogicAnd || p.Next.OperatorLogicOr {
			return nil
		}
	}

	_, err = p.Stdout.Write(b)
	return err
}

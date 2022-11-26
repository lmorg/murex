package expressions

import (
	"github.com/lmorg/murex/builtins/core/expressions/primitives"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("exp", cmdExpressions, types.Any)
}

func Execute(p *lang.Process, expression []byte) (*primitives.DataType, error) {
	tree := newExpTree(expression)
	tree.p = p
	err := tree.parse()
	if err != nil {
		return nil, err
	}

	return tree.execute()
}

func cmdExpressions(p *lang.Process) error {
	result, err := Execute(p, p.Parameters.ByteAll())
	if err != nil {
		return err
	}

	dt := result.DataType()
	p.Stdout.SetDataType(dt)

	b, err := lang.MarshalData(p, dt, result.Value)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

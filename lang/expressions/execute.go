package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
)

func ExecuteExpr(p *lang.Process, expression []rune) (*primitives.DataType, error) {
	tree := newExpTree(p, expression)

	err := tree.parse(true)
	if err != nil {
		return nil, err
	}

	return tree.executeExpr()
}

func (tree *expTreeT) ParseStatement(exec bool) error {
	tree.statement = new(StatementT)
	tree.charPos = 0

	err := tree.parseStatement(exec)
	if err != nil {
		return err
	}

	if !exec {
		return tree.statement.Validate()
	}

	return nil
}

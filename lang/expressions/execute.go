package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
)

func ExecuteExpr(p *lang.Process, expression []rune) (*primitives.DataType, error) {
	tree := NewParser(p, expression, 0)

	err := tree.parseExpression(true)
	if err != nil {
		return nil, err
	}

	return tree.executeExpr()
}

func (tree *ParserT) ParseStatement(exec bool) error {
	tree.statement = new(StatementT)
	tree.charPos = 0

	err := tree.parseStatement(exec)
	if err != nil {
		return err
	}

	return nil
}

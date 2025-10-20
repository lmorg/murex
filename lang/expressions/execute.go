package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
)

func ExecuteExpr(p *lang.Process, expression []rune) (*primitives.DataType, error) {
	tree := NewParser(p, expression, 0)

	err := tree.parseExpression(true, true)
	if err != nil {
		return nil, err
	}

	return tree.executeExpr()
}

func (tree *ParserT) ParseStatement(exec bool, opts ...StatementOpts) error {
	tree.statement = new(StatementT)
	tree.charPos = 0

	for _, opt := range opts {
		opt(tree)
	}

	err := tree.parseStatement(exec)
	if err != nil {
		return err
	}

	return nil
}

type StatementOpts func(tree *ParserT)

func WithCommand(cmd string) StatementOpts {
	return func(tree *ParserT) {
		tree.statement.SetCommand([]rune(cmd))
	}
}

func WithAutoEscapeLineFeed() StatementOpts {
	return func(tree *ParserT) {
		tree.ignoreLf = true
	}
}

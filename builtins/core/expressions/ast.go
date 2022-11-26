package expressions

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/builtins/core/expressions/primitives"
	"github.com/lmorg/murex/builtins/core/expressions/symbols"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

type astNodeT struct {
	key   symbols.Exp
	value []byte
	pos   int
	dt    *primitives.DataType
}

func (node *astNodeT) Value() string {
	return string(node.value)
}

type expTreeT struct {
	ast        []*astNodeT
	charPos    int
	charOffset int
	astPos     int
	expression []byte
	p          *lang.Process
}

func (tree *expTreeT) nextChar() byte {
	if tree.charPos+1 >= len(tree.expression) {
		return 0
	}
	return tree.expression[tree.charPos+1]
}

func (tree *expTreeT) appendAst(key symbols.Exp, value ...byte) {
	tree.ast = append(tree.ast, &astNodeT{
		key:   key,
		value: value,
		pos:   tree.charPos + tree.charOffset - len(value),
	})
}

func (tree *expTreeT) appendAstWithPrimitive(key symbols.Exp, dt *primitives.DataType) {
	tree.ast = append(tree.ast, &astNodeT{
		key: key,
		dt:  dt,
		pos: tree.charPos + tree.charOffset,
	})
}

func (tree *expTreeT) foldAst(new *astNodeT) error {
	switch {
	case tree.astPos <= 0:
		return fmt.Errorf("cannot fold when tree.astPos<%d> <= 0<%d> (%s)",
			tree.astPos, len(tree.ast), consts.IssueTrackerURL)

	case tree.astPos >= len(tree.ast)-1:
		return fmt.Errorf("cannot fold when tree.astPos<%d> >= len(tree.ast)-1<%d> (%s)",
			tree.astPos, len(tree.ast), consts.IssueTrackerURL)

	case len(tree.ast) == 3:
		tree.ast = []*astNodeT{new}

	case tree.astPos == 1:
		tree.ast = append([]*astNodeT{new}, tree.ast[3:]...)

	case tree.astPos == len(tree.ast)-2:
		tree.ast = append(tree.ast[:len(tree.ast)-3], new)

	default:
		start := append(tree.ast[:tree.astPos-1], new)
		end := tree.ast[tree.astPos+2:]
		tree.ast = append(start, end...)
	}

	return nil
}

// memory safe
func (tree *expTreeT) prevSymbol() *astNodeT {
	if tree.astPos-1 < 0 {
		return nil
	}

	return tree.ast[tree.astPos-1]
}

// memory safe
func (tree *expTreeT) currentSymbol() *astNodeT {
	if tree.astPos < 0 || tree.astPos >= len(tree.ast) {
		return nil
	}

	return tree.ast[tree.astPos]
}

// memory safe
func (tree *expTreeT) nextSymbol() *astNodeT {
	if tree.astPos+1 >= len(tree.ast) {
		return nil
	}

	return tree.ast[tree.astPos+1]
}

func (tree *expTreeT) getLeftAndRightSymbols() (*astNodeT, *astNodeT, error) {
	left := tree.prevSymbol()
	right := tree.nextSymbol()

	if left == nil {
		return nil, nil, raiseError(tree.ast[tree.astPos], "missing value left of operation")
	}

	if right == nil {
		return nil, nil, raiseError(tree.ast[tree.astPos], "missing value right of operation")
	}

	return left, right, nil
}

func node2primitive(node *astNodeT) (*primitives.DataType, error) {
	switch node.key {
	case symbols.Number:
		f, err := strconv.ParseFloat(node.Value(), 64)
		if err != nil {
			return nil, raiseError(node, err.Error())
		}
		return &primitives.DataType{
			Primitive: primitives.Number,
			Value:     f,
		}, nil

	case symbols.QuoteSingle:
		return &primitives.DataType{
			Primitive: primitives.String,
			Value:     node.Value(),
		}, nil

	case symbols.QuoteDouble:
		// TODO: expand vars
		return &primitives.DataType{
			Primitive: primitives.String,
			Value:     node.Value(),
		}, nil

	case symbols.Boolean:
		return &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     types.IsTrue(node.value, 0),
		}, nil

	case symbols.Bareword:
		return &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		}, nil
	}

	return nil, raiseError(node, fmt.Sprintf("unexpected error converting node to primitive (%s)", consts.IssueTrackerURL))
}

func newExpTree(expression []byte) *expTreeT {
	tree := new(expTreeT)
	tree.expression = expression
	return tree
}

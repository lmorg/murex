package expressions

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

type astNodeT struct {
	key    symbols.Exp
	value  []rune
	pos    int
	offset int
	dt     *primitives.DataType
}

func (node *astNodeT) Value() string {
	return string(node.value)
}

type ParserT struct {
	ast           []*astNodeT
	statement     *StatementT
	charPos       int
	charOffset    int
	astPos        int
	startRow      int
	endRow        int
	startCol      int
	endCol        int
	expression    []rune
	subExp        bool
	p             *lang.Process
	_strictTypes  interface{}
	_strictArrays interface{}
	_expandGlob   interface{}
}

func (tree *ParserT) nextChar() rune {
	if tree.charPos+1 >= len(tree.expression) {
		return 0
	}
	return tree.expression[tree.charPos+1]
}

func (tree *ParserT) prevChar() rune {
	if tree.charPos < 1 {
		return 0
	}
	return tree.expression[tree.charPos-1]
}

func (tree *ParserT) crLf() {
	tree.endRow++
	tree.endCol = tree.charPos
}

func (tree *ParserT) GetColumnN() int { return tree.charOffset - tree.startCol + 2 }
func (tree *ParserT) GetLineN() int   { return tree.startRow }

func (tree *ParserT) appendAst(key symbols.Exp, value ...rune) {
	tree.ast = append(tree.ast, &astNodeT{
		key:    key,
		value:  value,
		pos:    tree.charPos - len(value),
		offset: tree.charOffset,
	})
}

func (tree *ParserT) appendAstWithPrimitive(key symbols.Exp, dt *primitives.DataType, value ...rune) {
	tree.ast = append(tree.ast, &astNodeT{
		key:    key,
		value:  value,
		pos:    tree.charPos - len(value),
		offset: tree.charOffset,
		dt:     dt,
	})
}

func (tree *ParserT) foldAst(new *astNodeT) error {
	switch {
	case tree.astPos <= 0:
		return fmt.Errorf("cannot fold when tree.astPos<%d> <= 0<%d> (%s)",
			tree.astPos, len(tree.ast), consts.IssueTrackerURL)

	case tree.astPos >= len(tree.ast)-1:
		return fmt.Errorf("cannot fold when tree.astPos<%d> >= len(tree.ast)-1<%d> (%s)",
			tree.astPos, len(tree.ast)-1, consts.IssueTrackerURL)

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

func (tree *ParserT) foldLeftAst(new *astNodeT) error {
	switch {
	case tree.astPos <= 0:
		return fmt.Errorf("cannot fold when tree.astPos<%d> <= 0<%d> (%s)",
			tree.astPos, len(tree.ast), consts.IssueTrackerURL)

	case tree.astPos >= len(tree.ast):
		return fmt.Errorf("cannot fold when tree.astPos<%d> >= len(tree.ast)<%d> (%s)",
			tree.astPos, len(tree.ast), consts.IssueTrackerURL)

	case len(tree.ast) == 2:
		tree.ast = []*astNodeT{new}

	case tree.astPos == 1:
		tree.ast = append([]*astNodeT{new}, tree.ast[2:]...)

	case tree.astPos == len(tree.ast)-1:
		tree.ast = append(tree.ast[:len(tree.ast)-2], new)

	default:
		start := append(tree.ast[:tree.astPos-1], new)
		end := tree.ast[tree.astPos+1:]
		tree.ast = append(start, end...)
	}

	return nil
}

// memory safe
func (tree *ParserT) prevSymbol() *astNodeT {
	if tree.astPos-1 < 0 {
		return nil
	}

	return tree.ast[tree.astPos-1]
}

// memory safe
func (tree *ParserT) currentSymbol() *astNodeT {
	if tree.astPos < 0 || tree.astPos >= len(tree.ast) {
		return nil
	}

	return tree.ast[tree.astPos]
}

// memory safe
func (tree *ParserT) nextSymbol() *astNodeT {
	if tree.astPos+1 >= len(tree.ast) {
		return nil
	}

	return tree.ast[tree.astPos+1]
}

func (tree *ParserT) getLeftAndRightSymbols() (*astNodeT, *astNodeT, error) {
	left := tree.prevSymbol()
	right := tree.nextSymbol()

	if left == nil {
		return nil, nil, raiseError(tree.expression, tree.ast[tree.astPos], 0, "missing value left of operation")
	}

	if right == nil {
		return nil, nil, raiseError(tree.expression, tree.ast[tree.astPos], 0, "missing value right of operation")
	}

	return left, right, nil
}

func node2primitive(node *astNodeT) (*primitives.DataType, error) {
	switch node.key {
	case symbols.Number:
		f, err := strconv.ParseFloat(node.Value(), 64)
		if err != nil {
			return nil, raiseError(nil, node, 0, err.Error())
		}
		return primitives.NewPrimitive(primitives.Number, f), nil

	case symbols.QuoteSingle, symbols.QuoteDouble, symbols.QuoteParenthesis:
		return primitives.NewPrimitive(primitives.String, node.Value()), nil

	case symbols.Boolean:
		return primitives.NewPrimitive(
			primitives.Boolean,
			types.IsTrueString(string(node.value), 0),
		), nil

	case symbols.Bareword:
		return primitives.NewPrimitive(primitives.Bareword, nil), nil

	case symbols.Calculated, symbols.Scalar, symbols.SubExpressionBegin:
		return primitives.NewPrimitive(primitives.Null, nil), nil

	case symbols.Null:
		return primitives.NewPrimitive(primitives.Null, nil), nil

	}

	return nil, raiseError(nil, node, 0, fmt.Sprintf("unexpected error converting node to primitive (%s)", consts.IssueTrackerURL))
}

func (tree *ParserT) StrictTypes() bool {
	if tree._strictTypes != nil {
		return tree._strictTypes.(bool)
	}

	var err error
	tree._strictTypes, err = tree.p.Config.Get("proc", "strict-types", types.Boolean)
	if err != nil {
		panic(err)
	}

	return tree._strictTypes.(bool)
}

func (tree *ParserT) StrictArrays() bool {
	if tree._strictArrays != nil {
		return tree._strictArrays.(bool)
	}

	var err error
	tree._strictArrays, err = tree.p.Config.Get("proc", "strict-arrays", types.Boolean)
	if err != nil {
		panic(err)
	}

	return tree._strictArrays.(bool)
}

func (tree *ParserT) ExpandGlob() bool {
	if tree._expandGlob != nil {
		return tree._expandGlob.(bool)
	}

	var err error
	tree._expandGlob, err = tree.p.Config.Get("shell", "expand-globs", types.Boolean)
	if err != nil {
		panic(err)
	}

	tree._expandGlob = tree._expandGlob.(bool) && tree.p.Scope.Id == lang.ShellProcess.Id &&
		tree.p.Parent.Id == lang.ShellProcess.Id && !tree.p.Background.Get() && lang.Interactive

	return tree._expandGlob.(bool)
}

package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	fn "github.com/lmorg/murex/lang/expressions/functions"
	"github.com/lmorg/murex/utils/consts"
)

func init() {
	lang.ParseExpression = ExpressionParser
	lang.ParseStatementParameters = StatementParametersParser
}

// ExpressionParser is intended to be called from other parsers as a way of
// embedding this expressions library into other language syntaxes. This
// function just parses the expression and returns the end of the expression.
func ExpressionParser(expression []rune, offset int, exec bool) (int, error) {
	tree := NewParser(nil, expression, offset)

	err := tree.parseExpression(exec)
	if err != nil {
		return 0, err
	}

	err = tree.validateExpression()
	if err != nil {
		return 0, err
	}

	return tree.charPos, nil
}

// StatementParametersParser is intended to be called from other parsers as a
// way of parsing function parameters
func StatementParametersParser(expression []rune, p *lang.Process) error {
	if p.Name.String() == lang.ExpressionFunctionName {
		s := string(p.Parameters.PreParsed[0])
		p.Parameters.DefineParsed([]string{s})
		return nil
	}

	tree := NewParser(nil, expression, 0)
	tree.p = p
	err := tree.ParseStatement(true)
	if err != nil {
		return err
	}

	params := make([]string, len(tree.statement.parameters))
	for i := range params {
		params[i] = string(tree.statement.parameters[i])
	}

	p.Parameters.DefineParsed(params)
	return nil
}

func NewParser(p *lang.Process, expression []rune, offset int) *ParserT {
	tree := new(ParserT)
	tree.expression = expression
	tree.p = p
	tree.charOffset = offset
	return tree
}

func (tree *ParserT) preParser() (int, error) {
	expErr := tree.parseExpression(false)
	if expErr == nil {
		// if successful parse, then also validate.
		// no point validating if the parser has already failed
		expErr = tree.validateExpression()
	}

	if expErr == nil {
		return tree.charPos, nil
	}

	stErr := tree.ParseStatement(false)
	if stErr != nil {
		return 0, stErr
	}

	stErr = tree.statement.validate()

	if len(tree.statement.parameters) > 0 && tree.statement.parameters[0][0] == '=' ||
		stErr != nil {
		// i _still_ think this is probably an expression
		return 0, expErr
	}

	return tree.charPos, nil

}

var exprFuncName = []rune(lang.ExpressionFunctionName)

func (blk *BlockT) append(tree *ParserT, this fn.Property, next fn.Property) {
	switch {
	case tree == nil:
		// do nothing

	case tree.statement == nil:

		blk.Functions = append(blk.Functions, fn.FunctionT{
			Command:    exprFuncName,
			Parameters: [][]rune{tree.expression[tree.charOffset : tree.charPos+1]},
			Properties: blk.nextProperty | this,
			//blk.Functions[i].LineN = blk.lineN // TODO
			//blk.Functions[i].ColumnN = blk.columnN
		})

	default:
		blk.Functions = append(blk.Functions, fn.FunctionT{
			Command:    tree.statement.command,
			Parameters: tree.statement.parameters,
			NamedPipes: tree.statement.namedPipes,
			Raw:        tree.expression[tree.charOffset : tree.charPos+1],
			Properties: blk.nextProperty | this,
			//blk.Functions[i].LineN = blk.lineN // TODO
			//blk.Functions[i].ColumnN = blk.columnN
		})

	}

	blk.nextProperty = next
}

func (blk *BlockT) ParseBlock() error {
	var tree *ParserT

	for ; blk.charPos < len(blk.expression); blk.charPos++ {
		r := blk.expression[blk.charPos]

		switch r {
		case ' ', '\t', '\r':
			continue

		case '\n':
			blk.append(tree, 0, fn.P_NEW_CHAIN)
			tree = nil
			continue

		case ';':
			blk.append(tree, 0, fn.P_NEW_CHAIN)
			tree = nil

		case '&':
			if blk.nextChar() == '&' {
				blk.charPos++
				blk.append(tree, 0, fn.P_NEW_CHAIN|fn.P_LOGIC_AND)
				tree = nil

			} else {
				blk.panic('&', '&')
			}

		case '|':
			if blk.nextChar() == '|' {
				blk.charPos++
				blk.append(tree, 0, fn.P_NEW_CHAIN|fn.P_LOGIC_OR)
				tree = nil

			} else {
				blk.append(tree, fn.P_PIPE_OUT, fn.P_METHOD)
				tree = nil
			}

		case '-':
			if blk.nextChar() != '>' {
				blk.panic('-', '>')
			}
			blk.charPos++
			blk.append(tree, fn.P_PIPE_OUT, fn.P_METHOD)
			tree = nil

		case '?':
			blk.append(tree, fn.P_PIPE_ERR, fn.P_METHOD)
			tree = nil

		case '=':
			if tree.nextChar() != '>' {
				blk.panic('=', '>')
			}
			blk.charPos++
			blk.append(tree, fn.P_PIPE_OUT, fn.P_METHOD)
			tree = nil
			//panic("TODO") // TODO

		case '>':
			if tree.nextChar() != '>' {
				blk.panic('>', '>')
			}
			panic("TODO") // TODO

			/*default:
			if tree == nil {
				// this is probably just the first run in a new block. eg
				//   { -> do something }
				continue
			}
			blk.panic(blk.expression[blk.charPos], 0)*/

		default:
			tree := NewParser(nil, blk.expression[blk.charPos:], 0)
			newPos, err := tree.preParser()
			if err != nil {
				return err
			}

			blk.charPos += newPos

		}

	}

	if blk.charPos >= len(blk.expression) {
		blk.append(tree, 0, 0)
	}

	return nil
}

/*func (blk *BlockT) ParseBlock() error {
	parseBlock(blk, nil)
	for blk.charPos < len(blk.expression) {
		tree := NewParser(nil, blk.expression[blk.charPos:], 0)
		newPos, err := tree.preParser()

		blk.charPos += newPos + 1

		if newPos == 1 || (tree.statement != nil && len(tree.statement.command) == 0) {
			continue
		}
		if err != nil {
			return err
		}

		if blk.charPos >= len(blk.expression) {
			blk.append(tree, 0, 0)
			break
		}

		parseBlock(blk, tree)
	}

	return nil
}

func parseBlock(blk *BlockT, tree *ParserT) error {
	for ; blk.charPos < len(blk.expression)-1; blk.charPos++ {
		r := blk.expression[blk.charPos]
		if r == ' ' || r == '\t' || r == '\r' {
			continue
		}
		break
	}

	switch blk.expression[blk.charPos] {
	case '\n':
		blk.append(tree, 0, fn.P_NEW_CHAIN)

	case ';':
		blk.append(tree, 0, fn.P_NEW_CHAIN)

	case '&':
		if blk.nextChar() == '&' {
			blk.charPos++
			blk.append(tree, 0, fn.P_NEW_CHAIN|fn.P_LOGIC_AND)
		} else {
			blk.panic('&', '&')
		}

	case '|':
		if blk.nextChar() == '|' {
			blk.charPos++
			blk.append(tree, 0, fn.P_NEW_CHAIN|fn.P_LOGIC_OR)
		} else {
			blk.append(tree, fn.P_PIPE_OUT, fn.P_METHOD)
		}

	case '-':
		if blk.nextChar() != '>' {
			blk.panic('-', '>')
		}
		blk.charPos++
		blk.append(tree, fn.P_PIPE_OUT, fn.P_METHOD)

	case '?':
		blk.append(tree, fn.P_PIPE_ERR, fn.P_METHOD)

	case '=':
		if tree.nextChar() != '>' {
			blk.panic('=', '>')
		}
		blk.charPos++
		blk.append(tree, fn.P_PIPE_OUT, fn.P_METHOD)
		//panic("TODO") // TODO

	case '>':
		if tree.nextChar() != '>' {
			blk.panic('>', '>')
		}
		panic("TODO") // TODO

	default:
		if tree == nil {
			// this is probably just the first run in a new block. eg
			//   { -> do something }
			return nil
		}
		blk.panic(blk.expression[blk.charPos], 0)

	}

	return nil
}*/

func (blk *BlockT) panic(found rune, follows rune) {
	msg := "unexpected parser error: '%s' found"
	if follows == 0 {
		panic(fmt.Sprintf(msg+". "+consts.IssueTrackerURL, string(found)))
	}

	msg += " but no '%s' follows"
	panic(fmt.Sprintf(msg+". "+consts.IssueTrackerURL, string(found), string(follows)))
}

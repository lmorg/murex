package expressions

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	fn "github.com/lmorg/murex/lang/expressions/functions"
	"github.com/lmorg/murex/lang/expressions/node"
	"github.com/lmorg/murex/lang/types"
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
	tree := NewParser(nil, expression, offset, node.Nil)

	err := tree.parseExpression(exec, false)
	if err != nil {
		return 0, err
	}

	err = tree.validateExpression(exec)
	if err != nil {
		return 0, err
	}

	return tree.charPos, nil
}

// StatementParametersParser is intended to be called from other parsers as a
// way of parsing function parameters
func StatementParametersParser(expression []rune, p *lang.Process) (string, []string, error) {
	if p.Name.String() == lang.ExpressionFunctionName {
		return p.Name.String(), []string{string(p.Parameters.PreParsed[0])}, nil
	}

	tree := NewParser(nil, expression, 0, node.Nil)
	tree.p = p
	err := tree.ParseStatement(true)
	if err != nil {
		return "", nil, err
	}

	return tree.statement.String(), tree.statement.Parameters(), nil
}

func NewParser(p *lang.Process, expression []rune, offset int, syntaxTree node.SyntaxTreeT) *ParserT {
	return &ParserT{
		expression: expression,
		p:          p,
		charOffset: offset,
		syntaxTree: syntaxTree,
	}
}

func (blk *BlockT) preParser(tree *ParserT) (int, error) {
	expErr := tree.parseExpression(false, false)
	if expErr == nil {
		// if successful parse, then also validate.
		// no point validating if the parser has already failed
		expErr = tree.validateExpression(false)
	}

	if expErr == nil {
		blk.syntaxTree.Merge(tree.syntaxTree)
		return tree.charPos, nil
	}

	tree.syntaxTree = blk.syntaxTree.New()
	defer blk.syntaxTree.Merge(tree.syntaxTree)

	stErr := tree.ParseStatement(false)
	if stErr != nil {
		return 0, stErr
	}

	stErr = tree.statement.validate()

	if len(tree.statement.command) == 0 {
		return 0, errors.New("you cannot have zero length commands")
	}

	if stErr == nil && !tree.statement.asStatement &&
		len(tree.statement.parameters) > 0 && len(tree.statement.parameters[0]) > 0 && tree.statement.parameters[0][0] == '=' {
		// i _still_ think this is probably an expression
		return 0, expErr
	}

	return tree.charPos, nil

}

var expressionFunctionName = []rune(lang.ExpressionFunctionName)

func (blk *BlockT) append(tree *ParserT, this fn.Property, next fn.Property) error {
	switch {

	case tree == nil && blk.nextProperty.FollowOnFn():
		exprRune, exprPos := cropCodeInErrMsg(blk.expression, blk.charPos)
		return fmt.Errorf("invalid syntax at %d. Unexpected pipeline continuation token:\n> %s\n> %s\n> these tokens:     %s\n> shouldn't follow: %s",
			blk.charPos,
			string(exprRune), strings.Repeat(" ", exprPos)+"^",
			next.Decompose(),
			fn.Property(fn.P_PIPE_OUT|fn.P_PIPE_ERR|fn.P_LOGIC_AND|fn.P_LOGIC_OR).Decompose())

	case len(blk.Functions) > 0 && tree == nil && next.FollowOnFn():
		exprRune, exprPos := cropCodeInErrMsg(blk.expression, blk.charPos)
		return fmt.Errorf("invalid syntax at %d. Semi-colon or line break preceding a pipeline continuation token:\n> %s\n> %s\n> these tokens:     %s\n> shouldn't follow: %s",
			blk.charPos,
			string(exprRune), strings.Repeat(" ", exprPos)+"^",
			this.Decompose(),
			fn.Property(fn.P_NEW_CHAIN|fn.P_LOGIC_AND|fn.P_LOGIC_OR).Decompose())

	case tree == nil:
		// do nothing

	case tree.statement == nil:
		if tree.charPos+1 >= len(tree.expression) {
			tree.charPos = len(tree.expression) - 1
		}
		blk.Functions = append(blk.Functions, fn.FunctionT{
			Raw:        tree.expression[:tree.charPos+1],
			Command:    expressionFunctionName,
			Parameters: [][]rune{tree.expression[:tree.charPos+1]},
			Properties: blk.nextProperty | this,
			LineN:      blk.lineN + tree.GetLineN(),
			ColumnN:    tree.GetColumnN(),
		})

	default:
		blk.Functions = append(blk.Functions, fn.FunctionT{
			Raw:        tree.expression[:tree.charPos+1],
			Command:    tree.statement.command,
			Parameters: tree.statement.parameters,
			NamedPipes: tree.statement.namedPipes,
			Cast:       tree.statement.cast,
			Properties: blk.nextProperty | this,
			LineN:      blk.lineN + tree.GetLineN(),
			ColumnN:    tree.GetColumnN(),
		})

	}

	blk.nextProperty = next
	return nil
}

var formatGeneric = []rune("format " + types.Generic)

func (blk *BlockT) ParseBlock() error {
	var tree *ParserT
	//blk.syntaxTree.Add(node.H_COMMAND)

	for ; blk.charPos < len(blk.expression); blk.charPos++ {
		r := blk.expression[blk.charPos]

		switch r {
		case ' ', '\t', '\r':
			continue

		case '\n':
			if err := blk.append(tree, 0, fn.P_NEW_CHAIN); err != nil {
				return err
			}
			tree = nil
			blk.lineN++
			blk.offset = blk.charPos
			continue

		case '#':
			comment := NewParser(nil, blk.expression[blk.charPos:], 0, blk.syntaxTree)
			comment.parseComment()
			blk.charPos += comment.charPos

		case '/':
			switch {
			case blk.nextChar() == '#':
				comment := NewParser(nil, blk.expression[blk.charPos:], 0, blk.syntaxTree)
				if err := comment.parseCommentMultiLine(); err != nil {
					return err
				}
				blk.charPos += comment.charPos
			default:
				tree = NewParser(nil, blk.expression[blk.charPos:], blk.charPos-1, blk.syntaxTree.New())
				newPos, err := blk.preParser(tree)
				if err != nil {
					return err
				}
				blk.charPos += newPos
			}

		case ';':
			if err := blk.append(tree, 0, fn.P_NEW_CHAIN); err != nil {
				return err
			}
			tree = nil

		case '&':
			switch {
			case blk.nextChar() == '&':
				blk.charPos++
				if err := blk.append(tree, 0, fn.P_NEW_CHAIN|fn.P_FOLLOW_ON|fn.P_LOGIC_AND); err != nil {
					return err
				}
				tree = nil
			case tree == nil:
				tree = NewParser(nil, blk.expression[blk.charPos:], 0, blk.syntaxTree.New())
				newPos, err := blk.preParser(tree)
				if err != nil {
					return err
				}
				blk.charPos += newPos
			default:
				blk.panic('&', '&')
			}

		case '|':
			if blk.nextChar() == '|' {
				blk.charPos++
				if err := blk.append(tree, 0, fn.P_NEW_CHAIN|fn.P_FOLLOW_ON|fn.P_LOGIC_OR); err != nil {
					return err
				}
				tree = nil

			} else {
				if err := blk.append(tree, fn.P_PIPE_OUT, fn.P_FOLLOW_ON|fn.P_METHOD); err != nil {
					return err
				}
				tree = nil
			}

		case '-':
			switch {
			case blk.nextChar() == '>':
				blk.charPos++
				if err := blk.append(tree, fn.P_PIPE_OUT, fn.P_FOLLOW_ON|fn.P_METHOD); err != nil {
					return err
				}
				tree = nil
			case tree == nil:
				tree = NewParser(nil, blk.expression[blk.charPos:], 0, blk.syntaxTree.New())
				newPos, err := blk.preParser(tree)
				if err != nil {
					return err
				}
				blk.charPos += newPos
			default:
				blk.panic('-', '>')
			}

		case '?':
			if err := blk.append(tree, fn.P_PIPE_ERR, fn.P_FOLLOW_ON|fn.P_METHOD); err != nil {
				return err
			}
			tree = nil

		case '=':
			switch {
			case blk.nextChar() == '>':
				blk.charPos++
				if err := blk.append(tree, fn.P_PIPE_OUT, fn.P_FOLLOW_ON|fn.P_METHOD); err != nil {
					return err
				}
				tree = nil
				format := NewParser(nil, formatGeneric, 0, blk.syntaxTree.New())
				_, _ = blk.preParser(format)
				if err := blk.append(format, fn.P_PIPE_OUT, fn.P_FOLLOW_ON|fn.P_METHOD); err != nil {
					return err
				}

			case tree == nil:
				tree = NewParser(nil, blk.expression[blk.charPos:], 0, blk.syntaxTree.New())
				newPos, err := blk.preParser(tree)
				if err != nil {
					return err
				}
				blk.charPos += newPos
			default:
				blk.panic('=', '>')
			}

		case '>':
			switch {
			case blk.nextChar() == '>':
				/*if len(blk.Functions) > 0 &&
					len(blk.Functions[len(blk.Functions)-1].Raw) == 0 &&
					!blk.Functions[len(blk.Functions)-1].Properties.Method() {
					panic("ugh")
				}*/

				err := blk.append(tree, fn.P_PIPE_OUT, fn.P_FOLLOW_ON|fn.P_METHOD)
				if err != nil {
					return err
				}
				tree, err = blk.parseStatementWithKnownCommand('>', '>')
				if err != nil {
					return err
				}
			default:
				tree = NewParser(nil, blk.expression[blk.charPos:], blk.charPos-1, blk.syntaxTree.New())
				newPos, err := blk.preParser(tree)
				if err != nil {
					return err
				}
				blk.charPos += newPos
			}

		default:
			tree = NewParser(nil, blk.expression[blk.charPos:], blk.charPos-1, blk.syntaxTree.New())
			newPos, err := blk.preParser(tree)
			if err != nil {
				return err
			}
			blk.charPos += newPos

		}

	}

	if blk.charPos >= len(blk.expression) {
		if err := blk.append(tree, 0, 0); err != nil {
			return err
		}
	}

	return nil
}

func (blk *BlockT) parseStatementWithKnownCommand(command ...rune) (*ParserT, error) {
	tree := NewParser(nil, blk.expression[blk.charPos:], 0, node.Nil) // TODO: check this is correct
	tree.statement = new(StatementT)
	tree.statement.command = command
	tree.charPos = len(command)
	err := tree.parseStatement(false)
	if err != nil {
		return nil, err
	}
	blk.charPos += tree.charPos
	return tree, nil
}

func (blk *BlockT) panic(found rune, follows rune) {
	msg := "unexpected parser error: '%s' found"
	if follows == 0 {
		panic(fmt.Sprintf(msg+". "+consts.IssueTrackerURL, string(found)))
	}

	msg += " but no '%s' follows"
	panic(fmt.Sprintf(msg+". "+consts.IssueTrackerURL, string(found), string(follows)))
}

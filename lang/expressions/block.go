package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/functions"
	"github.com/lmorg/murex/lang/expressions/node"
)

type BlockT struct {
	Functions    []functions.FunctionT
	expression   []rune
	charPos      int
	lineN        int
	offset       int
	nextProperty functions.Property
	syntaxTree   node.SyntaxTreeT
}

func (blk *BlockT) nextChar() rune {
	if blk.charPos+1 >= len(blk.expression) {
		return 0
	}
	return blk.expression[blk.charPos+1]
}

func NewBlock(block []rune) *BlockT {
	blk := new(BlockT)
	blk.expression = block
	blk.nextProperty = functions.P_NEW_CHAIN
	blk.syntaxTree = node.Nil
	return blk
}

func ParseBlock(block []rune) (*[]functions.FunctionT, error) {
	blk := NewBlock(block)
	err := blk.ParseBlock()
	return &blk.Functions, err
}

func init() {
	lang.ParseBlock = ParseBlock
}

func SyntaxHighlight(block []rune) []rune {
	blk := NewBlock(block)
	blk.syntaxTree = node.NewHighlighter(&node.DefaultTheme)
	_ = blk.ParseBlock()
	return blk.syntaxTree.SyntaxHighlight()
}

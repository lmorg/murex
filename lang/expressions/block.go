package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/functions"
)

type BlockT struct {
	Functions    []functions.FunctionT
	expression   []rune
	charPos      int
	lineN        int
	offset       int
	nextProperty functions.Property
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

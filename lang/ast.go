package lang

import (
	"github.com/lmorg/murex/lang/proc/parameters"
)

type AstNode struct {
	Name        string
	ParamTokens [][]parameters.ParamToken
	NewChain    bool
	Method      bool
	PipeOut     bool
	PipeErr     bool
	LineNumber  int
	ColNumber   int
}

type AstNodes []AstNode

// Last node in the AST array
func (n *AstNodes) Last() *AstNode {
	if len(*n) == 0 {
		return &(*n)[0]
	}
	return &(*n)[len(*n)-1]
}

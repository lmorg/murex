package lang

import (
	"github.com/lmorg/murex/lang/parameters"
)

// AstNode is a tokenized struct for each command, including it's parameters
type AstNode struct {
	Name        string
	ParamTokens [][]parameters.ParamToken
	NewChain    bool
	Method      bool
	PipeOut     bool
	PipeErr     bool
	LineNumber  int
	ColNumber   int
	LogicAnd    bool
	LogicOr     bool
}

// AstNodes is the entire code block sequenced as an array of tokenized structs
type AstNodes []AstNode

// Last node in the AST array
func (n *AstNodes) Last() *AstNode {
	if len(*n) == 0 {
		return &(*n)[0]
	}
	return &(*n)[len(*n)-1]
}

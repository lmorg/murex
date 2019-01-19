package lang

import (
	"github.com/lmorg/murex/lang/proc/parameters"
)

type astNode struct {
	Name        string
	ParamTokens [][]parameters.ParamToken
	NewChain    bool
	Method      bool
	PipeOut     bool
	PipeErr     bool
	LineNumber  int
	ColNumber   int
}

type astNodes []astNode

// Last node in the AST array
func (n *astNodes) Last() *astNode {
	if len(*n) == 0 {
		return &(*n)[0]
	}
	return &(*n)[len(*n)-1]
}

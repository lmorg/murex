package lang

import "github.com/lmorg/murex/lang/proc"

type Node struct {
	Name       string
	Parameters proc.Parameters
	NewChain   bool
	Method     bool
	PipeOut    bool
	PipeErr    bool
	Process    proc.Process
	//Children   Nodes
}

type Nodes []Node

func (n *Nodes) Last() *Node {
	if len(*n) == 0 {
		return &(*n)[0]
	}
	return &(*n)[len(*n)-1]
}

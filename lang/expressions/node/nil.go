package node

type _nil struct{}

func (n *_nil) New() SyntaxTreeT        { return Nil }
func (n *_nil) Add(Symbol, ...rune)     {}
func (n *_nil) Append(...rune)          {}
func (n *_nil) ChangeSymbol(Symbol)     {}
func (n *_nil) Merge(child SyntaxTreeT) {}
func (n *_nil) _nodes() []*nodeT        { return nil }
func (n *_nil) SyntaxHighlight() []rune { return nil }

var Nil = new(_nil)

package node

type syntaxTree struct {
	theme *ThemeT
	nodes []*nodeT
}

type nodeT struct {
	symbol Symbol
	r      []rune
}

func NewHighlighter(theme *ThemeT) SyntaxTreeT {
	if theme == nil {
		return &syntaxTree{theme: &DefaultTheme}
	}
	return &syntaxTree{theme: theme}
}

func (st *syntaxTree) New() SyntaxTreeT {
	return &syntaxTree{theme: st.theme}
}

func (st *syntaxTree) Add(symbol Symbol, r ...rune) {
	node := &nodeT{
		symbol: symbol,
		r:      r,
	}
	st.nodes = append(st.nodes, node)
}

func (st *syntaxTree) Append(r ...rune) {
	if len(st.nodes) == 0 {
		st.Add(H_COMMAND)
	}
	st.nodes[len(st.nodes)-1].r = append(st.nodes[len(st.nodes)-1].r, r...)
}
func (st *syntaxTree) ChangeSymbol(symbol Symbol) {
	st.nodes[len(st.nodes)-1].symbol = symbol
}

func (st *syntaxTree) Merge(child SyntaxTreeT) {
	st.nodes = append(st.nodes, child._nodes()...)
}

func (st *syntaxTree) _nodes() []*nodeT {
	return st.nodes
}

func (st *syntaxTree) SyntaxHighlight() []rune {
	var r []rune
	for _, node := range st.nodes {
		r = append(r, st.theme.highlight(node.symbol, node.r...)...)
	}
	return r
}

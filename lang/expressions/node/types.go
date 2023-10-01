package node

type SyntaxTreeT interface {
	New() SyntaxTreeT
	Append(Symbol, []rune)
	Merge(SyntaxTreeT)
	_nodes() []*nodeT
	SyntaxHighlight() []rune
}

/*Clone() ThemeT
	BeginSubExpr()
	UpdateParent()
	Clear()
	ClearExpr()
	GetHighlighted() []rune
	Begin(int, ...rune)
	Append(...rune)
	End(int, ...rune)
	Highlight(int, ...rune)
	IsHighlighter() bool
}*/

type Symbol int

const (
	H_COMMAND Symbol = iota
	H_CMD_MODIFIER
	H_PARAMETER
	H_GLOB
	H_NUMBER
	H_BAREWORD
	H_BOOLEAN
	H_NULL
	H_VARIABLE
	H_MACRO
	H_ESCAPE
	H_QUOTED_STRING
	H_ARRAY_ITEM
	H_OBJECT_KEY
	H_OBJECT_VALUE
	H_OPERATOR
	H_PIPE
	H_COMMENT
	H_ERROR
	_H_BRACE

	_adjust Symbol = 10 + iota
	H_END_COMMAND
	H_END_CMD_MODIFIER
	H_END_PARAMETER
	H_END_GLOB
	H_END_NUMBER
	H_END_BAREWORD
	H_END_BOOLEAN
	H_END_NULL
	H_END_VARIABLE
	H_END_MACRO
	H_END_ESCAPE
	H_END_QUOTED_STRING
	H_END_ARRAY_ITEM
	H_END_OBJECT_KEY
	H_END_OBJECT_VALUE
	H_END_OPERATOR
	H_END_PIPE
	H_END_COMMENT
	H_END_ERROR

	_H_END_BRACE

	H_BRACE_OPEN
	H_BRACE_CLOSE
)

const ADJUST = _adjust + 1

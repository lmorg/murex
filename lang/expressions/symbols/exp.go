package symbols

//go:generate stringer -type=Exp
type Exp int

const (
	// The order of these constants is important.
	// murex saves a few instructions when calculating the order of operations
	// by checking the const ID against a range (rather than checking each
	// operation matches)

	/*
		ERRORS
	*/

	Undefined  Exp = 0 // only used to catch edge case bugs where an operator might be allocated but not assigned a role
	Unexpected Exp = iota + 1
	InvalidHyphen
	SubExpressionEnd
	ObjectEnd
	ArrayEnd

	/*
		DATA VALUES
	*/
	DataValues // not used. Just a title to name a range

	Bareword
	SubExpressionBegin
	ObjectBegin
	ArrayBegin
	QuoteSingle
	QuoteDouble
	QuoteParenthesis
	Number
	Boolean
	Null
	Scalar
	Calculated // data fields that are a result of a calculation

	/*
		OPERATIONS
	*/
	Operations // not used. Just a title to name a range

	// 15. Comma operator
	// 14. Assignment operators (right to left)
	Assign
	AssignUpdate
	AssignAndAdd
	AssignAndSubtract
	AssignAndDivide
	AssignAndMultiply
	AssignOrMerge

	// 13. Conditional expression (ternary)
	Elvis
	NullCoalescing

	// 12. Logical OR
	LogicalOr

	// 11. Logical AND
	LogicalAnd

	// 10. Bitwise inclusive (normal) OR
	// 09. Bitwise exclusive OR (XOR)
	// 08. Bitwise AND
	// 07. Comparisons: equal and not equal
	EqualTo
	NotEqualTo
	Like
	NotLike
	Regexp
	NotRegexp

	// 06. Comparisons: less-than and greater-than
	GreaterThan
	GreaterThanOrEqual
	LessThan
	LessThanOrEqual

	// 05. Bitwise shift left and right

	// 04.1 Marge
	Merge

	// 04. Addition and subtraction
	PlusPlus
	Add
	MinusMinus
	Subtract

	// 03. Multiplication, division, modulo
	Multiply
	Divide

	// 02. (most) unary operators, sizeof and type casts (right to left)
	// 01. Function call, scope, array/member access
)

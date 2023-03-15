// Code generated by "stringer -type=Exp"; DO NOT EDIT.

package symbols

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Undefined-0]
	_ = x[Unexpected-2]
	_ = x[InvalidHyphen-3]
	_ = x[SubExpressionEnd-4]
	_ = x[ObjectEnd-5]
	_ = x[ArrayEnd-6]
	_ = x[DataValues-7]
	_ = x[Bareword-8]
	_ = x[SubExpressionBegin-9]
	_ = x[ObjectBegin-10]
	_ = x[ArrayBegin-11]
	_ = x[QuoteSingle-12]
	_ = x[QuoteDouble-13]
	_ = x[QuoteParenthesis-14]
	_ = x[Number-15]
	_ = x[Boolean-16]
	_ = x[Null-17]
	_ = x[Scalar-18]
	_ = x[Calculated-19]
	_ = x[Operations-20]
	_ = x[Assign-21]
	_ = x[AssignAndAdd-22]
	_ = x[AssignAndSubtract-23]
	_ = x[AssignAndDivide-24]
	_ = x[AssignAndMultiply-25]
	_ = x[EqualTo-26]
	_ = x[NotEqualTo-27]
	_ = x[Like-28]
	_ = x[NotLike-29]
	_ = x[Regexp-30]
	_ = x[NotRegexp-31]
	_ = x[GreaterThan-32]
	_ = x[GreaterThanOrEqual-33]
	_ = x[LessThan-34]
	_ = x[LessThanOrEqual-35]
	_ = x[Add-36]
	_ = x[Subtract-37]
	_ = x[Multiply-38]
	_ = x[Divide-39]
}

const (
	_Exp_name_0 = "Undefined"
	_Exp_name_1 = "UnexpectedInvalidHyphenSubExpressionEndObjectEndArrayEndDataValuesBarewordSubExpressionBeginObjectBeginArrayBeginQuoteSingleQuoteDoubleQuoteParenthesisNumberBooleanNullScalarCalculatedOperationsAssignAssignAndAddAssignAndSubtractAssignAndDivideAssignAndMultiplyEqualToNotEqualToLikeNotLikeRegexpNotRegexpGreaterThanGreaterThanOrEqualLessThanLessThanOrEqualAddSubtractMultiplyDivide"
)

var (
	_Exp_index_1 = [...]uint16{0, 10, 23, 39, 48, 56, 66, 74, 92, 103, 113, 124, 135, 151, 157, 164, 168, 174, 184, 194, 200, 212, 229, 244, 261, 268, 278, 282, 289, 295, 304, 315, 333, 341, 356, 359, 367, 375, 381}
)

func (i Exp) String() string {
	switch {
	case i == 0:
		return _Exp_name_0
	case 2 <= i && i <= 39:
		i -= 2
		return _Exp_name_1[_Exp_index_1[i]:_Exp_index_1[i+1]]
	default:
		return "Exp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

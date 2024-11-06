package parameters

import "sync"

// Parameters is the parameter object
type Parameters struct {
	mutex     sync.RWMutex
	params    []string
	PreParsed [][]rune
}

func (param *Parameters) Dump() interface{} {
	dump := make(map[string]interface{})

	param.mutex.Lock()
	params := make([]string, len(param.params))
	pre := make([]string, len(param.PreParsed))

	copy(params, param.params)

	for i := range pre {
		pre[i] = string(param.PreParsed[i])
	}

	param.mutex.Unlock()

	dump["Parsed"] = params
	dump["PreParsed"] = pre

	return dump
}

// GetRawishExpression should only be used for expressions.
// Added to pass TestExpressionsMultipleParams():
// builtins/core/expressions/expressions_test.go:220
func (param *Parameters) GetRawishExpression() []rune {
	var r []rune

	for i := range param.PreParsed {
		r = append(r, param.PreParsed[i]...)
	}

	return r
}

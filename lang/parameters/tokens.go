package parameters

import "sync"

// Parameters is the parameter object
type Parameters struct {
	mutex     sync.RWMutex
	params    []string
	PreParsed [][]rune
}

func (param *Parameters) Dump() any {
	dump := make(map[string]any)

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

func (param *Parameters) Raw() []rune {
	var r []rune

	for i := range param.PreParsed {
		r = append(r, ' ')
		r = append(r, param.PreParsed[i]...)
	}

	return r
}

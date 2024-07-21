//go:build trace
// +build trace

package lang

import (
	"fmt"
	"regexp"
	"runtime"
)

var rxTrace = regexp.MustCompile(`^.*murex/`)

type traceT struct {
	Line string
	Func string
}

func trace(p *Process) {
	for i := 2; i < 22; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			return
		}

		file = rxTrace.ReplaceAllString(file, "")

		p.Trace = append(p.Trace, traceT{
			Line: fmt.Sprintf("%s:%d", file, line),
			Func: runtime.FuncForPC(pc).Name() + "(...)",
		})
	}
}

package functions

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestFunctionsBitwise(t *testing.T) {
	count.Tests(t, 2)

	fn := new(FunctionT)
	fn.Properties = P_NEW_CHAIN | P_PIPE_OUT | P_LOGIC_AND

	if !fn.Properties.NewChain() || fn.Properties.Method() ||
		!fn.Properties.PipeOut() || fn.Properties.PipeErr() ||
		!fn.Properties.LogicAnd() || fn.Properties.LogicOr() {
		t.Error("test 0 failed")
	}

	fn.Properties = 0
	fn.Properties = P_METHOD | P_PIPE_ERR | P_LOGIC_OR

	if fn.Properties.NewChain() || !fn.Properties.Method() ||
		fn.Properties.PipeOut() || !fn.Properties.PipeErr() ||
		fn.Properties.LogicAnd() || !fn.Properties.LogicOr() {
		t.Error("test 1 failed")
	}
}

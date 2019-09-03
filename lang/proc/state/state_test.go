package state

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestStateStringer tests stringer has ran
func TestStateStringer(t *testing.T) {
	count.Tests(t, 9, "TestStateStringer")

	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error("Not all constants have been stringified")
		}
	}()

	t.Log(Undefined.String())
	t.Log(MemAllocated.String())
	t.Log(Assigned.String())
	t.Log(Starting.String())
	t.Log(Executing.String())
	t.Log(Executed.String())
	t.Log(Terminating.String())
	t.Log(AwaitingGC.String())
	t.Log(Stopped.String())
}

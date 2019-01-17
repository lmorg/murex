package state

import (
	"testing"
)

// TestState tests stringer has ran
func TestState(t *testing.T) {
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

package runmode

import (
	"testing"
)

// TestRunmode tests stringer has ran
func TestRunmode(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error("Not all constants have been stringified")
		}
	}()

	t.Log(Normal.String())
	t.Log(Shell.String())
	t.Log(Try.String())
	t.Log(TryPipe.String())
	t.Log(Evil.String())
}

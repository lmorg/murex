package runmode

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestRunmodeStringer tests stringer has ran
func TestRunmodeStringer(t *testing.T) {
	count.Tests(t, 5, "TestRunmodeStringer")

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

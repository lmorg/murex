package process_test

import (
	"os/exec"
	"testing"

	"github.com/lmorg/murex/lang/process"
	"github.com/lmorg/murex/test/count"
)

func TestExec(t *testing.T) {
	count.Tests(t, 4)

	e := new(process.Exec)

	e.Set(13, new(exec.Cmd))

	pid, cmd := e.Get()
	if pid != 13 {
		t.Errorf("Set and/or Get failed. Didn't return 13")
	}
	if cmd == nil {
		t.Errorf("Set and/or Get failed. Returned nil")
	}

	if e.Pid() != 13 {
		t.Errorf("Set and/or Pid failed. Didn't return 13")
	}
}

//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang_test

import (
	"testing"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func TestProcessExecStruct(t *testing.T) {
	count.Tests(t, 2)

	lang.InitEnv()

	fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)

	go fork.Execute([]rune(`exec: sleep 3`))
	time.Sleep(1 * time.Second)

	p := lang.ForegroundProc.Get()
	name := p.Name.String()
	param, _ := p.Parameters.String(0)
	if name != "exec" || param != "sleep" {
		t.Error("Invalid foreground process!")
		t.Logf("  Expected name:  exec")
		t.Logf("  Actual name:    %s", name)
		t.Logf("  Expected param: sleep")
		t.Logf("  Actual param:   %s", param)
		return
	}

	if p.SystemProcess == nil {
		t.Errorf("Expecting a non-nil p.SystemProcess")
		return
	}

	pid := p.SystemProcess.Pid()
	if p.SystemProcess.Pid() == 0 {
		t.Errorf("Expecting a non-zero pid, instead found %d", pid)
	}
}

package process_test

import (
	"errors"
	"os"
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

type sysProcTest1T struct{}

func (sp *sysProcTest1T) Pid() int                   { return 0 }
func (sp *sysProcTest1T) ExitNum() int               { return 1 }
func (sp *sysProcTest1T) Kill() error                { return nil }
func (sp *sysProcTest1T) Signal(sig os.Signal) error { return errors.New("3") }
func (sp *sysProcTest1T) State() *os.ProcessState    { return nil }
func (sp *sysProcTest1T) ForcedTTY() bool            { return false }

type sysProcTest2T struct{}

func (sp *sysProcTest2T) Pid() int                   { return 4 }
func (sp *sysProcTest2T) ExitNum() int               { return 5 }
func (sp *sysProcTest2T) Kill() error                { return errors.New("6") }
func (sp *sysProcTest2T) Signal(sig os.Signal) error { return nil }
func (sp *sysProcTest2T) State() *os.ProcessState    { return nil }
func (sp *sysProcTest2T) ForcedTTY() bool            { return true }

func TestSystemProcess(t *testing.T) {
	count.Tests(t, (5*2)+1)

	p := lang.NewTestProcess()

	if p.SystemProcess.External() {
		t.Errorf("p.SystemProcess.External() returned true, should be false")
		return
	}

	var success bool

	p.SystemProcess.Set(&sysProcTest1T{})
	switch {
	case !p.SystemProcess.External():
		t.Errorf("invalid return for: p.SystemProcess.External() in test 1")

	case p.SystemProcess.Pid() != 0:
		t.Errorf("invalid return for: p.SystemProcess.Pid() in test 1")

	case p.SystemProcess.ExitNum() != 1:
		t.Errorf("invalid return for: p.SystemProcess.ExitNum() in test 1")

	case p.SystemProcess.Kill() != nil:
		t.Errorf("invalid return for: p.SystemProcess.Kill() in test 1")

	case p.SystemProcess.Signal(nil).Error() != "3":
		t.Errorf("invalid return for: p.SystemProcess.Signal(nil).Error() in test 1")

	default:
		success = true
	}

	if !success {
		return
	}

	success = false

	p.SystemProcess.Set(&sysProcTest2T{})
	switch {
	case !p.SystemProcess.External():
		t.Errorf("invalid return for: p.SystemProcess.External() in test 2")

	case p.SystemProcess.Pid() != 4:
		t.Errorf("invalid return for: p.SystemProcess.Pid() in test 2")

	case p.SystemProcess.ExitNum() != 5:
		t.Errorf("invalid return for: p.SystemProcess.ExitNum() in test 2")

	case p.SystemProcess.Kill().Error() != "6":
		t.Errorf("invalid return for: p.SystemProcess.Kill() in test 2")

	case p.SystemProcess.Signal(nil) != nil:
		t.Errorf("invalid return for: p.SystemProcess.Signal(nil).Error() in test 2")

	default:
		success = true
	}
}

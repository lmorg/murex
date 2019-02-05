// +build ignore

package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func initTest() *Process {
	InitEnv()

	p := NewTestProcess()
	p.Name = "test"

	return p
}

// TestBaseline tests that basline doesn't return empty results which might cause the tests to pass when they shouldn't
func TestBaseline(t *testing.T) {
	p := initTest()
	fork := p.Fork(0)

	if p.Name == "" {
		t.Error("No value set for child process (Name)")
		t.Log("  Expected:", "test")
		t.Log("  Actual:  ", p.Name)
	}

	if fork.parent.Name == "" {
		t.Error("No value set for fork process (Name)")
		t.Log("  Expected:", "")
		t.Log("  Actual:  ", fork.parent.Name)
	}

}

// TestForkF_SHELL tests the F_SHELL flag
func TestForkF_SHELL(t *testing.T) {

	// negative test

	p := initTest()
	fork := p.Fork(0)

	if fork.parent.Name == ShellProcess.Name {
		t.Error("Forked process has inherited shell's process name")
	}

	// positive test

	p = initTest()
	fork = p.Fork(F_SHELL)
	if fork.parent.Name != ShellProcess.Name {
		t.Error("Forked process hasn't inherited shell's process name")
		t.Log("  Expected:", ShellProcess.Name)
		t.Log("  Actual:  ", fork.parent.Name)
	}
}

// TestForkF_NEW_MODULE tests the F_NEW_MODULE flag
func TestForkF_NEW_MODULE(t *testing.T) {
	p := initTest()

	fork := p.Fork(F_NEW_MODULE)
	if fork.parent.Module == ShellProcess.Module {
		t.Error("Forked process has inherited shell's process module name")
		//t.Log("  Expected:", ShellProcess.Module)
		//t.Log("  Actual:  ", fork.parent.Name)
	}
}

// TestForkF_PARENT_VARTABLE tests the F_PARENT_VARTABLE flag
func TestForkF_PARENT_VARTABLE(t *testing.T) {
	InitEnv()

	// negative test
	/*
		block := []rune(`
			set foo=bar; source{$foo}
		`)

		fork := ShellProcess.Fork(F_CREATE_STDOUT | F_PARENT_VARTABLE)
		i, err := fork.Execute(block)
		if i != 0 {
			t.Errorf("None zero exit number: %d", i)
		}
		if err != nil {
			t.Errorf("Error executing block: %s", err.Error())
		}

		b, err := fork.Stdout.ReadAll()
		if err != nil {
			t.Errorf("Error returned from Readall: %s", err.Error())
		}

		if string(b) != "bar" {
			t.Error("Shell should have been set")
			t.Log("  Expected: bar")
			t.Log("  Actual:   ", string(b))
			t.Log("  []byte(): ", b)
		}*/

	// positive test
	p := initTest()

	fork := p.Fork(F_PARENT_VARTABLE)
	if err := fork.Variables.Set("foo", "bar", types.String); err != nil {
		t.Error(err)
		return
	}

	if fork.registerFid {
		t.Error("Forked process should not be a registered FID if the variables are not forked")
	}
}

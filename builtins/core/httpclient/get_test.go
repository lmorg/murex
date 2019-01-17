package httpclient

import (
	"testing"

	"github.com/lmorg/murex/lang/proc"
)

// TestGet tests the get function
func TestGet(t *testing.T) {
	proc.InitEnv()

	p := proc.NewTestProcess()
	p.Config = proc.ShellProcess.Config
	p.Parameters.Params = []string{"https://github.com"}

	err := cmdGet(p)
	if err != nil {
		t.Error(err)
	}
}

// TestGetFile tests the getfile function
func TestGetFile(t *testing.T) {
	proc.InitEnv()

	p := proc.NewTestProcess()
	p.Config = proc.ShellProcess.Config
	p.Parameters.Params = []string{"https://github.com"}

	err := cmdGetFile(p)
	if err != nil {
		t.Error(err)
	}
}

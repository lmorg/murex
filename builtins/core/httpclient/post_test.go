package httpclient

import (
	"testing"

	"github.com/lmorg/murex/lang/proc"
)

// TestPost tests the post function
func TestPost(t *testing.T) {
	proc.InitEnv()

	p := proc.NewTestProcess()
	p.Config = proc.ShellProcess.Config
	p.Parameters.Params = []string{"https://github.com"}

	err := cmdPost(p)
	if err != nil {
		t.Error(err)
	}
}

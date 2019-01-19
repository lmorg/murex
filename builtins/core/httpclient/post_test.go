package httpclient

import (
	"testing"

	"github.com/lmorg/murex/lang"
)

// TestPost tests the post function
func TestPost(t *testing.T) {
	lang.InitEnv()

	p := lang.NewTestProcess()
	p.Config = lang.ShellProcess.Config
	p.Parameters.Params = []string{"https://github.com"}

	err := cmdPost(p)
	if err != nil {
		t.Error(err)
	}
}

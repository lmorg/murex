package httpclient

import (
	"testing"

	"github.com/lmorg/murex/lang"
)

// TestPost tests the post function
func TestPost(t *testing.T) {
	lang.InitEnv()
	addr := StartHTTPServer(t)

	p := lang.NewTestProcess()
	p.Config = lang.ShellProcess.Config
	p.Parameters.Params = []string{addr}

	err := cmdPost(p)
	if err != nil {
		t.Error(err)
	}
}

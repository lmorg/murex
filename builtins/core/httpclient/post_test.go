package httpclient

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

// TestPost tests the post function
func TestPost(t *testing.T) {
	count.Tests(t, 1)

	lang.InitEnv()
	addr := StartHTTPServer(t)

	p := lang.NewTestProcess()
	p.Config = lang.ShellProcess.Config
	p.Parameters.DefineParsed([]string{addr})

	err := cmdPost(p)
	if err != nil {
		t.Error(err)
	}
}

package httpclient

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

// TestGet tests the get function
func TestGet(t *testing.T) {
	count.Tests(t, 1)

	lang.InitEnv()
	addr := StartHTTPServer(t)

	p := lang.NewTestProcess()
	p.Config = lang.ShellProcess.Config
	p.Parameters.Params = []string{addr}

	err := cmdGet(p)
	if err != nil {
		t.Error(err)
	}
}

// TestGetFile tests the getfile function
func TestGetFile(t *testing.T) {
	count.Tests(t, 1)

	lang.InitEnv()
	addr := StartHTTPServer(t)

	p := lang.NewTestProcess()
	p.Config = lang.ShellProcess.Config
	p.Parameters.Params = []string{addr}

	err := cmdGetFile(p)
	if err != nil {
		t.Error(err)
	}
}

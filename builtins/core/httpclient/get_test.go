package httpclient

import (
	"testing"

	"github.com/lmorg/murex/lang"
)

// TestGet tests the get function
func TestGet(t *testing.T) {
	lang.InitEnv()

	p := lang.NewTestProcess()
	p.Config = lang.ShellProcess.Config
	p.Parameters.Params = []string{"https://github.com"}

	err := cmdGet(p)
	if err != nil {
		t.Error(err)
	}
}

// TestGetFile tests the getfile function
func TestGetFile(t *testing.T) {
	lang.InitEnv()

	p := lang.NewTestProcess()
	p.Config = lang.ShellProcess.Config
	p.Parameters.Params = []string{"https://github.com"}

	err := cmdGetFile(p)
	if err != nil {
		t.Error(err)
	}
}

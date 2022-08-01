package open

import (
	"os"
	"testing"

	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

func TestHttp(t *testing.T) {
	if os.Getenv("MUREX_TEST_NO_HTTP") != "" {
		t.Skip("Env var MUREX_TEST_NO_HTTP set so skipping this test")
	}

	count.Tests(t, 1)

	t.Logf("This test requires an active internet connection and thus might fail even when the code is correct. run `export MUREX_TEST_NO_HTTP=true` to skip this test")

	p := lang.NewTestProcess()

	_, dt, err := http(p, "https://api.github.com/repos/lmorg/murex/issues")
	if err != nil {
		t.Error(err.Error())
	}

	if dt != types.Json {
		t.Errorf("API not recognized as JSON: `%s`", dt)
	}
}

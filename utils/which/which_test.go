package which

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestPath(t *testing.T) {
	count.Tests(t, 1, "TestPath")

	path := os.Getenv("PATH")
	if path == "" {
		t.Error("$PATH is empty. This will undoubtably cause problems running murex and will likely cause other tests to fail as well")
	}
}
func TestWhich(t *testing.T) {
	count.Tests(t, 1, "TestWhich")

	if Which("go") == "" {
		t.Error("Which() couldn't find the `go` executable in your $PATH")
		t.Log("$PATH: " + os.Getenv("PATH"))
	}
}

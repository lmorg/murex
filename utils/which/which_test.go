package which

import (
	"os"
	"testing"
)

func TestPath(t *testing.T) {
	path := os.Getenv("PATH")
	if path == "" {
		t.Error("$PATH is empty. This will undoubtably cause problems running murex and will likely cause other tests to fail as well")
	}
}
func TestWhich(t *testing.T) {
	if Which("go") == "" {
		t.Error("Which() couldn't find the `go` executable in your $PATH")
		t.Log("$PATH: " + os.Getenv("PATH"))
	}
}

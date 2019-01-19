package gopath

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test"
)

// TestGoPath checks for the existence of GOPATH and warns the user if it is not present
func TestGoPath(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		t.Error("GOPATH environmental variable is not set")
		t.Log("GOPATH should to be set for murex to do accurate testing otherwise the path is guessed from your current working directory. This might cause some tests to fail if you're working directory is not the root of murex's Go source tree")

		pwd, err := os.Getwd()
		t.Log("Current working directory:", pwd)
		if err != nil {
			t.Log("Error running Getwd():", err)
			t.Log("A failing Getwd means we'll just have to guess where murex's source is :(")
		}

		t.Log("Murex source path assumed to be:", Source([]string{}))
		t.Log("If this isn't correct then expect other `go test`'s to fail")
	}
}

// TestSource just does a quick check that Source() does return the root of
// murex's source tree
func TestSource(t *testing.T) {
	path := Source([]string{})
	t.Log("Source returns:", path)
	test.Exists(t, path+"main.go")
}

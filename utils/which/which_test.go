package which

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestPath(t *testing.T) {
	count.Tests(t, 1)

	path := os.Getenv("PATH")
	if path == "" {
		t.Error("$PATH is empty. This will undoubtably cause problems running murex and will likely cause other tests to fail as well")
	}
}
func TestWhich(t *testing.T) {
	count.Tests(t, 1)

	if Which("go") == "" {
		t.Error("Which() couldn't find the `go` executable in your $PATH")
		t.Log("$PATH: " + os.Getenv("PATH"))
	}
}

func TestWhichDirInCwd(t *testing.T) {
	count.Tests(t, 1)

	tmp := t.TempDir()
	// make a dir with the same name as a real executable
	if err := os.Mkdir(tmp+"/go", 0755); err != nil {
		t.Fatal(err)
	}

	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(old)

	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	result := Which("go")

	if result == "go" || result == "" {
		t.Errorf("Which(\"go\") returned %q; expected the resultolved $PATH path", result)
	}
}

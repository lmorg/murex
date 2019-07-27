package gen

import (
	"os"
	"testing"

	docgen "github.com/lmorg/murex/utils/docgen/api"
)

type logger struct {
	t *testing.T
}

func (l logger) Write(b []byte) (int, error) {
	l.t.Log(string(b))
	return len(b), nil
}

// TestDocgenConfigTemplates tests the config YAML and template files are all
// valid and the project can render
func TestDocgenConfigTemplates(t *testing.T) {
	if _, err := os.Stat("gen/docgen.yaml"); os.IsNotExist(err) {
		os.Chdir("..")
	}

	l := logger{t: t}
	docgen.SetLogger(l)

	docgen.ReadOnly = true
	err := docgen.ReadConfig("gen/docgen.yaml")
	if err != nil {
		t.Error(err)
	}

	err = docgen.Render()
	if err != nil {
		t.Error(err)
	}
}

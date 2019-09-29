package shell

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestAspellInstalled(t *testing.T) {
	test.InstalledDepsTest(t, "aspell")
}

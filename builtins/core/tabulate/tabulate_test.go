package tabulate

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/utils/json"
)

func TestTabulateHelp(t *testing.T) {
	b, err := json.Marshal(desc, false)
	if err != nil {
		t.Fatal(err)
	}

	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		"",
		types.Generic,
		[]string{"--help"},
		string(b),
		nil,
	)
}

package cmdautocomplete

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestMxJsonBug(t *testing.T) {
	block := `[
        {
            "DynamicDesc": ({
                systemctl: --help -> @[..Unit Commands:]s -> tabulate: --column-wraps --map --key-inc-hint --split-space
            }),
            "Optional": true,
            "AllowMultiple": false
        },
        {
            "DynamicDesc": ({
                systemctl: --help -> @[Unit Commands:..]s -> tabulate: --column-wraps --map --key-inc-hint
            }),
            "Optional": false,
            "AllowMultiple": false,
            "FlagValues": {
            }
        }
    ]`

	test.RunMethodTest(
		t, cmdAutocomplete, "autocomplete",
		block, types.Null,
		[]string{"set", "systemctl"},
		"", nil,
	)
}

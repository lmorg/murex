package lang_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParserAtBug(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				autocomplete set systemctl ({[
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
				]})`,
		},
		/*{
			Block: `
				autocomplete set systemctl ({[
					{
						"DynamicDesc": ({
							systemctl: --help -> @foo[..Unit Commands:]s -> tabulate: --column-wraps --map --key-inc-hint --split-space
						}),
						"Optional": true,
						"AllowMultiple": false
					},
					{
						"DynamicDesc": ({
							systemctl: --help -> @bar[Unit Commands:..]s -> tabulate: --column-wraps --map --key-inc-hint
						}),
						"Optional": false,
						"AllowMultiple": false,
						"FlagValues": {
						}
					}
				]})`,
		},*/
	}

	test.RunMurexTests(tests, t)
}

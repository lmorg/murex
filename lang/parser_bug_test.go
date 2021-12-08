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

// https://github.com/lmorg/murex/issues/379
func TestParserBug379(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				out: one\
				     two \
					 three\ # test
					 four \ # test
					 five
				out: six`,
			Stdout: "one two three four five\nsix\n",
		},
		{
			Block: `
				out: one\
				     two \
					 three # test \
					 four # test  \
					 five
				out: six`,
			Stdout: "one two three four five\nsix\n",
		},
	}

	test.RunMurexTests(tests, t)
}

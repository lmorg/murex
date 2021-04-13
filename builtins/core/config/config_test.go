package cmdconfig_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestConfigCLI(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				config: define config test {
					"Description": "This is only a test",
					"DataType": "bool",
					"Global": false,
					"Default": false
				}

				config: get config test
				function TestConfigCLI {
					config: get config test

					config: set config test true
					config: get config test
				}

				config: get config test
				TestConfigCLI
				`,
			Stdout: `false\nfalse\ntrue\nfalse`,
		},
	}

	test.RunMurexTests(tests, t)
}

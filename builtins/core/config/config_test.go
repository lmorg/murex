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
					"Global": true,
					"Default": false
				}

				config: get config test
				
				config: set config test true

				config: get config test
				`,
			Stdout: "false\ntrue\n",
		},
		{
			Block: `
				config: define config test {
					"Description": "This is only a test",
					"DataType": "bool",
					"Global": false,
					"Default": false
				}

				config: get config test
				
				config: set config test true

				config: get config test
				`,
			Stdout: "false\ntrue\n",
		},
		{
			Block: `
				config: define config test {
					"Description": "This is only a test",
					"DataType": "bool",
					"Global": true,
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
				config: get config test
				`,
			Stdout: "false\nfalse\nfalse\ntrue\ntrue\n",
		},
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
				config: get config test
				`,
			Stdout: "false\nfalse\nfalse\ntrue\nfalse\n",
		},
		{
			Block: `
				config: define config test {
					"Description": "This is only a test",
					"DataType": "bool",
					"Global": false,
					"Default": false
				}

				config: get config test
				
				private TestConfigCLI {
					config: get config test

					config: set config test true

					config: get config test
				}

				config: get config test
				TestConfigCLI
				config: get config test
				`,
			Stdout: "false\nfalse\nfalse\ntrue\nfalse\n",
		},
	}

	test.RunMurexTests(tests, t)
}

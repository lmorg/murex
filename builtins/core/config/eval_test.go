package cmdconfig_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestEval(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				config define test-eval 00 %{
					Description: "This is only a test"
					DataType: json
					Global: true
					Default: %{foo: bar, hello: world}
				}`,
		},
		{
			Block:  `config get test-eval 00`,
			Stdout: `{"foo":"bar","hello":"world"}`,
		},
		{
			Block: `
				config eval test-eval 00 { -> alter /foo baz }
				config get test-eval 00`,
			Stdout: `{"foo":"baz","hello":"world"}`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

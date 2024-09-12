package cmdconfig_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestDynamic(t *testing.T) {
	const (
		readMsg = "returned from dynamic reader"
	)

	tests := []test.MurexTest{
		{
			Block: `
				config define test-dynamic 00 %{
					Description: "This is only a test"
					DataType: str
					Global: true
					Dynamic: {
						Read: '{
							out "` + readMsg + `"
						}'
						Write: '{
							<stdin> -> regexp m/^foobar$/
						}'
					}
					Default: "default string"
				}`,
		},
		{
			Block:  `config get test-dynamic 00`,
			Stdout: readMsg + "\n\n",
		},
		{
			Block: `config set test-dynamic 00 foobar`,
		},
		{
			Block: `config set test-dynamic 00 baz`,
			Stderr: `Error`,
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

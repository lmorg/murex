package lang_test

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/test"
)

/* Bug fix:
» function abc { config: set proc echo true ; out: testing; a: [1..3] }
» abc
panic: interface conversion: interface {} is string, not bool

goroutine 922 [running]:
github.com/lmorg/murex/lang.executeProcess(0xc0001082b0)
		/Users/laurencemorgan/go/src/github.com/lmorg/murex/lang/process.go:196 +0x158d
created by github.com/lmorg/murex/lang.runModeNormal
		/Users/laurencemorgan/go/src/github.com/lmorg/murex/lang/interpreter.go:180 +0x7e
murex-dev»
*/
func TestBugFix(t *testing.T) {
	config.InitConf.Define("proc", "echo", config.Properties{
		Description: "Echo shell functions",
		Default:     false,
		DataType:    types.Boolean,
	})
	lang.InitEnv()
	defaults.Defaults(lang.ShellProcess.Config, false)
	shell.SignalHandler(false)

	tests := []test.MurexTest{
		{
			Block: `
				function TestBugFix {
					config: set proc echo true
					out: testing
					a: [1..3]
				}
				TestBugFix
			`,
			Stdout: "testing\n1\n2\n3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func FuzzParseBlock(f *testing.F) {
	tests := []string{"out: hello world", "", "bg { err: abc 123 }"}
	for _, tc := range tests {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		lang.ParseBlock([]rune(orig))
		// we are just testing we can't cause an unhandled panic
	})
}

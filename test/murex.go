package test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/typemgmt"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
)

// MurexTest is a basic framework to test murex code
type MurexTest struct {
	Block   string
	Stdout  string
	Stderr  string
	ExitNum int
}

// RunMurexTests runs through all the test cases for MurexTest
func RunMurexTests(tests []MurexTest, t *testing.T) {
	defaults.Defaults(proc.InitConf, false)
	proc.InitEnv()

	for i := range tests {
		stdout := streams.NewStdin()
		stderr := streams.NewStdin()

		exitNum, err := lang.RunBlockShellNamespace([]rune(tests[i].Block), nil, stdout, stderr)
		if err != nil {
			t.Error(err.Error())
		}

		b, err := stderr.ReadAll()
		if err != nil {
			t.Error(tests[i].Block, "- unable to read from stderr: "+err.Error())
		}

		if string(b) != tests[i].Stderr {
			t.Error(tests[i].Block, "- stderr doesn't match exected error message")
		}

		b, err = stdout.ReadAll()
		if err != nil {
			t.Error(tests[i].Block, "- unable to read from stdout: "+err.Error())
		}

		if string(b) != tests[i].Stdout {
			t.Error(tests[i].Block, "- stdout doesn't match exected output")
		}

		if exitNum != tests[i].ExitNum {
			t.Error(tests[i].Block, "- exit number doesn't match expected exit number")
		}
	}
}

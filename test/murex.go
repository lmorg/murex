package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test/count"
)

// MurexTest is a basic framework to test murex code.
// Please note this shouldn't be confused with the murex scripting language's inbuilt testing framework!
type MurexTest struct {
	Block   string
	Stdout  string
	Stderr  string
	ExitNum int
}

// RunMurexTests runs through all the test cases for MurexTest where STDOUT/ERR are literal strings
func RunMurexTests(tests []MurexTest, t *testing.T) {
	t.Helper()
	count.Tests(t, len(tests))

	defaults.Config(config.InitConf, false)
	lang.InitEnv()

	for i := range tests {
		hasError := false

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
		go func(testNum int) {
			time.Sleep(30 * time.Second)
			if !fork.Process.HasCancelled() {
				panic(fmt.Sprintf("timeout in %s() test %d", t.Name(), testNum))
			}
		}(i)

		fork.Name.Set("RunMurexTests()")
		fork.FileRef = &ref.File{Source: &ref.Source{Module: fmt.Sprintf("murex/%s-%d", t.Name(), i)}}
		exitNum, err := fork.Execute([]rune(tests[i].Block))
		if err != nil {
			t.Errorf("Cannot execute script on test %d", i)
			t.Log("  ", err)
			continue
		}

		bErr, err := fork.Stderr.ReadAll()
		if err != nil {
			t.Errorf("Cannot ReadAll() from Stderr on test %d", i)
			t.Log("  ", err)
			continue
		}

		if string(bErr) != tests[i].Stderr {
			hasError = true
		}

		bOut, err := fork.Stdout.ReadAll()
		if err != nil {
			t.Errorf("Cannot ReadAll() from Stdout on test %d", i)
			t.Log("  ", err)
			continue
		}

		if string(bOut) != tests[i].Stdout {
			hasError = true
		}

		if exitNum != tests[i].ExitNum {
			hasError = true
		}

		if hasError {
			t.Errorf("Code block doesn't return expected values in test %d", i)
			t.Log("  Code block:      ", tests[i].Block)

			t.Log("  Expected Stdout: ", strings.Replace(tests[i].Stdout, "\n", `\n`, -1))
			t.Log("  Actual Stdout:   ", strings.Replace(string(bOut), "\n", `\n`, -1))
			t.Log("  Exp. out bytes:  ", []byte(tests[i].Stdout))
			t.Log("  Act. out bytes:  ", bOut)

			t.Log("  Expected Stderr: ", strings.Replace(tests[i].Stderr, "\n", `\n`, -1))
			t.Log("  Actual Stderr:   ", strings.Replace(string(bErr), "\n", `\n`, -1))
			t.Log("  Exp. err bytes:  ", []byte(tests[i].Stderr))
			t.Log("  Act. err bytes:  ", bErr)

			t.Log("  Expected exitnum:", tests[i].ExitNum)
			t.Log("  Actual exitnum:  ", exitNum)
		}
	}
}

// RunMurexTestsRx runs through all the test cases for MurexTest where STDOUT/ERR are regexp expressions
func RunMurexTestsRx(tests []MurexTest, t *testing.T) {
	t.Helper()
	count.Tests(t, len(tests))

	defaults.Config(config.InitConf, false)
	lang.InitEnv()

	for i := range tests {
		hasError := false

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
		fork.Name.Set("RunMurexTestsRx()")
		fork.FileRef = &ref.File{Source: &ref.Source{Module: fmt.Sprintf("murex/%s-%d", t.Name(), i)}}
		exitNum, err := fork.Execute([]rune(tests[i].Block))
		if err != nil {
			t.Errorf("Cannot execute script on test %d", i)
			t.Log("  ", err)
			continue
		}

		bErr, err := fork.Stderr.ReadAll()
		if err != nil {
			t.Errorf("Cannot ReadAll() from Stderr on test %d", i)
			t.Log("  ", err)
			continue
		}

		rxErr, err := regexp.Compile(tests[i].Stderr)
		if err != nil {
			t.Errorf("Cannot compile regexp expression for STDERR on test %d", i)
			t.Log("  ", err)
			continue
		}
		if !rxErr.MatchString(string(bErr)) {
			hasError = true
		}

		bOut, err := fork.Stdout.ReadAll()
		if err != nil {
			t.Errorf("Cannot ReadAll() from Stdout on test %d", i)
			t.Log("  ", err)
			continue
		}

		rxOut, err := regexp.Compile(tests[i].Stdout)
		if err != nil {
			t.Errorf("Cannot compile regexp expression for STDOUT on test %d", i)
			t.Log("  ", err)
			continue
		}
		if !rxOut.MatchString(string(bOut)) {
			hasError = true
		}

		if exitNum != tests[i].ExitNum {
			hasError = true
		}

		if hasError {
			t.Errorf("Code block doesn't return expected values in test %d", i)
			t.Log("  Code block:      ", tests[i].Block)

			t.Log("  Expected out rx: ", tests[i].Stdout)
			t.Log("  Actual Stdout:   ", strings.Replace(string(bOut), "\n", `\n`, -1))
			t.Log("  Act. out bytes:  ", bOut)

			t.Log("  Expected err rx: ", tests[i].Stderr)
			t.Log("  Actual Stderr:   ", strings.Replace(string(bErr), "\n", `\n`, -1))
			t.Log("  Act. err bytes:  ", bErr)

			t.Log("  Expected exitnum:", tests[i].ExitNum)
			t.Log("  Actual exitnum:  ", exitNum)
		}
	}
}

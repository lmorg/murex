package proc

import (
	"errors"
	"fmt"
	"regexp"
	"sync"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

// TestProperties are the values prescribed to an individual test case
type TestProperties struct {
	Name       string
	out        *TestChecks
	err        *TestChecks
	exitNumPtr *int
	exitNum    int
}

type TestChecks struct {
	stdio    stdio.Io
	Regexp   *regexp.Regexp
	Block    []rune
	RunBlock func(*Process, []rune) ([]byte, error)
}

type TestResults struct {
	Passed     bool
	TestName   string
	Message    string
	Exec       string
	Params     []string
	LineNumber int
	ColNumber  int
}

type Tests struct {
	mutex   sync.Mutex
	test    []*TestProperties
	Results []TestResults
}

// NewTests creates a new testing scope
func NewTests() (tests *Tests) {
	tests = new(Tests)
	return
}

func (tests *Tests) Define(name string, out *TestChecks, err *TestChecks, exitNum int) error {
	tests.mutex.Lock()

	var i int
	for ; i < len(tests.test); i++ {
		if tests.test[i].Name == name {
			goto define
		}
	}

	tests.test = append(tests.test, &TestProperties{
		Name:    name,
		out:     out,
		err:     err,
		exitNum: exitNum,
	})

	tests.mutex.Unlock()
	return nil

define:
	tests.mutex.Unlock()
	return errors.New("Test already defined for '" + name + "' in this scope.")
}

func (tests *Tests) SetStreams(name string, stdout, stderr stdio.Io, exitNumPtr *int) error {
	tests.mutex.Lock()

	var i int
	for ; i < len(tests.test); i++ {
		if tests.test[i].Name == name {
			goto set
		}
	}

	tests.mutex.Unlock()
	return errors.New("Test named but there is no test defined for '" + name + "'.")

set:
	tests.test[i].out.stdio = stdout
	tests.test[i].err.stdio = stderr
	tests.test[i].exitNumPtr = exitNumPtr
	tests.mutex.Unlock()
	return nil
}

func (tests *Tests) CloseTest(name string) {
	tests.mutex.Lock()

	var i int
	for ; i < len(tests.test); i++ {
		if tests.test[i].Name == name {
			goto set
		}
	}

	tests.mutex.Unlock()
	return //errors.New("Test named but there is no test defined for '" + name + "'.")

set:
	tests.test[i].out.stdio.Close()
	tests.test[i].err.stdio.Close()
	tests.mutex.Unlock()
	return //nil
}

func (tests *Tests) AddResult(test *TestProperties, p *Process, passed bool, message string) {
	//tests.Results.murex.Lock()
	tests.Results = append(tests.Results, TestResults{
		TestName:   test.Name,
		Exec:       p.Name,
		Params:     p.Parameters.StringArray(),
		LineNumber: p.LineNumber,
		ColNumber:  p.ColNumber,
		Passed:     passed,
		Message:    message,
	})
	//tests.Results.murex.Unlock()
}

func (tests *Tests) WriteResults(config *config.Config, pipe stdio.Io) error {
	tests.mutex.Lock()
	defer tests.mutex.Unlock()

	if len(tests.Results) == 0 {
		return nil
	}

	reportType, err := config.Get("test", "report-format", types.String)
	if err != nil {
		return err
	}

	switch reportType.(string) {
	case "json":
		b, err := json.Marshal(tests.Results, pipe.IsTTY())
		if err != nil {
			return err
		}

		tests.Results = make([]TestResults, 0)

		_, err = pipe.Writeln(b)
		return err

	case "table":
		s := "[STATUS] Line Col. Function   Message"
		pipe.Writeln([]byte(s))
		for i := range tests.Results {
			s = fmt.Sprintf(" %-4d %-4d %-10s %s\n",
				tests.Results[i].LineNumber,
				tests.Results[i].ColNumber,
				tests.Results[i].Exec,
				tests.Results[i].Message,
			)

			if allowAnsi() {
				if tests.Results[i].Passed {
					pipe.Write([]byte("\x1b[32m" + passFail(tests.Results[i].Passed) + "\x1b[0m"))
				} else {
					pipe.Write([]byte("\x1b[31m" + passFail(tests.Results[i].Passed) + "\x1b[0m"))
				}
			} else {
				pipe.Write([]byte(passFail(tests.Results[i].Passed)))
			}
			pipe.Write([]byte(s))

		}
		return nil

	default:
		return errors.New("Invalid report type requested via `config set test report-format`.")
	}
}

func passFail(passed bool) string {
	if passed {
		return "[PASSED]"
	}
	return "[FAILED]"
}

func allowAnsi() bool {
	v, err := ShellProcess.Config.Get("shell", "add-colour", types.Boolean)
	if err != nil {
		return false
	}
	return v.(bool)
}

func (tests *Tests) Dump() []string {
	tests.mutex.Lock()

	names := make([]string, 0)

	for _, ptr := range tests.test {
		names = append(names, ptr.Name)
	}

	tests.mutex.Unlock()
	return names
}

func (tests *Tests) Compare(name string, p *Process) {
	tests.mutex.Lock()

	var i int
	for ; i < len(tests.test); i++ {
		if tests.test[i].Name == name {
			goto compare
		}
	}

	tests.mutex.Unlock()
	return //errors.New("Test named but there is no test defined for '" + name + "'.")

compare:

	var failed bool
	test := tests.test[i]
	tests.mutex.Unlock()

	// read stdout
	stdout, err := test.out.stdio.ReadAll()
	if err != nil {
		failed = true
		tests.AddResult(test, p, !failed, "Cannot read from stdout.")
	}
	stdout = utils.CrLfTrim(stdout)

	// read stderr
	stderr, err := test.err.stdio.ReadAll()
	if err != nil {
		failed = true
		tests.AddResult(test, p, !failed, "Cannot read from stderr.")
	}
	stderr = utils.CrLfTrim(stderr)

	// compare stdout
	if len(test.out.Block) > 0 {
		b, err := test.out.RunBlock(p, test.out.Block)
		if err != nil {
			failed = true
			tests.AddResult(test, p, !failed, err.Error())
		}
		if string(b) != string(stdout) {
			failed = true
			tests.AddResult(test, p, !failed,
				fmt.Sprintf("stdout: wanted '%s' got '%s'.",
					b, stdout))
		}

	} else if test.out.Regexp != nil {
		if !test.out.Regexp.Match(stdout) {
			failed = true
			tests.AddResult(test, p, !failed,
				fmt.Sprintf("stdout: regexp did not match '%s'.",
					stdout))
		}
	}

	// compare stderr
	if len(test.err.Block) > 0 {
		b, err := test.err.RunBlock(p, test.err.Block)
		if err != nil {
			failed = true
			tests.AddResult(test, p, !failed, err.Error())
		}
		if string(b) != string(stderr) {
			failed = true
			tests.AddResult(test, p, !failed,
				fmt.Sprintf("stderr: wanted '%s' got '%s'.",
					b, stderr))
		}

	} else if test.err.Regexp != nil {
		if !test.err.Regexp.Match(stderr) {
			failed = true
			tests.AddResult(test, p, !failed,
				fmt.Sprintf("stderr: regexp did not match '%s'.",
					stderr))
		}
	}

	// test exit number
	if test.exitNum != *test.exitNumPtr {
		failed = true
		tests.AddResult(test, p, !failed,
			fmt.Sprintf("exit number: wanted %d got %d.",
				test.exitNum, *test.exitNumPtr))
	}

	// if not failed, log a success result
	if !failed {
		tests.AddResult(test, p, true, "")
	}
}

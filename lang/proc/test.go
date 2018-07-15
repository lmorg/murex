package proc

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/json"
)

// TestProperties are the values prescribed to an individual test case
type TestProperties struct {
	Name       string
	out        *TestChecks
	err        *TestChecks
	exitNumPtr *int
	exitNum    int
	HasRan     bool
}

type TestChecks struct {
	stdio    stdio.Io
	Regexp   *regexp.Regexp
	Block    []rune
	RunBlock func(*Process, []rune) ([]byte, error)
}

type TestResults struct {
	Status     TestStatus
	TestName   string
	Message    string
	Exec       string
	Params     []string
	LineNumber int
	ColNumber  int
}

type TestStatus string

const (
	TestPassed TestStatus = "PASSED"
	TestFailed TestStatus = "FAILED"
	TestError  TestStatus = "ERROR"
	TestMissed TestStatus = "MISSED"
)

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

func (tests *Tests) AddResult(test *TestProperties, p *Process, status TestStatus, message string) {
	//tests.Results.murex.Lock()
	tests.Results = append(tests.Results, TestResults{
		TestName:   test.Name,
		Exec:       p.Name,
		Params:     p.Parameters.StringArray(),
		LineNumber: p.LineNumber,
		ColNumber:  p.ColNumber,
		Status:     status,
		Message:    message,
	})
	//tests.Results.murex.Unlock()
}

func (tests *Tests) WriteResults(config *config.Config, pipe stdio.Io) error {
	allowAnsi := func() bool {
		v, err := ShellProcess.Config.Get("shell", "add-colour", types.Boolean)
		if err != nil {
			return false
		}
		return v.(bool)
	}

	params := func(exec string, params []string) (s string) {
		s = exec + " '" + strings.Join(params, "' '") + "'"
		if len(s) > 50 {
			s = s[:49] + "…"
		}
		return
	}

	tests.mutex.Lock()
	defer tests.mutex.Unlock()

	if len(tests.Results) == 0 {
		return nil
	}

	reportType, err := config.Get("test", "report-format", types.String)
	if err != nil {
		return err
	}

	reportPipe, err := config.Get("test", "report-pipe", types.String)
	if err != nil {
		reportPipe = ""
	}

	if reportPipe.(string) != "" {
		pipe, err = GlobalPipes.Get(reportPipe.(string))
		if err != nil {
			return err
		}
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
		if reportPipe.(string) == "" {
			pipe.Writeln([]byte(consts.TestTableHeadings))
		}
		for i := range tests.Results {
			if allowAnsi() {
				switch tests.Results[i].Status {
				case TestPassed:
					pipe.Write([]byte("[\x1b[32m"))
				case TestFailed, TestError:
					pipe.Write([]byte("[\x1b[31m"))
				case TestMissed:
					pipe.Write([]byte("[\x1b[34m"))
				}
			}

			s := fmt.Sprintf("%s\x1b[0m] %-4d %-4d %-50s %s\n",
				tests.Results[i].Status,
				tests.Results[i].LineNumber,
				tests.Results[i].ColNumber,
				params(tests.Results[i].Exec, tests.Results[i].Params),
				tests.Results[i].Message,
			)

			pipe.Write([]byte(s))

		}
		return nil

	default:
		return errors.New("Invalid report type requested via `config set test report-format`.")
	}
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
	tests.AddResult(tests.test[i], p, TestError, "Test named but there is no test defined.")
	return //errors.New("Test named but there is no test defined for '" + name + "'.")

compare:

	var failed bool
	test := tests.test[i]
	test.HasRan = true
	tests.mutex.Unlock()

	left := func(b []byte) []byte {
		crop, err := p.Config.Get("test", "crop-message", types.Integer)
		if err != nil || crop.(int) == 0 {
			return b
		}

		if len(b) < crop.(int) {
			return b
		}

		return append(b[:crop.(int)-1], []byte(string([]rune{'…'}))...)
	}

	// read stdout
	stdout, err := test.out.stdio.ReadAll()
	if err != nil {
		failed = true
		tests.AddResult(test, p, TestError, "Cannot read from stdout.")
	}
	stdout = utils.CrLfTrim(stdout)

	// read stderr
	stderr, err := test.err.stdio.ReadAll()
	if err != nil {
		failed = true
		tests.AddResult(test, p, TestError, "Cannot read from stderr.")
	}
	stderr = utils.CrLfTrim(stderr)

	// compare stdout
	if len(test.out.Block) > 0 {
		b, err := test.out.RunBlock(p, test.out.Block)
		if err != nil {
			failed = true
			tests.AddResult(test, p, TestError, err.Error())
		}
		if string(b) != string(stdout) {
			failed = true
			tests.AddResult(test, p, TestFailed,
				fmt.Sprintf("stdout: wanted '%s' got '%s'.",
					left(b), left(stdout)))
		}

	} else if test.out.Regexp != nil {
		if !test.out.Regexp.Match(stdout) {
			failed = true
			tests.AddResult(test, p, TestFailed,
				fmt.Sprintf("stdout: regexp did not match '%s'.",
					left(stdout)))
		}
	}

	// compare stderr
	if len(test.err.Block) > 0 {
		b, err := test.err.RunBlock(p, test.err.Block)
		if err != nil {
			failed = true
			tests.AddResult(test, p, TestError, err.Error())
		}
		if string(b) != string(stderr) {
			failed = true
			tests.AddResult(test, p, TestFailed,
				fmt.Sprintf("stderr: wanted '%s' got '%s'.",
					left(b), left(stderr)))
		}

	} else if test.err.Regexp != nil {
		if !test.err.Regexp.Match(stderr) {
			failed = true
			tests.AddResult(test, p, TestFailed,
				fmt.Sprintf("stderr: regexp did not match '%s'.",
					left(stderr)))
		}
	}

	// test exit number
	if test.exitNum != *test.exitNumPtr {
		failed = true
		tests.AddResult(test, p, TestFailed,
			fmt.Sprintf("exit number: wanted %d got %d.",
				test.exitNum, *test.exitNumPtr))
	}

	// if not failed, log a success result
	if !failed {
		tests.AddResult(test, p, TestPassed, "All test conditions were met.")
	}
}

func (tests *Tests) ReportMissedTests(p *Process) {
	tests.mutex.Lock()

	for _, test := range tests.test {
		if test.HasRan {
			continue
		}

		tests.AddResult(test, p, TestMissed, "Test was defined but no function ran against that test pipe.")
	}

	tests.mutex.Unlock()
}

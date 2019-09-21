package lang

/*
	This test library relates to the testing framework within the murex
	language itself rather than Go's test framework within the murex project.

	The naming convention here is basically the inverse of Go's test naming
	convention. ie Go source files will be named "test_unit.go" (because
	calling it unit_test.go would mean it's a Go test rather than murex test)
	and the code is named UnitTestPlans (etc) rather than TestUnitPlans (etc)
	because the latter might suggest they would be used by `go test`. This
	naming convention is a little counterintuitive but it at least avoids
	naming conflicts with `go test`.
*/

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/json"
)

// SetStreams is called when a particular test case is run. eg
//
//     out <test_example> "Run this test"
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

// AddResult is called after the test has run so the result can be recorded
func (tests *Tests) AddResult(test *TestProperties, p *Process, status TestStatus, message string) {
	tests.Results.Add(&TestResult{
		TestName:   test.Name,
		Exec:       p.Name,
		Params:     p.Parameters.StringArray(),
		LineNumber: p.FileRef.Line,
		ColNumber:  p.FileRef.Column,
		Status:     status,
		Message:    message,
	})
}

// WriteResults is the reporting tool
func (tests *Tests) WriteResults(config *config.Config, pipe stdio.Io) error {
	params := func(exec string, params []string) (s string) {
		if len(params) > 1 {
			//s = exec + " '" + strings.Join(params, "' '") + "'"
			s = exec + " " + strings.Join(params, " ")
		} else {
			s = exec
		}
		if len(s) > 50 {
			s = s[:49] + "…"
		}
		return
	}

	escape := func(s string) string {
		s = strings.ReplaceAll(s, "\n", `\n`)
		s = strings.ReplaceAll(s, "\r", `\r`)
		s = strings.ReplaceAll(s, "\t", `\t`)
		return s
	}

	left := func(s string) string {
		crop, err := config.Get("test", "crop-message", types.Integer)
		if err != nil || crop.(int) == 0 {
			return s
		}

		if len(s) < crop.(int) {
			return s
		}

		return s[:crop.(int)-1] + "…"
	}

	tests.mutex.Lock()
	defer tests.mutex.Unlock()

	if tests.Results.Len() == 0 {
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

	verbose, err := config.Get("test", "verbose", types.Boolean)
	if err != nil {
		verbose = false
	}

	ansiColour, err := config.Get("shell", "color", types.Boolean)
	if err != nil {
		ansiColour = false
	}

	if reportPipe.(string) != "" {
		pipe, err = GlobalPipes.Get(reportPipe.(string))
		if err != nil {
			return err
		}
	}

	defer func() {
		tests.Results.results = make([]*TestResult, 0)
	}()

	switch reportType.(string) {
	case "json":
		pipe.SetDataType(types.Json)

		b, err := json.Marshal(tests.Results.results, pipe.IsTTY())
		if err != nil {
			return err
		}

		_, err = pipe.Writeln(b)
		return err

	case "table":
		pipe.SetDataType(types.Generic)

		if reportPipe.(string) == "" {
			pipe.Writeln([]byte(consts.TestTableHeadings))
		}

		var colour, reset string
		if ansiColour.(bool) {
			reset = "\x1b[0m"
		}
		for _, r := range tests.Results.results {
			if !verbose.(bool) && (r.Status == TestMissed || r.Status == TestInfo) {
				continue
			}

			if ansiColour.(bool) {
				switch r.Status {
				case TestPassed:
					colour = "\x1b[32m"
				case TestFailed, TestError:
					colour = "\x1b[31m"
				case TestMissed, TestInfo:
					colour = "\x1b[34m"
				case TestState:
					colour = "\x1b[33m"
				}
			}

			s := fmt.Sprintf("[%s%-6s%s] %-10s %-50s %-4d %-4d %s",
				colour, r.Status, reset,
				r.TestName,
				params(r.Exec, r.Params),
				r.LineNumber,
				r.ColNumber,
				left(escape(r.Message)),
			)

			pipe.Writeln([]byte(s))

		}
		return nil

	case "csv":
		pipe.SetDataType("csv")
		s := fmt.Sprintf(`%s %-13s %-53s %-7s %-7s %s`,
			`"Status",`,
			`"Test Name",`,
			`"Process",`,
			`"Line",`,
			`"Col.",`,
			`"Message"`,
		)
		pipe.Writeln([]byte(s))

		for _, r := range tests.Results.results {
			if !verbose.(bool) && (r.Status == TestMissed || r.Status == TestInfo) {
				continue
			}

			s = fmt.Sprintf(`%-9s %-13s %-53s %6d, %6d, %s`,
				`"`+r.Status+`",`,
				`"`+r.TestName+`",`,
				`"`+params(r.Exec, r.Params)+`",`,
				r.LineNumber,
				r.ColNumber,
				`"`+strings.ReplaceAll(escape(r.Message), `"`, `""`)+`"`,
			)

			pipe.Writeln([]byte(s))

		}
		return nil

	default:
		return errors.New("Invalid report type requested via `config set test report-format`")
	}
}

package lang

import (
	"runtime/debug"
	"testing"

	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type parserTestConditions struct {
	Block    string
	Expected AstNodes
}

func testParser(t *testing.T, tests []parserTestConditions) {
	var j, i int

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Panic caught in test %d, node %d: %s\n%s\n%s", j, i, tests[j].Block, r, string(debug.Stack()))
		}
	}()

	t.Helper()
	count.Tests(t, len(tests))

	InitEnv()

	for j = range tests {
		exp := tests[j].Expected

		nodes, pErr := parser([]rune(tests[j].Block))
		if pErr.Code != 0 {
			jsonExp, _ := json.Marshal(exp, false)
			t.Error("Unable to parse valid command line:")
			t.Logf("  Test #:    %d", j)
			t.Logf("  Block:     %s", tests[j].Block)
			t.Logf("  Error Msg: %s", pErr.Message)
			t.Logf("  Expected:  %s", string(jsonExp))
			t.Logf("  Actual:    n/a")
			continue
		}

		// compile just in case any parser validation is deferred to the compiler
		_, errNo := compile(nodes, ShellProcess)
		if errNo != 0 {
			jsonExp, _ := json.Marshal(exp, false)
			t.Error("Unable to compile valid command line:")
			t.Logf("  Test #:    %d", j)
			t.Logf("  Block:     %s", tests[j].Block)
			t.Logf("  Error Msg: %s", errMessages[errNo])
			t.Logf("  Expected:  %s", string(jsonExp))
			t.Logf("  Actual:    n/a")
			continue
		}

		if len(*nodes) != len(exp) {
			t.Error("Parsing failed; number of nodes expected did not match returned:")
			jsonNodes, _ := json.Marshal(nodes, false)
			jsonExp, _ := json.Marshal(exp, false)
			t.Logf("  Test #:     %d", j)
			t.Logf("  Block:      %s", tests[j].Block)
			t.Logf("  Node count: %d exp, %d actual", len(exp), len(*nodes))
			t.Logf("  Expected:   %s", string(jsonExp))
			t.Logf("  Actual:     %s", string(jsonNodes))
			continue
		}

		for i = range *nodes {
			switch {

			case (*nodes)[i].NewChain != exp[i].NewChain:
				t.Error("Parsing failed; NewChain mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v", exp[i].NewChain)
				t.Logf("  Actual:   %v", (*nodes)[i].NewChain)

			case (*nodes)[i].PipeOut != exp[i].PipeOut:
				t.Error("Parsing failed; PipeOut mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v", exp[i].PipeOut)
				t.Logf("  Actual:   %v", (*nodes)[i].PipeOut)

			case (*nodes)[i].PipeErr != exp[i].PipeErr:
				t.Error("Parsing failed; PipeErr mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v", exp[i].PipeErr)
				t.Logf("  Actual:   %v", (*nodes)[i].PipeErr)

			case (*nodes)[i].Method != exp[i].Method:
				t.Error("Parsing failed; Method mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v", exp[i].Method)
				t.Logf("  Actual:   %v", (*nodes)[i].Method)

			case (*nodes)[i].Name != exp[i].Name:
				t.Error("Parsing failed; Name mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %s", exp[i].Name)
				t.Logf("  Actual:   %s", (*nodes)[i].Name)

			default:
				jsonNodes, _ := json.Marshal((*nodes)[i].ParamTokens, false)
				jsonExp, _ := json.Marshal(exp[i].ParamTokens, false)
				if string(jsonNodes) != string(jsonExp) {
					t.Error("Parsing failed; Parameter mismatch:")
					t.Logf("  Test #:   %d", j)
					t.Logf("  Block:    %s", tests[j].Block)
					t.Logf("  Node #:   %d", i)
					t.Logf("  Expected: %s", jsonExp)
					t.Logf("  Actual:   %s", jsonNodes)
				}
			}
		}
	}
}

func queryParser(t *testing.T, block string) string {
	t.Helper()

	nodes, pErr := parser([]rune(block))
	if pErr.Code != 0 {
		t.Errorf("Error at %d: %s", pErr.EndByte, pErr.Message)
	}

	b, err := json.Marshal(nodes, true)
	if err != nil {
		t.Errorf("Error marshalling nodes as JSON: %s", err)
	}

	return string(b)
}

type parserTestSimpleConditions struct {
	Block    string
	Expected []parserTestSimpleExpected
	Error    bool
}

type parserTestSimpleExpected struct {
	Name       string
	Parameters []string
	Method     int
}

const (
	TEST_NEW_PIPE = 1 << iota
	TEST_PIPE_OUT
	TEST_PIPE_ERR
	TEST_METHOD
)

func testParserSimple(t *testing.T, tests []parserTestSimpleConditions) {
	var j, i int

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Panic caught in test %d, node %d: %s\n%s\n%s", j, i, tests[j].Block, r, string(debug.Stack()))
		}
	}()

	t.Helper()
	count.Tests(t, len(tests))

	InitEnv()

	for j = range tests {
		exp := tests[j].Expected

		nodes, pErr := parser([]rune(tests[j].Block))
		if pErr.Code != 0 {
			if !tests[j].Error {
				jsonExp, _ := json.Marshal(exp, false)
				t.Error("Unable to parse valid command line:")
				t.Logf("  Test #:    %d", j)
				t.Logf("  Block:     %s", tests[j].Block)
				t.Logf("  Error Msg: %s", pErr.Message)
				t.Logf("  Expected:  %s", string(jsonExp))
				t.Logf("  Actual:    n/a")
			}
			continue
		}

		// compile just in case any parser validation is deferred to the compiler
		_, errNo := compile(nodes, ShellProcess)
		if errNo != 0 {
			if !tests[j].Error {
				jsonExp, _ := json.Marshal(exp, false)
				t.Error("Unable to compile valid command line:")
				t.Logf("  Test #:    %d", j)
				t.Logf("  Block:     %s", tests[j].Block)
				t.Logf("  Error Msg: %s", errMessages[errNo])
				t.Logf("  Expected:  %s", string(jsonExp))
				t.Logf("  Actual:    n/a")
			}
			continue
		}

		if tests[j].Error {
			jsonExp, _ := json.Marshal(exp, false)
			jsonAct, _ := json.Marshal(nodes, false)
			t.Error("Error expected, no parser nor compiler error raised:")
			t.Logf("  Test #:    %d", j)
			t.Logf("  Block:     %s", tests[j].Block)
			t.Logf("  len(exp):  %d", len(exp))
			t.Logf("  len(nodes):%d", len(*nodes))
			t.Logf("  Expected:  %s", string(jsonExp))
			t.Logf("  Actual:    %s", string(jsonAct))
			continue
		}

		if len(*nodes) != len(exp) {
			jsonExp, _ := json.Marshal(exp, false)
			jsonAct, _ := json.Marshal(nodes, false)
			t.Error("Number of nodes expected different to actual:")
			t.Logf("  Test #:    %d", j)
			t.Logf("  Block:     %s", tests[j].Block)
			t.Logf("  len(exp):  %d", len(exp))
			t.Logf("  len(nodes):%d", len(*nodes))
			t.Logf("  Expected:  %s", string(jsonExp))
			t.Logf("  Actual:    %s", string(jsonAct))
			continue
		}

		for i = range *nodes {
			switch {

			case (*nodes)[i].NewChain != (exp[i].Method&TEST_NEW_PIPE != 0):
				t.Error("Parsing failed; NewChain mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v (%d)", exp[i].Method&TEST_NEW_PIPE != 0, exp[i].Method)
				t.Logf("  Actual:   %v", (*nodes)[i].NewChain)

			case (*nodes)[i].PipeOut != (exp[i].Method&TEST_PIPE_OUT != 0):
				t.Error("Parsing failed; PipeOut mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v (%d)", exp[i].Method&TEST_PIPE_OUT != 0, exp[i].Method)
				t.Logf("  Actual:   %v", (*nodes)[i].PipeOut)

			case (*nodes)[i].PipeErr != (exp[i].Method&TEST_PIPE_ERR != 0):
				t.Error("Parsing failed; PipeErr mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v (%d)", exp[i].Method&TEST_PIPE_ERR != 0, exp[i].Method)
				t.Logf("  Actual:   %v", (*nodes)[i].PipeErr)

			case (*nodes)[i].Method != (exp[i].Method&TEST_METHOD != 0):
				t.Error("Parsing failed; Method mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v (%d)", exp[i].Method&TEST_METHOD != 0, exp[i].Method)
				t.Logf("  Actual:   %v", (*nodes)[i].Method)

			case (*nodes)[i].Name != exp[i].Name:
				t.Error("Parsing failed; Name mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %s", exp[i].Name)
				t.Logf("  Actual:   %s", (*nodes)[i].Name)

			default:
				params := parameters.Parameters{Tokens: (*nodes)[i].ParamTokens}
				ParseParameters(ShellProcess, &params)

				if params.Len() != len(exp[i].Parameters) {
					var jsonExp, jsonAct string
					b, err := json.Marshal(exp[i].Parameters, true)
					if err != nil {
						jsonExp = err.Error()
					} else {
						jsonExp = string(b)
					}
					b, err = json.Marshal(params.StringArray(), true)
					if err != nil {
						jsonAct = err.Error()
					} else {
						jsonAct = string(b)
					}
					t.Error("Parsing failed; Invalid param len():")
					t.Logf("  Test #:   %d", j)
					t.Logf("  Block:    %s", tests[j].Block)
					t.Logf("  Node #:   %d", i)
					t.Logf("  Expected: %d (%s)", len(exp[i].Parameters), jsonExp)
					t.Logf("  Actual:   %d (%s)", params.Len(), jsonAct)
					continue
				}

				for p, actual := range params.StringArray() {
					if actual != exp[i].Parameters[p] {
						t.Error("Parsing failed; Invalid parameter:")
						t.Logf("  Test #:   %d", j)
						t.Logf("  Block:    %s", tests[j].Block)
						t.Logf("  Node #:   %d", i)
						t.Logf("  Param #:  %d", p)
						t.Logf("  Expected: %s", exp[i].Parameters[p])
						t.Logf("  Actual:   %s", actual)
						continue
					}
				}
			}
		}
	}
}

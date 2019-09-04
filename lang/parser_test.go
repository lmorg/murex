package lang

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func testParser(t *testing.T, tests []parserTestConditions) {
	count.Tests(t, len(tests), "testParser")
	for j := range tests {
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

		if len(nodes) != len(exp) {
			t.Error("Parsing failed; number of nodes expected did not match returned:")
			jsonNodes, _ := json.Marshal(nodes, false)
			jsonExp, _ := json.Marshal(exp, false)
			t.Logf("  Test #:     %d", j)
			t.Logf("  Block:      %s", tests[j].Block)
			t.Logf("  Node count: %d exp, %d actual", len(exp), len(nodes))
			t.Logf("  Expected:   %s", string(jsonExp))
			t.Logf("  Actual:     %s", string(jsonNodes))
			continue
		}

		for i := range nodes {
			switch {

			case nodes[i].NewChain != exp[i].NewChain:
				t.Error("Parsing failed; NewChain mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v", exp[i].NewChain)
				t.Logf("  Actual:   %v", nodes[i].NewChain)

			case nodes[0].PipeOut != exp[i].PipeOut:
				t.Error("Parsing failed; PipeOut mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v", exp[i].PipeOut)
				t.Logf("  Actual:   %v", nodes[i].PipeOut)

			case nodes[0].PipeErr != exp[i].PipeErr:
				t.Error("Parsing failed; PipeErr mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v", exp[i].PipeErr)
				t.Logf("  Actual:   %v", nodes[i].PipeErr)

			case nodes[0].Method != exp[i].Method:
				t.Error("Parsing failed; Method mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %v", exp[i].Method)
				t.Logf("  Actual:   %v", nodes[i].Method)

			case nodes[0].Name != exp[i].Name:
				t.Error("Parsing failed; Name mismatch:")
				t.Logf("  Test #:   %d", j)
				t.Logf("  Block:    %s", tests[j].Block)
				t.Logf("  Node #:   %d", i)
				t.Logf("  Expected: %s", exp[i].Name)
				t.Logf("  Actual:   %s", nodes[i].Name)

			default:
				jsonNodes, _ := json.Marshal(nodes[i].ParamTokens, false)
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

package jsonlines_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestUnmarshalTableMxCode(t *testing.T) {
	for i := range tableJsonl {
		test.RunMurexTests([]test.MurexTest{{
			Block: `
				tout jsonl (` + tableJsonl[i] + `) -> struct-keys -> format str -> msort
			`,
			ExitNum: 0,
			Stdout:  tableExpected,
			Stderr:  ``,
		}}, t)
	}
}

var tableJsonl = []string{
	// no whitespace
	//`["Name", "Session", "Score", "Completed"]["Gilbert", "2013", 24, true]["Alexa", "2013", 29, true]["May", "2012B", 14, false]["Deloise", "2012A", 19, true]`,
	// space no CRLF
	//`["Name", "Session", "Score", "Completed"] ["Gilbert", "2013", 24, true] ["Alexa", "2013", 29, true] ["May", "2012B", 14, false] ["Deloise", "2012A", 19, true]`,
	// CRLF
	`["Name", "Session", "Score", "Completed"]
["Gilbert", "2013", 24, true]
["Alexa", "2013", 29, true]
["May", "2012B", 14, false]
["Deloise", "2012A", 19, true]`,
	// double CRLF
	/*`["Name", "Session", "Score", "Completed"]

	["Gilbert", "2013", 24, true]

	["Alexa", "2013", 29, true]

	["May", "2012B", 14, false]

	["Deloise", "2012A", 19, true]`,
	// trailing CRLF
	`["Name", "Session", "Score", "Completed"]

	["Gilbert", "2013", 24, true]

	["Alexa", "2013", 29, true]

	["May", "2012B", 14, false]

	["Deloise", "2012A", 19, true]

	`,*/
}

var tableExpected = `/0
/0/0
/0/1
/0/2
/0/3
/1
/1/0
/1/1
/1/2
/1/3
/2
/2/0
/2/1
/2/2
/2/3
/3
/3/0
/3/1
/3/2
/3/3
/4
/4/0
/4/1
/4/2
/4/3
`

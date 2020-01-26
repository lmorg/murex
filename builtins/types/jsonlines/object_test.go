package jsonlines_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestUnmarshalObjectMxCode(t *testing.T) {
	for i := range objectJsonl {
		test.RunMurexTests([]test.MurexTest{{
			Block: `
				tout jsonl (` + objectJsonl[i] + `) -> struct-keys -> format str -> msort
			`,
			ExitNum: 0,
			Stdout:  objectExpected,
			Stderr:  ``,
		}}, t)
	}
}

var objectJsonl = []string{
	// no whitespace
	`{ "name": "Gilbert", "wins": [ [ "straight", "7♣" ], [ "one pair", "10♥" ] ] }{ "name": "Alexa", "wins": [ [ "two pair", "4♠" ], [ "two pair", "9♠" ] ] }`,
	// space no CRLF
	//`{ "name": "Gilbert", "wins": [ [ "straight", "7♣" ], [ "one pair", "10♥" ] ] } { "name": "Alexa", "wins": [ [ "two pair", "4♠" ], [ "two pair", "9♠" ] ] }`,
	// CRLF
	`{ "name": "Gilbert", "wins": [ [ "straight", "7♣" ], [ "one pair", "10♥" ] ] }
	{ "name": "Alexa", "wins": [ [ "two pair", "4♠" ], [ "two pair", "9♠" ] ] }`,
	// double CRLF
	`{ "name": "Gilbert", "wins": [ [ "straight", "7♣" ], [ "one pair", "10♥" ] ] }

	{ "name": "Alexa", "wins": [ [ "two pair", "4♠" ], [ "two pair", "9♠" ] ] }`,
	// no whitespace
	/*`{
	  "name": "Gilbert",
	  "wins": [
	    [
	      "straight",
	      "7♣"
	    ],
	    [
	      "one pair",
	      "10♥"
	    ]
	  ]
	}{
	  "name": "Alexa",
	  "wins": [
	    [
	      "two pair",
	      "4♠"
	    ],
	    [
	      "two pair",
	      "9♠"
	    ]
	  ]
	}`,*/
	// space no CRLF
	/*`{
	  "name": "Gilbert",
	  "wins": [
	    [
	      "straight",
	      "7♣"
	    ],
	    [
	      "one pair",
	      "10♥"
	    ]
	  ]
	}{
	  "name": "Alexa",
	  "wins": [
	    [
	      "two pair",
	      "4♠"
	    ],
	    [
	      "two pair",
	      "9♠"
	    ]
	  ]
	}`,*/
	// CRLF
	/*`{
	  "name": "Gilbert",
	  "wins": [
	    [
	      "straight",
	      "7♣"
	    ],
	    [
	      "one pair",
	      "10♥"
	    ]
	  ]
	}
	{
	  "name": "Alexa",
	  "wins": [
	    [
	      "two pair",
	      "4♠"
	    ],
	    [
	      "two pair",
	      "9♠"
	    ]
	  ]
	}`,*/
	// double CRLF
	/*`{
	  "name": "Gilbert",
	  "wins": [
	    [
	      "straight",
	      "7♣"
	    ],
	    [
	      "one pair",
	      "10♥"
	    ]
	  ]
	}

	{
	  "name": "Alexa",
	  "wins": [
	    [
	      "two pair",
	      "4♠"
	    ],
	    [
	      "two pair",
	      "9♠"
	    ]
	  ]
	}`,
	// trailing CRLF
	`{
	  "name": "Gilbert",
	  "wins": [
	    [
	      "straight",
	      "7♣"
	    ],
	    [
	      "one pair",
	      "10♥"
	    ]
	  ]
	}

	{
	  "name": "Alexa",
	  "wins": [
	    [
	      "two pair",
	      "4♠"
	    ],
	    [
	      "two pair",
	      "9♠"
	    ]
	  ]
	}

	`,*/
}

var objectExpected = `/0
/0/name
/0/wins
/0/wins/0
/0/wins/0/0
/0/wins/0/1
/0/wins/1
/0/wins/1/0
/0/wins/1/1
/1
/1/name
/1/wins
/1/wins/0
/1/wins/0/0
/1/wins/0/1
/1/wins/1
/1/wins/1/0
/1/wins/1/1
`

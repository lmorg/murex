package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/proc/parameters"
)

func TestParserColon(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example --flag`},
		{Expected: nodes, Block: `example       --flag`},
		{Expected: nodes, Block: `example:--flag`},
		{Expected: nodes, Block: `example: --flag`},
		{Expected: nodes, Block: `example:      --flag`},
	}

	testParser(t, tests)
}

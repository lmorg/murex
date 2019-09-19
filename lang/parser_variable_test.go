package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/proc/parameters"
)

func TestParserVariable1(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-", Type: parameters.TokenTypeValue},
		{Key: "var", Type: parameters.TokenTypeString},
		{Key: "-", Type: parameters.TokenTypeValue},
	}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example -$var-`},
		{Expected: nodes, Block: `example "-$var-"`},
	}

	testParser(t, tests)
}

func TestParserVariable2(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-$var-", Type: parameters.TokenTypeValue},
	}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example '-$var-'`},
		{Expected: nodes, Block: `example -\$var-`},
		{Expected: nodes, Block: `example "-\$var-"`},
	}

	testParser(t, tests)
}

func TestParserVariable3(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-\\$var-", Type: parameters.TokenTypeValue},
	}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example '-\$var-'`},
	}

	testParser(t, tests)
}

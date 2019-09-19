package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/proc/parameters"
)

func TestParserVariableBlockString1(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-", Type: parameters.TokenTypeValue},
		{Key: "block", Type: parameters.TokenTypeBlockString},
		{Key: "-", Type: parameters.TokenTypeValue},
	}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example -${block}-`},
		{Expected: nodes, Block: `example "-${block}-"`},
	}

	testParser(t, tests)
}

func TestParserVariableBlockString2(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-${block}-", Type: parameters.TokenTypeValue},
	}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example '-${block}-'`},
		{Expected: nodes, Block: `example -\${block}-`},
		{Expected: nodes, Block: `example "-\${block}-"`},
	}

	testParser(t, tests)
}

func TestParserVariableBlockString3(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-\\${block}-", Type: parameters.TokenTypeValue},
	}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example '-\${block}-'`},
	}

	testParser(t, tests)
}

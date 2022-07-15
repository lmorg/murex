package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/parameters"
)

func TestParserVariableBlockString1(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-", Type: parameters.TokenTypeValue},
		{Key: "block", Type: parameters.TokenTypeVarBlockString},
		{Key: "-", Type: parameters.TokenTypeValue},
	}}

	nodes := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example -${block}-`},
		{Expected: nodes, Block: `example "-${block}-"`},
		{Expected: nodes, Block: `example (-${block}-)`},
	}

	testParser(t, tests)
}

func TestParserVariableBlockString2(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-${block}-", Type: parameters.TokenTypeValue},
	}}

	nodes := AstNodes{{
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

	nodes := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example '-\${block}-'`},
	}

	testParser(t, tests)
}

func TestParserVariableBlockString4(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-", Type: parameters.TokenTypeValue},
		{Key: "{block}", Type: parameters.TokenTypeVarBlockString},
		{Key: "-", Type: parameters.TokenTypeValue},
	}}

	nodes := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example -${{block}}-`},
		{Expected: nodes, Block: `example "-${{block}}-"`},
		{Expected: nodes, Block: `example (-${{block}}-)`},
	}

	testParser(t, tests)
}

// This currently fails. Need to rewrite parser :(
/*func TestParserVariableBlockString5(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "-", Type: parameters.TokenTypeValue},
		{Key: "{block}", Type: parameters.TokenTypeBlockString},
		{Key: "-", Type: parameters.TokenTypeValue},
	}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example -${block\}}-`},
		{Expected: nodes, Block: `example -${block'}'}-`},
		{Expected: nodes, Block: `example -${block"}"}-`},
		{Expected: nodes, Block: `example -${block(})}-`},
		{Expected: nodes, Block: `example "-${block\}}-"`},
		//{Expected: nodes, Block: `example (-${block\}}-)`},
	}

	testParser(t, tests)
}*/

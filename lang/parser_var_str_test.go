package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/proc/parameters"
)

func TestParserVariableString1(t *testing.T) {
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

func TestParserVariableString2(t *testing.T) {
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

func TestParserVariableString3(t *testing.T) {
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

/*func TestParserVariableString4(t *testing.T) {
	params := [][]parameters.ParamToken{
		{{Key: "-", Type: parameters.TokenTypeValue}},
		{{Key: "var", Type: parameters.TokenTypeString}},
		{{Key: "-", Type: parameters.TokenTypeValue}},
	}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example - $var -`},
		{Expected: nodes, Block: `example -  $var  -`},
		{Expected: nodes, Block: `example - "$var" -`},
		{Expected: nodes, Block: `example -  "$var"  -`},
	}

	testParser(t, tests)
}*/

// fix bug with parser hanging
func TestParserParenthesisHung(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `out test $[foobar]`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       `out`,
					Parameters: []string{`test`, `[foobar]`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out test \$[foobar]`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       `out`,
					Parameters: []string{`test`, `$[foobar]`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out test @[foobar]`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       `out`,
					Parameters: []string{`test`, `@[foobar]`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out test \@[foobar]`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       `out`,
					Parameters: []string{`test`, `@[foobar]`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

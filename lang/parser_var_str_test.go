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
func TestParserParenthesisHungBug(t *testing.T) {
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

/*func TestParserParenthesisBug(t *testing.T) {
	block := `
    private autocomplete.systemctl {
        systemctl: list-unit-files -> !regexp m/unit files listed/ -> [:0] -> cast str
    }

    function autocomplete.systemctl.flags {
        systemctl: --help -> @[Unit Commands:..]s -> regexp m/(NAME|PATTERN)/ -> tabulate: --map --key-inc-hint -> formap key val {
            out ("$key": [{
                "Dynamic": ({ autocomplete.systemctl }),
                "ListView": true,
                "Optional": false,
                "AllowMultiple": true
            }],)
        }
        out ("": [{}]) # dummy value so there's no trailing comma
    }

    autocomplete set systemctl ({[
        {
            "DynamicDesc": ({
                systemctl: --help -> @[..Unit Commands:]s -> tabulate: --column-wraps --map --key-inc-hint --split-space
            }),
            "Optional": true,
            "AllowMultiple": false
        },
        {
            "DynamicDesc": ({
                systemctl: --help -> @[Unit Commands:..]s -> tabulate: --column-wraps --map --key-inc-hint
            }),
            "Optional": false,
            "AllowMultiple": false,
            "FlagValues": {
                ${ autocomplete.systemctl.flags }
            }
        }
    ]})`

	t.Error(queryParser(t, block))
}*/

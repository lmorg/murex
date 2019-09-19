package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/proc/parameters"
)

func TestParserNamedPiped1(t *testing.T) {
	pipeParams := [][]parameters.ParamToken{{{
		Key:  "<pipe>",
		Type: parameters.TokenTypeNamedPipe,
	}}}

	valParams := [][]parameters.ParamToken{{{
		Key:  "<notapipe>",
		Type: parameters.TokenTypeValue,
	}}}

	pipeNodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: pipeParams,
	}}

	valNodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: valParams,
	}}

	var tests = []parserTestConditions{
		{Expected: pipeNodes, Block: `example <pipe>`},

		{Expected: valNodes, Block: `example \<notapipe>`},
		{Expected: valNodes, Block: `example '<notapipe>'`},
		{Expected: valNodes, Block: `example "<notapipe>"`},
		{Expected: valNodes, Block: `example (<notapipe>)`},
	}

	testParser(t, tests)
}

func TestParserNamedPiped2(t *testing.T) {
	pipeParams := [][]parameters.ParamToken{
		{{Key: "<pipe>", Type: parameters.TokenTypeNamedPipe}},
		{{Key: "<pipe>", Type: parameters.TokenTypeNamedPipe}},
		{{Key: "<pipe>", Type: parameters.TokenTypeNamedPipe}},
		{{Key: "<pipe>", Type: parameters.TokenTypeNamedPipe}},
	}

	valParams := [][]parameters.ParamToken{
		{{Key: "<notapipe>", Type: parameters.TokenTypeValue}},
		{{Key: "<notapipe>", Type: parameters.TokenTypeValue}},
		{{Key: "<notapipe>", Type: parameters.TokenTypeValue}},
		{{Key: "<notapipe>", Type: parameters.TokenTypeValue}},
	}

	pipeNodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: pipeParams,
	}}

	valNodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: valParams,
	}}

	var tests = []parserTestConditions{
		{Expected: pipeNodes, Block: `example <pipe> <pipe> <pipe> <pipe>`},

		{Expected: valNodes, Block: `example \<notapipe> \<notapipe> \<notapipe> \<notapipe>`},
		{Expected: valNodes, Block: `example '<notapipe>' '<notapipe>' '<notapipe>' '<notapipe>'`},
		{Expected: valNodes, Block: `example "<notapipe>" "<notapipe>" "<notapipe>" "<notapipe>"`},
		{Expected: valNodes, Block: `example (<notapipe>) (<notapipe>) (<notapipe>) (<notapipe>)`},
	}

	testParser(t, tests)
}

func TestParserNamedPiped3(t *testing.T) {
	params := [][]parameters.ParamToken{
		{{Key: "<pipe>", Type: parameters.TokenTypeNamedPipe}},
		{{Key: "<notapipe>", Type: parameters.TokenTypeValue}},
		{{Key: "<pipe>", Type: parameters.TokenTypeNamedPipe}},
		{{Key: "<notapipe>", Type: parameters.TokenTypeValue}},
	}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example <pipe> \<notapipe> <pipe> \<notapipe>`},
		{Expected: nodes, Block: `example <pipe> '<notapipe>' <pipe> '<notapipe>'`},
		{Expected: nodes, Block: `example <pipe> "<notapipe>" <pipe> "<notapipe>"`},
		{Expected: nodes, Block: `example <pipe> (<notapipe>) <pipe> (<notapipe>)`},
	}

	testParser(t, tests)
}

func TestParserNamedPiped4(t *testing.T) {
	params := [][]parameters.ParamToken{
		{{Key: "<", Type: parameters.TokenTypeNamedPipe}},
		{{Key: "notapipe", Type: parameters.TokenTypeValue}},
		{{Key: ">", Type: parameters.TokenTypeValue}},
		{{Key: "<badpipe", Type: parameters.TokenTypeNamedPipe}},
	}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example < notapipe > <badpipe`},
	}

	testParser(t, tests)
}

func TestParserNamedPiped5(t *testing.T) {
	params := [][]parameters.ParamToken{{
		{Key: "<", Type: parameters.TokenTypeNamedPipe},
		{Key: "notapipe", Type: parameters.TokenTypeString},
		{Key: ">", Type: parameters.TokenTypeValue},
	}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example <$notapipe>`},
	}

	testParser(t, tests)
}

func TestParserNamedPiped6(t *testing.T) {
	params0 := [][]parameters.ParamToken{
		{{Key: "<$notapipe>", Type: parameters.TokenTypeNamedPipe}},
	}

	nodes0 := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params0,
	}}

	params1 := [][]parameters.ParamToken{
		{{Key: "<$notapipe>", Type: parameters.TokenTypeValue}},
	}

	nodes1 := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params1,
	}}

	params2 := [][]parameters.ParamToken{{
		{Key: "<", Type: parameters.TokenTypeValue},
		{Key: "notapipe", Type: parameters.TokenTypeString},
		{Key: ">", Type: parameters.TokenTypeValue},
	}}

	nodes2 := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params2,
	}}

	params3 := [][]parameters.ParamToken{{
		{Key: "<", Type: parameters.TokenTypeNamedPipe},
		{Key: "notapipe", Type: parameters.TokenTypeString},
		{Key: ">", Type: parameters.TokenTypeValue},
	}}

	nodes3 := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params3,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes0, Block: `example <\$notapipe>`},
		{Expected: nodes1, Block: `example <'$notapipe'>`},
		{Expected: nodes2, Block: `example <"$notapipe">`},
		{Expected: nodes3, Block: `example <($notapipe)>`},
	}

	testParser(t, tests)
}

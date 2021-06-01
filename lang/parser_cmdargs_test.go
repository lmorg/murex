package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/parameters"
)

// Some of the tests in this file might appear duplications upon first glance.
// They are not. They are testing whitespace, quotations, colons and every
// variation of the aforementioned - ensuring that the parser can correctly
// identify the command and parameter fields in every edge case.
func TestParserColon(t *testing.T) {
	null := []parameters.ParamToken{{
		Key:  "",
		Type: parameters.TokenTypeNil,
	}}

	param := []parameters.ParamToken{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}

	nodes1 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{param},
	}}

	nodes2 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{null, param},
	}}

	var tests = []parserTestConditions{
		{Expected: nodes1, Block: `example --flag`},
		{Expected: nodes1, Block: `example       --flag`},
		{Expected: nodes1, Block: `example:--flag`},
		{Expected: nodes2, Block: `example: --flag`},
		{Expected: nodes2, Block: `example:      --flag`},
	}

	testParser(t, tests)
}

func TestParserSpace1(t *testing.T) {
	null := []parameters.ParamToken{{
		Key:  "",
		Type: parameters.TokenTypeNil,
	}}
	param1 := []parameters.ParamToken{{
		Key:  "--flag1",
		Type: parameters.TokenTypeValue,
	}}
	param2 := []parameters.ParamToken{{
		Key:  "--flag2",
		Type: parameters.TokenTypeValue,
	}}

	nodes1 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{param1, param2},
	}}
	nodes2 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{null, param1, param2},
	}}

	var tests = []parserTestConditions{
		{Expected: nodes1, Block: `example --flag1 --flag2`},
		{Expected: nodes1, Block: `example       --flag1 --flag2`},
		{Expected: nodes1, Block: `example:--flag1 --flag2`},
		{Expected: nodes2, Block: `example: --flag1 --flag2`},
		{Expected: nodes2, Block: `example:      --flag1 --flag2`},

		{Expected: nodes1, Block: `example --flag1  --flag2`},
		{Expected: nodes1, Block: `example       --flag1  --flag2`},
		{Expected: nodes1, Block: `example:--flag1  --flag2`},
		{Expected: nodes2, Block: `example: --flag1  --flag2`},
		{Expected: nodes2, Block: `example:      --flag1  --flag2`},

		{Expected: nodes1, Block: `example --flag1    --flag2`},
		{Expected: nodes1, Block: `example       --flag1    --flag2`},
		{Expected: nodes1, Block: `example:--flag1    --flag2`},
		{Expected: nodes2, Block: `example: --flag1    --flag2`},
		{Expected: nodes2, Block: `example:      --flag1    --flag2`},

		{Expected: nodes1, Block: `example --flag1	--flag2`},
		{Expected: nodes1, Block: `example       --flag1	--flag2`},
		{Expected: nodes1, Block: `example:--flag1	--flag2`},
		{Expected: nodes2, Block: `example: --flag1	--flag2`},
		{Expected: nodes2, Block: `example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserSpace2(t *testing.T) {
	null := []parameters.ParamToken{{
		Key:  "",
		Type: parameters.TokenTypeNil,
	}}
	param1 := []parameters.ParamToken{{
		Key:  "--flag1",
		Type: parameters.TokenTypeValue,
	}}
	param2 := []parameters.ParamToken{{
		Key:  "--flag2",
		Type: parameters.TokenTypeValue,
	}}

	nodes1 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{param1, param2},
	}}
	nodes2 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{null, param1, param2},
	}}

	var tests = []parserTestConditions{
		{Expected: nodes1, Block: ` example --flag1 --flag2`},
		{Expected: nodes1, Block: ` example       --flag1 --flag2`},
		{Expected: nodes1, Block: ` example:--flag1 --flag2`},
		{Expected: nodes2, Block: ` example: --flag1 --flag2`},
		{Expected: nodes2, Block: ` example:      --flag1 --flag2`},

		{Expected: nodes1, Block: ` example --flag1  --flag2`},
		{Expected: nodes1, Block: ` example       --flag1  --flag2`},
		{Expected: nodes1, Block: ` example:--flag1  --flag2`},
		{Expected: nodes2, Block: ` example: --flag1  --flag2`},
		{Expected: nodes2, Block: ` example:      --flag1  --flag2`},

		{Expected: nodes1, Block: ` example --flag1    --flag2`},
		{Expected: nodes1, Block: ` example       --flag1    --flag2`},
		{Expected: nodes1, Block: ` example:--flag1    --flag2`},
		{Expected: nodes2, Block: ` example: --flag1    --flag2`},
		{Expected: nodes2, Block: ` example:      --flag1    --flag2`},

		{Expected: nodes1, Block: ` example --flag1	--flag2`},
		{Expected: nodes1, Block: ` example       --flag1	--flag2`},
		{Expected: nodes1, Block: ` example:--flag1	--flag2`},
		{Expected: nodes2, Block: ` example: --flag1	--flag2`},
		{Expected: nodes2, Block: ` example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserSpace3(t *testing.T) {
	null := []parameters.ParamToken{{
		Key:  "",
		Type: parameters.TokenTypeNil,
	}}
	param1 := []parameters.ParamToken{{
		Key:  "--flag1",
		Type: parameters.TokenTypeValue,
	}}
	param2 := []parameters.ParamToken{{
		Key:  "--flag2",
		Type: parameters.TokenTypeValue,
	}}

	nodes1 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{param1, param2},
	}}
	nodes2 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{null, param1, param2},
	}}

	var tests = []parserTestConditions{
		{Expected: nodes1, Block: `  example --flag1 --flag2`},
		{Expected: nodes1, Block: `  example       --flag1 --flag2`},
		{Expected: nodes1, Block: `  example:--flag1 --flag2`},
		{Expected: nodes2, Block: `  example: --flag1 --flag2`},
		{Expected: nodes2, Block: `  example:      --flag1 --flag2`},

		{Expected: nodes1, Block: `  example --flag1  --flag2`},
		{Expected: nodes1, Block: `  example       --flag1  --flag2`},
		{Expected: nodes1, Block: `  example:--flag1  --flag2`},
		{Expected: nodes2, Block: `  example: --flag1  --flag2`},
		{Expected: nodes2, Block: `  example:      --flag1  --flag2`},

		{Expected: nodes1, Block: `  example --flag1    --flag2`},
		{Expected: nodes1, Block: `  example       --flag1    --flag2`},
		{Expected: nodes1, Block: `  example:--flag1    --flag2`},
		{Expected: nodes2, Block: `  example: --flag1    --flag2`},
		{Expected: nodes2, Block: `  example:      --flag1    --flag2`},

		{Expected: nodes1, Block: `  example --flag1	--flag2`},
		{Expected: nodes1, Block: `  example       --flag1	--flag2`},
		{Expected: nodes1, Block: `  example:--flag1	--flag2`},
		{Expected: nodes2, Block: `  example: --flag1	--flag2`},
		{Expected: nodes2, Block: `  example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserSpace4(t *testing.T) {
	null := []parameters.ParamToken{{
		Key:  "",
		Type: parameters.TokenTypeNil,
	}}
	param1 := []parameters.ParamToken{{
		Key:  "--flag1",
		Type: parameters.TokenTypeValue,
	}}
	param2 := []parameters.ParamToken{{
		Key:  "--flag2",
		Type: parameters.TokenTypeValue,
	}}

	nodes1 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{param1, param2},
	}}
	nodes2 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{null, param1, param2},
	}}

	var tests = []parserTestConditions{
		{Expected: nodes1, Block: `	example --flag1 --flag2`},
		{Expected: nodes1, Block: `	example       --flag1 --flag2`},
		{Expected: nodes1, Block: `	example:--flag1 --flag2`},
		{Expected: nodes2, Block: `	example: --flag1 --flag2`},
		{Expected: nodes2, Block: `	example:      --flag1 --flag2`},

		{Expected: nodes1, Block: `	example --flag1  --flag2`},
		{Expected: nodes1, Block: `	example       --flag1  --flag2`},
		{Expected: nodes1, Block: `	example:--flag1  --flag2`},
		{Expected: nodes2, Block: `	example: --flag1  --flag2`},
		{Expected: nodes2, Block: `	example:      --flag1  --flag2`},

		{Expected: nodes1, Block: `	example --flag1    --flag2`},
		{Expected: nodes1, Block: `	example       --flag1    --flag2`},
		{Expected: nodes1, Block: `	example:--flag1    --flag2`},
		{Expected: nodes2, Block: `	example: --flag1    --flag2`},
		{Expected: nodes2, Block: `	example:      --flag1    --flag2`},

		{Expected: nodes1, Block: `	example --flag1	--flag2`},
		{Expected: nodes1, Block: `	example       --flag1	--flag2`},
		{Expected: nodes1, Block: `	example:--flag1	--flag2`},
		{Expected: nodes2, Block: `	example: --flag1	--flag2`},
		{Expected: nodes2, Block: `	example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserSpace5(t *testing.T) {
	null := []parameters.ParamToken{{
		Key:  "",
		Type: parameters.TokenTypeNil,
	}}
	param1 := []parameters.ParamToken{{
		Key:  "--flag1",
		Type: parameters.TokenTypeValue,
	}}
	param2 := []parameters.ParamToken{{
		Key:  "--flag2",
		Type: parameters.TokenTypeValue,
	}}

	nodes1 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{param1, param2},
	}}
	nodes2 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{null, param1, param2},
	}}

	var tests = []parserTestConditions{
		{Expected: nodes1, Block: `		example --flag1 --flag2`},
		{Expected: nodes1, Block: `		example       --flag1 --flag2`},
		{Expected: nodes1, Block: `		example:--flag1 --flag2`},
		{Expected: nodes2, Block: `		example: --flag1 --flag2`},
		{Expected: nodes2, Block: `		example:      --flag1 --flag2`},

		{Expected: nodes1, Block: `		example --flag1  --flag2`},
		{Expected: nodes1, Block: `		example       --flag1  --flag2`},
		{Expected: nodes1, Block: `		example:--flag1  --flag2`},
		{Expected: nodes2, Block: `		example: --flag1  --flag2`},
		{Expected: nodes2, Block: `		example:      --flag1  --flag2`},

		{Expected: nodes1, Block: `		example --flag1    --flag2`},
		{Expected: nodes1, Block: `		example       --flag1    --flag2`},
		{Expected: nodes1, Block: `		example:--flag1    --flag2`},
		{Expected: nodes2, Block: `		example: --flag1    --flag2`},
		{Expected: nodes2, Block: `		example:      --flag1    --flag2`},

		{Expected: nodes1, Block: `		example --flag1	--flag2`},
		{Expected: nodes1, Block: `		example       --flag1	--flag2`},
		{Expected: nodes1, Block: `		example:--flag1	--flag2`},
		{Expected: nodes2, Block: `		example: --flag1	--flag2`},
		{Expected: nodes2, Block: `		example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserCmdQuotes1(t *testing.T) {
	null := []parameters.ParamToken{{
		Key:  "",
		Type: parameters.TokenTypeNil,
	}}

	param := []parameters.ParamToken{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}

	nodes1 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{param},
	}}

	nodes2 := AstNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: [][]parameters.ParamToken{null, param},
	}}

	var tests = []parserTestConditions{
		{Expected: nodes1, Block: `'example' --flag`},
		{Expected: nodes1, Block: `'example'       --flag`},
		{Expected: nodes1, Block: `'example':--flag`},
		{Expected: nodes2, Block: `'example': --flag`},
		{Expected: nodes2, Block: `'example':      --flag`},

		{Expected: nodes1, Block: `"example" --flag`},
		{Expected: nodes1, Block: `"example"       --flag`},
		{Expected: nodes1, Block: `"example":--flag`},
		{Expected: nodes2, Block: `"example": --flag`},
		{Expected: nodes2, Block: `"example":      --flag`},
	}

	testParser(t, tests)
}

func TestParserCmdQuotes2(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := AstNodes{{
		NewChain:    true,
		Name:        "example ",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `'example ' --flag`},
		{Expected: nodes, Block: `'example '       --flag`},

		{Expected: nodes, Block: `"example " --flag`},
		{Expected: nodes, Block: `"example "       --flag`},
	}

	testParser(t, tests)
}

func TestParserCmdQuotes3(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := AstNodes{{
		NewChain:    true,
		Name:        "example:",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `'example:' --flag`},
		{Expected: nodes, Block: `'example:'      --flag`},

		{Expected: nodes, Block: `"example:" --flag`},
		{Expected: nodes, Block: `"example:"       --flag`},
	}

	testParser(t, tests)
}

func TestParserCmdQuotes4(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag2",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := AstNodes{{
		NewChain:    true,
		Name:        "example:--flag1",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `'example:'--flag1 --flag2`},
		{Expected: nodes, Block: `"example:"--flag1 --flag2`},
	}

	testParser(t, tests)
}

func TestParserUnexpectedColons(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `func p1 p2 p3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1", "p2", "p3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func: p1 p2 p3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1", "p2", "p3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		// func no colon
		{
			Block: `func p1: p2 p3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1:", "p2", "p3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func p1 p2: p3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1", "p2:", "p3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func p1 p2 p3:`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1", "p2", "p3:"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func p1: p2: p3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1:", "p2:", "p3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func p1: p2: p3:`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1:", "p2:", "p3:"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		// func colon
		{
			Block: `func: p1: p2 p3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1:", "p2", "p3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func: p1 p2: p3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1", "p2:", "p3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func: p1 p2 p3:`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1", "p2", "p3:"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func: p1: p2: p3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1:", "p2:", "p3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `func: p1: p2: p3:`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "func",
					Parameters: []string{"p1:", "p2:", "p3:"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

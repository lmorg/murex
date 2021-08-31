package lang

import (
	"testing"
)

// Some of these tests are commented out. This is because those characters
// cannot be used with the wrapper test functions and will instead require
// dedicated test plans written specifically for them

func TestParserFuncEscOrQuote(t *testing.T) {
	tests := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m" /*"n",*/, "o", "p", "q" /*"r",*/ /*"s",*/ /*"t",*/, "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"!" /*`"`,*/, "£", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", ";", ":" /*"'",*/, "@", "#", "~", "|", ",", "<", ".", ">", "/", "?",
		" ", "\t", "\r", /*"\n",*/
		"世", "界",
		"😀", "😍", "🙏",
	}

	for i := range tests {
		testParserFuncWrapper(t, tests[i])
	}
}

func testParserFuncWrapper(t *testing.T, condition string) {
	tests := []parserTestSimpleConditions{

		// function escape

		{
			Block: `\` + condition,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` \` + condition,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `\` + condition + ` `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` \` + condition + ` `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `\` + condition + ` foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{"foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` \` + condition + ` foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{"foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		// function single quote

		{
			Block: `'` + condition + `'`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` '` + condition + `'`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `'` + condition + `' `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` '` + condition + `' `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `'` + condition + `' foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{"foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` '` + condition + `' foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{"foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		// function double quote

		{
			Block: `"` + condition + `"`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` "` + condition + `"`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `"` + condition + `" `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` "` + condition + `" `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `"` + condition + `" foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{"foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` "` + condition + `" foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       condition,
					Parameters: []string{"foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

func TestParserParamEscOrQuote(t *testing.T) {
	tests := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m" /*"n",*/, "o", "p", "q" /*"r",*/ /*"s",*/ /*"t",*/, "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"!" /*`"`,*/, "£" /*"$",*/, "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", ";", ":" /*"'",*/ /*"@",*/, "#" /*"~",*/, "|", ",", "<", ".", ">", "/", "?",
		//" ", "\t", "\r", /*"\n",*/
		"世", "界",
		"😀", "😍", "🙏",
	}

	for i := range tests {
		testParserParamWrapper(t, tests[i])
	}
}

func testParserParamWrapper(t *testing.T, condition string) {
	tests := []parserTestSimpleConditions{

		// parameter escape

		{
			Block: `foobar \` + condition,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar  \` + condition,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foobar \` + condition + ` `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar  \` + condition + ` `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foobar \` + condition + ` foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition, "foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar \` + condition + ` foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition, "foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		// parameter single quote

		{
			Block: `foobar '` + condition + `'`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar '` + condition + `'`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foobar '` + condition + `' `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar '` + condition + `' `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foobar '` + condition + `' foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition, "foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar '` + condition + `' foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition, "foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		// parameter double quote

		{
			Block: `foobar "` + condition + `"`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar "` + condition + `"`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foobar "` + condition + `" `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar "` + condition + `" `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foobar "` + condition + `" foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition, "foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` foobar "` + condition + `" foobar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foobar",
					Parameters: []string{condition, "foobar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

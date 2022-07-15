package lang

import (
	"testing"
)

func TestParserEscapes(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `out \s`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{" "},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out \t`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"\t"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out \r`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"\r"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out \n`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"\n"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out \\`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{`\`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		{
			Block: `out "\"`,
			Error: true,
		},
		{
			Block: `out "\\"`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{`\`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out '\'`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{`\`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out (\)`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{`\`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out (\))`,
			Error: true,
		},
		{
			Block: `out (\)\)`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{`\)`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out {\}`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{`{\}`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `out {\}}`,
			Error: true,
		},
		{
			Block: `out {\}\}`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{`{\}}`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

func TestParserBackTicks(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: "out `",
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"'"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		{
			Block: "out \\`",
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"`"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: "out '`'",
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"`"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: "out \"`\"",
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"`"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: "out (`)",
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"`"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: "out {`}",
			Expected: []parserTestSimpleExpected{
				{
					Name:       "out",
					Parameters: []string{"{`}"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

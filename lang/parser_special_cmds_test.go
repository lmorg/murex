package lang

import (
	"testing"
)

func TestParserEqu(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `= a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` =a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` = a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `=a +b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"a", "+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `=a + b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"a", "+", "b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		// negative tests

		{
			Block: `a=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a=a+b",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` a=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a=a+b",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `a= a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a=",
					Parameters: []string{"a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` a= a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a=",
					Parameters: []string{"a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `a =a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a",
					Parameters: []string{"=a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `a = a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a",
					Parameters: []string{"=", "a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `a =a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a",
					Parameters: []string{"=a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` a = a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a",
					Parameters: []string{"=", "a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

func TestParserAutoGlob(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `@g =a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "@g",
					Parameters: []string{"=a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` @g =a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "@g",
					Parameters: []string{"=a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `@g=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "@g=a+b",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` @g=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "@g=a+b",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		// negative testing

		{
			Block: `a \@g =a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a",
					Parameters: []string{"@g", "=a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` a \@g =a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a",
					Parameters: []string{"@g", "=a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `a \@g=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a",
					Parameters: []string{"@g=a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` a \@g=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a",
					Parameters: []string{"@g=a+b"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `a@g=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a@g=a+b",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: ` a@g=a+b`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "a@g=a+b",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

package lang

import (
	"testing"
)

func TestParserParenthesis(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `foo (bar)`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       `foo`,
					Parameters: []string{`bar`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foo b(ar)`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       `foo`,
					Parameters: []string{`b(ar)`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foo (ba)r`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       `foo`,
					Parameters: []string{`bar`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foo b(a)r`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       `foo`,
					Parameters: []string{`b(a)r`},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

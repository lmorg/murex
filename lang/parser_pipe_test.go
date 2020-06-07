package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestParserPipeIn(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `|foo`,
			Error: true,
		},
		{
			Block: `| foo`,
			Error: true,
		},
		{
			Block: `  |  foo`,
			Error: true,
		},

		{
			Block: `->foo`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},
		{
			Block: `-> foo`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},
		{
			Block: `  ->  foo`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		/*{
			Block: `=>foo`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},
		{
			Block: `=> foo`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},
		{
			Block: `  =>  foo`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},*/
	}
	testParserSimple(t, tests)
}

func TestParserPipeOut(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `foo|bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `foo->bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `foo=>bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "format",
					Parameters: []string{types.Generic},
					Method:     TEST_METHOD | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `foo | bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `foo -> bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `foo => bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "format",
					Parameters: []string{types.Generic},
					Method:     TEST_METHOD | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `  foo  |  bar  `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `  foo  ->  bar  `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `  foo  =>  bar  `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "format",
					Parameters: []string{types.Generic},
					Method:     TEST_METHOD | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		// with more parameters

		{
			Block: `foo 1 2 3 | bar 1 2 3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `foo 1 2 3 -> bar 1 2 3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `foo 1 2 3 => bar 1 2 3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_NEW_PIPE | TEST_PIPE_OUT,
				},
				{
					Name:       "format",
					Parameters: []string{types.Generic},
					Method:     TEST_METHOD | TEST_PIPE_OUT,
				},
				{
					Name:       "bar",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_METHOD,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

func TestParserPipeErr(t *testing.T) {
	tests := []parserTestSimpleConditions{
		{
			Block: `foo?bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo?bar",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		{
			Block: `foo ?bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_ERR,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `foo ? bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_ERR,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `  foo  ?  bar  `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{},
					Method:     TEST_NEW_PIPE | TEST_PIPE_ERR,
				},
				{
					Name:       "bar",
					Parameters: []string{},
					Method:     TEST_METHOD,
				},
			},
		},

		// with more parameters

		{
			Block: `foo 1 2 3?bar 1 2 3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"1", "2", "3?bar", "1", "2", "3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		{
			Block: `foo 1 2 3? bar 1 2 3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"1", "2", "3?", "bar", "1", "2", "3"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		{
			Block: `foo 1 2 3 ? bar 1 2 3`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_NEW_PIPE | TEST_PIPE_ERR,
				},
				{
					Name:       "bar",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_METHOD,
				},
			},
		},

		{
			Block: `  foo 1 2 3   ?   bar 1 2 3  `,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_NEW_PIPE | TEST_PIPE_ERR,
				},
				{
					Name:       "bar",
					Parameters: []string{"1", "2", "3"},
					Method:     TEST_METHOD,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

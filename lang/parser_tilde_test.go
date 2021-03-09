package lang

import (
	"testing"

	"github.com/lmorg/murex/utils/home"
)

func TestParserTilde(t *testing.T) {
	tests := []parserTestSimpleConditions{

		/// cmd == foo
		{
			Block: `foo =~ bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"=" + home.MyDir, "bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foo:=~ bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"=" + home.MyDir, "bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foo: =~ bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"=" + home.MyDir, "bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		///
		{
			Block: `foo = ~ bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"=", home.MyDir, "bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foo:= ~ bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"=", home.MyDir, "bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `foo: = ~ bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "foo",
					Parameters: []string{"=", home.MyDir, "bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},

		/// cmd == eval
		{
			Block: `= foo=~bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"foo=~bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `= foo =~ bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"foo", "=~", "bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		///
		{
			Block: `= foo= ~`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"foo=", home.MyDir},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
		{
			Block: `= foo = ~ bar`,
			Expected: []parserTestSimpleExpected{
				{
					Name:       "=",
					Parameters: []string{"foo", "=", home.MyDir, "bar"},
					Method:     TEST_NEW_PIPE,
				},
			},
		},
	}

	testParserSimple(t, tests)
}

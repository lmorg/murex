package expressions

import (
	"testing"
)

func TestParseCommentMultiLineStatement(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `echo /# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo /# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
		///
		{
			Statement: `echo/# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo/# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
		///
		{
			Statement: `echo /# this is an infixed comment #/foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo /# this is an infixed comment #/foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
	}

	testParseStatement(t, tests)
}

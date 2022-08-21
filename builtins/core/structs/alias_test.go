package structs_test

import (
	"fmt"
	"math/rand"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestAliasParamParsing(t *testing.T) {
	alias := fmt.Sprintf("GoTest-alias-%d-", rand.Int())

	tests := []test.MurexTest{
		// errors

		{
			Block:  fmt.Sprintf(`alias %s%d`, alias, -1),
			Stderr: "no command supplied",
			ExitNum: 1,
		},
		{
			Block:  fmt.Sprintf(`alias %s%d `, alias, -2),
			Stderr: "no command supplied",
			ExitNum: 1,
		},
		{
			Block:  fmt.Sprintf(`alias %s%d=`, alias, -3),
			Stderr: "no command supplied",
			ExitNum: 1,
		},
		{
			Block:  fmt.Sprintf(`alias %s%d= `, alias, -4),
			Stderr: "no command supplied",
			ExitNum: 1,
		},
		{
			Block:  fmt.Sprintf(`alias %s%d =`, alias, -1),
			Stderr: "no command supplied",
			ExitNum: 1,
		},
		{
			Block:  fmt.Sprintf(`alias %s%d foobar`, alias, -1),
			Stderr: "invalid syntax",
			ExitNum: 1,
		},

		// no space

		{
			Block:  fmt.Sprintf(`alias %s%d=foo bar; alias -> [%s%d]`, alias, 0, alias, 0),
			Stdout: "[\"foo\",\"bar\"]",
		},

		// 1 space

		{
			Block:  fmt.Sprintf(`alias %s%d= foo bar; alias -> [%s%d]`, alias, 1, alias, 1),
			Stdout: "[\"foo\",\"bar\"]",
		},
		{
			Block:  fmt.Sprintf(`alias %s%d =foo bar; alias -> [%s%d]`, alias, 2, alias, 2),
			Stdout: "[\"foo\",\"bar\"]",
		},
		{
			Block:  fmt.Sprintf(`alias %s%d = foo bar; alias -> [%s%d]`, alias, 3, alias, 3),
			Stdout: "[\"foo\",\"bar\"]",
		},

		// 2 spaces

		{
			Block:  fmt.Sprintf(`alias %s%d=  foo bar; alias -> [%s%d]`, alias, 4, alias, 4),
			Stdout: "[\"foo\",\"bar\"]",
		},
		{
			Block:  fmt.Sprintf(`alias %s%d  =foo bar; alias -> [%s%d]`, alias, 5, alias, 5),
			Stdout: "[\"foo\",\"bar\"]",
		},
		{
			Block:  fmt.Sprintf(`alias %s%d  =  foo bar; alias -> [%s%d]`, alias, 6, alias, 6),
			Stdout: "[\"foo\",\"bar\"]",
		},

		// 1 tab

		{
			Block:  fmt.Sprintf(`alias %s%d=	foo bar; alias -> [%s%d]`, alias, 7, alias, 7),
			Stdout: "[\"foo\",\"bar\"]",
		},
		{
			Block:  fmt.Sprintf(`alias %s%d	=foo bar; alias -> [%s%d]`, alias, 8, alias, 8),
			Stdout: "[\"foo\",\"bar\"]",
		},
		{
			Block:  fmt.Sprintf(`alias %s%d	=	foo bar; alias -> [%s%d]`, alias, 9, alias, 9),
			Stdout: "[\"foo\",\"bar\"]",
		},

		// 2 tabs

		{
			Block:  fmt.Sprintf(`alias %s%d=		foo bar; alias -> [%s%d]`, alias, 10, alias, 10),
			Stdout: "[\"foo\",\"bar\"]",
		},
		{
			Block:  fmt.Sprintf(`alias %s%d		=foo bar; alias -> [%s%d]`, alias, 11, alias, 11),
			Stdout: "[\"foo\",\"bar\"]",
		},
		{
			Block:  fmt.Sprintf(`alias %s%d		=		foo bar; alias -> [%s%d]`, alias, 12, alias, 12),
			Stdout: "[\"foo\",\"bar\"]",
		},


	}

	test.RunMurexTestsRx(tests, t)
}

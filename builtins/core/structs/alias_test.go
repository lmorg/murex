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
	const errMissingCommand = "missing command to alias"

	tests := []test.MurexTest{
		// errors

		{
			Block:   fmt.Sprintf(`alias %s%d`, alias, -1),
			Stderr:  errMissingCommand,
			ExitNum: 1,
		},
		{
			Block:   fmt.Sprintf(`alias %s%d `, alias, -2),
			Stderr:  errMissingCommand,
			ExitNum: 1,
		},
		{
			Block:   fmt.Sprintf(`alias %s%d=`, alias, -3),
			Stderr:  errMissingCommand,
			ExitNum: 1,
		},
		{
			Block:   fmt.Sprintf(`alias %s%d= `, alias, -4),
			Stderr:  errMissingCommand,
			ExitNum: 1,
		},
		{
			Block:   fmt.Sprintf(`alias %s%d =`, alias, -5),
			Stderr:  errMissingCommand,
			ExitNum: 1,
		},
		{
			Block:   fmt.Sprintf(`alias %s%d foobar`, alias, -6),
			Stderr:  "invalid syntax",
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

func TestAliasCopy(t *testing.T) {
	token := fmt.Sprintf("GoTest-alias-copy-%d-", rand.Int())

	tests := []test.MurexTest{
		{
			Block: fmt.Sprintf(`
				summary %[1]s-foo-%[2]d foobar-%[2]d
				alias %[1]s-bar-%[2]d = %[1]s-foo-%[2]d
				runtime --summaries -> [%[1]s-bar-%[2]d]
			`, token, 0),
			Stderr:  "not found",
			ExitNum: 1,
		},
		{
			Block: fmt.Sprintf(`
				summary %[1]s-foo-%[2]d foobar-%[2]d
				alias --copy %[1]s-bar-%[2]d = %[1]s-foo-%[2]d
				runtime --summaries -> [%[1]s-bar-%[2]d]
			`, token, 1),
			Stdout: "^foobar-1$",
		},
		{
			Block: fmt.Sprintf(`
				autocomplete set %[1]s-foo-%[2]d %%[{Flags: [ "foobar-%[2]d" ]}]
				alias --copy %[1]s-bar-%[2]d = %[1]s-foo-%[2]d
				runtime --autocomplete -> [[ /%[1]s-foo-%[2]d/FlagValues/0/Flags/0 ]]
			`, token, 2),
			Stdout: "^foobar-2$",
		},
	}

	test.RunMurexTestsRx(tests, t)
}

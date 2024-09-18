package management_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestExitNum(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `exitnum`,
			Stdout: "0\n",
		},
		{
			Block:  `err bob; exitnum`,
			Stdout: "1\n",
			Stderr: "bob\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestBuiltinExists(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `bexists: bob`,
			Stdout:  `{"Installed":null,"Missing":["bob"]}` + "\n",
			ExitNum: 1,
		},
		{
			Block:  `bexists: cd`,
			Stdout: `{"Installed":["cd"],"Missing":null}` + "\n",
		},
		{
			Block:  `bexists: cd bexists`,
			Stdout: `{"Installed":["cd","bexists"],"Missing":null}` + "\n",
		},
		{
			Block:   `bexists: bob1 bob2`,
			Stdout:  `{"Installed":null,"Missing":["bob1","bob2"]}` + "\n",
			ExitNum: 2,
		},
		{
			Block:   `bexists: cd bob1 bexists bob2`,
			Stdout:  `{"Installed":["cd","bexists"],"Missing":["bob1","bob2"]}` + "\n",
			ExitNum: 2,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestBuiltinExistsErrors(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `bexists`,
			Stderr:  `missing parameters`,
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

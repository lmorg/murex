package generic_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestGenericIndex(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `a: [0..4] -> cast * -> [0]`,
			Stdout: `^0\n1\n2\n3\n4\n$`,
		},
		{
			Block:  `a: [0..4] -> cast * -> [:0]`,
			Stdout: `^0\n1\n2\n3\n4\n$`,
		},
		{
			Block:  `a: [0..4] -> cast * -> [0:]`,
			Stdout: `^0\n$`,
		},
		///
		{
			Block:  `a: [0..4] -> cast generic -> [0]`,
			Stdout: `^0\n1\n2\n3\n4\n$`,
		},
		{
			Block:  `a: [0..4] -> cast generic -> [:0]`,
			Stdout: `^0\n1\n2\n3\n4\n$`,
		},
		{
			Block:  `a: [0..4] -> cast generic -> [0:]`,
			Stdout: `^0\n$`,
		},

		//////

		{
			Block:  `a: [Mon..Fri] -> cast * -> [Mon]`,
			Stdout: `^Mon\nTue\nWed\nThu\nFri\n$`,
		},
		{
			Block:   `a: [Mon..Fri] -> cast * -> [:Mon]`,
			Stderr:  `^Error`,
			ExitNum: 1,
		},
		{
			Block:  `a: [Mon..Fri] -> cast * -> [:0]`,
			Stdout: `^Mon\nTue\nWed\nThu\nFri\n$`,
		},
		{
			Block:   `a: [Mon..Fri] -> cast * -> [Mon:]`,
			Stderr:  `^Error`,
			ExitNum: 1,
		},
		{
			Block:  `a: [Mon..Fri] -> cast * -> [0:]`,
			Stdout: `^Mon\n$`,
		},
		///
		{
			Block:  `a: [Mon..Fri] -> cast generic -> [Mon]`,
			Stdout: `^Mon\nTue\nWed\nThu\nFri\n$`,
		},
		{
			Block:   `a: [Mon..Fri] -> cast generic -> [:Mon]`,
			Stderr:  `^Error`,
			ExitNum: 1,
		},
		{
			Block:  `a: [Mon..Fri] -> cast generic -> [:0]`,
			Stdout: `^Mon\nTue\nWed\nThu\nFri\n$`,
		},
		{
			Block:   `a: [Mon..Fri] -> cast generic -> [Mon:]`,
			Stderr:  `^Error`,
			ExitNum: 1,
		},
		{
			Block:  `a: [Mon..Fri] -> cast generic -> [0:]`,
			Stdout: `^Mon\n$`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

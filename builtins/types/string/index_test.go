package string_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestStringIndex(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `a: [0..4] -> [0]`,
			Stdout: "^0\n$",
		},
		{
			Block:   `a: [0..4] -> [:0]`,
			Stderr:  "Parameter, `:0`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [0..4] -> [0:]`,
			Stderr:  "Parameter, `0:`, isn't an integer",
			ExitNum: 1,
		},
		///
		{
			Block:  `a: [0..4] -> cast str -> [0]`,
			Stdout: "^0\n$",
		},
		{
			Block:   `a: [0..4] -> cast str -> [:0]`,
			Stderr:  "Parameter, `:0`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [0..4] -> cast str -> [0:]`,
			Stderr:  "Parameter, `0:`, isn't an integer",
			ExitNum: 1,
		},
		///
		{
			Block:  `a: [0..4] -> cast string -> [0]`,
			Stdout: "^0\n$",
		},
		{
			Block:   `a: [0..4] -> cast string -> [:0]`,
			Stderr:  "Parameter, `:0`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [0..4] -> cast string -> [0:]`,
			Stderr:  "Parameter, `0:`, isn't an integer",
			ExitNum: 1,
		},

		//////////

		{
			Block:   `a: [Mon..Fri] -> [Mon]`,
			Stderr:  "Parameter, `Mon`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> [:Mon]`,
			Stderr:  "Parameter, `:Mon`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> [Mon:]`,
			Stderr:  "Parameter, `Mon:`, isn't an integer",
			ExitNum: 1,
		},
		///
		{
			Block:   `a: [Mon..Fri] -> cast str -> [Mon]`,
			Stderr:  "Parameter, `Mon`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> cast str -> [:Mon]`,
			Stderr:  "Parameter, `:Mon`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> cast str -> [Mon:]`,
			Stderr:  "Parameter, `Mon:`, isn't an integer",
			ExitNum: 1,
		},
		///
		{
			Block:   `a: [Mon..Fri] -> cast string -> [Mon]`,
			Stderr:  "Parameter, `Mon`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> cast string -> [:Mon]`,
			Stderr:  "Parameter, `:Mon`, isn't an integer",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> cast string -> [Mon:]`,
			Stderr:  "Parameter, `Mon:`, isn't an integer",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

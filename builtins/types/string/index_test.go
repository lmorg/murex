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
			Stdout: "0\n",
		},
		{
			Block:   `a: [0..4] -> [:0]`,
			Stderr:  "Error in `[` ( 1,14): Parameter, `:0`, isn't an integer. strconv.Atoi: parsing \":0\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [0..4] -> [0:]`,
			Stderr:  "Error in `[` ( 1,14): Parameter, `0:`, isn't an integer. strconv.Atoi: parsing \"0:\": invalid syntax\n",
			ExitNum: 1,
		},
		///
		{
			Block:  `a: [0..4] -> cast str -> [0]`,
			Stdout: "0\n",
		},
		{
			Block:   `a: [0..4] -> cast str -> [:0]`,
			Stderr:  "Error in `[` ( 1,26): Parameter, `:0`, isn't an integer. strconv.Atoi: parsing \":0\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [0..4] -> cast str -> [0:]`,
			Stderr:  "Error in `[` ( 1,26): Parameter, `0:`, isn't an integer. strconv.Atoi: parsing \"0:\": invalid syntax\n",
			ExitNum: 1,
		},
		///
		{
			Block:  `a: [0..4] -> cast string -> [0]`,
			Stdout: "0\n",
		},
		{
			Block:   `a: [0..4] -> cast string -> [:0]`,
			Stderr:  "Error in `[` ( 1,29): Parameter, `:0`, isn't an integer. strconv.Atoi: parsing \":0\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [0..4] -> cast string -> [0:]`,
			Stderr:  "Error in `[` ( 1,29): Parameter, `0:`, isn't an integer. strconv.Atoi: parsing \"0:\": invalid syntax\n",
			ExitNum: 1,
		},

		//////////

		{
			Block:   `a: [Mon..Fri] -> [Mon]`,
			Stderr:  "Error in `[` ( 1,18): Parameter, `Mon`, isn't an integer. strconv.Atoi: parsing \"Mon\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> [:Mon]`,
			Stderr:  "Error in `[` ( 1,18): Parameter, `:Mon`, isn't an integer. strconv.Atoi: parsing \":Mon\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> [Mon:]`,
			Stderr:  "Error in `[` ( 1,18): Parameter, `Mon:`, isn't an integer. strconv.Atoi: parsing \"Mon:\": invalid syntax\n",
			ExitNum: 1,
		},
		///
		{
			Block:   `a: [Mon..Fri] -> cast str -> [Mon]`,
			Stderr:  "Error in `[` ( 1,30): Parameter, `Mon`, isn't an integer. strconv.Atoi: parsing \"Mon\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> cast str -> [:Mon]`,
			Stderr:  "Error in `[` ( 1,30): Parameter, `:Mon`, isn't an integer. strconv.Atoi: parsing \":Mon\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> cast str -> [Mon:]`,
			Stderr:  "Error in `[` ( 1,30): Parameter, `Mon:`, isn't an integer. strconv.Atoi: parsing \"Mon:\": invalid syntax\n",
			ExitNum: 1,
		},
		///
		{
			Block:   `a: [Mon..Fri] -> cast string -> [Mon]`,
			Stderr:  "Error in `[` ( 1,33): Parameter, `Mon`, isn't an integer. strconv.Atoi: parsing \"Mon\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> cast string -> [:Mon]`,
			Stderr:  "Error in `[` ( 1,33): Parameter, `:Mon`, isn't an integer. strconv.Atoi: parsing \":Mon\": invalid syntax\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [Mon..Fri] -> cast string -> [Mon:]`,
			Stderr:  "Error in `[` ( 1,33): Parameter, `Mon:`, isn't an integer. strconv.Atoi: parsing \"Mon:\": invalid syntax\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}

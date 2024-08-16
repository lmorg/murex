package lists

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestListJoinMethod(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `%[] -> list.join '|'`,
			Stdout: "",
		},
		{
			Block:  `%[Mon..Fri] -> list.join '|'`,
			Stdout: "Mon|Tue|Wed|Thu|Fri",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestListJoinFunction(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `list.join '|'`, // empty array
			Stdout: "",
		},
		{
			Block:  `list.join '|' @{ %[Mon..Fri] }`,
			Stdout: "Mon|Tue|Wed|Thu|Fri",
		},
	}

	test.RunMurexTests(tests, t)
}

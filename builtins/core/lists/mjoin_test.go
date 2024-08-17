package lists

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestListJoinMethod(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `%[] -> mjoin '|'`,
			Stdout: "",
		},
		{
			Block:  `%[Mon..Fri] -> mjoin '|'`,
			Stdout: "Mon|Tue|Wed|Thu|Fri",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestListJoinFunction(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `mjoin '|'`, // empty array
			Stdout: "",
		},
		{
			Block:  `mjoin '|' @{ %[Mon..Fri] }`,
			Stdout: "Mon|Tue|Wed|Thu|Fri",
		},
	}

	test.RunMurexTests(tests, t)
}

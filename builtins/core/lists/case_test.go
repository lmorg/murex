package lists

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestListUpperMethod(t *testing.T) {
	tests := []test.MurexTest{
		/*{
			Block:  `%[] -> list.case upper`,
			Stdout: "",
		},*/
		{
			Block:  `%[Mon..Fri] -> list.case upper`,
			Stdout: `["MON","TUE","WED","THU","FRI"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestListUpperFunction(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `list.case upper`, // empty array
			Stdout: "[]",
		},
		{
			Block:  `list.case upper @{ %[Mon..Fri] }`,
			Stdout: `["MON","TUE","WED","THU","FRI"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestListTitleMethod(t *testing.T) {
	tests := []test.MurexTest{
		/*{
			Block:  `%[] -> list.case title`,
			Stdout: "",
		},*/
		{
			Block:  `%[MON..FRI] -> list.case title`,
			Stdout: `["Mon","Tue","Wed","Thu","Fri"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestListTitleFunction(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `list.case title`, // empty array
			Stdout: "[]",
		},
		{
			Block:  `list.case title @{ %[MON..FRI] }`,
			Stdout: `["Mon","Tue","Wed","Thu","Fri"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestListLowerMethod(t *testing.T) {
	tests := []test.MurexTest{
		/*{
			Block:  `%[] -> list.case lower`,
			Stdout: "",
		},*/
		{
			Block:  `%[MON..FRI] -> list.case lower`,
			Stdout: `["mon","tue","wed","thu","fri"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestListLowerFunction(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `list.case lower`, // empty array
			Stdout: "[]",
		},
		{
			Block:  `list.case lower @{ %[MON..FRI] }`,
			Stdout: `["mon","tue","wed","thu","fri"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

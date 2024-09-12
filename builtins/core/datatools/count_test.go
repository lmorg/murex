package datatools_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestCountTotal(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `%[1..5,2,3] -> count`,
			Stdout: `7`,
		},
		{
			Block:  `%[1..5,2,3] -> count --total`,
			Stdout: `7`,
		},
		{
			Block:  `%[1..5,2,3] -> count -t`,
			Stdout: `7`,
		},
		/////
		{
			Block:  `%{a:1, b:2, c:3, d:1, e:2} -> count`,
			Stdout: `5`,
		},
		{
			Block:  `%{a:1, b:2, c:3, d:1, e:2} -> count --total`,
			Stdout: `5`,
		},
		{
			Block:  `%{a:1, b:2, c:3, d:1, e:2} -> count -t`,
			Stdout: `5`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestCountDuplication(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `%[1..5,2,3] -> count --duplications`,
			Stdout: `^{"1":1,"2":2,"3":2,"4":1,"5":1}$`,
		},
		{
			Block:  `%[1..5,2,3] -> count -d`,
			Stdout: `^{"1":1,"2":2,"3":2,"4":1,"5":1}$`,
		},
		/////
		{
			Block:   `%{a:1, b:2, c:3, d:1, e:2} -> count --duplications`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
		{
			Block:   `%{a:1, b:2, c:3, d:1, e:2} -> count -d`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestCountSum(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `%[1..5,2,3] -> count --sum`,
			Stdout: `20`,
		},
		{
			Block:  `%[1..5,2,3] -> count -s`,
			Stdout: `20`,
		},
		/////
		{
			Block:  `%{a:1, b:2, c:3, d:1, e:2} -> count --sum`,
			Stdout: `9`,
		},
		{
			Block:  `%{a:1, b:2, c:3, d:1, e:2} -> count -s`,
			Stdout: `9`,
		},
		/////
		{
			Block:  `%[1..5,2,3] -> count --sum-strict`,
			Stdout: `20`,
		},
		{
			Block:  `%{a:1, b:2, c:3, d:1, e:2} -> count --sum-strict`,
			Stdout: `9`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestCountBytes(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `out "举手之劳就可以使办公室更加环保，比如，使用再生纸" -> count --bytes`,
			Stdout: `73`,
		},
		{
			Block:  `out "举手之劳就可以使办公室更加环保，比如，使用再生纸" -> count -b`,
			Stdout: `73`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestCountRunes(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `out "举手之劳就可以使办公室更加环保，比如，使用再生纸" -> count --runes`,
			Stdout: `25`,
		},
		{
			Block:  `out "举手之劳就可以使办公室更加环保，比如，使用再生纸" -> count -r`,
			Stdout: `25`,
		},
	}

	test.RunMurexTests(tests, t)
}

package structs_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestForEachParallel(t *testing.T) {
	tests := []test.MurexTest{
		// test --parallel runs tasks in parallel by added pauses and checking execution time
		{
			Block:  `time { a [1..6] -> foreach ! --parallel 0 {sleep 1} }`,
			Stdout: `^$`,
			Stderr: `^1\.[0-9]+$`,
		},
		{
			Block:  `time { a [1..6] -> foreach ! --parallel 1 {sleep 1} }`,
			Stdout: `^$`,
			Stderr: `^6\.[0-9]+$`,
		},
		{
			Block:  `time { a [1..6] -> foreach ! --parallel 2 {sleep 1} }`,
			Stdout: `^$`,
			Stderr: `^3\.[0-9]+$`,
		},
		{
			Block:  `time { a [1..6] -> foreach ! --parallel 3 {sleep 1} }`,
			Stdout: `^$`,
			Stderr: `^2\.[0-9]+$`,
		},
		{
			Block:  `time { a [1..6] -> foreach ! --parallel 4 {sleep 1} }`,
			Stdout: `^$`,
			Stderr: `^2\.[0-9]+$`,
		},
		{
			Block:  `time { a [1..6] -> foreach ! --parallel 6 {sleep 1} }`,
			Stdout: `^$`,
			Stderr: `^1\.[0-9]+$`,
		},
		{
			Block:  `time { a [1..6] -> foreach ! --parallel 7 {sleep 1} }`,
			Stdout: `^$`,
			Stderr: `^1\.[0-9]+$`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestForEachParallelVars(t *testing.T) {
	tests := []test.MurexTest{
		// test parallel passes vars to fork instead of parent (shouldn't error)
		{
			Block:  `a [1..10] -> foreach -p 1 i { echo $i }`,
			Stdout: `^1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n$`,
			Stderr: `^$`,
		},
		{
			Block:  `a [1..10] -> foreach -p 1 i { echo $.i }`,
			Stdout: `^0\n1\n2\n3\n4\n5\n6\n7\n8\n9\n$`,
			Stderr: `^$`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

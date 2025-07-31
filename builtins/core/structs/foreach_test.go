package structs_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestForEach(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `a: [Monday..Friday] -> foreach day { out "$day is the best day" }`,
			Stdout: "Monday is the best day\nTuesday is the best day\nWednesday is the best day\nThursday is the best day\nFriday is the best day\n",
		},
		{
			Block:  `a: [Monday..Friday] -> foreach day { out "$day is the best day" }`,
			Stdout: "Monday is the best day\nTuesday is the best day\nWednesday is the best day\nThursday is the best day\nFriday is the best day\n",
		},
		{
			Block:  `a: [Mon..Fri] -> foreach { -> suffix "day" }`,
			Stdout: "Monday\nTueday\nWedday\nThuday\nFriday\n",
		},
		{
			Block:  `ja: [Mon..Fri] -> foreach { -> suffix "day" }`,
			Stdout: "Monday\nTueday\nWedday\nThuday\nFriday\n",
		},
		{
			Block:  `a: [Mon..Fri] -> foreach --jmap day { $day } { out $day"day" }`,
			Stdout: `{"Fri":"Friday","Mon":"Monday","Thu":"Thuday","Tue":"Tueday","Wed":"Wedday"}`,
		},
		{
			Block:  `ja: [Mon..Fri] -> foreach --jmap day { $day } { out $day"day" }`,
			Stdout: `{"Fri":"Friday","Mon":"Monday","Thu":"Thuday","Tue":"Tueday","Wed":"Wedday"}`,
		},
		{
			Block:  `a: [Mon..Fri] -> foreach { out nothing } -> debug -> [[ /Data-Type/Murex  ]]`,
			Stdout: types.String,
		},
		{
			Block:  `ja: [Mon..Fri] -> foreach { out nothing } -> debug -> [[ /Data-Type/Murex  ]]`,
			Stdout: types.JsonLines,
		},
		{
			Block:  `a: [Mon..Fri] -> foreach { null } -> debug -> [[ /Data-Type/Murex  ]]`,
			Stdout: types.String,
		},
		{
			Block:  `ja: [Mon..Fri] -> foreach { null } -> debug -> [[ /Data-Type/Murex  ]]`,
			Stdout: types.JsonLines,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestForEachParallel(t *testing.T) {
	tests := []test.MurexTest{
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

func TestForEachStep(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `%[1..12] -> foreach --step 3 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3]\nIteration 2: [4,5,6]\nIteration 3: [7,8,9]\nIteration 4: [10,11,12]\n",
		},
		{
			Block:  `%[1..13] -> foreach --step 3 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3]\nIteration 2: [4,5,6]\nIteration 3: [7,8,9]\nIteration 4: [10,11,12]\nIteration 5: [13]\n",
		},
		{
			Block:  `%[1..14] -> foreach --step 3 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3]\nIteration 2: [4,5,6]\nIteration 3: [7,8,9]\nIteration 4: [10,11,12]\nIteration 5: [13,14]\n",
		},
		{
			Block:  `%[1..15] -> foreach --step 3 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3]\nIteration 2: [4,5,6]\nIteration 3: [7,8,9]\nIteration 4: [10,11,12]\nIteration 5: [13,14,15]\n",
		},
		{
			Block:  `%[1..16] -> foreach --step 3 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3]\nIteration 2: [4,5,6]\nIteration 3: [7,8,9]\nIteration 4: [10,11,12]\nIteration 5: [13,14,15]\nIteration 6: [16]\n",
		},
		/////
		{
			Block:  `%[1..10] -> foreach --step 5 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3,4,5]\nIteration 2: [6,7,8,9,10]\n",
		},
		{
			Block:  `%[1..11] -> foreach --step 5 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3,4,5]\nIteration 2: [6,7,8,9,10]\nIteration 3: [11]\n",
		},
		{
			Block:  `%[1..12] -> foreach --step 5 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3,4,5]\nIteration 2: [6,7,8,9,10]\nIteration 3: [11,12]\n",
		},
		{
			Block:  `%[1..13] -> foreach --step 5 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3,4,5]\nIteration 2: [6,7,8,9,10]\nIteration 3: [11,12,13]\n",
		},
		{
			Block:  `%[1..14] -> foreach --step 5 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3,4,5]\nIteration 2: [6,7,8,9,10]\nIteration 3: [11,12,13,14]\n",
		},
		{
			Block:  `%[1..15] -> foreach --step 5 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3,4,5]\nIteration 2: [6,7,8,9,10]\nIteration 3: [11,12,13,14,15]\n",
		},
		{
			Block:  `%[1..16] -> foreach --step 5 value { out "Iteration $.i: $value" }`,
			Stdout: "Iteration 1: [1,2,3,4,5]\nIteration 2: [6,7,8,9,10]\nIteration 3: [11,12,13,14,15]\nIteration 4: [16]\n",
		},
	}

	test.RunMurexTests(tests, t)
}

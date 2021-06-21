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

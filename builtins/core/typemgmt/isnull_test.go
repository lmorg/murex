package typemgmt_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

// https://github.com/lmorg/murex/issues/781
func TestIsNull(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `TestIsNull0 = "null"; is-null <null> TestIsNull0`,
			ExitNum: 1,
		},
		{
			Block:   `TestIsNull1 = null; is-null <null> TestIsNull1`,
			ExitNum: 0,
		},
		{
			Block:   `TestIsNull2 = "foobar"; is-null <null> TestIsNull2`,
			ExitNum: 1,
		},
		{
			Block:   `is-null <null> TestIsNull3`,
			ExitNum: 0,
		},
		/////
		{
			Block:   `TestIsNull4 = "null"; is-null <null> $TestIsNull4`,
			ExitNum: 1,
		},
		{
			Block:   `TestIsNull5 = null; is-null <null> $TestIsNull5`,
			ExitNum: 0,
		},
		{
			Block:   `TestIsNull6 = "foobar"; is-null <null> $TestIsNull6`,
			ExitNum: 1,
		},
		{
			Block:   `is-null <null> $TestIsNull7`,
			ExitNum: 0,
		},
		/////
		{
			Block: `
				%{ "foo": null } -> set TestIsNull8
				if { is-null $TestIsNull8['foo'] } then {
					out "foo is null"
				} else {
					out "foo is NOT null"
				}`,
			Stdout:  "foo is null\n",
			ExitNum: 0,
		},
		{
			Block: `
				%{ "foo": null } -> set TestIsNull9
				$TestIsNull9 -> formap k v {
					if { is-null v } then {
						out "$k is null"
					} else {
						out "$k is NOT null"
					}
				}`,
			Stdout:  "foo is null\n",
			ExitNum: 0,
		},
	}

	test.RunMurexTests(tests, t)
}

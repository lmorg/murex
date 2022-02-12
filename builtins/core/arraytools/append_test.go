package arraytools

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/generic"
	_ "github.com/lmorg/murex/builtins/types/json"
	_ "github.com/lmorg/murex/builtins/types/jsonlines"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

/*
	ADDHEADING
*/

func TestAddheadingJsonl(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout: jsonl '["Bob", 23, true]' -> addheading name age active`,
			Stdout: "[\"name\",\"age\",\"active\"]\n[\"Bob\",\"23\",\"true\"]\n",
		},
	}

	test.RunMurexTests(tests, t)
}

/*
	PREPEND
*/

// TestPrependJsonStr tests the p with a JSON array of strings
func TestPrependJsonStr(t *testing.T) {
	test.RunMethodTest(t,
		cmdPrepend, "prepend",
		`["a","b","c"]`,
		types.Json,
		[]string{"new"},
		`["new","a","b","c"]`,
		nil,
	)
}

// TestPrependJsonInt tests the prepend method with a JSON array of integers
func TestPrependJsonInt(t *testing.T) {
	test.RunMethodTest(t,
		cmdPrepend, "prepend",
		`[1,2,3]`,
		types.Json,
		[]string{"9"},
		`["9","1","2","3"]`,
		nil,
	)
}

// TestPrependJsonMixed tests the prepend method with a JSON array of mixed data types
func TestPrependJsonMixed(t *testing.T) {
	test.RunMethodTest(t,
		cmdPrepend, "prepend",
		`[1,2,3]`,
		types.Json,
		[]string{"new"},
		`["new","1","2","3"]`,
		nil,
	)
}

// TestPrependGenericStr tests the prepend method with a generic list of strings
func TestPrependGenericStr(t *testing.T) {
	test.RunMethodTest(t,
		cmdPrepend, "prepend",
		"a\nb\nc",
		types.Generic,
		[]string{"new"},
		"new\na\nb\nc\n",
		nil,
	)
}

// TestPrependGenericInt tests the prepend method with a generic list of integers
func TestPrependGenericInt(t *testing.T) {
	test.RunMethodTest(t,
		cmdPrepend, "prepend",
		"1\n2\n3",
		types.Generic,
		[]string{"9"},
		"9\n1\n2\n3\n",
		nil,
	)
}

// TestPrependGenericMixed tests the prepend method with a generic list of mixed data types
func TestPrependGenericMixed(t *testing.T) {
	test.RunMethodTest(t,
		cmdPrepend, "prepend",
		"1\n2\n3",
		types.Generic,
		[]string{"new"},
		"new\n1\n2\n3\n",
		nil,
	)
}

/*
	APPEND
*/

// TestAppendJsonStr tests the append method with a JSON array of strings
func TestAppendJsonStr(t *testing.T) {
	test.RunMethodTest(t,
		cmdAppend, "append",
		`["a","b","c"]`,
		types.Json,
		[]string{"new"},
		`["a","b","c","new"]`,
		nil,
	)
}

// TestAppendJsonInt tests the append method with a JSON array of integers
func TestAppendJsonInt(t *testing.T) {
	test.RunMethodTest(t,
		cmdAppend, "append",
		`[1,2,3]`,
		types.Json,
		[]string{"9"},
		`["1","2","3","9"]`,
		nil,
	)
}

// TestAppendJsonMixed tests the append method with a JSON array of mixed data types
func TestAppendJsonMixed(t *testing.T) {
	test.RunMethodTest(t,
		cmdAppend, "append",
		`[1,2,3]`,
		types.Json,
		[]string{"new"},
		`["1","2","3","new"]`,
		nil,
	)
}

// TestAppendGenericStr tests the append method with a generic list of strings
func TestAppendGenericStr(t *testing.T) {
	test.RunMethodTest(t,
		cmdAppend, "append",
		"a\nb\nc",
		types.Generic,
		[]string{"new"},
		"a\nb\nc\nnew\n",
		nil,
	)
}

// TestAppendGenericInt tests the append method with a generic list of integers
func TestAppendGenericInt(t *testing.T) {
	test.RunMethodTest(t,
		cmdAppend, "append",
		"1\n2\n3",
		types.Generic,
		[]string{"9"},
		"1\n2\n3\n9\n",
		nil,
	)
}

// TestAppendGenericMixed tests the append method with a generic list of mixed data types
func TestAppendGenericMixed(t *testing.T) {
	test.RunMethodTest(t,
		cmdAppend, "append",
		"1\n2\n3",
		types.Generic,
		[]string{"new"},
		"1\n2\n3\nnew\n",
		nil,
	)
}

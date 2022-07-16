package index_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	_ "github.com/lmorg/murex/builtins/types/json"
	_ "github.com/lmorg/murex/builtins/types/jsonlines"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/utils/json"
)

var table = [][]string{
	{"a", "s", "l"},
	{"21", "m", "london"},
	{"32", "f", "spain"},
	{"43", "m", "italy"},
	{"54", "f", "france"},
	{"65", "m", "london"},
}
var jTable = json.LazyLogging(table)

var object = map[string]map[string][]int{
	"london": {"m": {21, 65}},
	"spain":  {"f": {32}},
	"italy":  {"m": {43}},
	"france": {"f": {54}},
}
var jObject = json.LazyLogging(object)

func TestIndexObject(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`tout json (%s) -> [spain] -> [f]`, jObject),
			Stdout: "[32]\n",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s) -> [ [/spain/f] ]`, jObject),
			Stdout: "[32]\n",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s) -> [spain] -> [f] -> [0]`, jObject),
			Stdout: "32",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s) -> [ [/spain/f/0] ]`, jObject),
			Stdout: "32",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s) -> [london] -> [m]`, jObject),
			Stdout: "[21,65]\n",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s) -> [ [/london/m] ]`, jObject),
			Stdout: "[21,65]\n",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s) -> [london] -> [m] -> [1]`, jObject),
			Stdout: "65",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s) -> [ [/london/m/1] ]`, jObject),
			Stdout: "65",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestIndexTable(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`tout json (%s)->format jsonl->[1]`, jTable),
			Stdout: `["21","m","london"]` + "\n",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s)->format jsonl->[l]`, jTable),
			Stdout: "[\"l\"]\n[\"london\"]\n[\"spain\"]\n[\"italy\"]\n[\"france\"]\n[\"london\"]\n",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s)->format jsonl->[:2]`, jTable),
			Stdout: "[\"l\"]\n[\"london\"]\n[\"spain\"]\n[\"italy\"]\n[\"france\"]\n[\"london\"]\n",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s)->format jsonl->[2:]`, jTable),
			Stdout: "[\"32\",\"f\",\"spain\"]\n",
		},
		{
			Block:  fmt.Sprintf(`tout json (%s)->format jsonl->[2]`, jTable),
			Stdout: "[\"32\",\"f\",\"spain\"]\n",
		},
	}

	test.RunMurexTests(tests, t)
}

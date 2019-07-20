package mkarray

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/test"
)

func TestArrayOct(t *testing.T) {
	tests := []test.MurexTest{
		// Octal defined by x
		{
			Block:  `a: [1..10x8]`,
			Stdout: "1\n2\n3\n4\n5\n6\n7\n10\n",
		},
		{
			Block:  `a: [01..10x8]`,
			Stdout: "01\n02\n03\n04\n05\n06\n07\n10\n",
		},
		{
			Block:  `a: [001..10x8]`,
			Stdout: "001\n002\n003\n004\n005\n006\n007\n010\n",
		},
		// Octal defined by period
		{
			Block:  `a: [1..10.8]`,
			Stdout: "1\n2\n3\n4\n5\n6\n7\n10\n",
		},
		{
			Block:  `a: [01..10.8]`,
			Stdout: "01\n02\n03\n04\n05\n06\n07\n10\n",
		},
		{
			Block:  `a: [001..10.8]`,
			Stdout: "001\n002\n003\n004\n005\n006\n007\n010\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestArrayDec(t *testing.T) {
	tests := []test.MurexTest{
		// Decimal
		{
			Block:  `a: [1..10]`,
			Stdout: "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n",
		},
		{
			Block:  `a: [01..10]`,
			Stdout: "01\n02\n03\n04\n05\n06\n07\n08\n09\n10\n",
		},
		{
			Block:  `a: [001..10]`,
			Stdout: "001\n002\n003\n004\n005\n006\n007\n008\n009\n010\n",
		},
		// Decimal defined by x
		{
			Block:  `a: [1..10x10]`,
			Stdout: "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n",
		},
		{
			Block:  `a: [01..10x10]`,
			Stdout: "01\n02\n03\n04\n05\n06\n07\n08\n09\n10\n",
		},
		{
			Block:  `a: [001..10x10]`,
			Stdout: "001\n002\n003\n004\n005\n006\n007\n008\n009\n010\n",
		},
		// Decimal defined by period
		{
			Block:  `a: [1..10.10]`,
			Stdout: "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n",
		},
		{
			Block:  `a: [01..10.10]`,
			Stdout: "01\n02\n03\n04\n05\n06\n07\n08\n09\n10\n",
		},
		{
			Block:  `a: [001..10.10]`,
			Stdout: "001\n002\n003\n004\n005\n006\n007\n008\n009\n010\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestArrayHex(t *testing.T) {
	tests := []test.MurexTest{
		// Hexadecimal defined by x
		{
			Block:  `a: [1..10x16]`,
			Stdout: "1\n2\n3\n4\n5\n6\n7\n8\n9\na\nb\nc\nd\ne\nf\n10\n",
		},
		{
			Block:  `a: [01..10x16]`,
			Stdout: "01\n02\n03\n04\n05\n06\n07\n08\n09\n0a\n0b\n0c\n0d\n0e\n0f\n10\n",
		},
		{
			Block:  `a: [001..10x16]`,
			Stdout: "001\n002\n003\n004\n005\n006\n007\n008\n009\n00a\n00b\n00c\n00d\n00e\n00f\n010\n",
		},
		// Hexadecimal defined by period
		{
			Block:  `a: [1..10.16]`,
			Stdout: "1\n2\n3\n4\n5\n6\n7\n8\n9\na\nb\nc\nd\ne\nf\n10\n",
		},
		{
			Block:  `a: [01..10.16]`,
			Stdout: "01\n02\n03\n04\n05\n06\n07\n08\n09\n0a\n0b\n0c\n0d\n0e\n0f\n10\n",
		},
		{
			Block:  `a: [001..10.16]`,
			Stdout: "001\n002\n003\n004\n005\n006\n007\n008\n009\n00a\n00b\n00c\n00d\n00e\n00f\n010\n",
		},
	}

	test.RunMurexTests(tests, t)
}

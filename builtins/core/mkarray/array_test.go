package mkarray

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/json"
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

func TestArrayDecTyped(t *testing.T) {
	tests := []test.MurexTest{
		// Decimal, typed
		{
			Block:  `ta: json [0..10]`,
			Stdout: "[0,1,2,3,4,5,6,7,8,9,10]",
		},
		{
			Block:  `ta: json [01..10]`,
			Stdout: `["01","02","03","04","05","06","07","08","09","10"]`,
		},
		{
			Block:  `ta: json [00..10]`,
			Stdout: `["00","01","02","03","04","05","06","07","08","09","10"]`,
		},
		{
			Block:  `ja: [0..10]`,
			Stdout: "[0,1,2,3,4,5,6,7,8,9,10]",
		},
		{
			Block:  `ja: [01..10]`,
			Stdout: `["01","02","03","04","05","06","07","08","09","10"]`,
		},
		{
			Block:  `ja: [00..10]`,
			Stdout: `["00","01","02","03","04","05","06","07","08","09","10"]`,
		},
		/////
		{
			Block:  `ta: json [10..0]`,
			Stdout: "[10,9,8,7,6,5,4,3,2,1,0]",
		},
		{
			Block:  `ta: json [10..01]`,
			Stdout: `["10","09","08","07","06","05","04","03","02","01"]`,
		},
		{
			Block:  `ta: json [10..00]`,
			Stdout: `["10","09","08","07","06","05","04","03","02","01","00"]`,
		},
		{
			Block:  `ja: [10..0]`,
			Stdout: "[10,9,8,7,6,5,4,3,2,1,0]",
		},
		{
			Block:  `ja: [10..01]`,
			Stdout: `["10","09","08","07","06","05","04","03","02","01"]`,
		},
		{
			Block:  `ja: [10..00]`,
			Stdout: `["10","09","08","07","06","05","04","03","02","01","00"]`,
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

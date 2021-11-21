package mkarray

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/test"
)

func TestRangeMonth(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `a: [June..October]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [June..January]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\nJanuary\n",
		},
		{
			Block:  `a: [December..June]`,
			Stdout: "December\nJanuary\nFebruary\nMarch\nApril\nMay\nJune\n",
		},
		{
			Block:   `a: [..June]`,
			Stdout:  "",
			Stderr:  "Error in `a` ( 1,1): unable to auto-detect range in `..June`\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [June..]`,
			Stdout:  "",
			Stderr:  "Error in `a` ( 1,1): unable to auto-detect range in `June..`\n",
			ExitNum: 1,
		},
		// lowercase
		{
			Block:  `a: [june..october]`,
			Stdout: "june\njuly\naugust\nseptember\noctober\n",
		},
		{
			Block:  `a: [june..january]`,
			Stdout: "june\njuly\naugust\nseptember\noctober\nnovember\ndecember\njanuary\n",
		},
		{
			Block:  `a: [december..june]`,
			Stdout: "december\njanuary\nfebruary\nmarch\napril\nmay\njune\n",
		},
		// uppercase
		{
			Block:  `a: [JUNE..OCTOBER]`,
			Stdout: "JUNE\nJULY\nAUGUST\nSEPTEMBER\nOCTOBER\n",
		},
		{
			Block:  `a: [JUNE..JANUARY]`,
			Stdout: "JUNE\nJULY\nAUGUST\nSEPTEMBER\nOCTOBER\nNOVEMBER\nDECEMBER\nJANUARY\n",
		},
		{
			Block:  `a: [DECEMBER..JUNE]`,
			Stdout: "DECEMBER\nJANUARY\nFEBRUARY\nMARCH\nAPRIL\nMAY\nJUNE\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRange(t *testing.T) {
	tests := []test.MurexTest{
		// single range
		{
			Block:  `a: [00..03]`,
			Stdout: "00\n01\n02\n03\n",
		},
		{
			Block:  `a: [01,03]`,
			Stdout: "01\n03\n",
		},
		// multiple ranges
		{
			Block:  `a: [01..02][01..02]`,
			Stdout: "0101\n0102\n0201\n0202\n",
		},
		{
			Block:  `a: [01,03][01..02]`,
			Stdout: "0101\n0102\n0301\n0302\n",
		},
		{
			Block:  `a: [01..02][01,03]`,
			Stdout: "0101\n0103\n0201\n0203\n",
		},
		{
			Block:  `a: [01,03][01,03]`,
			Stdout: "0101\n0103\n0301\n0303\n",
		},
		// multiple ranges with non-ranged data
		{
			Block:  `a: .[01..02].[01..02].`,
			Stdout: ".01.01.\n.01.02.\n.02.01.\n.02.02.\n",
		},
		{
			Block:  `a: .[01,03].[01..02].`,
			Stdout: ".01.01.\n.01.02.\n.03.01.\n.03.02.\n",
		},
		{
			Block:  `a: .[01..02].[01,03].`,
			Stdout: ".01.01.\n.01.03.\n.02.01.\n.02.03.\n",
		},
		{
			Block:  `a: .[01,03].[01,03].`,
			Stdout: ".01.01.\n.01.03.\n.03.01.\n.03.03.\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRangeDown(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `a: [10..7]`,
			Stdout: "10\n9\n8\n7\n",
		},
		{
			Block:  `a: [10..07]`,
			Stdout: "10\n09\n08\n07\n",
		},
		{
			Block:  `a: [010..007]`,
			Stdout: "010\n009\n008\n007\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRangeAlpha(t *testing.T) {
	tests := []test.MurexTest{
		// up
		{
			Block:  `a: [a..d]`,
			Stdout: "a\nb\nc\nd\n",
		},
		{
			Block:  `a: [A..D]`,
			Stdout: "A\nB\nC\nD\n",
		},
		// down
		{
			Block:  `a: [d..a]`,
			Stdout: "d\nc\nb\na\n",
		},
		{
			Block:  `a: [D..A]`,
			Stdout: "D\nC\nB\nA\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRangeAltNumberBase(t *testing.T) {
	tests := []test.MurexTest{
		// up
		{
			Block:  `a: [00..10x8]`,
			Stdout: "00\n01\n02\n03\n04\n05\n06\n07\n10\n",
		},
		{
			Block:  `a: [00..10x16]`,
			Stdout: "00\n01\n02\n03\n04\n05\n06\n07\n08\n09\n0a\n0b\n0c\n0d\n0e\n0f\n10\n",
		},
		// down
		{
			Block:  `a: [10..00x8]`,
			Stdout: "10\n7\n6\n5\n4\n3\n2\n1\n0\n",
		},
		{
			Block:  `a: [010..0x8]`,
			Stdout: "010\n007\n006\n005\n004\n003\n002\n001\n000\n",
		},
		{
			Block:  `a: [10..00x16]`,
			Stdout: "10\nf\ne\nd\nc\nb\na\n9\n8\n7\n6\n5\n4\n3\n2\n1\n0\n",
		},
		{
			Block:  `a: [010..0x16]`,
			Stdout: "010\n00f\n00e\n00d\n00c\n00b\n00a\n009\n008\n007\n006\n005\n004\n003\n002\n001\n000\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRangeDate(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `a: [25-Jun-20..05-Jul-20]`,
			Stdout: "25-Jun-20\n26-Jun-20\n27-Jun-20\n28-Jun-20\n29-Jun-20\n30-Jun-20\n01-Jul-20\n02-Jul-20\n03-Jul-20\n04-Jul-20\n05-Jul-20\n",
		},
		{
			Block:  `a: [03 Jan 20..28 Dec 19]`,
			Stdout: "03 Jan 20\n02 Jan 20\n01 Jan 20\n31 Dec 19\n30 Dec 19\n29 Dec 19\n28 Dec 19\n",
		},
		{
			Block:  `a: [25/June/2020..05/July/2020]`,
			Stdout: "25/June/2020\n26/June/2020\n27/June/2020\n28/June/2020\n29/June/2020\n30/June/2020\n01/July/2020\n02/July/2020\n03/July/2020\n04/July/2020\n05/July/2020\n",
		},
		{
			Block:  `a: [03-January-2020..28-December-2019]`,
			Stdout: "03-January-2020\n02-January-2020\n01-January-2020\n31-December-2019\n30-December-2019\n29-December-2019\n28-December-2019\n",
		},
		// lowercase
		{
			Block:  `a: [25-jun-20..05-jul-20]`,
			Stdout: "25-jun-20\n26-jun-20\n27-jun-20\n28-jun-20\n29-jun-20\n30-jun-20\n01-jul-20\n02-jul-20\n03-jul-20\n04-jul-20\n05-jul-20\n",
		},
		{
			Block:  `a: [03-jan-20..28-dec-19]`,
			Stdout: "03-jan-20\n02-jan-20\n01-jan-20\n31-dec-19\n30-dec-19\n29-dec-19\n28-dec-19\n",
		},
		{
			Block:  `a: [25-june-2020..05-july-2020]`,
			Stdout: "25-june-2020\n26-june-2020\n27-june-2020\n28-june-2020\n29-june-2020\n30-june-2020\n01-july-2020\n02-july-2020\n03-july-2020\n04-july-2020\n05-july-2020\n",
		},
		{
			Block:  `a: [03-january-2020..28-december-2019]`,
			Stdout: "03-january-2020\n02-january-2020\n01-january-2020\n31-december-2019\n30-december-2019\n29-december-2019\n28-december-2019\n",
		},
		// uppercase
		{
			Block:  `a: [03-JAN-20..28-DEC-19]`,
			Stdout: "03-JAN-20\n02-JAN-20\n01-JAN-20\n31-DEC-19\n30-DEC-19\n29-DEC-19\n28-DEC-19\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRangeDateFormats(t *testing.T) {
	tests := []test.MurexTest{
		// dd mm yy
		{
			Block:  `a: [01-jan-20..03-jan-20]`,
			Stdout: "01-jan-20\n02-jan-20\n03-jan-20\n",
		},
		{
			Block:  `a: [03-jan-20..01-jan-20]`,
			Stdout: "03-jan-20\n02-jan-20\n01-jan-20\n",
		},
		{
			Block:  `a: [01 jan 20..03 jan 20]`,
			Stdout: "01 jan 20\n02 jan 20\n03 jan 20\n",
		},
		{
			Block:  `a: [03 jan 20..01 jan 20]`,
			Stdout: "03 jan 20\n02 jan 20\n01 jan 20\n",
		},
		{
			Block:  `a: [01/jan/20..03/jan/20]`,
			Stdout: "01/jan/20\n02/jan/20\n03/jan/20\n",
		},
		{
			Block:  `a: [03/jan/20..01/jan/20]`,
			Stdout: "03/jan/20\n02/jan/20\n01/jan/20\n",
		},
		{
			Block:  `a: [01-january-20..03-january-20]`,
			Stdout: "01-january-20\n02-january-20\n03-january-20\n",
		},
		{
			Block:  `a: [03-january-20..01-january-20]`,
			Stdout: "03-january-20\n02-january-20\n01-january-20\n",
		},
		{
			Block:  `a: [01 january 20..03 january 20]`,
			Stdout: "01 january 20\n02 january 20\n03 january 20\n",
		},
		{
			Block:  `a: [03 january 20..01 january 20]`,
			Stdout: "03 january 20\n02 january 20\n01 january 20\n",
		},
		{
			Block:  `a: [01/january/20..03/january/20]`,
			Stdout: "01/january/20\n02/january/20\n03/january/20\n",
		},
		{
			Block:  `a: [03/january/20..01/january/20]`,
			Stdout: "03/january/20\n02/january/20\n01/january/20\n",
		},

		// mm dd yy

		{
			Block:  `a: [jan-01-20..jan-03-20]`,
			Stdout: "jan-01-20\njan-02-20\njan-03-20\n",
		},
		{
			Block:  `a: [jan-03-20..jan-01-20]`,
			Stdout: "jan-03-20\njan-02-20\njan-01-20\n",
		},
		{
			Block:  `a: [jan 01 20..jan 03 20]`,
			Stdout: "jan 01 20\njan 02 20\njan 03 20\n",
		},
		{
			Block:  `a: [jan 03 20..jan 01 20]`,
			Stdout: "jan 03 20\njan 02 20\njan 01 20\n",
		},
		{
			Block:  `a: [jan/01/20..jan/03/20]`,
			Stdout: "jan/01/20\njan/02/20\njan/03/20\n",
		},
		{
			Block:  `a: [jan/03/20..jan/01/20]`,
			Stdout: "jan/03/20\njan/02/20\njan/01/20\n",
		},
		{
			Block:  `a: [january-01-20..january-03-20]`,
			Stdout: "january-01-20\njanuary-02-20\njanuary-03-20\n",
		},
		{
			Block:  `a: [january-03-20..january-01-20]`,
			Stdout: "january-03-20\njanuary-02-20\njanuary-01-20\n",
		},
		{
			Block:  `a: [january 01 20..january 03 20]`,
			Stdout: "january 01 20\njanuary 02 20\njanuary 03 20\n",
		},
		{
			Block:  `a: [january 03 20..january 01 20]`,
			Stdout: "january 03 20\njanuary 02 20\njanuary 01 20\n",
		},
		{
			Block:  `a: [january/01/20..january/03/20]`,
			Stdout: "january/01/20\njanuary/02/20\njanuary/03/20\n",
		},
		{
			Block:  `a: [january/03/20..january/01/20]`,
			Stdout: "january/03/20\njanuary/02/20\njanuary/01/20\n",
		},

		// dd mm

		{
			Block:  `a: [01-jan..03-jan]`,
			Stdout: "01-jan\n02-jan\n03-jan\n",
		},
		{
			Block:  `a: [03-jan..01-jan]`,
			Stdout: "03-jan\n02-jan\n01-jan\n",
		},
		{
			Block:  `a: [01 jan..03 jan]`,
			Stdout: "01 jan\n02 jan\n03 jan\n",
		},
		{
			Block:  `a: [03 jan..01 jan]`,
			Stdout: "03 jan\n02 jan\n01 jan\n",
		},
		{
			Block:  `a: [01/jan..03/jan]`,
			Stdout: "01/jan\n02/jan\n03/jan\n",
		},
		{
			Block:  `a: [03/jan..01/jan]`,
			Stdout: "03/jan\n02/jan\n01/jan\n",
		},
		{
			Block:  `a: [01-january..03-january]`,
			Stdout: "01-january\n02-january\n03-january\n",
		},
		{
			Block:  `a: [03-january..01-january]`,
			Stdout: "03-january\n02-january\n01-january\n",
		},
		{
			Block:  `a: [01 january..03 january]`,
			Stdout: "01 january\n02 january\n03 january\n",
		},
		{
			Block:  `a: [03 january..01 january]`,
			Stdout: "03 january\n02 january\n01 january\n",
		},
		{
			Block:  `a: [01/january..03/january]`,
			Stdout: "01/january\n02/january\n03/january\n",
		},
		{
			Block:  `a: [03/january..01/january]`,
			Stdout: "03/january\n02/january\n01/january\n",
		},
	}

	test.RunMurexTests(tests, t)
}

//go:build ignore
// +build ignore

package csvbad

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/csv"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test"
)

/*func TestReadArray(t *testing.T) {
	input := []byte("England,Scotland,Ireland\na,b,c\n1,2,3\n")

	expected := []string{
		"England,Scotland,Ireland",
		"a,b,c",
		"1,2,3",
	}

	test.ReadArrayTest(t, typeName, input, expected)
}*/

func TestReadMap(t *testing.T) {
	input := []byte("England,Scotland,Ireland\na,b,c\n1,2,3\n")

	expected := []test.ReadMapExpected{
		{
			Key:   "England",
			Value: "a",
			Last:  false,
		},
		{
			Key:   "Scotland",
			Value: "b",
			Last:  false,
		},
		{
			Key:   "Ireland",
			Value: "c",
			Last:  true,
		},
		// -- next row
		{
			Key:   "England",
			Value: "1",
			Last:  false,
		},
		{
			Key:   "Scotland",
			Value: "2",
			Last:  false,
		},
		{
			Key:   "Ireland",
			Value: "3",
			Last:  true,
		},
	}

	lang.InitConf.Set("csv", "headings", true)

	test.ReadMapUnorderedTest(t, typeName, input, expected, config.InitConf)
}

/*func TestArrayWriter(t *testing.T) {
	input := []string{
		"England,Scotland,Ireland",
		"a,b,c",
		"1,2,3",
	}

	output := "England,Scotland,Ireland\na,b,c\n1,2,3\n"
	test.ArrayWriterTest(t, typeName, input, output)
}*/

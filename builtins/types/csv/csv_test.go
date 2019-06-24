package csv

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test"
)

/*func TestReadArray(t *testing.T) {
	input := []byte(`
	[
		"foo",
		"bar"
	]`)

	expected := []string{
		"foo",
		"bar",
	}

	test.ReadArrayTest(t,typeName, input, expected)
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

	test.ReadMapOrderedTest(t, typeName, input, expected, lang.InitConf)
}

/*func TestArrayWriter(t *testing.T) {
	input := []string{"foo", "bar"}
	output := `["foo","bar"]`
	test.ArrayWriterTest(t, typeName, input, output)
}
*/

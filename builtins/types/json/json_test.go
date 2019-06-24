package json

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestReadArray(t *testing.T) {
	input := []byte(`
	[
		"foo",
		"bar"
	]`)

	expected := []string{
		"foo",
		"bar",
	}

	test.ReadArrayTest(t, types.Json, input, expected)
}

func TestReadMap(t *testing.T) {
	config := config.NewConfiguration()

	input := []byte(`
	{
		"foo": "oof",
		"bar": "rab"
	}`)

	expected := []test.ReadMapExpected{
		{
			Key:   "foo",
			Value: "oof",
			Last:  true,
		},
		{
			Key:   "bar",
			Value: "rab",
			Last:  false,
		},
	}

	test.ReadMapUnorderedTest(t, types.Json, input, expected, config)
}

func TestArrayWriter(t *testing.T) {
	input := []string{"foo", "bar"}
	output := `["foo","bar"]`
	test.ArrayWriterTest(t, types.Json, input, output)
}

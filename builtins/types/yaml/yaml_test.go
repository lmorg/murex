package yaml

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/test"
)

func TestReadArray(t *testing.T) {
	input := []byte("- foo\n- bar\n")

	expected := []string{
		"foo",
		"bar",
	}

	test.ReadArrayTest(t, typeName, input, expected)
}

func TestReadMap(t *testing.T) {
	config := config.NewConfiguration()

	input := []byte("foo: oof\nbar: rab\n")

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

	test.ReadMapOrderedTest(t, typeName, input, expected, config)
}

func TestArrayWriter(t *testing.T) {
	input := []string{"foo", "bar"}
	output := "- foo\n- bar\n"
	test.ArrayWriterTest(t, typeName, input, output)
}

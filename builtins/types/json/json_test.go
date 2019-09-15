package json

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
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

func TestMarshalArrayString(t *testing.T) {
	count.Tests(t, 1, "TestMarshalArrayString")

	input := []string{"e", "d", "c", "b", "a"} // lets prove the output retains sorting
	output := `["e","d","c","b","a"]`

	lang.InitEnv()
	fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)

	b, err := marshal(fork.Process, input)
	if err != nil {
		t.Error(err)
		return
	}

	if string(b) != output {
		t.Error("Marshaller output doesn't match expected:")
		t.Logf("  Input:    %v", input)
		t.Logf("  Expected: '%s'", output)
		t.Logf("  Actual:   '%s'", b)
	}
}

func TestMarshalArrayInt(t *testing.T) {
	count.Tests(t, 1, "TestMarshalArrayInt")

	input := []int{5, 4, 3, 2, 1} // lets prove the output retains sorting
	output := `[5,4,3,2,1]`

	lang.InitEnv()
	fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)

	b, err := marshal(fork.Process, input)
	if err != nil {
		t.Error(err)
		return
	}

	if string(b) != output {
		t.Error("Marshaller output doesn't match expected:")
		t.Logf("   Input:     %v", input)
		t.Logf("  Expected: '%s'", output)
		t.Logf("  Actual:   '%s'", b)
	}
}

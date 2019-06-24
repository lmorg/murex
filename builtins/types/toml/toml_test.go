package toml

import (
	"testing"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/test"
)

func TestReadMap(t *testing.T) {
	config := config.NewConfiguration()

	input := []byte("bar = \"rab\"\nfoo = \"oof\"\n")

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

	test.ReadMapUnorderedTest(t, typeName, input, expected, config)
}

func TestArrayWriter(t *testing.T) {
	stdout := streams.NewStdin()

	_, err := stdout.WriteArray(typeName)
	switch {
	case err == nil:
		t.Error("Missing error condition! This test should produce an error and no error was raised")

	case err != errNakedArrays:
		t.Errorf("Error raised was unexpected: %s", err)
		t.Logf("  (Expected error was: %s)", errNakedArrays)
	}
}

package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/config"
)

// ArrayWriterTest is an easy template for testing stdio.ArrayWriter methods in murex types
func ArrayWriterTest(t *testing.T, dataType string, input []string, expected string) {
	stdout := streams.NewStdin()

	w, err := stdout.WriteArray(dataType)
	if err != nil {
		t.Fatalf("Unable to create ArrayWriter: %s", err)
	}

	for _, s := range input {
		err = w.WriteString(s)
		if err != nil {
			t.Fatalf("Unable to write %s: %s", s, err)
		}
	}

	err = w.Close()
	if err != nil {
		t.Fatalf("Unable to close ArrayWriter: %s", err)
	}

	b, err := stdout.ReadAll()
	if err != nil {
		t.Fatalf("Unable to ReadAll from stdout: %s", err)
	}

	compExp := strings.Replace(expected, "\n", `\n`, -1)
	compAct := strings.Replace(string(b), "\n", `\n`, -1)

	if compExp != compAct {
		t.Error("Unexpected output in ArrayWriter:")
		t.Logf("  Expected: %s", compExp)
		t.Logf("  Actual:   %s", compAct)
	}
}

// ReadArrayTest is an easy template for testing stdio.ReadArray methods in murex types
func ReadArrayTest(t *testing.T, dataType string, input []byte, expected []string) {
	stdout := streams.NewStdin()
	stdout.SetDataType(dataType)
	_, err := stdout.Write(input)

	if err != nil {
		t.Fatalf("Unable to Write to stdout: %s", err)
	}

	actual := make([]string, 0)
	err = stdout.ReadArray(func(b []byte) {
		actual = append(actual, string(b))
	})
	if err != nil {
		t.Fatalf("Unable to ReadArray from stdout: %s", err)
	}

	if len(expected) != len(actual) {
		t.Error("Unexpected output in ReadArray:")
		t.Logf("  Expected records: %d", len(expected))
		t.Logf("  Actual records:   %d", len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Unexpected output in ReadArray index: %d", i)
			t.Logf("  Expected: %s", expected[i])
			t.Logf("  Actual:   %s", actual[i])
		}
	}
}

// ReadMapExpected is used to list the expected output from stdio.ReadMap for the MapTest() test
type ReadMapExpected struct {
	Key   string
	Value string
	Last  bool
}

func (m ReadMapExpected) String() string {
	esc := func(s string) string {
		return strings.Replace(s, "\n", `\n`, -1)
	}
	return fmt.Sprintf("`%s`: `%s` (%t)", esc(m.Key), esc(m.Value), m.Last)
}

// ReadMapTest is an easy template for testing stdio.ReadMap methods in murex types
func ReadMapTest(t *testing.T, dataType string, input []byte, expected []ReadMapExpected, config *config.Config) {
	stdout := streams.NewStdin()
	stdout.SetDataType(dataType)
	_, err := stdout.Write(input)

	if err != nil {
		t.Fatalf("Unable to Write to stdout: %s", err)
	}

	actual := make([]ReadMapExpected, 0)
	err = stdout.ReadMap(config, func(key, value string, last bool) {
		actual = append(actual, ReadMapExpected{key, value, last})
	})
	if err != nil {
		t.Fatalf("Unable to ReadMap from stdout: %s", err)
	}

	if len(expected) != len(actual) {
		t.Error("Unexpected output in ReadMap:")
		t.Logf("  Expected records: %d", len(expected))
		t.Logf("  Actual records:   %d", len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Unexpected output in ReadMap index: %d", i)
			t.Logf("  Expected: %s", expected[i])
			t.Logf("  Actual:   %s", actual[i])
		}
	}
}

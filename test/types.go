package test

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/types"
)

// ArrayWriterTest is an easy template for testing stdio.ArrayWriter methods in murex types
func ArrayWriterTest(t *testing.T, dataType, expected string) {
	stdout := streams.NewStdin()
	//stdout.SetDataType(types.Null)

	w, err := stdout.WriteArray(types.Null)
	if err != nil {
		t.Fatalf("Unable to create ArrayWriter: %s", err)
	}

	err = w.WriteString("foo")
	if err != nil {
		t.Fatalf("Unable to write foo: %s", err)
	}

	err = w.WriteString("bar")
	if err != nil {
		t.Fatalf("Unable to write bar: %s", err)
	}

	err = w.Close()
	if err != nil {
		t.Fatalf("Unable to close ArrayWriter: %s", err)
	}

	b, err := stdout.ReadAll()
	if err != nil {
		t.Fatalf("Unable to ReadAll from stdout: %s", err)
	}

	if len(b) != 0 {
		t.Error("Unexpected output in ArrayWriter:")
		t.Logf("  Expected: %s", strings.Replace(expected, "\n", `\n`, -1))
		t.Logf("  Actual:   %s", strings.Replace(string(b), "\n", `\n`, -1))
	}
}

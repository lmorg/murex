package textmanip

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/generic"
	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestPrettyDefault(t *testing.T) {
	input := `{"foo":"bar"}`
	output := "{\n    \"foo\": \"bar\"\n}"

	test.RunMethodTest(t, cmdPretty, "pretty", input, types.Json, []string{}, output, nil)
	test.RunMethodTest(t, cmdPretty, "pretty", input, types.Generic, []string{}, output, nil)
}

func TestPrettyStrict(t *testing.T) {
	input := `{"foo":"bar"}`
	output := "{\n    \"foo\": \"bar\"\n}"

	test.RunMethodTest(t, cmdPretty, "pretty", input, types.Json, []string{"--strict"}, output, nil)
	test.RunMethodTest(t, cmdPretty, "pretty", input, types.Generic, []string{"--strict"}, input, nil)
}

package lists

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestJsplit(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		`hello world`,
		types.String,
		[]string{" "},
		`["hello","world"]`,
		nil,
	)
}

func TestJsplitCrLF1(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		"hello world\n",
		types.String,
		[]string{" "},
		`["hello","world"]`,
		nil,
	)
}

func TestJsplitCrLF2(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		"hello world\r\n",
		types.String,
		[]string{" "},
		`["hello","world"]`,
		nil,
	)
}

func TestJsplitCrLF3(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		"hello world\n\n",
		types.String,
		[]string{" "},
		`["hello","world"]`,
		nil,
	)
}

func TestJsplitCrLF4(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		"hello world\r\n\r\n",
		types.String,
		[]string{" "},
		`["hello","world"]`,
		nil,
	)
}

func TestJsplitCrLF5(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		"hello\nworld\n",
		types.String,
		[]string{"\n"},
		`["hello","world"]`,
		nil,
	)
}

func TestJsplitCrLF6(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		"hello\r\nworld\r\n",
		types.String,
		[]string{"\n"},
		`["hello","world"]`,
		nil,
	)
}

func TestJsplitCrLF7(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		"hello\n\nworld\n\n",
		types.String,
		[]string{"\n"},
		`["hello","","world"]`,
		nil,
	)
}

func TestJsplitCrLF8(t *testing.T) {
	test.RunMethodTest(t,
		cmdJsplit, "jsplit",
		"hello\r\n\r\nworld\r\n\r\n",
		types.String,
		[]string{"\n"},
		`["hello","","world"]`,
		nil,
	)
}

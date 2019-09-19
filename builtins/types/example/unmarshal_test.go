package example

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func TestUnmarshal(t *testing.T) {
	count.Tests(t, 1)

	lang.InitEnv()
	fork := lang.ShellProcess.Fork(lang.F_CREATE_STDIN)

	s := `{"Foo":"Bar","Bar":"Foo"}`
	fork.Stdin.Write([]byte(s))

	v, err := unmarshal(fork.Process)
	if err != nil {
		t.Error(err)
	}

	if v.(map[string]interface{})["Foo"] != "Bar" ||
		v.(map[string]interface{})["Bar"] != "Foo" {
		t.Error("JSON unmarshal failed")
	}
}

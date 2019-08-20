package example

import (
	"testing"

	"github.com/lmorg/murex/lang"
)

func TestMarshal(t *testing.T) {
	lang.InitEnv()
	fork := lang.ShellProcess.Fork(lang.F_NO_STDOUT)

	v := map[string]string{
		"Foo": "Bar",
		"Bar": "Foo",
	}

	b, err := marshal(fork.Process, v)
	if err != nil {
		t.Error(err)
	}

	if string(b) != `{"Bar":"Foo","Foo":"Bar"}` {
		t.Error("JSON marshal failed")
	}
}

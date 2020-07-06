package management

import (
	"strings"
	"testing"

	_ "github.com/lmorg/murex/builtins/core/index"
	_ "github.com/lmorg/murex/builtins/core/io"
	_ "github.com/lmorg/murex/builtins/core/runtime"
	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test/count"
)

func TestSourceMethod(t *testing.T) {
	block := `tout block { out "Hello, world!" } -> source`

	lang.InitEnv()

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	testSourceFileRef(t, fork, block)
}

func TestSourceParameter(t *testing.T) {
	block := `source { out "Hello, world!" }`

	lang.InitEnv()

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	testSourceFileRef(t, fork, block)
}

func TestSourceFile(t *testing.T) {
	block := `source source_test.mx`

	lang.InitEnv()

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	testSourceFileRef(t, fork, block)
}

func testSourceFileRef(t *testing.T, fork *lang.Fork, block string) {
	count.Tests(t, 1)

	fork.Name = "(fork)"

	exitNum, err := fork.Execute([]rune(block))
	if exitNum != 0 {
		t.Errorf("fork.Execute().exitNum != 0: %d", exitNum)
	}
	if err != nil {
		t.Errorf("fork.Execute().err != nil: %s", err.Error())
	}

	b, err := fork.Stderr.ReadAll()
	if err != nil {
		t.Errorf("fork.Stderr.ReadAll().err != nil: %s", err.Error())
	}
	if len(b) != 0 {
		t.Errorf("fork.Stderr.ReadAll().b != ``: %s", string(b))
	}

	b, err = fork.Stdout.ReadAll()
	if err != nil {
		t.Errorf("fork.Stdout.ReadAll().err != nil: %s", err.Error())
	}
	if string(b) != "Hello, world!\n" {
		t.Errorf(`fork.Stdout.ReadAll().b != "Hello, world!\n": %s`, string(b))
	}

	v := ref.History.Dump()

	for _, d := range v {
		//if d.DateTime == nil{
		//	t.Errorf("Source DateTime == nil")
		//}
		if d.Filename == "" {
			t.Errorf("Source Filename == ``")
		}
		if d.Module == "" {
			t.Errorf("Source Module == ``")
		}
		if !strings.HasPrefix(d.Module, "source/") {
			t.Errorf("Source Module != `source/...`")
		}
		if d.Source == "" {
			t.Errorf("Source Source == ``")
		}
		if !strings.Contains(d.Source, `out "Hello, world!"`) {
			t.Errorf(`Source Source != out "Hello, world!"`)
		}
	}
}

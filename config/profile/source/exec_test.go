package source_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config/profile/source"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
)

func TestExecNilFileRef(t *testing.T) {
	defer func() {
		recover()
	}()

	lang.InitEnv()
	source.Exec([]rune("out success"), nil, false)
	t.Error("panic is expected")
}

func TestExecSuccess(t *testing.T) {
	lang.InitEnv()
	source.Exec([]rune("out success"), &ref.Source{}, false)
}

func TestExecError(t *testing.T) {
	lang.InitEnv()
	source.Exec([]rune("false"), &ref.Source{}, false)
}

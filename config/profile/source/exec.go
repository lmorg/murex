package source

import (
	"os"

	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/consts"
)

func Exec(source []rune, sourceRef *ref.Source, exitOnError bool) {
	if sourceRef == nil {
		panic("sourceRef is not defined")
	}

	if debug.Enabled {
		debug.Logf("Loading profile `%s`", sourceRef.Module)
	}

	var stdin int
	if os.Getenv(consts.EnvMethod) != consts.EnvTrue {
		stdin = lang.F_NO_STDIN
	}
	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | stdin)
	fork.Stdout = new(term.Out)
	fork.Stderr = term.NewErr(ansi.IsAllowed())
	fork.FileRef.Source = sourceRef
	fork.RunMode = lang.ShellProcess.RunMode
	exitNum, err := fork.Execute(source)

	if err != nil {
		if exitNum == 0 {
			exitNum = 1
		}
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		lang.Exit(exitNum)
	}

	if exitNum != 0 && exitOnError {
		lang.Exit(exitNum)
	}
}

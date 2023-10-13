package preview

import (
	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils/readline"
)

func CommandLine(block []rune, _ string, _ bool, size *readline.PreviewSizeT) ([]string, int, error) {
	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_BACKGROUND | lang.F_PREVIEW)
	fork.FileRef = ref.NewModule(app.ShellModule)
	fork.Stderr = fork.Stdout
	_, err := fork.Execute(block)
	if err != nil {
		return parse([]byte(err.Error()), size)
	}

	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return parse([]byte(err.Error()), size)
	}

	return parse(b, size)
}

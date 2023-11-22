package shell

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/readline"
)

func DynamicPreview(previewBlock string, exe string, params []string) readline.PreviewFuncT {
	block := []rune(previewBlock)
	return func(ctx context.Context, line []rune, item string, _ bool, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT) {
		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
		fork.Name.Set(exe)
		fork.Parameters.DefineParsed(params)
		fork.FileRef = autocomplete.ExesFlagsFileRef[exe]

		_, err := fork.Execute(block)
		if err != nil {
			s, _, err := previewError(err, size)
			callback(s, 0, err)
			return
		}

		b, err := fork.Stdout.ReadAll()
		if err != nil {
			s, _, err := previewError(err, size)
			callback(s, 0, err)
			return
		}

		s, _, err := previewParse(b, size)
		if err != nil {
			s, _, err = previewError(err, size)
			callback(s, 0, err)
			return
		}

		i := previewPos(s, item)

		callback(s, i, err)
	}
}

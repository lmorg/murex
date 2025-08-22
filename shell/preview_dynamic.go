package shell

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/cache"
	"github.com/lmorg/readline/v4"
)

func DynamicPreview(previewBlock string, exe string, params []string) readline.PreviewFuncT {
	block := []rune(previewBlock)

	return func(ctx context.Context, line []rune, item string, _ bool, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT) {
		hash := cache.CreateHash(string(line), block)

		var b []byte

		if !cache.Read(cache.PREVIEW_DYNAMIC, hash, &b) {

			fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
			fork.Name.Set(exe)
			fork.Parameters.DefineParsed(params)
			fork.FileRef = autocomplete.ExesFlagsFileRef[exe]

			_, err := fork.Execute(block)
			if err != nil {
				s, _, err := previewError(err, size)
				callback(s, 0, err)
				return
			}

			b, err = fork.Stdout.ReadAll()
			if err != nil {
				s, _, err := previewError(err, size)
				callback(s, 0, err)
				return
			}

			e, err := fork.Stderr.ReadAll()
			if err != nil {
				s, _, err := previewError(err, size)
				callback(s, 0, err)
				return
			}
			b = append(b, e...)

			ttl := autocomplete.ExesFlags[exe][0].CacheTTL
			cache.Write(cache.PREVIEW_DYNAMIC, hash, &b, cache.Seconds(ttl))
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

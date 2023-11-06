package shell

import (
	"context"

	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/readline"
)

func PreviewCommand(ctx context.Context, _ []rune, command string, _ bool, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT) {
	if lang.GlobalAliases.Exists(command) {
		alias := lang.GlobalAliases.Get(command)
		if len(alias) == 0 {
			return //nil, 0, nil
		}
		if alias[0] != command {
			//return previewCommand(ctx, nil, alias[0], false, size)
			PreviewCommand(ctx, nil, alias[0], false, size, callback)
			return
		}
	}

	if lang.MxFunctions.Exists(command) {
		r, err := lang.MxFunctions.Block(command)
		if err != nil {
			return //nil, 0, err
		}
		//return previewParse([]byte(string(r)), size)
		callback(previewParse([]byte(string(r)), size))
		return
	}

	syn := docs.Synonym[command]
	b := docs.Definition(syn)
	if len(b) != 0 {
		//return previewParse(b, size)
		callback(previewParse(b, size))
		return
	}

	if !(*autocomplete.GlobalExes.Get())[command] {
		//return []string{"not a valid command"}, 0, nil
		callback([]string{"not a valid command"}, 0, nil)
		return
	}

	b = manPage(command, size)
	var (
		lines []string
		err   error
	)

	if len(b) > 0 {
		//return previewParse(b, size)
		lines, _, err = previewParse(b, size)

		callback(lines, 0, err)
		lines = previewHr(lines, size)
	}

	//callEvents("preview")

	block := []rune(`
		config set http user-agent curl/1.0
		config set http timeout 2
		trypipe {
			get https://cheat.sh/$(COMMAND)?T -> [ Body ]
		}`,
	)
	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.Name.Set("(f1)")
	err = fork.Variables.Set(fork.Process, "COMMAND", command, types.String)
	if err != nil {
		//return previewError(err, size)
		s, _, err := previewError(err, size)
		callback(append(lines, s...), 0, err)
		return
	}
	_, err = fork.Execute(block)
	if err != nil {
		//return previewError(err, size)
		s, _, err := previewError(err, size)
		callback(append(lines, s...), 0, err)
		return
	}
	b, err = fork.Stdout.ReadAll()
	if err != nil {
		//return previewError(err, size)
		s, _, err := previewError(err, size)
		callback(append(lines, s...), 0, err)
		return
	}

	//return previewParse(b, size)
	s, _, err := previewParse(b, size)
	callback(append(lines, s...), 0, err)
}

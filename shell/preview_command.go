package shell

import (
	"context"

	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/builtins/events/onPreview/previewops"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/readline"
)

func PreviewCommand(ctx context.Context, cmdLine []rune, command string, _ bool, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT) {
	if lang.GlobalAliases.Exists(command) {
		alias := lang.GlobalAliases.Get(command)
		if len(alias) == 0 {
			return
		}
		if alias[0] != command {
			PreviewCommand(ctx, cmdLine, alias[0], false, size, callback)
			return
		}
	}

	if lang.MxFunctions.Exists(command) {
		r, err := lang.MxFunctions.Block(command)
		if err != nil {
			return
		}
		lines, _, err := previewParse([]byte(string(r)), size)
		callback(lines, 0, err)
		callEventsPreview(ctx, previewops.Function, command, cmdLine, lines, size, callback)
		return
	}

	if lang.GoFunctions[command] != nil {
		syn := docs.Synonym[command]
		b := docs.Definition(syn)
		if len(b) != 0 {
			lines, _, err := previewParse(b, size)
			callback(lines, 0, err)
			callEventsPreview(ctx, previewops.Builtin, command, cmdLine, lines, size, callback)
			return
		}
	}

	/*if !(*autocomplete.GlobalExes.Get())[command] {
		callback([]string{"not a valid command"}, 0, nil)
		return
	}

	b := manPage(command, size)
	var (
		lines []string
		err   error
	)

	if len(b) > 0 {
		lines, _, err = previewParse(b, size)
		callback(lines, 0, err)
	}*/

	callEventsPreview(ctx, previewops.Exec, command, cmdLine, []string{}, size, callback)
}

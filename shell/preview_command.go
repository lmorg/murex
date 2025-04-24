package shell

import (
	"context"

	"github.com/lmorg/murex/builtins/events/onPreview/previewops"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/readline/v4"
)

func PreviewCommand(ctx context.Context, _cmdLine []rune, command string, _ bool, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT) {
	if command == "" {
		callback(previewParse([]byte("Nothing to preview"), size))
		return
	}

	cmdLine := utils.CrLfTrimRune(_cmdLine)

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

	switch {
	case lang.MxFunctions.Exists(command):
		// Murex function
		callEventsPreview(ctx, previewops.Function, command, cmdLine, []string{}, size, callback)

	case lang.GoFunctions[command] != nil:
		// murex builtin
		callEventsPreview(ctx, previewops.Builtin, command, cmdLine, []string{}, size, callback)

	default:
		// external executable
		callEventsPreview(ctx, previewops.Exec, command, cmdLine, []string{}, size, callback)
	}
}

package shell

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/murex/utils/readline"
)

func PreviewParameter(ctx context.Context, block []rune, parameter string, incImages bool, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT) {
	parameter = variables.ExpandString(parameter)
	if utils.Exists(parameter) {
		PreviewFile(ctx, nil, parameter, incImages, size, callback)
		return
	}

	pt, _ := parser.Parse(block, 0)

	parameterCallback := func(lines []string, pos int, err error) {
		if err != nil {
			callback(lines, 0, err)
			return
		}

		if parameter == "" || lang.GoFunctions[pt.FuncName] != nil {
			callback(lines, 0, err)
			return
		}

		i := previewPos(lines, parameter)

		callback(lines, i, nil)
	}

	PreviewCommand(ctx, block, pt.FuncName, false, size, parameterCallback)
}

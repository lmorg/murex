package shell

import (
	"context"
	"strings"

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

		for i := range lines {
			switch {
			case strings.HasPrefix(parameter, "--"):
				switch {
				case strings.Contains(lines[i], ", "+parameter):
					callback(lines, i, nil)
					return
				case strings.Contains(lines[i], "  "+parameter):
					callback(lines, i, nil)
					return
				default:
					continue
				}
			default:
				if strings.Contains(lines[i], "  "+parameter) {
					callback(lines, i, nil)
					return
				}
			}
		}

		callback(lines, 0, nil)
	}

	PreviewCommand(ctx, nil, pt.FuncName, false, size, parameterCallback)
}

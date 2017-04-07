package main

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	_ "github.com/lmorg/murex/lang/builtins"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"os"
	"runtime/trace"
)

func main() {
	var (
		exitNum int
		err     error
	)

	readFlags()

	if fTrace != "" {
		f, err := os.Create(fTrace)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		err = trace.Start(f)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()
	}

	switch {
	case fCommand != "":
		exitNum, err = lang.ProcessNewBlock(
			[]rune(fCommand),
			nil,
			nil,
			nil,
			types.Null,
		)

		if err != nil {
			if exitNum == 0 {
				exitNum = 1
			}
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
			os.Exit(exitNum)
		}

		if exitNum != 0 {
			os.Exit(exitNum)
		}

	case fStdin:
		os.Stderr.WriteString("Not implimented yet.\n")

	default:
		shell.Start()
	}

	debug.Log("[FIN]")
}

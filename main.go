package main

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	_ "github.com/lmorg/murex/lang/builtins"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"os"
	"runtime/trace"
)

func main() {
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
		lang.ProcessNewBlock(
			[]rune(fCommand),
			nil,
			nil,
			nil,
			types.Null,
		)

	case fStdin:
		os.Stderr.WriteString("Not implimented yet.\n")

	default:
		shell.Start()
	}

	debug.Log("[FIN]")
}

package lang

import (
	"errors"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

var ShellExitNum int // for when running murex in interactive shell mode

func ProcessNewBlock(block []rune, stdin, stdout, stderr streams.Io, gpName string) (exitNum int, err error) {
	grandParent := new(proc.Process)
	grandParent.Name = gpName
	grandParent.Parent = nil
	if gpName == "shell" {
		grandParent.ExitNum = ShellExitNum
	}

	if stdin != nil {
		grandParent.Stdin = stdin
	} else {
		grandParent.Stdin = streams.NewStdin()
		grandParent.Stdin.SetDataType(types.Null)
		grandParent.Stdin.Close()
	}

	if stdout != nil {
		grandParent.Stdout = stdout
	} else {
		grandParent.Stdout = new(streams.TermOut)
	}
	grandParent.Stdout.MakeParent()

	if stderr != nil {
		grandParent.Stderr = stderr
	} else {
		grandParent.Stderr = new(streams.TermErr)
	}
	grandParent.Stderr.MakeParent()

	tree, pErr := ParseBlock(block)
	if pErr.Code != 0 {
		grandParent.Stderr.Writeln([]byte(pErr.Message))
		debug.Json("ParseBlock returned:", pErr)
		err = errors.New(pErr.Message)
		return 1, err
	}

	compile(&tree, grandParent)

	// Support for different run modes:
	switch {
	case grandParent.Name == "try":
		exitNum = runHyperSensitive(&tree)
	default:
		exitNum = runNormal(&tree)
		//exitNum = runHyperSensitive(&tree)
	}

	// This will just unlock the parent lock. Stdxxx.Close() will still have to be called.
	grandParent.Stdout.UnmakeParent()
	grandParent.Stderr.UnmakeParent()

	debug.Json("Finished running &tree", tree)
	return
}

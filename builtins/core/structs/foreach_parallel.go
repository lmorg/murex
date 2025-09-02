package structs

import (
    "sync"
    "sync/atomic"

    "github.com/lmorg/murex/lang"
    "github.com/lmorg/murex/lang/expressions/functions"
)

const MAX_INT = int(^uint(0) >> 1)

func cmdForEachParallel(p *lang.Process, flags map[string]string, additional []string) error {
    block, varName, err := forEachInitializer(p, additional)
    if err != nil {
        return err
    }

    // Pre-parse the foreach block once
    var tree *[]functions.FunctionT
    if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
        trimmed := block[1 : len(block)-1]
        t, err := lang.ParseBlock(trimmed)
        if err != nil {
            return err
        }
        tree = t
    } else {
        t, err := lang.ParseBlock(block)
        if err != nil {
            return err
        }
        tree = t
    }

	parallel, err := getFlagValueInt(flags, foreachParallel)
	if err != nil {
		return err
	}

	if parallel < 1 {
		parallel = MAX_INT
	}

	var (
		iteration = int64(-1)
		wg        = new(sync.WaitGroup)
		wait      = make(chan struct{}, parallel)
	)

    err = p.Stdin.ReadArrayWithType(p.Context, func(varValue any, dataType string) {
        i := atomic.AddInt64(&iteration, 1)
        wait <- struct{}{}
        wg.Add(1)
        go func() {
            forEachParallelInnerLoopPreparsed(p, tree, varName, varValue, dataType, int(i))
            wg.Done()
            <-wait
        }()
    })

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func forEachParallelInnerLoop(p *lang.Process, block []rune, varName string, varValue any, dataType string, iteration int) {
	var b []byte
	b, err := convertToByte(varValue)
	if err != nil {
		p.Done()
		return
	}

	if len(b) == 0 || p.HasCancelled() {
		return
	}

	fork := p.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_CREATE_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	fork.Name.Set("foreach--parallel")
	fork.FileRef = p.FileRef

	if varName != "!" {
		err = fork.Variables.Set(fork.Process, varName, varValue, dataType)
		if err != nil {
			p.Stderr.Writeln([]byte("error: " + err.Error()))
			p.Done()
			return
		}
	}

	if !setMetaValues(fork.Process, iteration) {
		return
	}

	fork.Stdin.SetDataType(dataType)
	_, err = fork.Stdin.Writeln(b)
	if err != nil {
		p.Stderr.Writeln([]byte("error: " + err.Error()))
		p.Done()
		return
	}
	_, err = fork.Execute(block)
	if err != nil {
		p.Stderr.Writeln([]byte("error: " + err.Error()))
		p.Done()
		return
	}
	// Aggregate child output to parent to avoid shared-stream contention during execution
	if out, rerr := fork.Stdout.ReadAll(); rerr == nil && len(out) > 0 {
		_, _ = p.Stdout.Write(out)
	}
	if errb, rerr := fork.Stderr.ReadAll(); rerr == nil && len(errb) > 0 {
		_, _ = p.Stderr.Write(errb)
	}
}

// forEachParallelInnerLoopPreparsed uses a pre-parsed tree for each worker iteration.
func forEachParallelInnerLoopPreparsed(p *lang.Process, tree *[]functions.FunctionT, varName string, varValue any, dataType string, iteration int) {
    var b []byte
    b, err := convertToByte(varValue)
    if err != nil {
        p.Done()
        return
    }

    if len(b) == 0 || p.HasCancelled() {
        return
    }

	fork := p.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_CREATE_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
    fork.Name.Set("foreach--parallel")
    fork.FileRef = p.FileRef

    if varName != "!" {
        err = fork.Variables.Set(fork.Process, varName, varValue, dataType)
        if err != nil {
            p.Stderr.Writeln([]byte("error: " + err.Error()))
            p.Done()
            return
        }
    }

    if !setMetaValues(fork.Process, iteration) {
        return
    }

    fork.Stdin.SetDataType(dataType)
    _, err = fork.Stdin.Writeln(b)
    if err != nil {
        p.Stderr.Writeln([]byte("error: " + err.Error()))
        p.Done()
        return
    }
	_, err = fork.ExecuteTree(tree)
	if err != nil {
		p.Stderr.Writeln([]byte("error: " + err.Error()))
		p.Done()
		return
	}
	// Aggregate child output to parent to avoid shared-stream contention during execution
	if out, rerr := fork.Stdout.ReadAll(); rerr == nil && len(out) > 0 {
		_, _ = p.Stdout.Write(out)
	}
	if errb, rerr := fork.Stderr.ReadAll(); rerr == nil && len(errb) > 0 {
		_, _ = p.Stderr.Write(errb)
	}
}

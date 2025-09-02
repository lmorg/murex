package structs

import (
    "runtime"
    "sync"
    "sync/atomic"

    "github.com/lmorg/murex/lang"
    "github.com/lmorg/murex/lang/expressions/functions"
    "github.com/lmorg/murex/lang/types"
)

// No explicit hard cap: user-set --parallel is honored.
// When --parallel <= 0, default to a multiple of CPUs for safety.

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
        parallel = runtime.NumCPU() * 8
        if parallel < 1 { parallel = 1 }
    }

	var (
		iteration = int64(-1)
		wg        = new(sync.WaitGroup)
		wait      = make(chan struct{}, parallel)
	)

    // Ordering: default ordered unless explicitly disabled
    ordered := true
    if flags[foreachUnordered] == types.TrueString {
        ordered = false
    }
    if flags[foreachOrdered] == types.TrueString {
        ordered = true
    }
    // result channel for aggregator
    type result struct{ idx int; out, err []byte }
    resCh := make(chan result, parallel*2)
    aggDone := make(chan struct{})

    go func() {
        if ordered {
            pending := make(map[int]result)
            next := 0
            for r := range resCh {
                pending[r.idx] = r
                for {
                    if v, ok := pending[next]; ok {
                        if len(v.out) > 0 { _, _ = p.Stdout.Write(v.out) }
                        if len(v.err) > 0 { _, _ = p.Stderr.Write(v.err) }
                        delete(pending, next)
                        next++
                    } else {
                        break
                    }
                }
            }
        } else {
            for r := range resCh {
                if len(r.out) > 0 { _, _ = p.Stdout.Write(r.out) }
                if len(r.err) > 0 { _, _ = p.Stderr.Write(r.err) }
            }
        }
        close(aggDone)
    }()

    err = p.Stdin.ReadArrayWithType(p.Context, func(varValue any, dataType string) {
        i := atomic.AddInt64(&iteration, 1)
        wait <- struct{}{}
        wg.Add(1)
        go func() {
            // run worker
            forkOut, forkErr := forEachParallelWorkerPreparsed(p, tree, varName, varValue, dataType, int(i))
            // send to aggregator
            resCh <- result{idx: int(i), out: forkOut, err: forkErr}
            wg.Done()
            <-wait
        }()
    })

	if err != nil {
		return err
	}

    wg.Wait()
    close(resCh)
    <-aggDone
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
    if out, rerr := fork.Stdout.ReadAll(); rerr == nil && len(out) > 0 {
        _, _ = p.Stdout.Write(out)
    }
    if errb, rerr := fork.Stderr.ReadAll(); rerr == nil && len(errb) > 0 {
        _, _ = p.Stderr.Write(errb)
    }
}

// forEachParallelWorkerPreparsed runs a single iteration and returns captured stdout/stderr for aggregation.
func forEachParallelWorkerPreparsed(p *lang.Process, tree *[]functions.FunctionT, varName string, varValue any, dataType string, iteration int) (stdout, stderr []byte) {
    var b []byte
    b, err := convertToByte(varValue)
    if err != nil || len(b) == 0 || p.HasCancelled() {
        return nil, nil
    }
    fork := p.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_CREATE_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
    fork.Name.Set("foreach--parallel")
    fork.FileRef = p.FileRef
    if varName != "!" {
        if e := fork.Variables.Set(fork.Process, varName, varValue, dataType); e != nil { return nil, []byte("error: "+e.Error()) }
    }
    if !setMetaValues(fork.Process, iteration) { return nil, nil }
    fork.Stdin.SetDataType(dataType)
    if _, e := fork.Stdin.Writeln(b); e != nil { return nil, []byte("error: "+e.Error()) }
    if _, e := fork.ExecuteTree(tree); e != nil { return nil, []byte("error: "+e.Error()) }
    out, _ := fork.Stdout.ReadAll()
    errb, _ := fork.Stderr.ReadAll()
    return out, errb
}

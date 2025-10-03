package dag

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/alter"
)

func init() {
	lang.DefineMethod("dag", cmdDag, types.Unmarshal, types.Marshal)
}

const usage = `
Usage: dag [ --datatype | --append ] { pipeline } { pipeline } ...`

const (
	fDataType = "--datatype"
	fAppend   = "--append"
)

var args = &parameters.Arguments{
	AllowAdditional:    true,
	IgnoreInvalidFlags: false,
	Flags: map[string]string{
		fDataType: types.String,
		"-t":      fDataType,
		fAppend:   types.Boolean,
		"-a":      fAppend,
	},
}

func cmdDag(p *lang.Process) error {
	flags, blocks, err := p.Parameters.ParseFlags(args)
	if err != nil {
		return err
	}

	dt := flags[fDataType]
	if dt == "" {
		if p.IsMethod {
			dt = p.Stdin.GetDataType()
		} else {
			dt = types.String
		}
	}
	if dt == types.Generic {
		dt = types.String
	}
	p.Stdout.SetDataType(dt)

	if len(blocks) == 0 {
		return fmt.Errorf("missing graphs. %s", usage)
	}

	for i := range blocks {
		if !types.IsBlock([]byte(blocks[i])) {
			return fmt.Errorf("parameter not a code block: %s%s", blocks[i], usage)
		}
	}

	var (
		wg      sync.WaitGroup
		fStdin  = lang.F_NO_STDIN
		bStdin  []byte
		stdouts = make([]stdio.Io, len(blocks))
		errs    = make([]error, len(blocks))
		exitNum atomic.Int32
	)

	if p.IsMethod {
		fStdin = lang.F_CREATE_STDIN
		bStdin, err = p.Stdin.ReadAll()
		if err != nil {
			return fmt.Errorf("reading from stdin: %v", err)
		}
	}

	for i := range blocks {
		fork := p.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | fStdin | lang.F_CREATE_STDOUT)

		if p.IsMethod {
			fork.Stdin.SetDataType(dt)
			_, err = fork.Stdin.Write(bStdin)
			if err != nil {
				return fmt.Errorf("writing to node %d's stdin: %v", i+1, err)
			}
		}
		stdouts[i] = fork.Stdout
		wg.Add(1)
		go func() {
			var exit int
			exit, errs[i] = fork.Execute([]rune(blocks[i]))
			exitNum.Add(int32(exit))
			wg.Done()
		}()
	}

	wg.Wait()

	if flags[fAppend] == types.TrueString {
		for i := range blocks {
			if errs[i] != nil {
				return fmt.Errorf("error returned from pipeline %d: %v", i+1, err)
			}

			_, err = io.Copy(p.Stdout, stdouts[i])
			if errs[i] != nil {
				return fmt.Errorf("cannot write pipeline %d to stdout: %v", i+1, err)
			}
		}

	} else {
		var (
			merged any
			b      []byte
		)

		for i := range blocks {
			if errs[i] != nil {
				return fmt.Errorf("error returned from pipeline %d: %v", i+1, err)
			}

			b, err = stdouts[i].ReadAll()
			if err != nil {
				return fmt.Errorf("cannot read pipeline %d output: %v", i+1, err)
			}

			v, err := lang.UnmarshalDataBuffered(p, b, dt)
			if err != nil {
				return fmt.Errorf("cannot convert pipeline %d output to %s: %v", i+1, dt, err)
			}
			merged, err = alter.Merge(p.Context, merged, nil, v)
			if err != nil {
				return fmt.Errorf("cannot merge pipeline %d output: %v", i+1, err)
			}
		}

		b, err = lang.MarshalData(p, dt, merged)
		if err != nil {
			return fmt.Errorf("cannot marshal merged pipelines: %v", err)
		}
		_, err = p.Stdout.Write(b)
		if err != nil {
			return err
		}
	}

	p.ExitNum = int(exitNum.Load())
	return nil
}

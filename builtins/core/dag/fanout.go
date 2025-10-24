package dag

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/alter"
)

const fanout = "fanout"

func init() {
	lang.DefineMethod(fanout, cmdFanout, types.Unmarshal, types.Marshal)
}

var usage = fmt.Sprintf(`
Usage: %s [ %s | %s ] %s   { code block } { code block } ...
       %s [ %s | %s ] %s { { code block } { code block } ... }`,
	fanout, fDataType, fConcatenate, strings.Repeat(" ", len(fParse)),
	fanout, fDataType, fConcatenate, fParse)

const (
	fDataType    = "--datatype"
	fConcatenate = "--concat"
	fParse       = "--parse"
)

var args = &parameters.Arguments{
	AllowAdditional:    true,
	IgnoreInvalidFlags: false,
	Flags: map[string]string{
		fDataType:    types.String,
		"-t":         fDataType,
		fConcatenate: types.Boolean,
		"-c":         fConcatenate,
		fParse:       types.Boolean,
		"-p":         fParse,
	},
}

func cmdFanout(p *lang.Process) error {
	flags, sBlocks, err := p.Parameters.ParseFlags(args)
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

	parse := flags[fParse] == types.TrueString

	switch {
	case len(sBlocks) == 0:
		return fmt.Errorf("missing vertices in parameters.\n%s", usage)
	case parse && len(sBlocks) > 1:
		return fmt.Errorf("multiple parameters supplied with %s flag. Vertices should be included inside one block.\n%s", fParse, usage)
	}

	rBlocks := make([][]rune, len(sBlocks))
	for i := range sBlocks {
		rBlocks[i] = []rune(sBlocks[i])
	}

	if parse {
		if !types.IsBlockRune(rBlocks[0]) {
			return fmt.Errorf("parameter should be a block\n%s", usage)
		}

		tree := expressions.NewParser(p, types.BlockStripCurlyBrace(rBlocks[0]), 0)
		err := tree.ParseStatement(true, expressions.WithAutoEscapeLineFeed(), expressions.WithCommand(fanout))
		if err != nil {
			return fmt.Errorf("error parsing block for vertices: %v\n%s", err, usage)
		}

		rBlocks = tree.StatementParametersUnsafe()
		if len(rBlocks) == 0 {
			return fmt.Errorf("missing vertices in block.\n%s", usage)
		}
	}

	for i := range rBlocks {
		if !types.IsBlockRune(rBlocks[i]) {
			return fmt.Errorf("vertex %d is not a code block:\n%s\n%s", i+1, string(rBlocks[i]), usage)
		}
	}

	var (
		wg      sync.WaitGroup
		fStdin  = lang.F_NO_STDIN
		bStdin  []byte
		stdouts = make([]stdio.Io, len(rBlocks))
		errs    = make([]error, len(rBlocks))
		exitNum atomic.Int32
	)

	if p.IsMethod {
		fStdin = lang.F_CREATE_STDIN
		bStdin, err = p.Stdin.ReadAll()
		if err != nil {
			return fmt.Errorf("error reading from stdin: %v", err)
		}
	}

	for i := range rBlocks {
		fork := p.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | fStdin | lang.F_CREATE_STDOUT)

		if p.IsMethod {
			fork.Stdin.SetDataType(dt)
			_, err = fork.Stdin.Write(bStdin)
			if err != nil {
				return fmt.Errorf("error writing to vertex %d's stdin:\n%v", i+1, err)
			}
		}
		stdouts[i] = fork.Stdout
		wg.Add(1)
		go func() {
			var exit int
			exit, errs[i] = fork.Execute(rBlocks[i])
			exitNum.Add(int32(exit))
			wg.Done()
		}()
	}

	wg.Wait()

	if flags[fConcatenate] == types.TrueString {
		for i := range rBlocks {
			if errs[i] != nil {
				return fmt.Errorf("error returned from vertex %d:\n%v", i+1, err)
			}

			_, err = io.Copy(p.Stdout, stdouts[i])
			if errs[i] != nil {
				return fmt.Errorf("cannot write vertex %d to stdout:\n%v", i+1, err)
			}
		}

	} else {
		var (
			merged any
			b      []byte
		)

		for i := range rBlocks {
			if errs[i] != nil {
				return fmt.Errorf("error returned from vertex %d:\n%v", i+1, err)
			}

			b, err = stdouts[i].ReadAll()
			if err != nil {
				return fmt.Errorf("cannot read vertex %d output:\n%v", i+1, err)
			}

			v, err := lang.UnmarshalDataBuffered(p, b, dt)
			if err != nil {
				return fmt.Errorf("cannot convert vertex %d output to %s:\n%v", i+1, dt, err)
			}
			merged, err = alter.Merge(p.Context, merged, nil, v)
			if err != nil {
				return fmt.Errorf("cannot merge vertex %d output:\n%v", i+1, err)
			}
		}

		b, err = lang.MarshalData(p, dt, merged)
		if err != nil {
			return fmt.Errorf("cannot marshal merged vertices:\n%v", err)
		}
		_, err = p.Stdout.Write(b)
		if err != nil {
			return err
		}
	}

	p.ExitNum = int(exitNum.Load())
	return nil
}

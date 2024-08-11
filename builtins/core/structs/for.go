package structs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("for", cmdFor, types.Generic)
}

const (
	_VARIABLE    = 0
	_CONDITIONAL = 1
	_INCREMENTAL = 2
)

const _FOR_WARNING = "The syntax for `for` has changed and the old syntax is no longer valid.\nFor more information on it's usage, either:\n* Visit https://murex.rocks/commands/for.html in your preferred browser\n* or run `murex-docs for` from the command line "

// Example usage:
// for { $i=1; $i<6; $i=$i+1 } { echo $i }
func cmdFor(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	cblock, err := p.Parameters.Block(0)
	if err != nil {
		lang.FeatureWarning(_FOR_WARNING, p.FileRef)
		return fmt.Errorf(_FOR_WARNING)
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	parameters := strings.Split(string(cblock), ";")
	if len(parameters) != 3 {
		return errors.New("invalid syntax. Must be ( variable; conditional; incremental )")
	}

	_, err = p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_NO_STDOUT).Execute([]rune(parameters[_VARIABLE]))
	if err != nil {
		return err
	}

	rConditional := []rune(parameters[_CONDITIONAL])
	rIncremental := []rune(parameters[_INCREMENTAL])

	for {
		if p.HasCancelled() {
			return errors.New(errCancelled)
		}

		fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		i, err := fork.Execute(rConditional)
		if err != nil {
			return err
		}

		b, err := fork.Stdout.ReadAll()
		if err != nil {
			return err
		}
		if !types.IsTrue(b, i) {
			return nil
		}

		// Execute block
		p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN).Execute(block)

		// Increment counter
		_, err = p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_NO_STDOUT).Execute(rIncremental)
		if err != nil {
			return err
		}
	}

	//return nil
}

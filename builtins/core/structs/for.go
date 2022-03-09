package structs

import (
	"errors"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("for", cmdFor, types.Generic)
}

// Example usage:
// for ( i=1; i<6; i++ ) { echo $i }
func cmdFor(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	cblock, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	parameters := strings.Split(string(cblock), ";")
	if len(parameters) != 3 {
		return errors.New("invalid syntax. Must be ( variable; conditional; incremental )")
	}

	variable := "let " + parameters[0]
	conditional := "= " + parameters[1]
	incremental := "let " + parameters[2]

	_, err = p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_NO_STDOUT).Execute([]rune(variable))
	if err != nil {
		return err
	}

	rConditional := []rune(conditional)
	rIncremental := []rune(incremental)

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

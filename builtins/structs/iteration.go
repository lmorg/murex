package structs

import (
	"errors"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"strings"
)

func init() {
	proc.GoFunctions["for"] = proc.GoFunction{Func: cmdFor, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["foreach"] = proc.GoFunction{Func: cmdForEach, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["formap"] = proc.GoFunction{Func: cmdForMap, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["while"] = proc.GoFunction{Func: cmdWhile, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["!while"] = proc.GoFunction{Func: cmdWhile, TypeIn: types.Null, TypeOut: types.Generic}
}

// Example usage:
// for { i=1; i<6; i++ } { echo $i }
func cmdFor(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	cblock, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	parameters := strings.Split(string(cblock), ";")
	if len(parameters) != 3 {
		return errors.New("Invalid syntax. Must be { variable; conditional; incremental }")
	}

	variable := "let " + parameters[0]
	conditional := "eval " + parameters[1]
	incremental := "let " + parameters[2]

	_, err = lang.ProcessNewBlock([]rune(variable), nil, nil, p.Stderr, types.Null)
	if err != nil {
		return err
	}

	for {
		stdout := streams.NewStdin()
		i, err := lang.ProcessNewBlock([]rune(conditional), nil, stdout, p.Stderr, types.Null)
		stdout.Close()
		if err != nil {
			return err
		}

		if !types.IsTrue(stdout.ReadAll(), i) {
			return nil
		}

		lang.ProcessNewBlock(block, nil, p.Stdout, p.Stderr, types.Null)

		_, err = lang.ProcessNewBlock([]rune(incremental), nil, nil, p.Stderr, types.Null)
		if err != nil {
			return err
		}
	}

	return nil
}

func cmdForEach(p *proc.Process) (err error) {
	//p.Stdout.SetDataType(types.Generic)
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	var (
		block   []rune
		varName string
	)

	switch p.Parameters.Len() {
	case 1:
		block, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

	case 2:
		block, err = p.Parameters.Block(1)
		if err != nil {
			return err
		}

		varName, err = p.Parameters.String(0)
		if err != nil {
			return err
		}

	default:
		return errors.New("Invalid number of parameters.")
	}

	p.Stdin.ReadArray(func(b []byte) {
		if len(b) == 0 {
			return
		}

		if varName != "" {
			proc.GlobalVars.Set(varName, string(b), dt)
		}

		stdin := streams.NewStdin()
		stdin.SetDataType(dt)
		stdin.Writeln(b)
		stdin.Close()

		lang.ProcessNewBlock(block, stdin, p.Stdout, p.Stderr, p.Previous.Name)
	})

	return nil
}

func cmdForMap(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)
	dt := p.Stdin.GetDataType()
	//p.Stdout.SetDataType(dt)

	block, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}

	varKey, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	varVal, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	err = p.Stdin.ReadMap(&proc.GlobalConf, func(key, value string, last bool) {
		proc.GlobalVars.Set(varKey, key, types.String)
		proc.GlobalVars.Set(varVal, value, dt)

		lang.ProcessNewBlock(block, nil, p.Stdout, p.Stderr, p.Previous.Name)
		//_, err := lang.ProcessNewBlock(block, nil, p.Stdout, p.Stderr, p.Previous.Name)
		//if err != nil {
		//	p.Stderr.Writeln([]byte(err.Error()))
		//}

	})

	return err
}

func cmdWhile(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)

	switch p.Parameters.Len() {
	case 1:
		// Condition is taken from the while loop.
		block, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		for {
			stdout := streams.NewStdin()
			i, err := lang.ProcessNewBlock(block, nil, stdout, p.Stderr, types.Null)
			stdout.Close()
			if err != nil {
				return err
			}
			b := stdout.ReadAll()

			_, err = p.Stdout.Write(b)
			if err != nil {
				return err
			}

			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

		}

	case 2:
		// Condition is first parameter, while loop is second.
		ifBlock, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		whileBlock, err := p.Parameters.Block(1)
		if err != nil {
			return err
		}

		for {
			stdout := streams.NewStdin()
			i, err := lang.ProcessNewBlock(ifBlock, nil, stdout, nil, types.Null)
			stdout.Close()
			if err != nil {
				return err
			}
			b := stdout.ReadAll()
			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

			lang.ProcessNewBlock(whileBlock, nil, p.Stdout, p.Stderr, types.Null)
		}

	default:
		// Error
		return errors.New("Invalid number of parameters. Please read usage notes.")
	}

	return errors.New("cmdWhile(p *proc.Process) unexpected escaped a switch with default case.")
}

func cmdR(p *proc.Process) error {
	// @{r: abc[1,2,3] [1..3]
	//r [1..10] -> foreach line {}
	return nil
}

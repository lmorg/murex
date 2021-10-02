package structs

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	//lang.GoFunctions["foreach"] = cmdForEach
	lang.DefineMethod("foreach", cmdForEach, types.ReadArrayWithType, types.Any)
}

func cmdForEach(p *lang.Process) (err error) {
	flag, _ := p.Parameters.String(0)
	switch flag {
	case "--jmap":
		return cmdForEachJmap(p)

	default:
		return cmdForEachDefault(p)
	}
}

func cmdForEachDefault(p *lang.Process) (err error) {
	dt := p.Stdin.GetDataType()
	if dt == types.Json {
		p.Stdout.SetDataType(types.JsonLines)
	} else {
		p.Stdout.SetDataType(dt)
	}

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
		return errors.New("Invalid number of parameters")
	}

	err = p.Stdin.ReadArrayWithType(func(b []byte, dt string) {
		if len(b) == 0 || p.HasCancelled() {
			return
		}

		if varName != "!" {
			p.Variables.Set(p, varName, string(b), dt)
		}

		fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_CREATE_STDIN)
		fork.Stdin.SetDataType(dt)
		fork.Stdin.Writeln(b)
		fork.Execute(block)
	})

	return
}

func cmdForEachJmap(p *lang.Process) error {
	//dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(types.Json)

	varName, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	blockKey, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}

	blockVal, err := p.Parameters.Block(3)
	if err != nil {
		return err
	}

	m := make(map[string]string)

	err = p.Stdin.ReadArrayWithType(func(b []byte, dt string) {
		if len(b) == 0 || p.HasCancelled() {
			return
		}

		if varName != "!" {
			p.Variables.Set(p, varName, string(b), dt)
		}

		forkKey := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		forkKey.Execute(blockKey)
		bKey, err := forkKey.Stdout.ReadAll()
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			p.Kill()
		}

		forkVal := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		forkVal.Execute(blockVal)
		bVal, err := forkVal.Stdout.ReadAll()
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			p.Kill()
		}

		m[string(utils.CrLfTrim(bKey))] = string(utils.CrLfTrim(bVal))
	})

	if err != nil {
		return err
	}

	b, err := json.Marshal(m, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

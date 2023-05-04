package structs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineMethod("formap", cmdForMap, types.ReadMap, types.Any)
}

func cmdForMap(p *lang.Process) (err error) {
	flag, _ := p.Parameters.String(0)
	switch flag {
	case "--jmap":
		return cmdForMapJmap(p)

	default:
		return cmdForMapDefault(p)
	}
}

func cmdForMapDefault(p *lang.Process) error {
	//dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(types.Generic)

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

	var varErr error
	err = p.Stdin.ReadMap(p.Config, func(readmap *stdio.Map) {
		if p.HasCancelled() {
			return
		}

		if varKey != "!" {
			varErr = p.Variables.Set(p, varKey, readmap.Key, types.String)
			if varErr != nil {
				p.Done()
				return
			}
		}

		if varVal != "!" {
			varErr = p.Variables.Set(p, varVal, readmap.Value, readmap.DataType)
			if varErr != nil {
				p.Done()
				return
			}
		}

		fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN)
		fork.Execute(block)
	})

	if varErr != nil {
		return varErr
	}
	return err
}

// Example usage:
// <stdin> -> formap --jmap k v { $k } { out: $v[summary] } -> <stdout>
func cmdForMapJmap(p *lang.Process) error {
	//dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(types.Json)

	blockKey, err := p.Parameters.Block(3)
	if err != nil {
		return err
	}

	blockVal, err := p.Parameters.Block(4)
	if err != nil {
		return err
	}

	varKey, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	varVal, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	m := make(map[string]string)

	err = p.Stdin.ReadMap(p.Config, func(readmap *stdio.Map) {
		if p.HasCancelled() {
			return
		}

		if varKey != "!" {
			p.Variables.Set(p, varKey, readmap.Key, types.String)
		}
		if varVal != "!" {
			p.Variables.Set(p, varVal, readmap.Value, readmap.DataType)
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

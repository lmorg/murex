package structs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func cmdForEachJmap(p *lang.Process) error {
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

	var (
		m         = make(map[string]string)
		iteration int
	)

	err = p.Stdin.ReadArrayWithType(p.Context, func(v any, dt string) {
		var b []byte
		b, err = convertToByte(v)
		if err != nil {
			p.Done()
			return
		}

		if len(b) == 0 || p.HasCancelled() {
			return
		}

		if varName != "!" {
			p.Variables.Set(p, varName, v, dt)
		}

		iteration++
		if !setMetaValues(p, iteration) {
			return
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

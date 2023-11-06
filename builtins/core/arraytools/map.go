package arraytools

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("map", mkmap, types.Json)
}

func mkmap(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	blockKey, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	blockValue, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	var aKeys, aValues []string

	forkKeys := p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
	_, errKeys := forkKeys.Execute(blockKey)

	forkValues := p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
	_, errValues := forkValues.Execute(blockValue)

	if errKeys != nil {
		return errKeys
	}
	if errValues != nil {
		return errValues
	}

	errKeys = forkKeys.Stdout.ReadArray(p.Context, func(b []byte) {
		aKeys = append(aKeys, string(b))
	})

	errValues = forkValues.Stdout.ReadArray(p.Context, func(b []byte) {
		aValues = append(aValues, string(b))
	})

	if errKeys != nil {
		return errKeys
	}
	if errValues != nil {
		return errValues
	}

	if len(aKeys) > len(aValues) {
		return errors.New("there are more keys than values (k > v)")
	}

	if len(aKeys) < len(aValues) {
		return errors.New("there are more values than keys (v > k)")
	}

	m := make(map[string]string)
	for i := range aKeys {
		m[aKeys[i]] = aValues[i]
	}

	b, err := json.Marshal(m, p.Stdout.IsTTY())
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}

package cmdconfig

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func evalConfig(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	app, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	key, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(3)
	if err != nil {
		return err
	}

	dataType := p.Config.DataType(app, key)

	v, err := p.Config.Get(app, key, dataType)
	if err != nil {
		return err
	}

	fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_CREATE_STDIN | lang.F_CREATE_STDOUT)
	fork.Stdin.SetDataType(dataType)

	_, err = fork.Stdin.Write([]byte(v.(string)))
	if err != nil {
		return fmt.Errorf("couldn't write to eval's stdin: %s", err.Error())
	}

	p.ExitNum, err = fork.Execute(block)
	if err != nil {
		return err
	}

	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return err
	}

	return p.Config.Set(app, key, string(b), p.FileRef)
}

package cmdconfig

import (
	"errors"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func defineConfig(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	app, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	key, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	var b []byte

	if p.IsMethod {
		b, err = p.Stdin.ReadAll()
		if err != nil {
			return err
		}

	} else {
		b, err = p.Parameters.Byte(3)
		if err != nil {
			return err
		}
	}

	var properties config.Properties
	err = json.UnmarshalMurex(b, &properties)
	if err != nil {
		return err
	}

	properties.FileRef = p.FileRef

	switch {
	case properties.DataType == "":
		return errors.New("`DataType` not defined")

	case properties.Description == "":
		return errors.New("`Description` not defined")

	case (properties.Dynamic.Read == "" && properties.Dynamic.Write != "") ||
		(properties.Dynamic.Read != "" && properties.Dynamic.Write == ""):
		return errors.New("When using dynamic values, both the `read` and `write` need to contain code blocks")

	case properties.Dynamic.Read != "" && !types.IsBlock([]byte(properties.Dynamic.Read)):
		return errors.New("Dynamic `Read` is not a valid code block")

	case properties.Dynamic.Write != "" && !types.IsBlock([]byte(properties.Dynamic.Write)):
		return errors.New("Dynamic `Write` is not a valid code block")

	case properties.Dynamic.Read != "" && !properties.Global:
		return errors.New("`Global` must be `true` when dynamic values are defined")
	}

	if properties.Dynamic.Read != "" {
		properties.Dynamic.GetDynamic = getDynamic(
			[]rune(properties.Dynamic.Read), p.Parameters.Params, p.FileRef)
		properties.Dynamic.SetDynamic = setDynamic(
			[]rune(properties.Dynamic.Write), p.Parameters.Params, p.FileRef, properties.DataType)
	}

	lang.ShellProcess.Config.Define(app, key, properties)
	return nil
}

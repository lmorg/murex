package datatools

import (
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/objectkeys"
)

func init() {
	lang.DefineMethod("struct-keys", cmdStructKeys, types.Unmarshal, types.Json)
}

func cmdStructKeys(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	flags, additional, err := p.Parameters.ParseFlags(&parameters.Arguments{
		AllowAdditional: true,
		Flags: map[string]string{
			"--depth":     "int",
			"-d":          "--depth",
			"--separator": "str",
			"-s":          "--separator",
		},
	})

	if err != nil {
		return err
	}

	separator := flags.GetValue("--separator").String()
	if separator == "" {
		separator = "/"
	}

	depth := flags.GetValue("--depth").Integer()
	if depth == 0 && len(additional) == 1 {
		depth, _ = strconv.Atoi(additional[0])
	}

	if depth == 0 {
		depth = -1
	}

	dt := p.Stdin.GetDataType()
	v, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(types.Json)

	aw, err := p.Stdout.WriteArray(types.Json)
	if err != nil {
		return err
	}

	err = objectkeys.Recursive(p.Context, "", v, dt, separator, aw.WriteString, depth)
	if err != nil {
		return err
	}
	return aw.Close()
}

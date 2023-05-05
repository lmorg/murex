package toml

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	toml "github.com/pelletier/go-toml"
)

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
	return lang.MapTemplate(typeName, toml.Marshal, toml.Unmarshal, read, callback)
}

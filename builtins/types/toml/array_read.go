package toml

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/pelletier/go-toml"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	return lang.ArrayTemplate(toml.Marshal, toml.Unmarshal, read, callback)
}

func readArrayWithType(read stdio.Io, callback func([]byte, string)) error {
	return lang.ArrayWithTypeTemplate(typeName, toml.Marshal, toml.Unmarshal, read, callback)
}

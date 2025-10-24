package toml

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/pelletier/go-toml"
)

const typeName = "toml"

var errNakedArrays = errors.New("the TOML specification doesn't support naked arrays")

func init() {
	lang.RegisterDataType(typeName, lang.DataTypeIsKeyValue)
	stdio.RegisterReadArray(typeName, readArray)
	stdio.RegisterReadArrayWithType(typeName, readArrayWithType)
	stdio.RegisterReadMap(typeName, readMap)
	stdio.RegisterWriteArray(typeName, func(_ stdio.Io) (stdio.ArrayWriter, error) {
		return nil, errNakedArrays
	})

	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex
	lang.RegisterMarshaller(typeName, marshal)
	lang.RegisterUnmarshaller(typeName, unmarshal)

	lang.SetMime(typeName,
		"application/toml", // this is preferred but we will include others since not everyone follows standards.
		"application/x-toml",
		"text/toml",
		"text/x-toml",
		"+toml",
	)

	lang.SetFileExtensions(typeName, "toml")
}

func readIndex(p *lang.Process, params []string) error {
	var jInterface any

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = toml.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	return lang.IndexTemplateObject(p, params, &jInterface, toml.Marshal)
}

func marshal(_ *lang.Process, v any) ([]byte, error) {
	return toml.Marshal(v)
}

func unmarshal(p *lang.Process) (v any, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = toml.Unmarshal(b, &v)
	return
}

package yaml

import (
	"context"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	yaml "gopkg.in/yaml.v3"
)

const typeName = "yaml"

func init() {
	lang.RegisterDataType(typeName, lang.DataTypeIsObject)
	stdio.RegisterReadArray(typeName, readArray)
	stdio.RegisterReadArrayWithType(typeName, readArrayWithType)
	stdio.RegisterReadMap(typeName, readMap)
	stdio.RegisterWriteArray(typeName, newArrayWriter)
	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex
	lang.RegisterMarshaller(typeName, marshal)
	lang.RegisterUnmarshaller(typeName, unmarshal)

	lang.SetMime(typeName,
		"application/yaml", // this is preferred but we will include others since not everyone follows standards.
		"application/x-yaml",
		"text/yaml",
		"text/x-yaml",
		"+yaml",
	)

	lang.SetFileExtensions(typeName, "yaml", "yml")
}

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	return lang.ArrayTemplate(ctx, yaml.Marshal, yaml.Unmarshal, read, callback)
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	return lang.ArrayWithTypeTemplate(ctx, typeName, yaml.Marshal, yaml.Unmarshal, read, callback)
}

func noCrLf(b []byte) []byte {
	if len(b) > 0 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	if len(b) > 0 && b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}

	return b
}

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
	return lang.MapTemplate(typeName, yaml.Marshal, yaml.Unmarshal, read, callback)
}

func readIndex(p *lang.Process, params []string) error {
	var jInterface any

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	return lang.IndexTemplateObject(p, params, &jInterface, yaml.Marshal)
}

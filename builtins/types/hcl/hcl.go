package hcl

import (
	"context"

	"github.com/hashicorp/hcl"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

const typeName = "hcl"

func init() {
	lang.RegisterDataType(typeName, lang.DataTypeIsObject)
	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex

	stdio.RegisterReadArray(typeName, readArray)
	stdio.RegisterReadArrayWithType(typeName, readArrayWithType)
	stdio.RegisterReadMap(typeName, readMap)
	stdio.RegisterWriteArray(typeName, newArrayWriter)

	lang.RegisterMarshaller(typeName, marshal)
	lang.RegisterUnmarshaller(typeName, unmarshal)

	// These are just guessed at as I couldn't find any formally named MIMEs
	lang.SetMime(typeName,
		"application/hcl",
		"application/x-hcl",
		"text/hcl",
		"text/x-hcl",
	)

	lang.SetFileExtensions(typeName, "hcl", "tf", "tfvars")
}

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	// Create a marshaller function to pass to ArrayTemplate
	marshaller := func(v any) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayTemplate(ctx, marshaller, hcl.Unmarshal, read, callback)
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v any) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayWithTypeTemplate(ctx, types.Json, marshaller, hcl.Unmarshal, read, callback)
}

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v any) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.MapTemplate(types.Json, marshaller, hcl.Unmarshal, read, callback)
}

func readIndex(p *lang.Process, params []string) error {
	var jInterface any

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = hcl.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	marshaller := func(iface any) ([]byte, error) {
		return json.Marshal(iface, p.Stdout.IsTTY())
	}

	return lang.IndexTemplateObject(p, params, &jInterface, marshaller)
}

func marshal(p *lang.Process, v any) ([]byte, error) {
	return json.Marshal(v, p.Stdout.IsTTY())
}

func unmarshal(p *lang.Process) (v any, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = hcl.Unmarshal(b, &v)
	return
}

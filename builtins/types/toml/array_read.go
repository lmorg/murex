package toml

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/pelletier/go-toml"
)

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	return lang.ArrayTemplate(ctx, toml.Marshal, toml.Unmarshal, read, callback)
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	return lang.ArrayWithTypeTemplate(ctx, typeName, toml.Marshal, toml.Unmarshal, read, callback)
}

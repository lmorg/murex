package bson

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"go.mongodb.org/mongo-driver/bson"
)

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	return lang.ArrayTemplate(ctx, bson.Marshal, bson.Unmarshal, read, callback)
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(interface{}, string)) error {
	return lang.ArrayWithTypeTemplate(ctx, types.Json, bson.Marshal, bson.Unmarshal, read, callback)
}

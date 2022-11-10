package sexp

import (
	"bytes"
	"context"

	"github.com/abesto/sexp"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArrayC(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	return readArray(ctx, read, callback, true)
}

func readArrayS(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	return readArray(ctx, read, callback, false)
}

func readArray(ctx context.Context, read stdio.Io, callback func([]byte), canonical bool) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j, err := sexp.Unmarshal(b)
	if err != nil {
		return err
	}

	for i := range j {
		select {
		case <-ctx.Done():
			return nil

		default:
			switch j[i].(type) {
			case string:
				callback(bytes.TrimSpace([]byte(j[i].(string))))

			default:
				jBytes, err := sexp.Marshal(j[i], canonical)
				if err != nil {
					return err
				}
				callback(jBytes)
			}
		}
	}

	return nil
}

func readArrayWithTypeC(ctx context.Context, read stdio.Io, callback func(interface{}, string)) error {
	return readArrayWithType(ctx, read, callback, true, sexpr)
}
func readArrayWithTypeS(ctx context.Context, read stdio.Io, callback func(interface{}, string)) error {
	return readArrayWithType(ctx, read, callback, false, sexpr)
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(interface{}, string), canonical bool, dataType string) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j, err := sexp.Unmarshal(b)
	if err != nil {
		return err
	}

	for i := range j {
		select {
		case <-ctx.Done():
			return nil

		default:
			switch j[i].(type) {
			case string:
				callback(bytes.TrimSpace([]byte(j[i].(string))), types.String)

			default:
				jBytes, err := sexp.Marshal(j[i], canonical)
				if err != nil {
					return err
				}
				callback(jBytes, dataType)
			}
		}
	}

	return nil
}

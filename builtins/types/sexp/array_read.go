package sexp

import (
	"bytes"

	"github.com/abesto/sexp"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArrayC(read stdio.Io, callback func([]byte)) error { return readArray(read, callback, true) }
func readArrayS(read stdio.Io, callback func([]byte)) error { return readArray(read, callback, false) }

func readArray(read stdio.Io, callback func([]byte), canonical bool) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j, err := sexp.Unmarshal(b)
	if err != nil {
		return err
	}

	for i := range j {
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

	return nil
}

func readArrayWithTypeC(read stdio.Io, callback func([]byte, string)) error {
	return readArrayWithType(read, callback, true, sexpr)
}
func readArrayWithTypeS(read stdio.Io, callback func([]byte, string)) error {
	return readArrayWithType(read, callback, false, sexpr)
}

func readArrayWithType(read stdio.Io, callback func([]byte, string), canonical bool, dataType string) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j, err := sexp.Unmarshal(b)
	if err != nil {
		return err
	}

	for i := range j {
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

	return nil
}

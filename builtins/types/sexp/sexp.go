package sexp

import (
	"bytes"
	"github.com/abesto/sexp"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types/define"
	"strconv"
)

const (
	sexpr = "sexp"
	csexp = "csexp"
)

func init() {
	streams.ReadArray[sexpr] = readArrayS
	streams.ReadMap[sexpr] = readMapS
	define.ReadIndexes[sexpr] = readIndexS
	define.Marshallers[sexpr] = marshalS
	define.Unmarshallers[sexpr] = unmarshal

	streams.ReadArray[csexp] = readArrayC
	streams.ReadMap[csexp] = readMapC
	define.ReadIndexes[csexp] = readIndexC
	define.Marshallers[csexp] = marshalC
	define.Unmarshallers[csexp] = unmarshal

	// These are just guessed at as I couldn't find any formally named MIMEs
	define.SetMime(sexpr,
		"application/sexp",
		"application/x-sexp",
		"text/sexp",
		"text/x-sexp",
	)

	define.SetFileExtensions(sexpr, "sexp")
}

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

func readMapC(read stdio.Io, config *config.Config, callback func(key, value string, last bool)) error {
	return readMap(read, config, callback, true)
}

func readMapS(read stdio.Io, config *config.Config, callback func(key, value string, last bool)) error {
	return readMap(read, config, callback, false)
}

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool), canonical bool) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	se, err := sexp.Unmarshal(b)
	if err == nil {

		for i := range se {
			j, err := sexp.Marshal(se[i], canonical)
			if err != nil {
				return err
			}
			callback(strconv.Itoa(i), string(j), i != len(se)-1)
		}

		return nil
	}
	return err
}

func readIndexC(p *proc.Process, params []string) error { return readIndex(p, params, true) }
func readIndexS(p *proc.Process, params []string) error { return readIndex(p, params, false) }

func readIndex(p *proc.Process, params []string, canonical bool) (err error) {
	var se interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	se, err = sexp.Unmarshal(b)
	if err != nil {
		return err
	}

	/*var seArray []interface{}

	for _, key := range params {
		i, err := strconv.Atoi(key)
		if err != nil {
			return err
		}

		if i < 0 {
			return errors.New("Cannot have negative keys in array.")
		}
		if i >= len(se) {
			return errors.New("Key '" + key + "' greater than number of items in array.")
		}

		if len(params) > 1 {
			seArray = append(seArray, se[i])

		} else {
			switch se[i].(type) {
			case string:
				p.Stdout.Write([]byte(se[i].(string)))
			default:
				b, err := sexp.Marshal(se[i], canonical)
				if err != nil {
					return err
				}
				p.Stdout.Writeln(b)
			}
		}
	}
	if len(seArray) > 0 {
		b, err := sexp.Marshal(seArray, canonical)
		if err != nil {
			return err
		}
		p.Stdout.Writeln(b)
	}
	return nil*/

	marshaller := func(iface interface{}) ([]byte, error) {
		return sexp.Marshal(iface, canonical)
	}

	return define.IndexTemplateObject(p, params, &se, marshaller)
}

func marshalC(_ *proc.Process, v interface{}) ([]byte, error) { return sexp.Marshal(v, true) }
func marshalS(_ *proc.Process, v interface{}) ([]byte, error) { return sexp.Marshal(v, false) }

func unmarshal(p *proc.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	v, err = sexp.Unmarshal(b)
	return
}

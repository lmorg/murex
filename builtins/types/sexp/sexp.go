package sexp

import (
	"github.com/abesto/sexp"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

const (
	sexpr = "sexp"
	csexp = "csexp"
)

func init() {
	stdio.RegisterReadArray(sexpr, readArrayS)
	stdio.RegisterReadArrayWithType(sexpr, readArrayWithTypeS)
	stdio.RegisterReadMap(sexpr, readMapS)
	stdio.RegisterWriteArray(sexpr, newArrayWriterS)
	lang.ReadIndexes[sexpr] = readIndexS
	lang.ReadNotIndexes[sexpr] = readIndexS
	lang.RegisterMarshaller(sexpr, marshalS)
	lang.RegisterUnmarshaller(sexpr, unmarshal)

	stdio.RegisterReadArray(csexp, readArrayC)
	stdio.RegisterReadArrayWithType(csexp, readArrayWithTypeC)
	stdio.RegisterReadMap(csexp, readMapC)
	stdio.RegisterWriteArray(csexp, newArrayWriterC)
	lang.ReadIndexes[csexp] = readIndexC
	lang.ReadNotIndexes[csexp] = readIndexC
	lang.RegisterMarshaller(csexp, marshalC)
	lang.RegisterUnmarshaller(csexp, unmarshal)

	// These are just guessed at as I couldn't find any formally named MIMEs
	lang.SetMime(sexpr,
		"application/sexp",
		"application/x-sexp",
		"text/sexp",
		"text/x-sexp",
	)

	lang.SetFileExtensions(sexpr, "sexp", "lisp")
}

func readIndexC(p *lang.Process, params []string) error { return readIndex(p, params, true) }
func readIndexS(p *lang.Process, params []string) error { return readIndex(p, params, false) }

func readIndex(p *lang.Process, params []string, canonical bool) (err error) {
	var se interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	se, err = sexp.Unmarshal(b)
	if err != nil {
		return err
	}

	marshaller := func(iface interface{}) ([]byte, error) {
		return sexp.Marshal(iface, canonical)
	}

	return lang.IndexTemplateObject(p, params, &se, marshaller)
}

func marshalC(_ *lang.Process, v interface{}) ([]byte, error) { return sexp.Marshal(v, true) }
func marshalS(_ *lang.Process, v interface{}) ([]byte, error) { return sexp.Marshal(v, false) }

func unmarshal(p *lang.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	v, err = sexp.Unmarshal(b)
	return
}

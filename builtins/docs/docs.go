package docs

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("murex-docs", cmdMurexDocs, types.String)
}

var (
	// Definition stores the definitions for builtins.
	Definition DocsFuncT

	// Summary stores a one line summary of each builtins.
	// This will be auto-populated by docgen
	Summary map[string]string

	// Synonym is used for builtins that might have more than one internal alias.
	// This will be auto-populated by docgen
	Synonym map[string]string
)

type DocsFuncT func(string) []byte

func cmdMurexDocs(p *lang.Process) error {
	cmd, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	switch cmd {
	case "--docs":
		p.Stdout.SetDataType(types.Json)
		b, err := json.Marshal(listDocs(), p.Stdout.IsTTY())
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		return err

	case "--digest", "--summary":
		p.Stdout.SetDataType(types.String)
		cmd, err := p.Parameters.String(1)
		if err != nil {
			return err
		}

		syn := Synonym[cmd]
		if Summary[syn] == "" {
			return fmt.Errorf("no summary found for command `%s`", cmd)
		}

		_, err = p.Stdout.Write([]byte(Summary[syn]))
		return err

	default:
		p.Stdout.SetDataType(types.String)
		syn := Synonym[cmd]
		if syn == "" {
			syn = cmd
		}
		b := Definition(syn)
		if len(b) == 0 {
			return fmt.Errorf("no documentation found for `%s`", cmd)
		}

		_, err = p.Stdout.Writeln(b)
		return err
	}
}

func listDocs() map[string]string {
	m := make(map[string]string)
	for k, v := range Synonym {
		m[k] = Summary[v]
	}
	return m
}

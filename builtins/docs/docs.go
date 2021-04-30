package docs

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["murex-docs"] = cmdMurexDocs
}

var (
	// Definition stores the definitions for builtins.
	Definition = make(map[string]string)

	// Summary stores a one line summary of each builtins.
	// This will be auto-populated by docgen
	Summary map[string]string

	// Synonym is used for builtins that might have more than one internal alias.
	// This will be auto-populated by docgen
	Synonym map[string]string
)

func cmdMurexDocs(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)
	cmd, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if cmd == "--digest" || cmd == "--summary" {
		cmd, err := p.Parameters.String(1)
		if err != nil {
			return err
		}

		syn := Synonym[cmd]
		if Summary[syn] == "" {
			return fmt.Errorf("No summary found for command `%s`", cmd)
		}

		_, err = p.Stdout.Write([]byte(Summary[syn]))
		return err
	}

	syn := Synonym[cmd]
	if Definition[syn] == "" {
		return fmt.Errorf("No documentation found for command `%s`", cmd)
	}

	_, err = p.Stdout.Writeln([]byte(Definition[syn]))
	return err
}

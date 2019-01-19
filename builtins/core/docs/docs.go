package docs

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["murex-docs"] = cmdMurexDocs
}

// Definition stores the definitions for builtins
var Definition = make(map[string]string)

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

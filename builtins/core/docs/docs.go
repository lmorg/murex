package docs

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["murex-docs"] = cmdMurexDocs
}

var Definition map[string]string = make(map[string]string)

func cmdMurexDocs(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)
	cmd, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if cmd == "--digest" {
		cmd, err := p.Parameters.String(1)
		if err != nil {
			return err
		}

		syn := Synonym[cmd]
		if Digest[syn] == "" {
			return fmt.Errorf("No digests found on command `%s`.", cmd)
		}

		_, err = p.Stdout.Write([]byte(Digest[syn]))
		return err
	}

	syn := Synonym[cmd]
	if Definition[syn] == "" {
		return fmt.Errorf("No documentation found on command `%s`.", cmd)
	}

	b, err := base64.StdEncoding.DecodeString(Definition[syn])
	if err != nil {
		return err
	}

	buf := bytes.NewReader(b)
	gz, err := gzip.NewReader(buf)
	defer gz.Close()
	if err != nil {
		return err
	}

	io.Copy(p.Stdout, gz)
	return err
}

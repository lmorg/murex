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

var docs map[string]string = make(map[string]string)

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

		syn := synonyms[cmd]
		if digests[syn] == "" {
			return fmt.Errorf("No digests found on command `%s`.", cmd)
		}

		_, err = p.Stdout.Write([]byte(digests[syn]))
		return err
	}

	syn := synonyms[cmd]
	if docs[syn] == "" {
		return fmt.Errorf("No documentation found on command `%s`.", cmd)
	}

	b, err := base64.StdEncoding.DecodeString(docs[syn])
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

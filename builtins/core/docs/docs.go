package docs

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
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

	if docs[cmd] == "" {
		return fmt.Errorf("No documentation found on command `%s`.", cmd)
	}

	b, err := base64.StdEncoding.DecodeString(docs[cmd])
	if err != nil {
		return err
	}

	buf := bytes.NewReader(b)
	gz, err := gzip.NewReader(buf)
	defer gz.Close()
	if err != nil {
		return err
	}

	_, err := io.copy(p.Stdout, gz)
	return err
}

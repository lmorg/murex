package textmanip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["pretty"] = cmdPretty
	lang.GoFunctions["sprintf"] = cmdSprintf
}

func cmdPretty(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, b, "", "    ")
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(prettyJSON.Bytes())
	return err
}

func cmdSprintf(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.Len() == 0 {
		return errors.New("Parameters missing")
	}

	s := p.Parameters.StringAll()
	var a []interface{}

	err := p.Stdin.ReadArray(func(b []byte) {
		a = append(a, string(b))
	})

	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(fmt.Sprintf(s, a...)))
	return err
}

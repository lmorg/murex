package textmanip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"net/url"
)

func init() {
	proc.GoFunctions["pretty"] = cmdPretty
	proc.GoFunctions["sprintf"] = cmdSprintf
	proc.GoFunctions["qs2json"] = cmdQs2Json
}

func cmdPretty(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

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

func cmdSprintf(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)

	if !p.IsMethod {
		return errors.New("I must be called as a method.")
	}

	if p.Parameters.Len() == 0 {
		return errors.New("Parameters missing.")
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

func cmdQs2Json(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)

	var b []byte

	if p.IsMethod {
		b, err = p.Stdin.ReadAll()
		if err != nil {
			return
		}

	} else {
		b = p.Parameters.ByteAll()
	}

	qs, err := url.ParseQuery(string(b))
	if err != nil {
		return
	}

	b, err = utils.JsonMarshal(qs, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Write(b)
	return
}

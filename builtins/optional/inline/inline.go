// +build ignore

package coreutils

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"os/exec"
)

func init() {
	lang.GoFunctions["inline"] = inline

	lang.GlobalConf.Define("inline", "languages", config.Properties{
		Description: "Map of supported languages and how to invoke their compilers.",
		Default:     "{}",
		DataType:    types.Json,
	})
}

func inline(p *lang.Process) error {
	p.Stdout.SetDataType(types.Generic)

	lang, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	code := p.Parameters.StringAllRange(1, -1)

	conf, err := lang.GlobalConf.Get("inline", "languages", types.Json)
	if err != nil {
		return err
	}

	compilers := make(map[string]string)
	err = json.Unmarshal([]byte(conf.(string)), &compilers)
	if err != nil {
		return err
	}

	if compilers[lang] == "" {
		return errors.New("No compiler found for language: '" + lang + "'.")
	}

	cmd := exec.Command(code)
	err = cmd.Start()
	return err
}

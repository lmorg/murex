package terminal

import (
	"encoding/json"
	"fmt"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/spellcheck"
)

const (
	seqApc = "\x1b_"
	seqST  = "\x1b\\"
)

func ttyphoonWriteApcJson(op, key string, value any) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = fmt.Printf("%s%s;%s;%s%s", seqApc, op, key, string(b), seqST)
	return err
}

type _mxapcConfigDictionaryT struct {
	Exclude []string `json:"history"`
}

// ttyphoonOutputBlockBegin reports the command line and supplied environment to Ttyphoon.
func ttyphoonOutputBlockBegin(commandLine string, envVars map[string]any) {
	if commandLine == "" {
		return
	}

	payload := struct {
		CmdLine string
		EnvVars map[string]any
	}{CmdLine: commandLine, EnvVars: envVars}
	if err := ttyphoonWriteApcJson("begin", "output-block", payload); err != nil {
		debug.Log("terminal: beginning Ttyphoon output block:", err)
	}
}

// ttyphoonOutputBlockEnd reports the command's exit status to Ttyphoon.
func ttyphoonOutputBlockEnd(exitNum int) {
	if !IsTtyphoon() {
		return
	}
	payload := struct {
		ExitNum int
	}{ExitNum: exitNum}
	if err := ttyphoonWriteApcJson("end", "output-block", payload); err != nil {
		debug.Log("terminal: ending Ttyphoon output block:", err)
	}
}

func PublishDictionary(p *lang.Process) {
	if !IsTtyphoon() {
		return
	}

	err := ttyphoonWriteApcJson("config", "dictionary", _mxapcConfigDictionaryT{Exclude: spellcheck.Exclusions(p)})
	if err != nil {
		debug.Log(err)
	}
}

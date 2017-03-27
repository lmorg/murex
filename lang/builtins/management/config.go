package management

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GlobalConf.Define("shell", "Prompt", config.Properties{
		Description: "Shell prompt",
		Default:     `{ out: "murex » " }`,
		DataType:    types.String,
	})
}

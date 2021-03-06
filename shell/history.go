// +build !js

package shell

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
)

func setPromptHistory() {
	h, err := history.New(home.MyDir + consts.PathSlash + ".murex_history")
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Error opening history file: " + err.Error()))
	} else {
		Prompt.History = h
	}
}

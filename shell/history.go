//go:build !js
// +build !js

package shell

import (
	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/readline/v4"
)

var promptHistory readline.History

func definePromptHistory() {
	h, err := history.New(profilepaths.HistoryPath())
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Error opening history file: " + err.Error()))
	} else {
		promptHistory = h
	}
}

func setPromptHistory() {
	Prompt.History = promptHistory
}

//go:build !js
// +build !js

package shell

import (
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils/readline"
)

var promptHistory readline.History

func definePromptHistory() {
	h, err := history.New(profile.HistoryPath())
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Error opening history file: " + err.Error()))
	} else {
		promptHistory = h
	}
}

func setPromptHistory() {
	Prompt.History = promptHistory
}

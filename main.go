package main

import (
	"compress/gzip"
	"encoding/json"
	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
	"io/ioutil"
	"os"
)

func main() {
	readFlags()

	proc.ShellProcess.State = state.Executing
	proc.ShellProcess.Name = os.Args[0]
	proc.ShellProcess.Parameters.Params = os.Args[1:]
	proc.ShellProcess.Scope = proc.ShellProcess
	proc.ShellProcess.Parent = proc.ShellProcess

	// Sets $SHELL to be murex
	shellEnv, err := utils.Executable()
	if err != nil {
		shellEnv = proc.ShellProcess.Name
	}
	os.Setenv("SHELL", shellEnv)

	// Pre-populate $PWDHIST with current working directory
	s, _ := os.Getwd()
	pwd := []string{s}
	if b, err := json.MarshalIndent(&pwd, "", "    "); err == nil {
		proc.GlobalVars.Set("PWDHIST", string(b), types.Json)
	}

	switch {
	case fCommand != "":
		config.Defaults(&proc.GlobalConf, false)
		execSource([]rune(fCommand))

	case len(fSource) > 0:
		shell.SigHandler()
		config.Defaults(&proc.GlobalConf, false)
		execSource(diskSource(fSource[0]))

	default:
		config.Defaults(&proc.GlobalConf, true)
		execSource([]rune(config.DefaultMurexProfile))
		execProfile()
		shell.Start()
	}

	debug.Log("[FIN]")
}

func diskSource(filename string) []rune {
	var b []byte

	file, err := os.Open(filename)
	if err != nil {
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		os.Exit(1)
	}

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			file.Close()
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
			os.Exit(1)
		}
		b, err = ioutil.ReadAll(gz)

		file.Close()
		gz.Close()

		if err != nil {
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
			os.Exit(1)
		}

	} else {
		b, err = ioutil.ReadAll(file)
		file.Close()
		if err != nil {
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
			os.Exit(1)
		}
	}

	return []rune(string(b))
}

func execSource(source []rune) {
	exitNum, err := lang.ProcessNewBlock(
		source,
		nil,
		nil,
		nil,
		proc.ShellProcess,
	)

	if err != nil {
		if exitNum == 0 {
			exitNum = 1
		}
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		os.Exit(exitNum)
	}

	if exitNum != 0 {
		os.Exit(exitNum)
	}
}

func execProfile() {
	profile := home.MyDir + consts.PathSlash + ".murex_profile"

	file, err := os.OpenFile(profile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		return
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		return
	}

	lang.ProcessNewBlock([]rune(string(b)), nil, nil, nil, proc.ShellProcess)
}

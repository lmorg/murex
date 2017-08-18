package main

import (
	"compress/gzip"
	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
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

	os.Setenv("SHELL", proc.ShellProcess.Name)

	switch {
	case fCommand != "":
		execSource([]rune(fCommand))

	case fStdin:
		os.Stderr.WriteString("Not implemented yet.\n")
		os.Exit(1)

	case len(fSource) > 0:
		shell.SigHandler()
		execSource(diskSource(fSource[0]))

	default:
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
	profile := home.MyDir + ".murex_profile"

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

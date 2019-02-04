package main

import (
	"compress/gzip"
	"io/ioutil"
	"os"

	_ "github.com/lmorg/murex/builtins"
	_ "github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
)

func main() {
	readFlags()
	lang.InitEnv()

	switch {
	case fCommand != "":
		// default config
		defaults.Defaults(lang.ShellProcess.Config, false)
		shell.SignalHandler(false)

		// load modules a profile
		if fLoadMods {
			profile.Execute()
		}

		// read block from command line parameters
		execSource([]rune(fCommand))

	case len(fSource) > 0:
		// default config
		defaults.Defaults(lang.ShellProcess.Config, false)
		shell.SignalHandler(false)

		// load modules a profile
		if fLoadMods {
			profile.Execute()
		}

		// read block from disk
		execSource(diskSource(fSource[0]))

	default:
		// default config
		defaults.Defaults(lang.ShellProcess.Config, true)

		// compiled profile
		execSource(defaults.DefaultMurexProfile())

		// load modules and profile
		profile.Execute()

		// start interactive shell
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
	exitNum, err := lang.RunBlockShellConfigSpace(source, nil, new(term.Out), term.NewErr(ansi.IsAllowed()))

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

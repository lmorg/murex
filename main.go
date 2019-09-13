package main

import (
	"compress/gzip"
	"io/ioutil"
	"os"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
)

const (
	interactive    bool = true
	nonInteractive bool = false
)

func main() {
	readFlags()
	lang.InitEnv()

	switch {
	case fRunTests:
		defaults.Defaults(lang.ShellProcess.Config, nonInteractive)
		shell.SignalHandler(nonInteractive)

		// compiled profile
		source := defaults.DefaultMurexProfile()
		ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
		execSource(defaults.DefaultMurexProfile(), ref)

		// enable tests
		if err := lang.ShellProcess.Config.Set("test", "enabled", true); err != nil {
			panic(err)
		}
		if err := lang.ShellProcess.Config.Set("test", "auto-report", true); err != nil {
			panic(err)
		}
		if err := lang.ShellProcess.Config.Set("test", "verbose", true); err != nil {
			panic(err)
		}

		// run unit tests
		if !lang.GlobalUnitTests.Run(lang.ShellProcess, "*") {
			os.Exit(1)
		}

	case fCommand != "":
		// default config
		defaults.Defaults(lang.ShellProcess.Config, nonInteractive)
		shell.SignalHandler(nonInteractive)

		// load modules and profile
		if fLoadMods {
			// compiled profile
			source := defaults.DefaultMurexProfile()
			ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
			execSource(defaults.DefaultMurexProfile(), ref)

			// local profile
			profile.Execute()
		}

		// read block from command line parameters
		execSource([]rune(fCommand), nil)

	case len(fSource) > 0:
		// default config
		defaults.Defaults(lang.ShellProcess.Config, nonInteractive)
		shell.SignalHandler(nonInteractive)

		// load modules a profile
		if fLoadMods {
			// compiled profile
			source := defaults.DefaultMurexProfile()
			ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
			execSource(defaults.DefaultMurexProfile(), ref)

			// local profile
			profile.Execute()
		}

		// read block from disk
		execSource(diskSource(fSource[0]), nil)

	default:
		// default config
		defaults.Defaults(lang.ShellProcess.Config, interactive)

		// compiled profile
		source := defaults.DefaultMurexProfile()
		ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
		execSource(defaults.DefaultMurexProfile(), ref)

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

func execSource(source []rune, sourceRef *ref.Source) {
	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN)
	fork.Stdout = new(term.Out)
	fork.Stderr = term.NewErr(ansi.IsAllowed())
	if sourceRef != nil {
		fork.FileRef.Source = sourceRef
	}
	exitNum, err := fork.Execute(source)

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

package profile

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/home"
)

const preloadMessage = `# This file is loaded before any murex modules. It should only contain
# environmental variables required for the modules to work eg:
#
#     export PATH=...
#
# Any other profile config belongs in your profile script instead:
# `

const (
	F_PRELOAD = 1 << iota
	F_MODULES
	F_PROFILE
)

// Execute runs the preload script, then murex modules followed by your murex profile
func Execute(flags int) {
	if flags == 0 {
		panic("no flags specified")
	}

	autocomplete.UpdateGlobalExeList()

	pwd, err := os.Getwd()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	if flags&F_PRELOAD != 0 {
		if err := profile(preloadFileName, PreloadPath()); err != nil {
			os.Stderr.WriteString("There were problems loading profile `" + PreloadPath() + "`:" + utils.NewLineString)
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	if flags&F_MODULES != 0 {
		if err := modules(ModulePath()); err != nil {
			os.Stderr.WriteString("There were problems loading modules `" + ModulePath() + "`:" + utils.NewLineString)
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	if flags&F_PROFILE != 0 {
		if err := profile(profileFileName, ProfilePath()); err != nil {
			os.Stderr.WriteString("There were problems loading profile `" + ProfilePath() + "`:" + utils.NewLineString)
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	err = os.Chdir(pwd)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

func profile(name, path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0640)
	if err != nil {
		return err
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if len(b) == 0 && path == PreloadPath() {
		err := file.Close()
		if err != nil {
			return err
		}
		file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0640)
		if err != nil {
			return err
		}
		_, err = file.WriteString(preloadMessage + ProfilePath() + strings.Repeat(utils.NewLineString, 3))
		if err != nil {
			return err
		}
	}

	err = os.Chdir(home.MyDir)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	block := []rune(string(b))

	os.Stderr.WriteString("Loading profile `" + name + "`" + utils.NewLineString)

	// lets redirect all output to STDERR just in case this thing gets piped for any strange reason
	fork := lang.ShellProcess.Fork(lang.F_NEW_MODULE | lang.F_NEW_TESTS | lang.F_NO_STDIN)
	fork.Stdout = term.NewErr(false)
	fork.Stderr = term.NewErr(ansi.IsAllowed())
	fork.FileRef.Source = ref.History.AddSource(path, "profile/"+name, b)

	_, err = fork.Execute(block)
	return err
}

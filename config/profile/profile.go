package profile

import (
	"io"
	"os"
	"strings"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/builtins/pipes/term"
	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
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
	F_DEFAULT = 1 << iota
	F_PRELOAD
	F_MOD_PRELOAD
	F_MODULES
	F_PROFILE
)

// Execute runs the preload script, then murex modules followed by your murex profile
func Execute(flags int) {
	if flags == 0 {
		panic("no flags specified")
	}

	pwd, err := os.Getwd()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	if flags&F_MOD_PRELOAD != 0 {
		if err := modules(profilepaths.ModulePath(), true); err != nil {
			os.Stderr.WriteString("There were problems loading modules `" + profilepaths.ModulePath() + "`:" + utils.NewLineString)
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	if flags&F_PRELOAD != 0 {
		if err := profile(profilepaths.PreloadFileName, profilepaths.PreloadPath()); err != nil {
			os.Stderr.WriteString("There were problems loading profile `" + profilepaths.PreloadPath() + "`:" + utils.NewLineString)
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	if flags&F_DEFAULT != 0 {
		defaultProfile()
		autocomplete.UpdateGlobalExeList()
	}

	if flags&F_MODULES != 0 {
		if err := modules(profilepaths.ModulePath(), false); err != nil {
			os.Stderr.WriteString("There were problems loading modules `" + profilepaths.ModulePath() + "`:" + utils.NewLineString)
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	if flags&F_PROFILE != 0 {
		if err := profile(profilepaths.ProfileFileName, profilepaths.ProfilePath()); err != nil {
			os.Stderr.WriteString("There were problems loading profile `" + profilepaths.ProfilePath() + "`:" + utils.NewLineString)
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	err = os.Chdir(pwd)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	lang.ShellProcess.FileRef = ref.NewModule(app.ShellModule)
}

func profile(name, path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0640)
	if err != nil {
		return err
	}

	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(b) == 0 && path == profilepaths.PreloadPath() {
		err := file.Close()
		if err != nil {
			return err
		}
		file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0640)
		if err != nil {
			return err
		}
		_, err = file.WriteString(preloadMessage + profilepaths.ProfilePath() + strings.Repeat(utils.NewLineString, 3))
		if err != nil {
			return err
		}
	}

	err = os.Chdir(home.MyDir)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	block := []rune(string(b))

	quiet, _ := lang.ShellProcess.Config.Get("shell", "quiet", types.Boolean)
	if v, ok := quiet.(bool); !ok || !v {
		os.Stderr.WriteString("Loading profile `" + name + "`" + utils.NewLineString)
	}

	// lets redirect all output to STDERR just in case this thing gets piped for any strange reason
	fork := lang.ShellProcess.Fork(lang.F_NEW_MODULE | lang.F_NEW_TESTS | lang.F_NO_STDIN)
	fork.Stdout = term.NewErr(false)
	fork.Stderr = term.NewErr(ansi.IsAllowed())
	moduleName := app.UserProfile + name
	fork.FileRef = &ref.File{Source: &ref.Source{Module: moduleName}}
	fork.FileRef.Source = ref.History.AddSource(path, moduleName, b)

	_, err = fork.Execute(block)

	autocomplete.UpdateGlobalExeList()

	return err
}

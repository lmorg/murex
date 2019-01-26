package profile

import (
	"io/ioutil"
	"os"

	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
)

var (
	// ProfilePath is the filename and path to the user profile
	ProfilePath = home.MyDir + consts.PathSlash + ".murex_profile"

	// ModulePath is the path to the modules directory
	ModulePath = home.MyDir + consts.PathSlash + ".murex_modules"
)

// Execute runs the murex modules followed by .murex_profile
func Execute() {
	pwd, err := os.Getwd()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	if err := modules(); err != nil {
		os.Stderr.WriteString("There were problems loading modules `" + ModulePath + "`:" + utils.NewLineString)
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
	}

	if err := profile(); err != nil {
		os.Stderr.WriteString("There were problems loading profile `" + ProfilePath + "`:" + utils.NewLineString)
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
	}

	err = os.Chdir(pwd)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

func profile() error {
	file, err := os.OpenFile(ProfilePath, os.O_RDONLY|os.O_CREATE, 0640)
	if err != nil {
		return err
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = os.Chdir(home.MyDir)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	// lets redirect all output to STDERR just in case this thing gets piped for any strange reason
	_, err = lang.RunBlockShellConfigSpace([]rune(string(b)), nil, term.NewErr(false), term.NewErr(ansi.IsAllowed()))
	return err
}

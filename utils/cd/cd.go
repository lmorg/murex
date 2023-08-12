package cd

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/cd/cache"
)

// GlobalVarName is the name of the path history variable that `cd` writes to
const GlobalVarName = "PWDHIST"

// Chdir changes the current working directory and updates the global working
// environment
func Chdir(p *lang.Process, path string) error {
	if pwd, _ := os.Getwd(); pwd == path {
		return nil
	}

	err := os.Chdir(path)
	if err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		p.Stderr.Writeln([]byte(err.Error()))
		pwd = path
	}

	if lang.Interactive {
		go cache.GatherFileCompletions(pwd)
	}

	// Update $PWD environmental variable for compatibility reasons
	err = os.Setenv("PWD", pwd)
	if err != nil {
		p.Stderr.Writeln([]byte(err.Error()))
	}

	// Update $PWDHIST murex variable - a more idiomatic approach to PWD
	pwdHist, err := lang.GlobalVariables.GetValue(GlobalVarName)
	if err != nil {
		return err
	}

	switch pwdHist.(type) {
	case []string:
		pwdHist = append(pwdHist.([]string), pwd)
	case []interface{}:
		pwdHist = append(pwdHist.([]interface{}), pwd)
	default:
		debug.Log(fmt.Sprintf("$%s has become corrupt (%T) so regenerating", GlobalVarName, pwdHist))
		pwdHist = []string{pwd}
	}

	err = lang.GlobalVariables.Set(p, GlobalVarName, pwdHist, types.Json)
	return err
}

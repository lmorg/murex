package cd

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
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

	go cacheFileCompletions(pwd)

	//ansititle.Icon([]byte(pwd))

	// Update $PWD environmental variable for compatibility reasons
	err = os.Setenv("PWD", pwd)
	if err != nil {
		p.Stderr.Writeln([]byte(err.Error()))
	}

	// Update $PWDHIST murex variable - a more idiomatic approach to PWD
	pwdhist := lang.GlobalVariables.GetValue(GlobalVarName)

	switch pwdhist.(type) {
	case []string:
		pwdhist = append(pwdhist.([]string), pwd)
	default:
		debug.Log(fmt.Sprintf("$%s has become corrupt (%t) so regenerating", GlobalVarName, pwdhist))
		pwdhist = []string{pwd}
	}

	lang.GlobalVariables.Set(p, GlobalVarName, pwdhist, types.Json)
	return err
}

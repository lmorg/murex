package cd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

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
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		pwd = path
	}

	// Update $PWD environmental variable for compatibility reasons
	err = os.Setenv("PWD", pwd)
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}

	// Update $PWDHIST murex variable - a more idiomatic approach to PWD
	hist := lang.ShellProcess.Variables.GetString("PWDHIST")
	if hist == "" {
		hist = "[]"
	}

	var v []string
	err = json.Unmarshal([]byte(hist), &v)
	if err != nil {
		return fmt.Errorf("Unable to unpack $PWDHIST: %s", err.Error())
	}

	v = append(v, pwd)
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return fmt.Errorf("Unable to repack $PWDHIST: %s", err.Error())
	}

	err = lang.ShellProcess.Variables.Set("PWDHIST", string(b), types.Json)

	return err
}

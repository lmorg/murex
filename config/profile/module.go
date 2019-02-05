package profile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/utils/ansi"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

// Module is the structure for each module within a module's directory.
// Each directory can have multiple modules - this is done so you can separate
// functionality into different logical modules but still keep them inside one
// git repository (or other source control). However I expect the typical usage
// would be one module per repository.
type Module struct {
	Name     string
	Summary  string
	Version  string
	Source   string
	Package  string
	Disabled bool
}

// Packages is a struct of all the modules
var Packages = make(map[string][]Module)

var disabled []string

func isDisabled(name string) bool {
	name = strings.ToLower(name)
	for i := range disabled {
		if strings.ToLower(disabled[i]) == name {
			return true
		}
	}

	return false
}

// Path returns the full path to the murex script that is sourced into your running shell
func (m *Module) Path() string {
	return ModulePath + m.Package + consts.PathSlash + m.Source
}

func (m *Module) validate() error {
	var message string
	if strings.TrimSpace(m.Name) == "" {
		message += `    Property "Name" is empty. This should contain the name of the module` + utils.NewLineString
	}

	if strings.TrimSpace(m.Summary) == "" {
		message += `    Property "Summary" is empty. This should contain a brief description of the module` + utils.NewLineString
	}

	if strings.TrimSpace(m.Version) == "" {
		message += `    Property "Version" is empty. This should contain a version number of this module` + utils.NewLineString
	}

	if strings.TrimSpace(m.Source) == "" {
		message += "    Property \"Source\" is empty. This should contain the name (or path) of the murex script to be `source`ed into your running shell as part of this module" + utils.NewLineString

	} else {
		fi, err := os.Stat(m.Path())

		if err != nil {
			message += fmt.Sprintf("    Unable to stat() script `%s`: %s%s", m.Path(), err.Error(), utils.NewLineString)

		} else if fi.IsDir() {
			message += fmt.Sprintf("    Script `%s` exists but is a directory%s", m.Path(), utils.NewLineString)
		}
	}

	if message != "" {
		return errors.New(message)
	}

	return nil
}

func (m *Module) execute() error {
	file, err := os.OpenFile(m.Path(), os.O_RDONLY, 0640)
	if err != nil {
		return err
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	block := []rune(string(b))

	os.Stderr.WriteString(fmt.Sprintf("Loading module `%s/%s`%s", m.Package, m.Name, utils.NewLineString))

	// lets redirect all output to STDERR just in case this thing gets piped for any strange reason
	fork := lang.ShellProcess.Fork(lang.F_NEW_MODULE | lang.F_FUNCTION | lang.F_NO_STDIN)
	fork.Stdout = term.NewErr(false)
	fork.Stderr = term.NewErr(ansi.IsAllowed())
	fork.Module = m.Package + "/" + m.Name
	_, err = fork.Execute(block)
	return err
}

package modules

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
)

func gitPackage(p *lang.Process) error {
	block := `exec git -C "$(MUREX_MODULE_PATH)/$(PACKAGE_NAME)" @PARAMS`

	if p.Parameters.Len() < 3 {
		return fmt.Errorf("expecting `murex-package git package parameters...`")
	}
	params := p.Parameters.StringArray()
	packageName := params[1]

	f, err := os.Stat(profile.ModulePath() + "/" + packageName)
	if err != nil {
		return err
	}
	if !f.IsDir() {
		return fmt.Errorf("package path doesn't exist or isn't a directory: %s", profile.ModulePath()+"/"+packageName)
	}

	fork := p.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN)
	fork.Name.Set("(git package)")
	fork.Parameters.DefineParsed(params[2:])
	fork.FileRef = ref.NewModule("shell/modules.gitPackage")
	fork.Variables.Set(p, "MUREX_MODULE_PATH", profile.ModulePath(), types.String)
	fork.Variables.Set(p, "PACKAGE_NAME", packageName, types.String)

	exitNum, err := fork.Execute([]rune(block))
	p.ExitNum = exitNum

	return err
}

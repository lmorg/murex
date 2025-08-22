package modules

import (
	"embed"

	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
)

//go:embed *.mx
var mxScripts embed.FS

func newPackage(p *lang.Process) error {
	b, err := mxScripts.ReadFile("cmd-new.mx")
	if err != nil {
		return err
	}
	block := []rune(string(b))

	fork := p.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN)
	fork.Name.Set("(new package)")
	fork.FileRef = ref.NewModule("shell/modules.newPackage")
	fork.Variables.Set(p, "MUREX_MODULE_PATH", profilepaths.ModulePath(), types.String)

	exitNum, err := fork.Execute(block)
	p.ExitNum = exitNum

	return err
}

package open

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/which"
)

func openSystemCommandGeneric(p *lang.Process, path string) error {
	cmd, err := openCommand()
	if err != nil {
		return fmt.Errorf("cannot open '%s': %s", path, err.Error())
	}

	fork := p.Fork(lang.F_DEFAULTS)
	block := fmt.Sprintf(`%s %s`, cmd, path)
	exitNum, err := fork.Execute([]rune(block))
	p.ExitNum = exitNum
	return err
}

var openCommands = []string{
	"open", "xdg-open",
}

func openCommand() (string, error) {
	for i := range openCommands {
		openPath := which.Which(openCommands[i])
		if openPath != "" {
			return openPath, nil
		}
	}

	return "", errors.New("cannot locate any external open handlers, eg `open` or `open-xdg`")
}

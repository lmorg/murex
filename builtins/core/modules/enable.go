package modules

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/config/profile"
	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/lang"
)

func enableModules(p *lang.Process) error {
	if p.Parameters.Len() < 2 {
		return errors.New(noModulesToAble + "enable")
	}

	var disabled []string
	err := profile.ReadJson(profilepaths.ModulePath()+profile.DisabledFile, &disabled)
	if err != nil {
		return err
	}

	defer writeDisabled(&disabled)

	for _, pack := range p.Parameters.StringArray()[1:] {
		switch strings.Count(pack, "/") {
		case 0:
			if err := enablePack(pack); err != nil {
				return err
			}
		case 1:
			if disabled, err = enableMod(pack, disabled); err != nil {
				return err
			}
		default:
			return fmt.Errorf("`%s` is not a valid package/module format", pack)
		}
	}

	return nil
}

func enablePack(pack string) error {
	modulePath := profilepaths.ModulePath()
	return os.Rename(modulePath+pack+profile.IgnoredExt, modulePath+pack)
}

func enableMod(mod string, disabled []string) ([]string, error) {
	for i := range disabled {
		if disabled[i] == mod {
			disabled[i] = disabled[len(disabled)-1]
			return disabled[:len(disabled)-1], nil
		}
	}

	return disabled, fmt.Errorf("`%s` does not exist or has already been enabled", mod)
}

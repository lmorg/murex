package modules

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

const noModulesToAble = "Missing package or package/module names to "

func disableModules(p *lang.Process) error {
	if p.Parameters.Len() < 2 {
		return errors.New(noModulesToAble + "disable")
	}

	var disabled []string
	if err := profile.ReadJson(profile.ModulePath()+profile.DisabledFile, &disabled); err != nil {
		return err
	}

	defer writeDisabled(&disabled)

	for _, pack := range p.Parameters.StringArray()[1:] {
		switch strings.Count(pack, "/") {
		case 0:
			if err := disablePack(pack); err != nil {
				return err
			}
		case 1:
			if err := disableMod(pack, &disabled); err != nil {
				return err
			}
		default:
			return fmt.Errorf("`%s` is not a valid package/module format", pack)
		}
	}

	return nil
}

func disablePack(pack string) error {
	modulePath := profile.ModulePath()
	return os.Rename(modulePath+pack, modulePath+pack+profile.IgnoredExt)
}

func disableMod(mod string, disabled *[]string) error {
	for i := range *disabled {
		if (*disabled)[i] == mod {
			return fmt.Errorf("`%s` has already been disabled", mod)
		}
	}

	*disabled = append(*disabled, mod)
	return nil
}

func writeDisabled(disabled *[]string) error {
	file, err := os.OpenFile(profile.ModulePath()+profile.DisabledFile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := json.Marshal(*disabled, true)
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	return err
}

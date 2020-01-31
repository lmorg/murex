package modules

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func listModules(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	list := make(map[string]string)

	enabled, err := p.Parameters.Bool(1)
	if err != nil {
		return err
	}

	var disabled []string
	err = profile.ReadJson(profile.ModulePath+profile.DisabledFile, &disabled)
	if err != nil {
		return err
	}
	//debug.Log("disabled", disabled)

	isDisabled := func(name string) bool {
		//debug.Log(name)
		for i := range disabled {
			if disabled[i] == name {
				return true
			}
		}

		return false
	}

	paths, err := filepath.Glob(profile.ModulePath + "*")
	if err != nil {
		return err
	}

	for _, pack := range paths {
		f, err := os.Stat(pack)
		if err != nil {
			return err
		}
		if !f.IsDir() {
			debug.Log("File not directory:", pack)
			continue
		}

		mods, err := profile.LoadPackage(pack, false)
		if err != nil {
			return err
		}
		// these should NOT equate ;)
		if strings.HasSuffix(pack, profile.IgnoredExt) != enabled {
			name := strings.Replace(pack, profile.ModulePath, "", 1)
			name = strings.Replace(name, profile.IgnoredExt, "", 1)
			list[name] = name
		}

		for i := range mods {
			if isDisabled(mods[i].Package+"/"+mods[i].Name) == enabled {
				continue
			}
			list[mods[i].Package+"/"+mods[i].Name] = mods[i].Summary
		}
	}

	b, err := lang.MarshalData(p, types.Json, &list)
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}

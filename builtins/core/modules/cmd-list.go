package modules

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

const expecting = `Expecting enabled | disabled | loaded | not-loaded | packages`

func listModules(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	flag, _ := p.Parameters.String(1)

	switch flag {
	case "enabled":
		return listModulesEnDis(p, true)

	case "disabled":
		return listModulesEnDis(p, false)

	case "loaded":
		return listModulesLoadNotLoad(p, true)

	case "not-loaded":
		return listModulesLoadNotLoad(p, false)

	case "packages":
		return listPackages(p)

	case "":
		return fmt.Errorf("Missing parameter. %s", expecting)

	default:
		return fmt.Errorf("Invalid parameter `%s`. %s", flag, expecting)
	}
}

func listModulesEnDis(p *lang.Process, enabled bool) error {
	var disabled []string
	err := profile.ReadJson(profile.ModulePath+profile.DisabledFile, &disabled)
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

	list := make(map[string]string)

	for _, pack := range paths {
		f, err := os.Stat(pack)
		if err != nil {
			return err
		}

		// only read directories
		if !f.IsDir() {
			debug.Log("File not directory:", pack)
			continue
		}

		// no hidden files nor empty strings
		if len(f.Name()) == 0 || f.Name()[0] == '.' {
			continue
		}

		mods, err := profile.LoadPackage(pack, false)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
		}

		// these should NOT equate ;)
		if strings.HasSuffix(f.Name(), profile.IgnoredExt) != enabled {
			name := cropIgnoreExt(f.Name())
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

func listModulesLoadNotLoad(p *lang.Process, loaded bool) error {
	list := make(map[string]string)

	for _, mods := range profile.Packages {
		for i := range mods {
			if mods[i].Loaded == loaded {
				list[mods[i].Package+"/"+mods[i].Name] = mods[i].Summary
			}
		}
	}

	b, err := lang.MarshalData(p, types.Json, &list)
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}

func listPackages(p *lang.Process) error {
	paths, err := filepath.Glob(profile.ModulePath + "*")
	if err != nil {
		return err
	}

	var list []string

	for _, pack := range paths {
		f, err := os.Stat(pack)
		if err != nil {
			return err
		}
		if !f.IsDir() {
			debug.Log("File not directory:", pack)
			continue
		}

		list = append(list, cropIgnoreExt(f.Name()))
	}

	b, err := lang.MarshalData(p, types.Json, &list)
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}

func cropIgnoreExt(name string) string {
	return strings.Replace(name, profile.IgnoredExt, "", 1)
}

package modules

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lmorg/murex/config/profile"
	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

const expecting = `Expecting enabled | disabled | loaded | not-loaded | packages`

func listModules(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	flag, _ := p.Parameters.String(1)

	switch flag {
	case "enabled":
		return listAndPrint(listModulesEnDis, p, true)

	case "disabled":
		return listAndPrint(listModulesEnDis, p, false)

	case "loaded":
		return listAndPrint(listModulesLoadNotLoad, p, true)

	case "not-loaded":
		return listAndPrint(listModulesLoadNotLoad, p, false)

	case "packages":
		return listPackages(p)

	case "":
		return fmt.Errorf("missing parameter. %s", expecting)

	default:
		return fmt.Errorf("invalid parameter `%s`. %s", flag, expecting)
	}
}

// listAndPrint is a wrapper function around the listModules...() functions. The
// rational behind this weird design is so that the code is concise and readable
// for normal execution but easily testable (ie the p.Stdout.Write code) is
// removed from the logic
func listAndPrint(fn func(*lang.Process, bool) (map[string]string, error), p *lang.Process, enabled bool) error {
	list, err := fn(p, enabled)
	if err != nil {
		return err
	}

	b, err := lang.MarshalData(p, types.Json, &list)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

// listModulesEnDis reads from disk rather than the package cache (like `runtime`)
// because the typical use for `murex-package list enabled|disabled` is to view
// which packages and modules will be loaded with murex. To get a view of what is
// currently loaded in a given session then use `loaded` / `not-loaded` instead of
// `enabled` / `disabled`
func listModulesEnDis(p *lang.Process, enabled bool) (map[string]string, error) {
	var disabled []string
	modulePath := profilepaths.ModulePath()

	err := profile.ReadJson(modulePath+profile.DisabledFile, &disabled)
	if err != nil {
		return nil, err
	}

	isDisabled := func(name string) bool {
		for i := range disabled {
			if disabled[i] == name {
				return true
			}
		}

		return false
	}

	paths, err := filepath.Glob(modulePath + "*")
	if err != nil {
		return nil, err
	}

	list := make(map[string]string)

	for _, pack := range paths {
		f, err := os.Stat(pack)
		if err != nil {
			return nil, err
		}

		// only read directories
		if !f.IsDir() {
			continue
		}

		// no hidden files nor empty strings
		if len(f.Name()) == 0 || f.Name()[0] == '.' {
			continue
		}

		mods, _ := profile.LoadPackage(pack, false)

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

	return list, nil
}

func listModulesLoadNotLoad(p *lang.Process, loaded bool) (map[string]string, error) {
	list := make(map[string]string)

	for _, mods := range profile.Packages {
		for i := range mods {
			if mods[i].Loaded == loaded {
				list[mods[i].Package+"/"+mods[i].Name] = mods[i].Summary
			}
		}
	}

	return list, nil
}

func listPackages(p *lang.Process) error {
	paths, err := filepath.Glob(profilepaths.ModulePath() + "*")
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

package profile

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

func modules() error {
	if !strings.HasSuffix(ModulePath, consts.PathSlash) {
		ModulePath += consts.PathSlash
	}

	// Check module path
	fi, err := os.Stat(ModulePath)
	if os.IsNotExist(err) {
		err = os.Mkdir(ModulePath, 0740)
		if err != nil {
			return err
		}

	} else if !fi.IsDir() {
		return errors.New("murex module path exists but is not a directory")
	}

	// Check module disable file
	if err = disabledFile(); err != nil {
		return errors.New(err.Error() + utils.NewLineString + "Skipping module loading for safety reasons")
	}

	paths, err := filepath.Glob(ModulePath + "*")
	if err != nil {
		return err
	}

	for i := range paths {
		err = readDir(paths[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func disabledFile() error {
	filename := ModulePath + "disabled.json"

	fi, err := os.Stat(filename)
	switch {
	case os.IsNotExist(err):
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0640)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString("[]")
		return err

	case fi.IsDir():
		return errors.New("disabled.json is a directory - it should be an ordinary file")

	case err != nil:
		return err

	default:
		return readJson(filename, &disabled)

	}
}

func readDir(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}

	// file is not a directory thus not a module
	if !f.IsDir() {
		return nil
	}

	// ignore directory (this goes further than disabling because it prevents
	// the manifest from even being read)
	if strings.HasSuffix(f.Name(), ".ignore") {
		return nil
	}

	var manifest []Manifest
	err = readJson(path+consts.PathSlash+"module.json", &manifest)
	if err != nil {
		return err
	}

	var message string

	for i := range manifest {
		manifest[i].path = f.Name()
		manifest[i].Disabled = manifest[i].Disabled || isDisabled(manifest[i].Name)
		err = manifest[i].validate()
		if err != nil {
			message += fmt.Sprintf(
				"Error loading module `%s` in path `%s`:%s%s%s",
				manifest[i].Name,
				manifest[i].Path(),
				utils.NewLineString,
				err.Error(),
				utils.NewLineString,
			)
			manifest[i].Disabled = true
		}

		if manifest[i].Disabled {
			continue
		}

		err = os.Chdir(path)
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}

		manifest[i].execute()
	}
	Modules[f.Name()] = manifest

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}

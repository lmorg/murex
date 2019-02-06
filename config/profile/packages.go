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

const (
	// DisabledFile is an array of disabled modules
	DisabledFile = "disabled.json"

	// PackagesFile is used by the package manager, `mpac`, but we auto-create it here for consistency
	PackagesFile = "packages.json"
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

	// Check package management file
	if err = packageFile(); err != nil {
		return errors.New(err.Error() + utils.NewLineString + "This will break murex's package manager, `mpac`, however modules will continue to work without it")
	}

	paths, err := filepath.Glob(ModulePath + "*")
	if err != nil {
		return err
	}

	var message string

	for i := range paths {
		err = LoadPackage(paths[i])
		if err != nil {
			message += err.Error() + utils.NewLineString
		}
	}

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}

func disabledFile() error {
	err := autoFile(DisabledFile)
	if err != nil {
		return err
	}

	return readJson(ModulePath+DisabledFile, &disabled)
}

func packageFile() error {
	return autoFile(PackagesFile)
}

func autoFile(name string) error {
	filename := ModulePath + name

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
		return errors.New(name + " is a directory - it should be an ordinary file")

	case err != nil:
		return err

	default:
		return nil
	}
}

// LoadPackage reads in the contents of the package and then validates and
// sources each module within
func LoadPackage(path string) error {
	//path := ModulePath + pack

	f, err := os.Stat(path)
	if err != nil {
		return err
	}

	// file is not a directory thus not a module
	if !f.IsDir() {
		return nil
	}

	// ignore hidden directories. eg version control (.git), IDE workspace
	// settings, OS X metadirectories and other guff.
	if strings.HasPrefix(f.Name(), ".") {
		return nil
	}

	// disable package directory (this goes further than disabling the module
	// because it prevents the modules from even being read)
	if strings.HasSuffix(f.Name(), ".disable") {
		return nil
	}

	var module []Module
	err = readJson(path+consts.PathSlash+"module.json", &module)
	if err != nil {
		return err
	}

	var message string

	for i := range module {
		module[i].Package = f.Name()
		module[i].Disabled = module[i].Disabled || isDisabled(module[i].Name)
		err = module[i].validate()
		if err != nil {
			message += fmt.Sprintf(
				"Error loading module `%s` in path `%s`:%s%s%s",
				module[i].Name,
				module[i].Path(),
				utils.NewLineString,
				err.Error(),
				utils.NewLineString,
			)
			module[i].Disabled = true
		}

		if module[i].Disabled {
			continue
		}

		err = os.Chdir(path)
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}

		err = module[i].execute()
		if err != nil {
			message += fmt.Sprintf(
				"Error sourcing module `%s` in path `%s`:%s%s%s",
				module[i].Name,
				module[i].Path(),
				utils.NewLineString,
				err.Error(),
				utils.NewLineString,
			)
		}
	}
	Packages[f.Name()] = module

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}

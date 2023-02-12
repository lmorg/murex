package profile

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lmorg/murex/lang/tty"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

const (
	// DisabledFile is an array of disabled modules
	DisabledFile = "disabled.json"

	// PackagesFile is used by the package manager, `murex-package`, but we auto-create
	// it here for consistency
	PackagesFile = "packages.json"

	// IgnoredExt is an file extension which can be used on package directories
	// to have them ignored during start up
	IgnoredExt = ".ignore"
)

func modules(modulePath string) error {
	// Check module path
	fi, err := os.Stat(modulePath)
	if os.IsNotExist(err) {
		err = os.Mkdir(modulePath, 0740)
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
		return errors.New(err.Error() + utils.NewLineString + "This will break murex's package manager, `murex-package`, however modules will continue to work without it")
	}

	paths, err := filepath.Glob(modulePath + "*")
	if err != nil {
		return err
	}

	var message string

	for i := range paths {
		_, err = LoadPackage(paths[i], true)
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

	return ReadJson(ModulePath()+DisabledFile, &disabled)
}

func packageFile() error {
	return autoFile(PackagesFile)
}

func autoFile(name string) error {
	filename := ModulePath() + name

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
// sources each module within. The path value should be an absolute path.
func LoadPackage(path string, execute bool) ([]Module, error) {
	// Because we are expecting an absolute path and any errors with it being
	// relative will have been compiled into the Go code, we want to raise a
	// panic here so those errors get caught during testing rather than buggy
	// code getting pushed back to the master branch and thus released.
	if !filepath.IsAbs(path) {
		panic("relative path used in LoadPackage")
	}

	f, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// file is not a directory thus not a module
	if !f.IsDir() {
		return nil, nil
	}

	// ignore hidden directories. eg version control (.git), IDE workspace
	// settings, OS X metadirectories and other guff.
	if strings.HasPrefix(f.Name(), ".") {
		return nil, nil
	}

	// disable package directory (this goes further than disabling the module
	// because it prevents the modules from even being read)
	if strings.HasSuffix(f.Name(), IgnoredExt) {
		return nil, nil
	}

	var module []Module
	err = ReadJson(path+consts.PathSlash+"module.json", &module)
	if err != nil {
		return nil, err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var message string

	for i := range module {
		module[i].Package = f.Name()
		module[i].Disabled = module[i].Disabled || isDisabled(module[i].Package+"/"+module[i].Name)
		err = module[i].validate()
		if err != nil && !module[i].Disabled {
			message += fmt.Sprintf(
				"Error loading module `%s` in path `%s`:%s%s%s",
				module[i].Name,
				module[i].Path(),
				utils.NewLineString,
				err.Error(),
				utils.NewLineString,
			)
			continue
		}

		if !execute || module[i].Disabled {
			continue
		}

		err = os.Chdir(path)
		if err != nil {
			tty.Stderr.WriteString(err.Error())
		}

		module[i].Loaded = true

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

	if execute {
		Packages[f.Name()] = module

		err = os.Chdir(pwd)
		if err != nil {
			message += err.Error() + utils.NewLineString
		}
	}

	if message != "" {
		return module, errors.New(strings.TrimSpace(message))
	}

	return module, nil
}

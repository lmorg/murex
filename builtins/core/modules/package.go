package modules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/utils/json"
)

type packageDb struct {
	Protocol string
	URI      string
	Package  string
	modules  []profile.Module
}

func readPackagesFile(path string) ([]packageDb, error) {
	var db []packageDb

	err := profile.ReadJson(path, &db)
	return db, err
}

func writePackagesFile(db *[]packageDb) error {
	path := profile.ModulePath() + profile.PackagesFile

	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := json.Marshal(db, true)
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	return err
}

func readPackageFile(path string) (profile.Package, error) {
	var pack profile.Package

	err := profile.ReadJson(path, &pack)
	return pack, err
}

func mvPackagePath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		panic("path should be absolute")
	}

	pack, err := readPackageFile(path + "/package.json")
	if err != nil {
		return path, err
	}

	err = os.Rename(path, profile.ModulePath()+pack.Name)
	if err != nil {
		return path, fmt.Errorf("Unable to do post-install tidy up: %s", err)
	}

	return pack.Name, nil
}

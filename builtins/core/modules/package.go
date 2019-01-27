package modules

import (
	"fmt"
	"io/ioutil"
	"os"

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

	file, err := os.OpenFile(path, os.O_RDONLY, 0640)
	if err != nil {
		return nil, fmt.Errorf("Cannot open `%s` for read: %s", path, err.Error())
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Cannot read contents of  `%s`: %s", path, err.Error())
	}

	err = json.UnmarshalMurex(b, &db)
	if err != nil {
		return nil, fmt.Errorf("Cannot unmarshal `%s`: %s", path, err.Error())
	}

	return db, nil
}

func writePackagesFile(db *[]packageDb) error {
	path := profile.ModulePath + profile.PackagesFile

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

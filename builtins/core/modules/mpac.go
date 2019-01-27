package modules

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/utils"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/readline"
)

func init() {
	lang.GoFunctions["mpac"] = cmdModuleAdmin

}

func cmdModuleAdmin(p *lang.Process) error {
	method, _ := p.Parameters.String(0)
	switch method {
	case "get":
		return getModule(p)

	case "update":
		return updateModules(p)

	case "import":
		return importModules(p)

	default:
		return errors.New("Missing or invalid parameters. Please use `get`, `update` or `import`")
	}
}

func getModule(p *lang.Process) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	db, err := readPackagesFile(profile.ModulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	protocol, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	uri, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	err = os.Chdir(profile.ModulePath)
	if err != nil {
		return fmt.Errorf("Unable to get package: %s", err.Error())
	}

	pack, err := getPackage(p, protocol, uri)
	if err != nil {
		return err
	}

	db = append(db, packageDb{
		Package:  pack,
		URI:      uri,
		Protocol: protocol,
	})

	var message string

	err = writePackagesFile(&db)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	err = profile.LoadPackage(pack)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	err = os.Chdir(pwd)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}

func getPackage(p *lang.Process, protocol, uri string) (pack string, err error) {
	p.Stderr.Writeln([]byte("Getting package from `" + uri + "`...."))

	switch protocol {
	case "git":
		pack, err = gitGet(p, uri)
		if err != nil {
			return "", fmt.Errorf("Unable to update package: %s", err.Error())
		}

	default:
		return "", fmt.Errorf("Unable to get package: Unknown protocol `%s`", protocol)
	}

	return
}

func updateModules(p *lang.Process) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	db, err := readPackagesFile(profile.ModulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	for i := range db {
		p.Stderr.Writeln(bytes.Repeat([]byte{'-'}, readline.GetTermWidth()))
		p.Stderr.Writeln([]byte("Updating package " + db[i].Package + "...."))

		err = os.Chdir(profile.ModulePath + db[i].Package)
		if err != nil {
			p.Stderr.Writeln([]byte(fmt.Sprintf(
				"Unable to update package `%s`: %s", db[i].Package, err.Error(),
			)))
			continue
		}

		switch db[i].Protocol {
		case "git":
			err = gitUpdate(p, &db[i])
			if err != nil {
				p.Stderr.Writeln([]byte(fmt.Sprintf(
					"Unable to update package `%s`: %s", db[i].Package, err.Error(),
				)))
			}

		default:
			p.Stderr.Writeln([]byte(fmt.Sprintf(
				"Unable to update package `%s`: Unknown protocol `%s`", db[i].Package, db[i].Protocol,
			)))
		}
	}

	return os.Chdir(pwd)
}

func importModules(p *lang.Process) error {
	path, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	if path == profile.ModulePath+profile.PackagesFile {
		return errors.New("You cannot import the same file as the master packages.json file")
	}

	importDb, err := readPackagesFile(path)
	if err != nil {
		return err
	}

	db, err := readPackagesFile(profile.ModulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	for i := range importDb {
		err = os.Chdir(profile.ModulePath)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			continue
		}

		p.Stderr.Writeln(bytes.Repeat([]byte{'-'}, readline.GetTermWidth()))
		p.Stderr.Writeln([]byte("Importing `" + importDb[i].Package + "`...."))
		err = packageDirExists(importDb[i].Package)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			continue
		}

		importDb[i].Package, err = getPackage(p, importDb[i].Protocol, importDb[i].URI)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			continue
		}

		db = append(db, importDb[i])

		err = profile.LoadPackage(profile.ModulePath + importDb[i].Package)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
		}
	}

	err = os.Chdir(pwd)
	if err != nil {
		p.Stderr.Writeln([]byte(err.Error()))
	}

	var message string

	err = writePackagesFile(&db)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	err = os.Chdir(pwd)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}

func packageDirExists(pack string) error {
	_, err := os.Stat(pack)
	if os.IsNotExist(err) {
		return nil
	}

	return errors.New("A file or directory already exists with that package name")
}

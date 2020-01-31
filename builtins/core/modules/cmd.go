package modules

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
)

const usage = `
Usage: murex-package install         uri
                     update
                     reload
                     enable|disable  package[/module]
                     import          [uri|local path]packages.json
					 status
					 list            enabled|disabled
`

func init() {
	lang.GoFunctions["murex-package"] = cmdModuleAdmin
}

func cmdModuleAdmin(p *lang.Process) error {
	method, _ := p.Parameters.String(0)
	switch method {
	case "install", "get":
		return getModule(p)

	case "update":
		return updateModules(p)

	case "reload":
		return reloadModules(p)

	case "import":
		return importModules(p)

	case "enable":
		return enableModules(p)

	case "disable":
		return disableModules(p)

	case "status":
		return statusModules(p)

	case "list":
		return listModules(p)

	default:
		return errors.New("Missing or invalid parameters." + usage)
	}
}

func detectProtocol(uri string) (string, error) {
	switch {
	case strings.HasPrefix(uri, "http://"):
		return "", errors.New("For security reasons, downloading packages is not allowed over plain text http. Please use `https://` instead")

	case strings.HasSuffix(uri, ".git"):
		return "git", nil

	case strings.HasPrefix(uri, "https://"):
		return "https", nil

	default:
		return "", errors.New("Unable to get package: Unable to auto-detect a suitable protocol for transferring the package")
	}
}

func updateModules(p *lang.Process) error {
	db, err := readPackagesFile(profile.ModulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	for i := range db {
		//p.Stderr.Writeln(bytes.Repeat([]byte{'-'}, readline.GetTermWidth()))
		p.Stderr.Writeln([]byte("Updating package " + db[i].Package + "...."))

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

	return nil
}

func reloadModules(p *lang.Process) error {
	profile.Execute()
	return nil
}

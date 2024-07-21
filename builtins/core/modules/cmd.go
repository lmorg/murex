package modules

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/cd"
)

const usage = `
Usage: murex-package install         uri
                     remove          package
                     update
                     reload
                     enable|disable  package[/module]
                     import          [uri|local path]packages.json
                     status
                     list            loaded|not-loaded|enabled|disabled|packages`

func init() {
	lang.DefineFunction("murex-package", cmdModuleAdmin, types.Json)
}

func cmdModuleAdmin(p *lang.Process) error {
	method, _ := p.Parameters.String(0)
	switch method {
	case "install":
		return getModule(p)

	case "remove":
		return removePackage(p)

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

	case "new":
		return newPackage(p)

	case "cd":
		return cdPackage(p)

	case "git":
		return gitPackage(p)

	default:
		return errors.New("missing or invalid parameters." + usage)
	}
}

func detectProtocol(uri string) (string, error) {
	switch {
	case strings.HasPrefix(uri, "http://"):
		return "", errors.New("for security reasons, downloading packages is not allowed over plain text http. Please use `https://` instead")

	case strings.HasSuffix(uri, ".git"):
		return "git", nil

	case strings.HasPrefix(uri, "https://"):
		return "https", nil

	default:
		return "", errors.New("unable to get package: Unable to auto-detect a suitable protocol for transferring the package")
	}
}

func cdPackage(p *lang.Process) error {
	modulePath := profile.ModulePath()

	path, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	f, err := os.Stat(modulePath + path)
	if err != nil {
		var err2 error
		f, err2 = os.Stat(modulePath + path + profile.IgnoredExt)
		if err2 != nil {
			return fmt.Errorf("unable to cd to package: %s: %s", err, err2)
		}
	}

	if !f.IsDir() {
		return fmt.Errorf("`%s` is not a directory", f.Name())
	}

	return cd.Chdir(p, modulePath+f.Name())
}

func updateModules(p *lang.Process) error {
	db, err := readPackagesFile(profile.ModulePath() + profile.PackagesFile)
	if err != nil {
		return err
	}

	for i := range db {
		if err := packageDirExists(profile.ModulePath() + "/" + db[i].Package); err == nil {
			write(p, "{BLUE}Skipping package '{BOLD}%s{RESET}{BLUE}'....{RESET}", db[i].Package)
			continue
		}

		write(p, "Updating package '{BOLD}%s{RESET}'....", db[i].Package)

		switch db[i].Protocol {
		case "git":
			err = gitUpdate(p, &db[i])
			if err != nil {
				write(p, "{RED}Unable to update package '{BOLD}%s{RESET}{RED}': %s{RESET}", db[i].Package, err.Error())
			}

		default:
			write(p, "{RED}Unable to update package '{BOLD}%s{RESET}{RED}': Unknown protocol '%s'{RESET}", db[i].Package, db[i].Protocol)
		}
	}

	return nil
}

func reloadModules(_ *lang.Process) error {
	profile.Execute(profile.F_MODULES)
	return nil
}

func write(p *lang.Process, format string, v ...any) {
	message := fmt.Sprintf("* "+ansi.ExpandConsts(format), v...)
	p.Stdout.Writeln([]byte(message))
}

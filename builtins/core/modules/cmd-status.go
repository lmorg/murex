package modules

import (
	"bytes"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/readline"
)

func statusModules(p *lang.Process) error {
	db, err := readPackagesFile(profile.ModulePath() + profile.PackagesFile)
	if err != nil {
		return err
	}

	for i := range db {
		p.Stdout.Writeln(bytes.Repeat([]byte{'-'}, readline.GetTermWidth()))

		if err := packageDirExists(profile.ModulePath() + "/" + db[i].Package); err == nil {
			write(p, "{BLUE}Skipping package '{BOLD}%s{RESET}{BLUE}'....{RESET}", db[i].Package)
			continue
		}

		write(p, "Package status '{BOLD}%s{RESET}'....", db[i].Package)

		switch db[i].Protocol {
		case "git":
			err = gitStatus(p, &db[i])
			if err != nil {
				write(p, "{RED}Unable to return status for package '{BOLD}%s{RESET}{RED}': %s", db[i].Package, err.Error())
			}

		default:
			write(p, "{RED}Unable to return status for package '{BOLD}%s{RESET}{RED}': Unknown protocol `%s`", db[i].Package, db[i].Protocol)
		}
	}

	return nil
}

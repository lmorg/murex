package modules

import (
	"bytes"
	"fmt"

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
		p.Stderr.Writeln(bytes.Repeat([]byte{'-'}, readline.GetTermWidth()))
		p.Stderr.Writeln([]byte("Package status " + db[i].Package + "...."))

		switch db[i].Protocol {
		case "git":
			err = gitStatus(p, &db[i])
			if err != nil {
				p.Stderr.Writeln([]byte(fmt.Sprintf(
					"unable to return package status `%s`: %s", db[i].Package, err.Error(),
				)))
			}

		default:
			p.Stderr.Writeln([]byte(fmt.Sprintf(
				"unable to return package status `%s`: Unknown protocol `%s`", db[i].Package, db[i].Protocol,
			)))
		}
	}

	return nil
}

package modules

import (
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/config/profile"
	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/lists"
)

func removePackage(p *lang.Process) error {
	modulePath := profilepaths.ModulePath()

	db, err := readPackagesFile(modulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	pack, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	if strings.Contains(pack, "/") {
		return fmt.Errorf("package '%s' contains invalid character '/'. You cannot remove modules, only packages.\nTo disable modules, use `murex-package disable %s` instead", pack, pack)
	}

	i := packageExistsInDb(pack, db)
	if i == -1 {
		return fmt.Errorf("package '%s' is either not installed or missing from 'packages.json'", pack)
	}

	db, err = lists.RemoveOrdered(db, i)
	if err != nil {
		return fmt.Errorf("cannot remove package from 'packages.json': %s", err.Error())
	}
	err = writePackagesFile(&db)
	if err != nil {
		return fmt.Errorf("cannot write 'packages.json': %s", err.Error())
	}

	sEnabled := modulePath + "/" + pack
	sDisabled := modulePath + "/" + pack + profile.IgnoredExt

	fEnabled, errEnabled := os.Stat(sEnabled)
	fDisabled, errDisabled := os.Stat(sDisabled)

	switch {
	case errEnabled == nil && fEnabled.IsDir():
		write(p, "Removing package '{BOLD}%s{RESET}'....", pack)
		write(p, "Please note that you'll have to restart murex for this to take effect")
		return os.RemoveAll(sEnabled)

	case errDisabled == nil && fDisabled.IsDir():
		write(p, "Removing package '{BOLD}%s{RESET}'....", pack)
		return os.RemoveAll(sDisabled)

	default:
		return fmt.Errorf("no package named '%s' was found in '%s'", pack, modulePath)
	}
}

// -1 == not found
// all other numbers are the index in []packageDb
func packageExistsInDb(pack string, db []packageDb) int {
	for i := range db {
		if db[i].Package == pack {
			return i
		}
	}
	return -1
}

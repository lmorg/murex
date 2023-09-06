package modules

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/cd"
)

func getModule(p *lang.Process) error {
	modulePath := profile.ModulePath()

	db, err := readPackagesFile(modulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	uri, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	err = cd.Chdir(p, modulePath)
	if err != nil {
		return fmt.Errorf("unable to get package: %s", err.Error())
	}

	pack, protocol, err := getPackage(p, uri)
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

	_, err = profile.LoadPackage(modulePath+pack, true)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}

func getPackage(p *lang.Process, uri string) (pack, protocol string, err error) {
	write(p, "Getting package from '%s'....", uri)

	protocol, err = detectProtocol(uri)
	if err != nil {
		return "", "", err
	}

	switch protocol {
	case "git":
		pack, err = gitGet(p, uri)
		if err != nil {
			return "", protocol, fmt.Errorf("unable to update package: %s", err.Error())
		}

	case "https":
		return "", protocol, errors.New("protocol handler for HTTPS has not yet been written. Please use git in the mean time (you can do this by specifying a git extension in the uri)")

	default:
		return "", "", fmt.Errorf("this is weird, protocol detected as `%s` but no handler has been written", protocol)
	}

	return
}

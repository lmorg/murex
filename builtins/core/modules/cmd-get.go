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
	db, err := readPackagesFile(profile.ModulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	uri, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	err = cd.Chdir(p, profile.ModulePath)
	if err != nil {
		return fmt.Errorf("Unable to get package: %s", err.Error())
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

	_, err = profile.LoadPackage(profile.ModulePath+pack, true)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}

func getPackage(p *lang.Process, uri string) (pack, protocol string, err error) {
	p.Stderr.Writeln([]byte("Getting package from `" + uri + "`...."))

	protocol, err = detectProtocol(uri)
	if err != nil {
		return "", "", err
	}

	switch protocol {
	case "git":
		pack, err = gitGet(p, uri)
		if err != nil {
			return "", protocol, fmt.Errorf("Unable to update package: %s", err.Error())
		}

	case "https":
		return "", protocol, errors.New("Protocol handler for HTTPS has not yet been written. Please use git in the mean time (you can do this by specifying a git extension in the uri)")

	default:
		return "", "", fmt.Errorf("This is weird, protocol detected as `%s` but no handler has been written", protocol)
	}

	return
}

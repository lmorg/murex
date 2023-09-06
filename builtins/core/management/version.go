package management

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("version", cmdVersion, types.String)

	defaults.AppendProfile(`
	autocomplete set version { [{
		"Flags": [ "--short", "--no-app-name", "--license", "--copyright" ]
	}] }
`)
}

var rxVersionNum = regexp.MustCompile(`^[0-9]+\.[0-9]+`)

func cmdVersion(p *lang.Process) error {
	s, _ := p.Parameters.String(0)

	switch s {

	case "--short":
		p.Stdout.SetDataType(types.Number)
		num := rxVersionNum.FindStringSubmatch(app.Version())
		if len(num) != 1 {
			return errors.New("unable to extract version number from string")
		}
		_, err := p.Stdout.Write([]byte(num[0]))
		return err

	case "--no-app-name":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.Version()))
		return err

	case "--license":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.License))
		return err

	case "--copyright":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.Copyright))
		return err

	case "":
		p.Stdout.SetDataType(types.String)
		v := fmt.Sprintf("%s: %s\n%s\n%s", app.Name, app.Version(), app.License, app.Copyright)
		_, err := p.Stdout.Writeln([]byte(v))
		return err

	default:
		return fmt.Errorf("not a valid parameter: %s", s)
	}

}

package management

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("version", cmdVersion, types.String)

	defaults.AppendProfile(`
	autocomplete set version %[{
		Flags: [
			--short     --no-app-name --semver
			--license   --license-full
			--copyright --build-date  --branch
		]
	}]
`)
}

func cmdVersion(p *lang.Process) error {
	s, _ := p.Parameters.String(0)

	switch s {

	case "--short":
		p.Stdout.SetDataType(types.Number)
		version := fmt.Sprintf("%d.%d", app.Major, app.Minor)
		_, err := p.Stdout.Write([]byte(version))
		return err

	case "--no-app-name":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.Version()))
		return err

	case "--semver":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(fmt.Sprintf("%d.%d.%d", app.Major, app.Minor, app.Revision)))
		return err

	case "--license":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.License))
		return err

	case "--license-full":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.GetLicenseFull()))
		return err

	case "--copyright":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.Copyright))
		return err

	case "--build-date":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.BuildDate))
		return err

	case "--branch":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.Branch))
		return err

	case "":
		p.Stdout.SetDataType(types.String)
		v := fmt.Sprintf(
			"%s: %s\nBuilt: %s\nLicense: %s\nCopyright: %s",
			strings.Title(app.Name), app.Version(),
			app.BuildDate,
			app.License,
			app.Copyright)
		_, err := p.Stdout.Writeln([]byte(v))
		return err

	default:
		return fmt.Errorf("not a valid parameter: %s", s)
	}

}

package typemgmt

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("datetime", cmdDateTime, types.Any, types.Any)
}

func cmdDateTime(p *lang.Process) error {
	flags, _, err := p.Parameters.ParseFlags(&parameters.Arguments{
		Flags: map[string]string{
			"--in":    types.String,
			"--out":   types.String,
			"--value": types.String,
		},
		AllowAdditional: false,
	})
	if err != nil {
		return err
	}

	if flags["--value"] == "" && flags["--in"] != "{now}" {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}

		flags["--value"] = string(b)
	}

	in := flags["--in"]
	out := flags["--out"]

	if in == "" {
		return errors.New("no datatime format provided for input value: `--in` required")
	}

	if out == "" {
		return errors.New("no datatime format provided for output value: `--out` required")
	}

	var datetime time.Time

	switch {
	case strings.HasPrefix(in, "{go}"):
		datetime, err = time.Parse(in[4:], flags["--value"])
		if err != nil {
			return err
		}

	case strings.HasPrefix(in, "{now}"):
		datetime = time.Now()
		if err != nil {
			return err
		}

	}

	switch {
	case strings.HasPrefix(out, "{go}"):
		_, err = p.Stdout.Write([]byte(datetime.Format(out[4:])))
		return err

	case strings.HasPrefix(out, "{unix}"):
		_, err = p.Stdout.Write([]byte(strconv.FormatInt(datetime.Unix(), 10)))
		return err
	}

	return errors.New("TODO!!!")
}

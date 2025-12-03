package time

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/escape"
)

func init() {
	lang.DefineMethod("datetime", cmdDateTime, types.Any, types.Any)

	defaults.AppendProfile(fmt.Sprintf(`
		autocomplete: set datetime %%[{
			FlagsDesc: {
				%[3]s: "(optional) Input data/time string to be parsed"
				%[1]s: "Formatting rules of input data/time string"
				%[2]s: "Formatting rules of output date/time string"
			}
			AllowMultiple: true
			AnyValue:      true
		}]`, _FLAG_IN, _FLAG_OUT, _FLAG_VAL))
}

const (
	_FLAG_IN  = "--in"
	_FLAG_OUT = "--out"
	_FLAG_VAL = "--value"

	_FORMAT_PYTHON = "{py}"
	_FORMAT_GOLANG = "{go}"
	_FORMAT_EPOCH  = "{unix}"
	_FORMAT_NOW    = "{now}"

	errTooManyParameters = "too many parameters without flags"
)

func cmdDateTime(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	flags, additional, err := p.Parameters.ParseFlags(&parameters.Arguments{
		Flags: map[string]string{
			_FLAG_IN:  types.String,
			_FLAG_OUT: types.String,
			_FLAG_VAL: types.String,
		},
		AllowAdditional: true,
	})
	if err != nil {
		return err
	}

	var (
		fIn    = flags.GetValue(_FLAG_IN).String()
		fOut   = flags.GetValue(_FLAG_OUT).String()
		fValue = flags.GetValue(_FLAG_VAL).String()
	)

	switch len(additional) {
	case 0:
		break

	case 1:
		if fIn != "" || fOut != "" || fValue != "" {
			escape.CommandLine(additional)
			return fmt.Errorf("%s: %s", errTooManyParameters, strings.Join(additional, " "))
		}
		fIn, fOut = _FORMAT_NOW, additional[0]
		goto skipFlagValidation

	default:
		escape.CommandLine(additional)
		return fmt.Errorf("%s: %s", errTooManyParameters, strings.Join(additional, " "))
	}

	if fValue == "" && fIn != _FORMAT_NOW {
		if err := p.ErrIfNotAMethod(); err != nil {
			return err
		}

		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		fValue = string(utils.CrLfTrim(b))
	}

	if fIn == "" {
		return fmt.Errorf("no datatime format provided for input value: `%s` required", _FLAG_IN)
	}

	if fOut == "" {
		return fmt.Errorf("no datatime format provided for output value: `%s` required", _FLAG_OUT)
	}

skipFlagValidation:

	var datetime time.Time

	// Parse --in

	switch {
	case strings.HasPrefix(fIn, _FORMAT_GOLANG):
		datetime, err = time.Parse(fIn[4:], fValue)
		if err != nil {
			return err
		}

	case strings.HasPrefix(fIn, _FORMAT_PYTHON):
		return errors.New("TODO! This feature hasn't yet been developed")

	case strings.HasPrefix(fIn, _FORMAT_NOW):
		datetime = time.Now()

	default:
		return fmt.Errorf("unknown or invalid input parser formatter, expecting `%s`, `%s` or `%s`", _FORMAT_GOLANG, _FORMAT_PYTHON, _FORMAT_NOW)
	}

	// Write --out

	switch {
	case strings.HasPrefix(fOut, _FORMAT_GOLANG):
		_, err = p.Stdout.Write([]byte(datetime.Format(fOut[4:])))
		return err

	case strings.HasPrefix(fOut, _FORMAT_PYTHON):
		s, err := pyParse(fOut[4:], datetime)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write([]byte(s))
		return err

	case strings.HasPrefix(fOut, _FORMAT_EPOCH):
		_, err = p.Stdout.Write([]byte(strconv.FormatInt(datetime.Unix(), 10)))
		return err

	default:
		return fmt.Errorf("unknown or invalid output parser formatter, expecting `%s`, `%s`, or `%s`", _FORMAT_GOLANG, _FORMAT_PYTHON, _FORMAT_EPOCH)
	}

}

func pyParse(format string, t time.Time) (string, error) {
	var (
		isDirective bool
		s           string
	)

	for _, c := range format {
		if !isDirective {
			if c == '%' {
				isDirective = true
			} else {
				s += string(c)
			}
			continue
		}

		isDirective = false

		switch c {
		case 'a':
			s += t.Format("Mon")
		case 'A':
			s += t.Format("Monday")
		case 'w':
			switch t.Format("Monday") {
			case "Sunday":
				s += "0"
			case "Monday":
				s += "1"
			case "Tuesday":
				s += "2"
			case "Wednesday":
				s += "3"
			case "Thursday":
				s += "4"
			case "Friday":
				s += "5"
			case "Saturday":
				s += "6"
			}
		case 'd':
			s += t.Format("02")
		case 'b':
			s += t.Format("Jan")
		case 'B':
			s += t.Format("January")
		case 'm':
			s += t.Format("01")
		case 'y':
			s += t.Format("06")
		case 'Y':
			s += t.Format("2006")
		case 'H':
			s += t.Format("15")
		case 'I':
			s += t.Format("03")
		case 'p':
			s += t.Format("PM")
		case 'M':
			s += t.Format("04")
		case 'S':
			s += t.Format("05")
		case 'f':
			s += t.Format("000000")
		case 'z':
			s += t.Format("-0700")
		case 'Z':
			s += t.Format("MST")
		case 'j':
			s += t.Format("002")
		case 'U':
			return "", errors.New("`%U` is currently unsupported")
		case 'W':
			return "", errors.New("`%W` is currently unsupported")
		case 'c':
			s += t.Format("Mon Jan 03 15:04:05 2006")
		case 'x':
			s += t.Format("01/02/06")
		case 'X':
			s += t.Format("15:04:05")
		case '%':
			s += "%"
		default:
			return "", fmt.Errorf("`%%%s` is not a valid Python datetime directive", string(c))
		}
	}

	return s, nil
}

package parameters

import (
	"errors"
	"github.com/lmorg/murex/lang/types"
	"strings"
)

type Arguments struct {
	AllowAdditional bool
	Flags           map[string]string
}

func (p *Parameters) ParseFlags(args *Arguments) (flags map[string]string, additional []string, err error) {
	return ParseFlags(p.Params, args)
}

func ParseFlags(params []string, args *Arguments) (flags map[string]string, additional []string, err error) {
	var previous string
	flags = make(map[string]string)

	for i := range params {
		switch {
		case strings.HasPrefix(params[i], "-"):
			switch {
			case previous != "":
				return nil, nil, errors.New("Invalid parameters! Flag found without value: " + previous)
			case args.Flags[params[i]] == types.Boolean:
				flags[params[i]] = types.TrueString
			case args.Flags[params[i]] != "":
				previous = params[i]
			default:
				return nil, nil, errors.New("Invalid parameters! Flag not recognised: " + params[i])
			}

		case previous != "":
			flags[previous] = params[i]
			previous = ""

		default:
			if !args.AllowAdditional {
				return nil, nil, errors.New("Invalid parameters! Parameter found without a flag: " + previous)
			}
			additional = append(additional, params[i])
		}
	}

	if previous != "" {
		return nil, nil, errors.New("Invalid parameters! Flag found without value: " + previous)
	}

	return
}

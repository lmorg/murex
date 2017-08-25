package parameters

import (
	"errors"
	"github.com/lmorg/murex/lang/types"
	"strings"
)

// Arguments is a struct which holds the allowed flags supported when parsing the flags (with ParseFlags)
type Arguments struct {
	AllowAdditional bool
	Flags           map[string]string
}

// ParseFlags parses the parameters and return which flags are set.
// `Arguments` is a list of supported flags taken as a struct to enable easy querying from within murex shell scripts.
// eg:
// 	   args {
// 	   	   "AllowAdditional": true,
// 	   	   "Flags": {
// 	 	     "--str": "str",
// 	 	     "--num": "num",
// 	 	     "--bool": "bool",
// 	 	     "-b": "--bool"
// 	   }
func ParseFlags(params []string, args *Arguments) (flags map[string]string, additional []string, err error) {
	var previous string
	flags = make(map[string]string)
	additional = make([]string, 0)

	for i := range params {
	scanFlags:
		switch {
		case strings.HasPrefix(params[i], "-"):
			switch {
			case strings.HasPrefix(args.Flags[params[i]], "-"):
				params[i] = args.Flags[params[i]]
				goto scanFlags
			case previous != "":
				return nil, nil, errors.New("Invalid parameters! Flag found without value: `" + previous + "`")
			case args.Flags[params[i]] == types.Boolean:
				flags[params[i]] = types.TrueString
			case args.Flags[params[i]] != "":
				previous = params[i]
			default:
				return nil, nil, errors.New("Invalid parameters! Flag not recognised: `" + params[i] + "`")
			}

		case previous != "":
			flags[previous] = params[i]
			previous = ""

		default:
			if !args.AllowAdditional {
				return nil, nil, errors.New("Invalid parameters! Parameter found without a flag: `" + params[i] + "`")
			}
			additional = append(additional, params[i])
		}
	}

	if previous != "" {
		return nil, nil, errors.New("Invalid parameters! Flag found without value: `" + previous + "`")
	}

	return
}

// ParseFlags - this instance of ParseFlags is a wrapper function for ParseFlags (above) so you can use inside your
// proc.Process.Parameters object
func (p *Parameters) ParseFlags(args *Arguments) (flags map[string]string, additional []string, err error) {
	return ParseFlags(p.Params, args)
}

package parameters

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

// Arguments is a struct which holds the allowed flags supported when parsing the flags (with ParseFlags)
type Arguments struct {
	AllowAdditional    bool
	IgnoreInvalidFlags bool
	Flags              map[string]string
}

const invalidParameters = "invalid parameters"

// ParseFlags parses the parameters and return which flags are set.
// `Arguments` is a list of supported flags taken as a struct to enable easy querying from within murex shell scripts.
//
//	  args {
//	  	   "allowadditional": true,
//	  	   "flags": {
//		     "--str": "str",
//		     "--num": "num",
//		     "--bool": "bool",
//		     "-b": "--bool"
//	  }
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
			//case previous != "":
			//	return nil, nil, fmt.Errorf("%s: flag found without value: `%s`", invalidParameters, previous)
			case args.Flags[params[i]] == types.Boolean:
				flags[params[i]] = types.TrueString
			case args.Flags[params[i]] != "":
				previous = params[i]
			case previous != "":
				flags[previous] = params[i]
				previous = ""
			case args.IgnoreInvalidFlags && args.AllowAdditional:
				additional = append(additional, params[i])
			default:
				return nil, nil, fmt.Errorf("%s: flag not recognised: `%s`", invalidParameters, params[i])
			}

		case previous != "":
			flags[previous] = params[i]
			previous = ""

		default:
			if !args.AllowAdditional {
				return nil, nil, fmt.Errorf("%s: parameter found without a flag: `%s`", invalidParameters, params[i])
			}
			additional = append(additional, params[i])
		}
	}

	if previous != "" {
		return nil, nil, fmt.Errorf("%s: flag found without value: `%s`", invalidParameters, previous)
	}

	return
}

// ParseFlags - this instance of ParseFlags is a wrapper function for ParseFlags (above) so you can use inside your
// lang.Process.Parameters object
func (p *Parameters) ParseFlags(args *Arguments) (flags map[string]string, additional []string, err error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return ParseFlags(p.params, args)
}

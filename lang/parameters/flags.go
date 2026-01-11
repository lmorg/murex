package parameters

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

type FlagValueT interface {
	String() string
	Integer() int
	Number() float64
	Boolean() bool
	Any() any
}

type flagValue struct {
	v  any
	dt string
}

func (fv *flagValue) String() string {
	if fv.dt == types.String {
		return fv.v.(string)
	}
	panic(fmt.Sprintf("cannot String() on %s", fv.dt))
}

func (fv *flagValue) Integer() int {
	if fv.dt == types.Integer {
		return fv.v.(int)
	}
	panic(fmt.Sprintf("cannot Integer() on %s", fv.dt))
}

func (fv *flagValue) Number() float64 {
	if fv.dt == types.Number || fv.dt == types.Float {
		return fv.v.(float64)
	}
	panic(fmt.Sprintf("cannot Number() on %s", fv.dt))
}

func (fv *flagValue) Boolean() bool {
	if fv.dt == types.Boolean {
		return fv.v.(bool)
	}
	panic(fmt.Sprintf("cannot Boolean() on %s", fv.dt))
}

func (fv *flagValue) Any() any {
	return fv.v
}

type nullValue struct{}

func (nv *nullValue) String() string  { return "" }
func (nv *nullValue) Integer() int    { return 0 }
func (nv *nullValue) Number() float64 { return 0 }
func (nv *nullValue) Boolean() bool   { return false }
func (nv *nullValue) Any() any        { return nil }

type FlagsT struct {
	flags map[string]FlagValueT
}

func (f *FlagsT) set(flag string, v any, dt string) error {
	v, err := types.ConvertGoType(v, dt)
	if err != nil {
		return fmt.Errorf("flag %s is not a %s\n%v", flag, dt, err)
	}
	f.flags[flag] = &flagValue{v: v, dt: dt}
	return nil
}

func (f *FlagsT) GetValue(flag string) FlagValueT {
	v, ok := f.flags[flag]
	if ok {
		return v
	}
	return &nullValue{}
}

func (f *FlagsT) GetNullable(flag string) (FlagValueT, bool) {
	v, ok := f.flags[flag]
	return v, ok
}

func (f *FlagsT) GetMap() map[string]any {
	m := make(map[string]any)
	for k, v := range f.flags {
		m[k] = v.Any()
	}
	return m
}

func (f *FlagsT) Len() int {
	return len(f.flags)
}

func newFlagsT() *FlagsT {
	return &FlagsT{flags: make(map[string]FlagValueT)}
}

// Arguments is a struct which holds the allowed flags supported when parsing the flags (with ParseFlags)
type Arguments struct {
	AllowAdditional     bool
	IgnoreInvalidFlags  bool
	StrictFlagPlacement bool
	Flags               map[string]string
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
//
// Returns:
// 1. map of flags,
// 2. additional parameters,
// 3. error
func ParseFlags(params []string, args *Arguments) (*FlagsT, []string, error) {
	var (
		previous    string
		flags       = newFlagsT()
		additional  = make([]string, 0)
		ignoreFlags bool
		i           int
	)

	for i = range params {
	scanFlags:
		switch {
		case ignoreFlags:
			additional = append(additional, params[i])

		case strings.HasPrefix(params[i], "-"):
			switch {
			case args.AllowAdditional && params[i] == "--":
				ignoreFlags = true
			case strings.HasPrefix(args.Flags[params[i]], "-"):
				params[i] = args.Flags[params[i]]
				goto scanFlags
			case args.Flags[params[i]] == types.Boolean:
				flags.set(params[i], true, types.Boolean)
			case args.Flags[params[i]] != "":
				previous = params[i]
			case previous != "":
				if err := flags.set(previous, params[i], args.Flags[previous]); err != nil {
					return nil, nil, err
				}
				previous = ""
			case args.IgnoreInvalidFlags && args.AllowAdditional:
				additional = append(additional, params[i])
			default:
				return nil, nil, fmt.Errorf("%s: flag not recognized: `%s`", invalidParameters, params[i])
			}

		case previous != "":
			if err := flags.set(previous, params[i], args.Flags[previous]); err != nil {
				return nil, nil, err
			}
			previous = ""

		default:
			if !args.AllowAdditional {
				return nil, nil, fmt.Errorf("%s: parameter found without a flag: `%s`", invalidParameters, params[i])
			}
			additional = append(additional, params[i])
			if args.StrictFlagPlacement {
				ignoreFlags = true
			}
		}
	}

	if previous != "" {
		return nil, nil, fmt.Errorf("%s: flag found without value: `%s`", invalidParameters, previous)
	}

	return flags, additional, nil
}

// ParseFlags - this instance of ParseFlags is a wrapper function for ParseFlags (above) so you can use inside your
// lang.Process.Parameters object
func (p *Parameters) ParseFlags(args *Arguments) (flags *FlagsT, additional []string, err error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return ParseFlags(p.params, args)
}

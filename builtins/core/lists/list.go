package lists

import (
	"github.com/lmorg/murex/config/defaults"
)

func init() {
	defaults.AppendProfile(`
		alias list.sort    = msort
		alias list.reverse = mtac
		alias list.prepend = prepend
		alias list.append  = append
		alias list.prefix  = prefix
		alias list.suffix  = suffix
		alias list.left    = left
		alias list.right   = right
		alias list.regex   = regexp
		alias list.string  = match
		alias list.split   = jsplit
	`)
}

/*func cmdList(p *lang.Process) error {
	cmd, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	p.Parameters.DefineParsed(p.Parameters.StringArray()[1:])

	switch cmd {
	case "sort":
		return cmdMSort(p)
	case "reverse":
		return cmdMtac(p)
	case "prepend":
		return cmdPrepend(p)
	case "append":
		return cmdAppend(p)
	case "prefix":
		return cmdPrefix(p)
	case "suffix":
		return cmdSuffix(p)
	case "left":
		return cmdPrefix(p)
	case "right":
		return cmdSuffix(p)
	case "regex":
		return cmdRegexp(p)
	case "string":
		return cmdMatch(p)
	case "split":
		return cmdJsplit(p)
	}

	return nil
}*/

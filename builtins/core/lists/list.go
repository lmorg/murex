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

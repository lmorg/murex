package datatools

import "github.com/lmorg/murex/config/defaults"

func init() {
	defaults.AppendProfile(`
		alias  struct.alter =  alter
		alias  struct.count =  count
		alias  struct.keys  =  struct-keys
	`)
}

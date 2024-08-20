package arraytools

import "github.com/lmorg/murex/config/defaults"

func init() {
	defaults.AppendProfile(`
		alias  struct.new.2darray =  2darray
		alias  struct.new.heading =  addheading
		alias  struct.new.map     =  map
	`)
}

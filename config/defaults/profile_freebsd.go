//go:build freebsd
// +build freebsd

package defaults

import (
	_ "embed"
)

//go:embed profile_freebsd.mx
var profileFreebsd []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_freebsd",
		Block: profileFreebsd,
	})
}

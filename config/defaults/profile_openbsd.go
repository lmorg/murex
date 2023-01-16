//go:build openbsd
// +build openbsd

package defaults

import (
	_ "embed"
)

//go:embed profile_openbsd.mx
var profileOpenbsd []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_openbsd.mx",
		Block: profileOpenbsd,
	})
}

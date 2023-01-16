//go:build netbsd
// +build netbsd

package defaults

import (
	_ "embed"
)

//go:embed profile_netbsd.mx
var profileNetbsd []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_netbsd.mx",
		Block: profileNetbsd,
	})
}

//go:build solaris
// +build solaris

package defaults

import (
	_ "embed"
)

//go:embed profile_solaris.mx
var profileSolaris []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_solaris",
		Block: profileSolaris,
	})
}

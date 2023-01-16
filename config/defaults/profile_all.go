package defaults

import (
	_ "embed"
)

//go:embed profile_all.mx
var profileAll []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_all.mx",
		Block: profileAll,
	})
}

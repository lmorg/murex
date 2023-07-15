package defaults

import (
	_ "embed"
)

//go:embed profile_any.mx
var profileAny []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_any",
		Block: profileAny,
	})
}

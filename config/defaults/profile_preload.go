package defaults

import (
	_ "embed"
)

//go:embed profile_preload.mx
var profilePreload []byte

func init() {
	// push to top of slice so it is first profile loaded
	DefaultProfiles = append([]*DefaultProfileT{{
		Name:  "profile_preload",
		Block: profilePreload,
	}}, DefaultProfiles...)
}

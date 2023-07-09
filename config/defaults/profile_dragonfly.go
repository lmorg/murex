//go:build dragonfly
// +build dragonfly

package defaults

import (
	_ "embed"
)

//go:embed profile_dragonfly.mx
var profileDragonfly []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_dragonfly",
		Block: profileDragonfly,
	})
}

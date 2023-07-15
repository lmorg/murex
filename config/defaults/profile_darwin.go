//go:build darwin
// +build darwin

package defaults

import (
	_ "embed"
)

//go:embed profile_darwin.mx
var profileDarwin []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_darwin",
		Block: profileDarwin,
	})
}

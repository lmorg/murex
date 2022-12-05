//go:build !windows && !plan9
// +build !windows,!plan9

package defaults

import (
	_ "embed"
)

//go:embed profile_posix.mx
var profilePosix []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_posix.mx",
		Block: profilePosix,
	})
}

//go:build plan9
// +build plan9

package defaults

import (
	_ "embed"
)

//go:embed profile_plan9.mx
var profilePlan9 []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_plan9",
		Block: profilePlan9,
	})
}

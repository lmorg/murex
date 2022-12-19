package defaults

import "strings"

var murexProfile []string

// AppendProfile is used as a way of creating a platform specific default
// profile generated at compile time
func AppendProfile(block string) {
	murexProfile = append(murexProfile, "\n"+block+"\n")
}

type DefaultProfileT struct {
	Name  string
	Block []byte
}

var DefaultProfiles []*DefaultProfileT

func AddMurexProfile() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile",
		Block: []byte(strings.Join(murexProfile, "\n\n")),
	})
}

package profile

import (
	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile/source"
	"github.com/lmorg/murex/integrations"
	"github.com/lmorg/murex/lang/ref"
)

func defaultProfile() {
	defaults.AddMurexProfile()

	for _, profile := range defaults.DefaultProfiles {
		ref := ref.History.AddSource("(builtin)", app.ShellProfile+profile.Name, profile.Block)
		source.Exec([]rune(string(profile.Block)), ref, false)
	}

	for _, profile := range integrations.Profiles() {
		ref := ref.History.AddSource("(builtin)", app.ShellProfile+"integrations_"+profile.Name, profile.Block)
		source.Exec([]rune(string(profile.Block)), ref, false)
	}
}

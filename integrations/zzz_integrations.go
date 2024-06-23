package integrations

import (
	"embed"
	"fmt"

	"github.com/lmorg/murex/config/defaults"
)

var resources map[string]*embed.FS

func init() {
	resources = make(map[string]*embed.FS)
}

func Profiles() []*defaults.DefaultProfileT {
	var profiles []*defaults.DefaultProfileT

	for resName, res := range resources {

		files, err := res.ReadDir(`.`)
		if err != nil {
			panic(fmt.Sprintf("err in res.ReadDir(`.`) against resources[%s]: %s", resName, err.Error()))
		}

		for _, f := range files {
			b, err := res.ReadFile(f.Name())
			if err != nil {
				panic(fmt.Sprintf("err in res.ReadFile(%s) against resources[%s]: %s", f.Name(), resName, err.Error()))
			}

			profiles = append(profiles, &defaults.DefaultProfileT{
				Name:  f.Name()[:len(f.Name())-3], // strip ".mx" from filename
				Block: b,
			})
		}

	}

	return profiles
}

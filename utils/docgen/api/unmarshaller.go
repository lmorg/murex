package docgen

import (
	"fmt"

	yaml "gopkg.in/yaml.v3"
)

func parseSourceFile(path string, structure interface{}) {
	f := fileReader(path)
	yml := yaml.NewDecoder(f)
	yml.KnownFields(true)
	err := yml.Decode(structure)
	if err != nil {
		panic(fmt.Sprintf("%s (%s)", err.Error(), path))
	}
}

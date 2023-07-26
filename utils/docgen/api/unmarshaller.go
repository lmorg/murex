package docgen

import (
	"fmt"

	yaml "gopkg.in/yaml.v3"
)

func parseSourceFile(path string, structure interface{}) {
	f := fileReader(path)
	b := readAll(f)

	err := yaml.Unmarshal(b, structure)
	if err != nil {
		panic(fmt.Sprintf("%s (%s)", err.Error(), path))
	}
}

/*func structuredMessage(message string, v interface{}) string {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err.Error())
	}

	return message + "\n" + string(b)
}
*/

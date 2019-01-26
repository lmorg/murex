package profile

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lmorg/murex/utils/json"
)

func readJson(path string, v interface{}) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0640)
	if err != nil {
		return fmt.Errorf("Cannot open `%s` for read: %s", path, err.Error())
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Cannot read contents of  `%s`: %s", path, err.Error())
	}

	err = json.UnmarshalMurex(b, v)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal `%s`: %s", path, err.Error())
	}

	return nil
}

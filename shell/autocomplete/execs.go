package autocomplete

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

var wslMounts []string

// Read returns an interface{} of the user dictionary.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func WslMountsGet() (interface{}, error) {
	return wslMounts, nil
}

// Write takes a JSON-encoded string and writes it to the dictionary slice.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WslMountsSet(v interface{}) error {
	switch v := v.(type) {
	case string:
		return json.Unmarshal([]byte(v), &wslMounts)

	default:
		return fmt.Errorf("invalid data-type. Expecting a %s encoded string", types.Json)
	}
}

func listExesWindows(path string, exes map[string]bool) {
	var showExts bool

	v, err := lang.ShellProcess.Config.Get("shell", "extensions-enabled", types.Boolean)
	if err != nil {
		showExts = false
	} else {
		showExts = v.(bool)
	}

	files, _ := os.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		name := strings.ToLower(f.Name())

		if len(name) < 5 {
			continue
		}

		ext := name[len(name)-4:]

		if ext == ".exe" || ext == ".com" || ext == ".bat" || ext == ".cmd" || ext == ".scr" {
			if showExts {
				exes[name] = true
			} else {
				exes[name[:len(name)-4]] = true
			}
		}
	}
}

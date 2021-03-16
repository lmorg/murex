package envvars

import (
	"os"
)

// All returns a map of environmental variables
func All() (map[string]string, error) {
	return all(os.Environ())
}

func all(envs []string) (map[string]string, error) {
	m := make(map[string]string)

	for _, s := range envs {
		if s == "" || s == "=" {
			m[""] = ""
			continue
		}

		var i int
		for ; i < len(s); i++ {
			if s[i] == '=' {
				break
			}
		}

		switch i {
		case 0:
			m[""] = s[1:]
		case len(s):
			m[s] = ""
		default:
			m[s[:i]] = s[i+1:]
		}
	}

	return m, nil
}

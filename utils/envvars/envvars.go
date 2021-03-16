package envvars

import (
	"os"
)

// All populates a map of environmental variables
func All(m map[string]interface{}) {
	all(os.Environ(), m)
}

func all(envs []string, m map[string]interface{}) {
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
}

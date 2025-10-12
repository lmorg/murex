package envvars

import (
	"os"
)

// All populates a map of environmental variables
func All(m map[string]any) {
	all(os.Environ(), m)
}

func all(envs []string, m map[string]any) {
	for _, env := range envs {
		key, val := Split(env)
		m[key] = val
	}
}

func Split(env string) (string, string) {
	if env == "" || env == "=" {
		return "", ""
	}

	var i int
	for ; i < len(env); i++ {
		if env[i] == '=' {
			break
		}
	}

	switch i {
	case 0:
		return "", env[1:]
	case len(env):
		return env, ""
	default:
		return env[:i], env[i+1:]
	}
}

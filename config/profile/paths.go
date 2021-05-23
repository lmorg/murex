package profile

import (
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
)

const (
	preloadFileName = ".murex_preload"
	moduleDirName   = ".murex_modules/"
	profileFileName = ".murex_profile"

	PreloadEnvVar = "MUREX_PRELOAD"
	ModuleEnvVar  = "MUREX_MODULES"
	ProfileEnvVar = "MUREX_PROFILE"
)

func PreloadPath() string {
	return validateProfilePath(PreloadEnvVar, preloadFileName)
}

func ProfilePath() string {
	return validateProfilePath(ProfileEnvVar, profileFileName)
}

func validateProfilePath(envvar, defaultFileName string) string {
	path := os.Getenv(envvar)
	if strings.TrimSpace(path) == "" {
		return home.MyDir + consts.PathSlash + defaultFileName
	}

	fi, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Override path specified in %s does not exist: '%s'!\n  Assuming this is intentional, a new file will be created.\n",
			envvar, path)
		return path
	}

	if fi.IsDir() {
		path += consts.PathSlash + defaultFileName
	}

	return path
}

func ModulePath() string {
	path := os.Getenv(ModuleEnvVar)
	if strings.TrimSpace(path) == "" {
		return home.MyDir + consts.PathSlash + moduleDirName
	}

	fi, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Override path specified in %s does not exist: '%s'!\n  Assuming this is intentional, a new directory will be created.\n",
			ModuleEnvVar, path)
		return addTrailingSlash(path)
	}

	if !fi.IsDir() {
		fmt.Fprintf(os.Stderr,
			"Override path specified in %s is a file, directory expected: '%s'!\n  Falling back to default path.\n",
			ModuleEnvVar, path)
		return home.MyDir + consts.PathSlash + moduleDirName
	}

	return addTrailingSlash(path)
}

func addTrailingSlash(path string) string {
	if strings.HasSuffix(path, consts.PathSlash) {
		return path
	}
	return path + consts.PathSlash
}

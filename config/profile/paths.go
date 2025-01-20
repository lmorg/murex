package profile

import (
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
)

const (
	// default locations
	preloadFileName = ".murex_preload"
	moduleDirName   = ".murex_modules/"
	profileFileName = ".murex_profile"
	fileNameCrop    = len(".murex_")

	// PreloadEnvVar environmental variable name for the preload profile override path. It's stored as a constant so typos are caught by the compiler
	PreloadEnvVar = "MUREX_PRELOAD"

	// ModuleEnvVar environmental variable name for the module override path. It's stored as a constant so typos are caught by the compiler
	ModuleEnvVar = "MUREX_MODULES"

	// ProfileEnvVar environmental variable name for the profile override path. It's stored as a constant so typos are caught by the compiler
	ProfileEnvVar = "MUREX_PROFILE"

	ConfigEnvVar = "MUREX_CONFIG_DIR"
)

// PreloadPath returns the path of the preload profile
func PreloadPath() string {
	return ValidateProfilePath(PreloadEnvVar, preloadFileName, false)
}

// ProfilePath returns the path of your murex profile
func ProfilePath() string {
	return ValidateProfilePath(ProfileEnvVar, profileFileName, false)
}

func ValidateProfilePath(envvar, defaultFileName string, isDir bool) string {
	path := os.Getenv(envvar)
	if strings.TrimSpace(path) != "" {
		return _validateProfilePathWithProfileEnvVar(envvar, defaultFileName, path, isDir)
	}

	path = os.Getenv(ConfigEnvVar)
	if path != "" {
		return _validateProfilePathWithConfigEnvVar(defaultFileName, path)
	}

	return _validateProfilePathWithHome(defaultFileName)
}

func _validateProfilePathWithProfileEnvVar(envvar, defaultFileName, path string, isDir bool) string {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Override path specified in %s does not exist: '%s'!\n  Assuming this is intentional, a new file will be created.\n",
			envvar, path)
		return path
	}

	if fi.IsDir() && !isDir {
		path += consts.PathSlash + defaultFileName
	} else if isDir {
		path += consts.PathSlash
	}

	return path
}

func _validateProfilePathWithConfigEnvVar(defaultFileName, path string) string {
	fi, err := os.Stat(path)
	pathFile := path + consts.PathSlash + defaultFileName[fileNameCrop:]
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Override path specified in %s does not exist: '%s'!\n  Assuming this is intentional, a new directory will be created.\n",
			ConfigEnvVar, path)

		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			home := _validateProfilePathWithHome(defaultFileName)
			fmt.Fprintf(os.Stderr, "!!! ERROR: %s\n!!! This path cannot be used so defaulting to $HOME: %s\n",
				err.Error(), home)
			return home
		}

		return pathFile
	}

	if fi.IsDir() {
		return pathFile
	}

	home := _validateProfilePathWithHome(defaultFileName)
	fmt.Fprintf(os.Stderr,
		"Override path specified in %s exists and is not a directory: '%s'!\n  This path cannot be used so defaulting to $HOME: %s\n",
		ConfigEnvVar, path, home)

	return home
}

func _validateProfilePathWithHome(defaultFileName string) string {
	return home.MyDir + consts.PathSlash + defaultFileName
}

// ModulePath returns the install path of the murex modules / packages
func ModulePath() string {
	path := ValidateProfilePath(ModuleEnvVar, moduleDirName, true)

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
		return _validateProfilePathWithHome(moduleDirName)
	}

	return addTrailingSlash(path)
}

func addTrailingSlash(path string) string {
	if strings.HasSuffix(path, consts.PathSlash) {
		return path
	}
	return path + consts.PathSlash
}

package lang

import (
	"strings"

	"github.com/lmorg/murex/lang/types"
)

// GetExtType gets the murex data type for a corresponding file extension
func GetExtType(extension string) (dt string) {
	dt = fileExts[strings.ToLower(extension)]
	if dt == "" {
		return types.Generic
	}
	return
}

// GetMimes returns MIME lookup table
func GetMimes() map[string]string {
	return mimes
}

// ReadMimes returns an any of mimes.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadMimes() (any, error) {
	return mimes, nil
}

// GetFileExts returns the file extension lookup table
func GetFileExts() map[string]string {
	return fileExts
}

// ReadFileExtensions returns an any of fileExts.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadFileExtensions() (any, error) {
	return fileExts, nil
}

// GetDefaultMimes returns default MIME lookup table
func GetDefaultMimes() map[string]string {
	return defaultMimes
}

// ReadDefaultMimes returns an any of default MIMEs.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadDefaultMimes() (any, error) {
	return defaultMimes, nil
}

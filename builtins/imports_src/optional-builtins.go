//go:build windows
// +build windows

package imports

// This is where you can add or remove built in functions. Imports here require underscoring.
// Each builtin package should include a README.md with details about its use. However the code should also be readable
// so take a look through the .go files if you're still undecided about whether to include a builtin or not.
//
// These are all optional functions
import (
	_ "github.com/lmorg/murex/builtins/optional/encoders" // base64, file archives, etc
	_ "github.com/lmorg/murex/builtins/optional/time"     // sleep
)

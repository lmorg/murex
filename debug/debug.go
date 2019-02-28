package debug

import (
	"encoding/json"
	"log"
)

// Enabled is a flag used for debugging murex code. This can be enabled at
// startup by a `--debug` flag or during runtime with `debug on`.
var Enabled bool

// Inspect is a flag used for debugging this Go source code. It can only be
// enabled at startup by `--inspect`. This is because it breaks a few security
// features of the shell by allowing `runtime` to query data outside of your
// normal scope.
var Inspect bool

// Log writes a debug message
func Log(data ...interface{}) {
	if Enabled {
		log.Println(data...)
	}
}

// Json encode an object then write it as a debug message
func Json(context string, data interface{}) {
	if Enabled {
		b, _ := json.MarshalIndent(data, "", "\t")
		Log(context, "JSON:"+string(b))
	}
}

// Dump is used for runtime output of the status of various debug modes
func Dump() interface{} {
	type status struct {
		Debug   bool
		Inspect bool
	}

	return status{
		Debug:   Enabled,
		Inspect: Inspect,
	}
}

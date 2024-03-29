package debug

import (
	"encoding/json"
	"log"
)

// Enabled is a flag used for debugging murex code. This can be enabled at
// startup by a `--debug` flag or during runtime with `debug on`.
var Enabled bool

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
		Debug bool
	}

	return status{
		Debug: Enabled,
	}
}

package debug

import (
	"encoding/json"
	"log"
)

// Enabled is a flag used for debugging information - enabled at startup by a
// `--debug` flag or during runtime with `debug on`
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

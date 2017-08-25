package debug

import (
	"encoding/json"
	"log"
)

// Enable debugging information - enabled at runtime by a flag or `debug on`
var Enable bool

// Log writes a debug message
func Log(data ...interface{}) {
	if Enable {
		log.Println(data...)
	}
}

// Json encode an object then write it as a debug message
func Json(context string, data interface{}) {
	if Enable {
		b, _ := json.MarshalIndent(data, "", "\t")
		Log(context, "JSON:"+string(b))
	}
}

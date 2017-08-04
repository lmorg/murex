package debug

import (
	"encoding/json"
	"log"
)

// Debugging information - enabled at runtime by a flag or `debug on`
var Enable bool

// Write a debug message
func Log(data ...interface{}) {
	if Enable {
		log.Println(data...)
	}
}

// JSON encode an object then write it as a debug message
func Json(context string, data interface{}) {
	if Enable {
		b, _ := json.MarshalIndent(data, "", "\t")
		Log(context, "JSON:"+string(b))
	}
}

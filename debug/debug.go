package debug

import (
	"encoding/json"
	"log"
)

var Enable bool

func Log(data ...interface{}) {
	if Enable {
		log.Println(data...)
	}
}

func Json(context string, data interface{}) {
	if Enable {
		b, _ := json.MarshalIndent(data, "", "\t")
		Log(context, "JSON:"+string(b))
	}
}

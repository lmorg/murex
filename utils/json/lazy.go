package json

import "encoding/json"

func LazyLogging(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func LazyLoggingPretty(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "    ")
	return string(b)
}

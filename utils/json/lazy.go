package json

import "encoding/json"

func LazyLogging(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func LazyLoggingPretty(v any) string {
	b, _ := json.MarshalIndent(v, "", "    ")
	return string(b)
}

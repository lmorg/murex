package types

// DataTypeFromInterface returns the Murex data-type expected to be associated
// with any specific Go data type.
func DataTypeFromInterface(v interface{}) (dataType string) {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		dataType = Integer

	case float64, float32:
		dataType = Number

	//case string, []byte, []rune:
	//	dataType = String

	case bool:
		dataType = Boolean

	//case []string, []int, []float32, []float64,
	//	map[string]string, map[interface{}]string, map[string]interface{}, map[interface{}]interface{}:
	//	dataType = Json

	default:
		dataType = Generic
	}

	return
}

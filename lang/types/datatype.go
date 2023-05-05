package types

// DataTypeFromInterface returns the Murex data-type expected to be associated
// with any specific Go data type.
func DataTypeFromInterface(v interface{}) string {
	switch v.(type) {
	case int:
		return Integer

	case float64:
		return Number

	case string, []byte, []rune:
		return String

	case bool:
		return Boolean

	case nil:
		return Null

	default:
		return Generic
	}
}

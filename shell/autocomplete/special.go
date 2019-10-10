package autocomplete

// isSpecialBuiltin identifies special builtins
func isSpecialBuiltin(s string) bool {
	switch s {
	case ">", ">>", "[", "[[", "@[", "=":
		return true
	default:
		return false
	}
}

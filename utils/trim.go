package utils

// CrLfTrim removes the trailing carridge return and line feed from a byte array
// however it is only concerned with one instance (ie you can still append a
// CRLF to the data if you have two instances of a \r\n terminator).
func CrLfTrim(b []byte) []byte {
	trim := make([]byte, len(b))
	copy(trim, b)

	if len(trim) > 0 && trim[len(trim)-1] == '\n' {
		trim = trim[:len(trim)-1]
	}

	if len(trim) > 0 && trim[len(trim)-1] == '\r' {
		trim = trim[:len(trim)-1]
	}

	return trim
}

// CrLfTrimRune removes the trailing carridge return and line feed from a rune
// array however it is only concerned with one instance (ie you can still append
// a CRLF to the data if you have two instances of a \r\n terminator).
func CrLfTrimRune(r []rune) []rune {
	trim := make([]rune, len(r))
	copy(trim, r)

	if len(trim) > 0 && trim[len(trim)-1] == '\n' {
		trim = trim[:len(trim)-1]
	}

	if len(trim) > 0 && trim[len(trim)-1] == '\r' {
		trim = trim[:len(trim)-1]
	}

	return trim
}

// CrLfTrimString removes the trailing carridge return and line feed from a
// string however it is only concerned with one instance (ie you can still
// append a CRLF to the data if you have two instances of a \r\n terminator).
func CrLfTrimString(s string) string {
	if len(s) > 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}

	if len(s) > 0 && s[len(s)-1] == '\r' {
		s = s[:len(s)-1]
	}

	return s
}

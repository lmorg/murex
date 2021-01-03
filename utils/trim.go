package utils

// CrLfTrim removes the trailing carridge return and line feed from a byte array
// however it is only concerned with one instance (ie you can still append a
// CRLF to the data if you have two instances of a \r\n terminator).
// CrLfTrim creates a copy of the byte slice.
func CrLfTrim(b []byte) []byte {
	i := len(b)

	if i != 0 && b[i-1] == '\n' {
		i--
	}

	if i != 0 && b[i-1] == '\r' {
		i--
	}

	if i == 0 {
		return []byte{}
	}

	trim := make([]byte, i)
	copy(trim, b[:i])

	return trim
}

// CrLfTrimRune removes the trailing carridge return and line feed from a rune
// array however it is only concerned with one instance (ie you can still append
// a CRLF to the data if you have two instances of a \r\n terminator).
// CrLfTrimRune creates a copy of the rune slice.
func CrLfTrimRune(r []rune) []rune {
	i := len(r)

	if i != 0 && r[i-1] == '\n' {
		i--
	}

	if i != 0 && r[i-1] == '\r' {
		i--
	}

	if i == 0 {
		return []rune{}
	}

	trim := make([]rune, i)
	copy(trim, r[:i])

	return trim
}

// CrLfTrimString removes the trailing carridge return and line feed from a
// string however it is only concerned with one instance (ie you can still
// append a CRLF to the data if you have two instances of a \r\n terminator).
func CrLfTrimString(s string) string {
	i := len(s)

	if i != 0 && s[i-1] == '\n' {
		i--
	}

	if i != 0 && s[i-1] == '\r' {
		i--
	}

	if i == 0 {
		return ""
	}

	return s[:i]
}

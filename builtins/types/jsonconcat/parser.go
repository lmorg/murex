package jsonconcat

import "fmt"

func parse(b []byte, callback func([]byte)) error {
	var (
		brace, last   int
		quote, escape bool
		open, close   byte
	)

	if len(b) == 0 {
		return nil
	}

	switch b[0] {
	case '{':
		open, close = '{', '}'
	case '[':
		open, close = '[', ']'
	default:
		return fmt.Errorf("invalid first character '%s'. This doesn't appear to be valid JSON", string([]byte{b[0]}))
	}

	for i := range b {
		if escape {
			escape = false
			continue
		}

		switch b[i] {
		case '\\':
			escape = !escape

		case '"':
			quote = !quote

		case open:
			brace++

		case close:
			brace--
			if brace == 0 {
				json := make([]byte, i-last+1)
				copy(json, b[last:i+1])
				callback(json)
				last = i + 1
			}
		}
	}

	if brace > 0 {
		return fmt.Errorf("reached end of document with %d missing `%s`", brace, string([]byte{close}))
	}

	return nil
}

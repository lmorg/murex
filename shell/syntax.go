package shell

import (
	"github.com/lmorg/murex/utils/inject"
	"github.com/lmorg/murex/utils/parser"
)

func syntaxCompletion(line []rune, change string, pos int) ([]rune, int) {
	if pos < 0 || len(line) < pos {
		return line, pos
	}

	var part parser.ParsedTokens
	full, _ := parse(line)
	if pos >= len(line)-1 {
		part = full
	} else {
		part, _ = parse(line[:pos+1])
	}

	if part.Escaped || part.Comment {
		return line, pos
	}

	var previousRune rune
	if pos > 0 {
		previousRune = line[pos-1] // character before current one
	}

	var nextRune rune
	if pos < len(line)-1 {
		nextRune = line[pos+1] // character after current one
	}

	if len(line) > 1 && pos > 0 && previousRune == '\\' {
		return line, pos
	}

	posEOL := pos == len(line)-1

	switch change {
	case "'":
		switch {
		case (part.NestedBlock > 0 || part.SquareBracket) && !posEOL:
			new, err := inject.Rune(line, []rune{'\''}, pos)
			if err != nil {
				return line, pos
			}
			return new, pos
		case part.QuoteDouble || part.QuoteBrace > 0:
			return line, pos // do nothing because QuoteSingle might be an apostrophe
		case !part.QuoteSingle && full.QuoteSingle && !posEOL && line[pos+1] == '\'':
			return append(line[:pos], line[pos+1:]...), pos
		case !part.QuoteSingle && full.QuoteSingle && full.LastCharacter == '\'':
			return line[:len(line)-1], pos
		case !part.QuoteSingle && !full.QuoteSingle:
			return line, pos
		case posEOL:
			return append(line, '\''), pos
		case part.QuoteSingle && full.LastCharacter == '\'':
			return line[:len(line)-1], pos
		default:
			return append(line, '\''), pos
		}

	case "\"":
		switch {
		case (part.NestedBlock > 0 || part.SquareBracket) && !posEOL:
			new, err := inject.Rune(line, []rune{'"'}, pos)
			if err != nil {
				return line, pos
			}
			return new, pos
		case part.QuoteSingle || part.QuoteBrace > 0:
			//	return line, pos
			new, err := inject.Rune(line, []rune{'"'}, pos)
			if err != nil {
				return line, pos
			}
			return new, pos
		case !part.QuoteDouble && full.QuoteDouble && pos < len(line) && line[pos+1] == '"':
			return append(line[:pos], line[pos+1:]...), pos
		case !part.QuoteDouble && full.QuoteDouble && full.LastCharacter == '"':
			return line[:len(line)-1], pos
		case !part.QuoteDouble && !full.QuoteDouble:
			return line, pos //
		case posEOL:
			return append(line, '"'), pos
		case part.QuoteDouble && full.LastCharacter == '"':
			return line[:len(line)-1], pos
		default:
			return append(line, '"'), pos
		}

	case "(":
		switch {
		case part.SquareBracket || part.NestedBlock > 0 ||
			full.SquareBracket || full.NestedBlock > 0:
			new, err := inject.Rune(line, []rune{')'}, pos+1)
			if err != nil {
				return line, pos
			}
			return new, pos
		case part.QuoteSingle || part.QuoteDouble:
			return line, pos
		case full.QuoteBrace == 1 && part.QuoteBrace == 1:
			return append(line, ')'), pos
		case part.QuoteBrace > 0:
			new, err := inject.Rune(line, []rune{')'}, pos+1)
			if err != nil {
				return line, pos
			}
			return new, pos
		case posEOL:
			return append(line, ')'), pos
		default:
			return line, pos
		}

	case ")":
		if full.QuoteBrace < 0 && part.QuoteBrace == 0 && full.LastCharacter == ')' {
			return line[:len(line)-1], pos
		}

	case "{":
		switch {
		case part.SquareBracket:
			new, err := inject.Rune(line, []rune{'}'}, pos+1)
			if err != nil {
				return line, pos
			}
			return new, pos
		case part.QuoteSingle || part.QuoteDouble:
			return line, pos
		case part.QuoteBrace > 0:
			new, err := inject.Rune(line, []rune{'}'}, pos+1)
			if err != nil {
				return line, pos
			}
			return new, pos
		case full.NestedBlock == 1 && part.NestedBlock == 1:
			return append(line, '}'), pos
		case part.NestedBlock > 0:
			new, err := inject.Rune(line, []rune{'}'}, pos+1)
			if err != nil {
				return line, pos
			}
			return new, pos
		case posEOL:
			return append(line, '}'), pos
		default:
			return line, pos
		}

	case "}":
		if full.NestedBlock < 0 && part.NestedBlock == 0 && full.LastCharacter == '}' {
			return line[:len(line)-1], pos
		}

	case "[":
		switch {
		case part.QuoteSingle || part.QuoteDouble || part.QuoteBrace > 0 || part.NestedBlock > 0:
			new, err := inject.Rune(line, []rune{']'}, pos+1)
			if err != nil {
				return line, pos
			}
			return new, pos
		case full.SquareBracket && full.NestedBlock == 0 && full.LastCharacter == '[':
			return append(line, ']'), pos
		case full.FuncName == "[[]" && change == "[" && line[pos+1] == ']':
			newLine := append(line[:pos+1], ' ', ' ', ']', ']')
			newLine = append(newLine, line[pos+1:]...)
			return newLine, pos + 1
		}

	case "]":
		if part.SquareBracket && full.SquareBracket && nextRune == ']' {
			return append(line[:pos], line[pos+1:]...), pos
		}
	}

	return line, pos
}

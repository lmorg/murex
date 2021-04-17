package shell

import (
	"github.com/lmorg/murex/utils/inject"
	"github.com/lmorg/murex/utils/parser"
)

func syntaxCompletion(line []rune, change string, pos int) ([]rune, int) {
	// This is lazy I know, but it's faster and less error prone than checking
	// the size of every array. Plus produces more readable code.
	/*defer func() {
		if debug.Enabled {
			return
		}
		if r := recover(); r != nil {
			newLine = line
			newPos = pos
		}
	}()*/

	var part parser.ParsedTokens
	full, _ := parse(line)
	if pos >= len(line) {
		part = full
	} else {
		part, _ = parse(line[:pos+1])
	}

	if len(line) > 1 && pos > 0 && line[pos-1] == '\\' {
		return line, pos
	}

	posEOL := pos == len(line)-1

	switch change {
	case "'":
		switch {
		case part.QuoteDouble || part.QuoteBrace > 0:
			return line, pos
		case posEOL:
			return append(line, '\''), pos
		case part.QuoteSingle && full.LastCharacter == '\'':
			return line[:len(line)-1], pos
		default:
			//new, err:=inject.Rune(line,[]rune{'\''},pos)
			return append(line, '\''), pos
		}

	case "\"":
		switch {
		case part.QuoteSingle || part.QuoteBrace > 0:
			return line, pos
		case posEOL:
			return append(line, '"'), pos
		case part.QuoteDouble && full.LastCharacter == '"':
			return line[:len(line)-1], pos
		default:
			//new, err:=inject.Rune(line,[]rune{'"'},pos)
			return append(line, '"'), pos
		}

	case "(":
		switch {
		case part.SquareBracket || part.NestedBlock > 0 ||
			full.SquareBracket || full.NestedBlock > 0:
			new, err := inject.Rune(line, []rune{')'}, pos+1)
			if err != nil {
				return line, pos + 1
			}
			return new, pos
		case part.QuoteSingle || part.QuoteDouble:
			return line, pos
		case full.QuoteBrace == 1 && part.QuoteBrace == 1:
			return append(line, ')'), pos
		case part.QuoteBrace > 0:
			new, err := inject.Rune(line, []rune{')'}, pos+1)
			if err != nil {
				return line, pos + 1
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
				return line, pos + 1
			}
			return new, pos
		case part.QuoteSingle || part.QuoteDouble:
			return line, pos
		case part.QuoteBrace > 0:
			new, err := inject.Rune(line, []rune{'}'}, pos+1)
			if err != nil {
				return line, pos + 1
			}
			return new, pos
		case full.NestedBlock == 1 && part.NestedBlock == 1:
			return append(line, '}'), pos
		case part.NestedBlock > 0:
			new, err := inject.Rune(line, []rune{'}'}, pos+1)
			if err != nil {
				return line, pos + 1
			}
			return new, pos
		case posEOL:
			return append(line, '}'), pos
		default:
			return line, pos
		}

	case "[":
		switch {
		case part.QuoteSingle || part.QuoteDouble || part.QuoteBrace > 0 || part.NestedBlock > 0:
			new, err := inject.Rune(line, []rune{']'}, pos+1)
			if err != nil {
				return line, pos + 1
			}
			return new, pos
		case full.SquareBracket && full.NestedBlock == 0 && full.LastCharacter == '[':
			return append(line, ']'), pos
		case full.FuncName == "[[]" && change == "[" && line[pos+1] == ']':
			newLine := append(line[:pos+1], ' ', ' ', ']', ']')
			newLine = append(newLine, line[pos+1:]...)
			return newLine, pos + 1
		}
	}

	return line, pos
}

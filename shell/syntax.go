package shell

import (
	"github.com/lmorg/murex/debug"
)

func syntaxCompletion(line []rune, pos int) (newLine []rune, newPos int) {
	// This is lazy I know, but it's faster and less error prone than checking
	// the size of every array. Plus produces more readable code.
	defer func() {
		if debug.Enabled {
			return
		}
		if r := recover(); r != nil {
			newLine = line
			newPos = pos
		}
	}()

	pt, _ := parse(line)

	switch {
	case pt.QuoteSingle && pt.QuoteBrace == 0 && pt.NestedBlock == 0:
		if pos < len(line)-1 && line[pos] == '\'' && line[len(line)-1] == '\'' {
			return line[:len(line)-1], pos
		}
		if pos < len(line)-1 || line[pos] != '\'' {
			return append(line, '\''), pos
		}

	case pt.QuoteDouble && pt.QuoteBrace == 0:
		if pt.NestedBlock == 0 {
			if pos < len(line)-1 && line[pos] == '"' && line[len(line)-1] == '"' {
				return line[:len(line)-1], pos
			}
			if pos < len(line)-1 || line[pos] != '"' {
				return append(line, '"'), pos
			}
		} else {
			// dont do anything
			return line, pos
		}

	case pt.QuoteBrace > 0 && pt.NestedBlock == 0:
		if pos < len(line)-1 || line[pos] != '(' {
			return append(line, ')'), pos
		}

	case pt.QuoteBrace < 0:
		if line[pos] == ')' && line[len(line)-1] == ')' && pos != len(line)-1 {
			return line[:len(line)-1], pos
		}

	case pt.NestedBlock > 0 && pt.QuoteBrace == 0:
		if (pos < len(line)-1 || line[pos] != '{') && !pt.SquareBracket {
			return append(line, '}'), pos
		}

	case pt.NestedBlock < 0:
		if line[pos] == '}' && line[len(line)-1] == '}' && pos != len(line)-1 {
			return line[:len(line)-1], pos
		}

	case pt.SquareBracket && pt.NestedBlock == 0 && pt.LastCharacter == '[':
		return append(line, ']'), pos

	}

	return line, pos
}

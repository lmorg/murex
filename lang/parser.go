package lang

import (
	"github.com/lmorg/murex/lang/proc/parameters"
)

func genEmptyParamTokens() (pt [][]parameters.ParamToken) {
	pt = make([][]parameters.ParamToken, 1)
	pt[0] = make([]parameters.ParamToken, 1)
	return
}

// ParseBlock parses a murex code block.
// Returns the abstract syntax tree (astNodes) or any syntax errors preventing
// a successful parse (ParserError)
func ParseBlock(block []rune) (nodes astNodes, pErr ParserError) {
	return AstCache.ParseCache(block)
}

func parser(block []rune) (nodes astNodes, pErr ParserError) {
	//defer debug.Json("Parser", nodes)

	var (
		// Current state
		lineNumber       int
		colNumber        int
		last             rune
		commentLine      bool
		escaped          bool
		quoteSingle      bool
		quoteDouble      bool
		quoteBrace       int
		braceCount       int
		unclosedIndex    bool
		ignoreWhitespace = true
		scanFuncName     = true

		// Parsed thus far
		node   = astNode{NewChain: true, ParamTokens: genEmptyParamTokens()}
		pop    = &node.Name
		pCount int // parameter count
		pToken = &node.ParamTokens[0][0]
	)
	//defer debug.Json("Last node", node)

	startParameters := func() {
		scanFuncName = false
		node.ParamTokens = genEmptyParamTokens()
		pop = &node.ParamTokens[0][0].Key
		pCount = 0
		pToken = &node.ParamTokens[pCount][0]
	}

	appendNode := func() {
		if len(node.ParamTokens) > 1 && len(node.ParamTokens[len(node.ParamTokens)-1]) == 0 {
			node.ParamTokens = node.ParamTokens[:len(node.ParamTokens)-1]
		}

		if node.Name != "" {
			nodes = append(nodes, node)
			// The below code is faster but much less readable and parsed
			// source should be cached so we don't really see a performance
			// improvement in the benchmarks thus cannot justify the code
			// complexity this brings. However more complex murex scripts with
			// fewer iteration blocks (ie fewer cacheable blocks) might see
			// a subtle performance improvement but the one might argue such
			// a shell script was poorly written from the outset. For now, I
			// shall leave this code here as an example of where not to
			// optimise so a future maintainer doesn't get gunghoe.
			/*m := len(nodes)
			//n := m + len(data)
			n := m + 1
			if n > cap(nodes) { // if necessary, reallocate
				// allocate double what's needed, for future growth.
				newSlice := make([]astNode, (m+2)*2)
				copy(newSlice, nodes)
				nodes = newSlice
			}
			nodes = nodes[0:n]
			//copy(nodes[m:n], node)
			nodes[m] = node*/
		}

		ignoreWhitespace = true
	}

	pUpdate := func(r rune) {
		if !scanFuncName && pToken.Type == parameters.TokenTypeNil {

			if r == '<' && last != '\\' &&
				!quoteSingle && !quoteDouble && quoteBrace == 0 {
				pToken.Type = parameters.TokenTypeNamedPipe
			} else {
				pToken.Type = parameters.TokenTypeValue
			}
		}

		if node.Name == "" {
			node.LineNumber = lineNumber
			node.ColNumber = colNumber
		}

		*pop += string(r)
	}

	for i, r := range block {
		colNumber++

		if commentLine {
			if r == '\n' {
				commentLine = false
				lineNumber++
				colNumber = 0
			}
			continue
		}

		if pToken.Type > parameters.TokenTypeValue {
			switch {
			case pToken.Type == parameters.TokenTypeIndex:
				if r != ']' {
					*pop += string(r)
					last = r
					continue
				}
				*pop += string(r)
				node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{})
				pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				pop = &pToken.Key
				unclosedIndex = false
				continue

			case pToken.Type == parameters.TokenTypeRange:
				if unclosedIndex {
					*pop += string(r)
					last = r
					continue
				}
				if r == ']' {
					unclosedIndex = false
					last = r
					*pop += string(r)
					continue
				}
				if !unclosedIndex && 'a' <= r && r <= 'z' {
					last = r
					*pop += string(r)
					continue
				}
				node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
				pCount++
				pToken = &node.ParamTokens[pCount][0]
				pop = &pToken.Key
				goto nextParser

			case pToken.Type == parameters.TokenTypeTilde &&
				(r == '_' || r == '-' || r == '.' ||
					('a' <= r && r <= 'z') ||
					('A' <= r && r <= 'Z') ||
					('0' <= r && r <= '9')):
				*pop += string(r)
				last = r
				continue

			case r == '_' ||
				('a' <= r && r <= 'z') ||
				('A' <= r && r <= 'Z') ||
				('0' <= r && r <= '9'):
				*pop += string(r)
				last = r
				continue

			case r == '}':
				braceCount--
				switch {
				case braceCount > 0:
					*pop += string(r)
				case pToken.Type == parameters.TokenTypeBlockString:
					node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{})
					pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
					pop = &pToken.Key
				case pToken.Type == parameters.TokenTypeBlockArray:
					node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
					pCount++
					pToken = &node.ParamTokens[pCount][0]
					pop = &pToken.Key
				default:
					*pop += string(r)
				}
				continue

			case r == '{' && last == '$':
				pToken.Type = parameters.TokenTypeBlockString
				braceCount++
				continue

			case r == '{' && last == '@':
				pToken.Type = parameters.TokenTypeBlockArray
				braceCount++
				continue

			case r == '{':
				braceCount++
				*pop += string(r)
				continue

			case r == '[' && pToken.Type == parameters.TokenTypeString && last != '$':
				pToken.Type = parameters.TokenTypeIndex
				*pop += string(r)
				last = r
				unclosedIndex = true
				continue

			case r == '[' && pToken.Type == parameters.TokenTypeArray && last != '@':
				pToken.Type = parameters.TokenTypeRange
				*pop += string(r)
				last = r
				unclosedIndex = true
				continue

			case braceCount > 0:
				*pop += string(r)
				continue

			default:
				if pToken.Type == parameters.TokenTypeString || pToken.Type == parameters.TokenTypeTilde {
					node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{})
					pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
					pop = &pToken.Key
				} else {
					node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
					pCount++
					pToken = &node.ParamTokens[pCount][0]
					pop = &pToken.Key
				}
			}
		}

	nextParser:
		switch r {
		case '#':
			switch {
			case escaped, quoteSingle, quoteDouble, quoteBrace > 0, braceCount > 0:
				pUpdate(r)
			default:
				commentLine = true
			}

		case '\\':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case quoteSingle:
				pUpdate(r)
			case quoteBrace > 0:
				pUpdate(r)
			case escaped:
				pUpdate(r)
				escaped = false
			default:
				escaped = true
			}

		case '\'':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle:
				quoteSingle = false
			case quoteDouble, quoteBrace > 0:
				pUpdate(r)
			default:
				pToken.Type = parameters.TokenTypeValue
				quoteSingle = true
			}

		case '"':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteBrace > 0:
				pUpdate(r)
			case quoteDouble:
				quoteDouble = false
			default:
				pToken.Type = parameters.TokenTypeValue
				quoteDouble = true
			}

		case '`':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			default:
				pUpdate('\'')
			}

		case '(':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case quoteBrace > 0:
				pUpdate(r)
				quoteBrace++
			case scanFuncName:
				pUpdate(r)
				startParameters()
				quoteBrace++
			default:
				quoteBrace++
			}

		case ')':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case quoteBrace == 0:
				pErr = raiseErr(ErrClosingBraceQuoteNoOpen, i)
				return
			case quoteBrace > 1:
				pUpdate(r)
				quoteBrace--
			case quoteBrace == 1:
				quoteBrace--
			default:
				pErr = raiseErr(ErrUnexpectedParsingError, i) //+" No case found for `switch ')' { ... }`.", i)
				return
			}

		case ':':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName:
				startParameters()
			default:
				pUpdate(r)
			}

		case '^':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName:
				pUpdate(r)
				startParameters()
			default:
				pUpdate(r)
			}

		case '=':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName:
				pUpdate(r)
				startParameters()
			default:
				pUpdate(r)
			}

		case '[':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName:
				pUpdate(r)
				if i < len(block) && block[i+1] != '[' {
					startParameters()
				}
			default:
				pUpdate(r)
			}

		case '{':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case scanFuncName:
				startParameters()
				pUpdate(r)
				braceCount++
			default:
				pUpdate(r)
				braceCount++
			}

		case '}':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case scanFuncName:
				pErr = raiseErr(ErrUnexpectedCloseBrace, i)
				return
			case braceCount == 0:
				pErr = raiseErr(ErrClosingBraceBlockNoOpen, i)
				return
			default:
				pUpdate(r)
				braceCount--
			}

		case ' ', '\t', '\r':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case !scanFuncName:
				node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
				pCount++
				pToken = &node.ParamTokens[pCount][0]
				pop = &pToken.Key
			case scanFuncName && !ignoreWhitespace:
				startParameters()
			default:
				// do nothing
			}

		case '\n':
			lineNumber++
			colNumber = 0
			switch {
			case escaped:
				//pUpdate(r)
				pUpdate(' ')
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName && ignoreWhitespace:
				// do nothing
			default:
				appendNode()
				node = astNode{NewChain: true}
				pop = &node.Name
				scanFuncName = true
			}

		case '|':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeToken, i)
				return
			default:
				node.PipeOut = true
				appendNode()
				node = astNode{}
				pop = &node.Name
				scanFuncName = true
			}

		case '?':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeToken, i)
				return
			case last == ' ' || last == '\t':
				node.PipeErr = true
				appendNode()
				node = astNode{}
				pop = &node.Name
				scanFuncName = true
			default:
				pUpdate(r)
			}

		case '<':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			default:
				pUpdate(r)
			}

		case '>':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case last == '-':
				*pop = (*pop)[:len(*pop)-1]
				if len(*pop) == 0 {
					pToken.Type = parameters.TokenTypeNil
				}
				node.PipeOut = true
				appendNode()
				node = astNode{Method: true}
				pop = &node.Name
				scanFuncName = true
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}

		case ';':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			default:
				appendNode()
				node = astNode{NewChain: true}
				pop = &node.Name
				scanFuncName = true
			}

		case '~':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case scanFuncName || braceCount > 0 || quoteSingle:
				pUpdate(r)
			default:
				node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{Type: parameters.TokenTypeTilde})
				pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				pop = &pToken.Key
			}

		case '$':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case scanFuncName || braceCount > 0 || quoteSingle:
				pUpdate(r)
			default:
				node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{Type: parameters.TokenTypeString})
				pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				pop = &pToken.Key
			}

		case '@':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case scanFuncName || braceCount > 0 || quoteSingle:
				pUpdate(r)
			case last != ' ' && last != '\t':
				pUpdate(r)
			default:
				node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
				pCount++
				pToken = &node.ParamTokens[pCount][0]
				pToken.Type = parameters.TokenTypeArray
				pop = &pToken.Key
			}

		case 's':
			switch {
			case escaped:
				pUpdate(' ')
				escaped = false
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}

		case 't':
			switch {
			case escaped:
				pUpdate('\t')
				escaped = false
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}

		case 'r':
			switch {
			case escaped:
				pUpdate('\r')
				escaped = false
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}

		case 'n':
			switch {
			case escaped:
				pUpdate('\n')
				escaped = false
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}

		default:
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}
		}

		last = r
	}

	switch {
	case unclosedIndex:
		return nil, raiseErr(ErrUnclosedIndex, 0)
	case escaped:
		return nil, raiseErr(ErrUnterminatedEscape, 0)
	case quoteSingle:
		return nil, raiseErr(ErrUnterminatedQuotesSingle, 0)
	case quoteDouble:
		return nil, raiseErr(ErrUnterminatedQuotesDouble, 0)
	case quoteBrace > 0:
		return nil, raiseErr(ErrUnterminatedBraceQuote, 0)
	case braceCount > 0:
		return nil, raiseErr(ErrUnterminatedBraceBlock, 0)
	}

	appendNode()

	return
}

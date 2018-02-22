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
// Returns the abstract syntax tree (astNodes) or any syntax errors preventing a successful parse (ParserError)
func ParseBlock(block []rune) (nodes astNodes, pErr ParserError) {
	//defer debug.Json("Parser", nodes)

	var (
		// Current state
		lineNumber               int
		colNumber                int
		last                     rune
		commentLine              bool
		escaped                  bool
		quoteSingle, quoteDouble bool
		braceCount               int
		unclosedIndex            bool
		ignoreWhitespace         bool = true
		scanFuncName             bool = true

		// Parsed thus far
		node   astNode                = astNode{NewChain: true, ParamTokens: genEmptyParamTokens()}
		pop    *string                = &node.Name
		pCount int                    // parameter count
		pToken *parameters.ParamToken = &node.ParamTokens[0][0]
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
		}

		ignoreWhitespace = true
	}

	pUpdate := func(r rune) {
		if !scanFuncName && pToken.Type == parameters.TokenTypeNil {
			pToken.Type = parameters.TokenTypeValue
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
			case pToken.Type == parameters.TokenTypeIndex && r != ']':
				*pop += string(r)
				last = r
				continue

			case pToken.Type == parameters.TokenTypeIndex && r == ']':
				*pop += string(r)
				//node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
				//pCount++
				//pToken = &node.ParamTokens[pCount][0]
				//pop = &pToken.Key
				node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{})
				pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				pop = &pToken.Key
				unclosedIndex = false
				continue

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

		switch r {
		case '#':
			switch {
			case escaped, quoteSingle, quoteDouble, braceCount > 0:
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
			case quoteDouble:
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
			case quoteSingle:
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
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			default:
				pUpdate('\'')
			}

		case ':':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
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
			case quoteSingle, quoteDouble:
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
			case quoteSingle, quoteDouble:
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
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName:
				pUpdate(r)
				startParameters()
			default:
				pUpdate(r)
			}

		case '{':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case scanFuncName:
				startParameters()
				//*pop += string(r)
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
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case scanFuncName:
				pErr = raiseErr(ErrUnexpectedCloseBrace, i)
				return
			case braceCount == 0:
				pErr = raiseErr(ErrClosingBraceNoOpen, i)
				return
			default:
				pUpdate(r)
				braceCount--
				/*if braceCount == 0 {
					appendNode()
					node = astNode{NewChain: true}
					pop = &node.Name
					scanFuncName = true
					//newLine = true
				}*/
			}

		case ' ', '\t', '\r':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case !scanFuncName:
				//if len(*pop) > 0 {
				//if pToken.Type != parameters.TokenTypeNil {
				node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
				pCount++
				pToken = &node.ParamTokens[pCount][0]
				pop = &pToken.Key
				//}
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
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case !scanFuncName:
				appendNode()
				node = astNode{NewChain: true}
				pop = &node.Name
				scanFuncName = true
				//newLine = true
			case scanFuncName && !ignoreWhitespace:
				startParameters()
			default:
				// do nothing
			}

		case '|':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeToken, i)
				return
			/*case newLine:
			newLine = false
			node.NewChain = false
			if len(nodes) > 0 {
				nodes.Last().PipeOut = true
			}*/
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
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeToken, i)
				return
			/*case newLine:
			newLine = false
			node.NewChain = false
			if len(nodes) > 0 {
				nodes.Last().PipeErr = true
			}*/
			case last == ' ' || last == '\t':
				node.PipeErr = true
				appendNode()
				node = astNode{}
				pop = &node.Name
				scanFuncName = true
			default:
				pUpdate(r)
			}

		case '>':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case last == '-':
				/*if len(node.Name) == 0 {
					pErr = raiseErr(ErrUnexpectedPipeToken, i)
					return
				}*/
				*pop = (*pop)[:len(*pop)-1]
				if len(*pop) == 0 {
					pToken.Type = parameters.TokenTypeNil
				}
				node.PipeOut = true
				appendNode()
				node = astNode{Method: true}
				pop = &node.Name
				scanFuncName = true

				/*if newLine {
					node.NewChain = false
					node.Method = true
					nodes.Last().PipeOut = true
					newLine = false
				}*/
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}

		case ';':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
				//case !scanFuncName:
			default:
				appendNode()
				node = astNode{NewChain: true}
				pop = &node.Name
				scanFuncName = true
				//default:
				// do nothing
			}

		case '~':
			if braceCount == 0 && !quoteSingle && !escaped {
				node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{Type: parameters.TokenTypeTilde})
				pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				pop = &pToken.Key
			} else {
				pUpdate(r)
			}

		case '$':
			if !scanFuncName && braceCount == 0 && !quoteSingle && !escaped {
				node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{Type: parameters.TokenTypeString})
				pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				pop = &pToken.Key
			} else {
				pUpdate(r)
			}

		case '@':
			if !scanFuncName && braceCount == 0 && !quoteSingle && !quoteDouble && !escaped && (last == ' ' || last == '\t') {
				node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
				pCount++
				pToken = &node.ParamTokens[pCount][0]
				pToken.Type = parameters.TokenTypeArray
				pop = &pToken.Key
			} else {
				pUpdate(r)
			}

		case 's':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case quoteSingle:
				pUpdate(r)
			case escaped:
				pUpdate(' ')
				escaped = false
			default:
				pUpdate(r)
			}

		case 't':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case quoteSingle:
				pUpdate(r)
			case escaped:
				pUpdate('\t')
				escaped = false
			default:
				pUpdate(r)
			}

		case 'r':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case quoteSingle:
				pUpdate(r)
			case escaped:
				pUpdate('\r')
				escaped = false
			default:
				pUpdate(r)
			}

		case 'n':
			switch {
			case braceCount > 0:
				pUpdate(r)
			case quoteSingle:
				pUpdate(r)
			case escaped:
				pUpdate('\n')
				escaped = false
			default:
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
				/*if b != '-' {
					newLine = false
				}*/
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
	case braceCount > 0:
		return nil, raiseErr(ErrUnterminatedBrace, 0)
	}

	appendNode()

	return
}

package lang

import (
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
)

func genEmptyParamTokens() (pt [][]parameters.ParamToken) {
	pt = make([][]parameters.ParamToken, 1)
	pt[0] = make([]parameters.ParamToken, 1)
	return
}

// DontCacheAst is an override for disabling the AST cache. This is enabled for
// some tests (particularly fuzz testing)
var DontCacheAst bool

// ParseBlock parses a murex code block.
// Returns the abstract syntax tree (AstNodes) or any syntax errors preventing
// a successful parse (ParserError)
func ParseBlock(block []rune) (nodes *AstNodes, pErr ParserError) {
	if DontCacheAst {
		return parser(block)
	}

	return AstCache.ParseCache(block)
}

func parser(block []rune) (*AstNodes, ParserError) {
	//defer debug.Json("Parser", nodes)

	var (
		nodes AstNodes
		pErr  ParserError

		// Current state
		i                int
		r                rune
		lineNumber       int
		colNumber        int
		last             rune
		commentLine      bool
		escaped          bool
		escapedBfComment bool
		quoteSingle      bool
		quoteDouble      bool
		quoteBrace       int
		quoteBraceHide   bool
		braceCount       int
		unclosedIndex    bool
		ignoreWhitespace = true
		scanFuncName     = true

		// Parsed thus far
		node   = AstNode{NewChain: true, ParamTokens: genEmptyParamTokens()}
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

	next := func(r rune) bool {
		if i+1 == len(block) {
			return false
		}

		if block[i+1] == r {
			return true
		}

		return false
	}

	nextAlphaNumeric := func() bool {
		if i+1 == len(block) {
			return false
		}

		if block[i+1] == '_' ||
			(block[i+1] >= 'a' && 'z' >= block[i+1]) ||
			(block[i+1] >= 'A' && 'Z' >= block[i+1]) ||
			(block[i+1] >= '0' && '9' >= block[i+1]) {
			return true
		}

		return false
	}

	for ; i < len(block); i++ {
		r = block[i]
		colNumber++

		// comment
		if commentLine {
			if r == '\n' {
				commentLine = false
				lineNumber++
				colNumber = 0
			}
			if escapedBfComment {
				escapedBfComment = false
				switch {
				case !scanFuncName && last != ' ' && last != ':':
					node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
					pCount++
					pToken = &node.ParamTokens[pCount][0]
					pop = &pToken.Key
				case scanFuncName && !ignoreWhitespace:
					startParameters()
				default:
					pUpdate(r)
				}
			}
			continue
		}

		// variable tokens
		if pToken.Type > parameters.TokenTypeValue {
			switch {
			case pToken.Type == parameters.TokenTypeIndex ||
				pToken.Type == parameters.TokenTypeElement ||
				pToken.Type == parameters.TokenTypeRange:

				*pop += string(r)

				if r != ']' {
					last = r
					continue
				}

				if next(']') {
					*pop += string(r)
					i++
				}

				node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{})
				pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				pop = &pToken.Key
				unclosedIndex = false
				continue

			case pToken.Type == parameters.TokenTypeRange:
				//if unclosedIndex {
				if r != ']' {
					*pop += string(r)
					last = r
					continue
				}
				/*if r == ']' {
					unclosedIndex = false
					last = r
					*pop += string(r)
					continue
				}*/
				/*if !unclosedIndex && 'a' <= r && r <= 'z' {
					last = r
					*pop += string(r)
					continue
				}*/
				node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
				pCount++
				pToken = &node.ParamTokens[pCount][0]
				pop = &pToken.Key
				//goto nextParser
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
				last = r
				continue

			/*case r == '(' && last == '$':
			pToken.Type = parameters.TokenTypeBlockString
			braceCount++
			continue*/

			case r == '{' && last == '@':
				pToken.Type = parameters.TokenTypeBlockArray
				braceCount++
				last = r
				continue

			case r == '{':
				braceCount++
				*pop += string(r)
				continue

			case r == '[' && pToken.Type == parameters.TokenTypeString && last != '$':
				if next('[') {
					pToken.Type = parameters.TokenTypeElement
					*pop += string(r)
					i++
				} else {
					pToken.Type = parameters.TokenTypeIndex
				}
				*pop += string(r)
				last = r
				unclosedIndex = true
				continue

			case r == '[' && pToken.Type == parameters.TokenTypeArray:
				//if last != '@' {
				pToken.Type = parameters.TokenTypeRange
				*pop += string(r)
				last = r
				unclosedIndex = true

				/*} else {
					pToken.Type = parameters.TokenTypeValue
					*pop += "@["
				}*/
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

		//nextParser:
		switch r {
		case '#':
			switch {
			case escaped:
				pUpdate(r)
				ignoreWhitespace = false
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0, braceCount > 0:
				pUpdate(r)
				ignoreWhitespace = false
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
				if !scanFuncName {
					pToken.Type = parameters.TokenTypeValue
				}
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
				if !scanFuncName {
					pToken.Type = parameters.TokenTypeValue
				}
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
				ignoreWhitespace = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
				ignoreWhitespace = false
			case quoteBrace > 0:
				pUpdate(r)
				quoteBrace++
			case scanFuncName && len(*pop) == 0:
				pUpdate(r)
				startParameters()
				quoteBrace++
			case len(*pop) > 0:
				pUpdate(r)
				quoteBrace++
				quoteBraceHide = true
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
				ignoreWhitespace = false
			case quoteSingle, quoteDouble:
				pUpdate(r)
				ignoreWhitespace = false
			/*case quoteBrace == 0:
			pErr = raiseErr(ErrClosingBraceQuoteNoOpen, i)
			return*/
			case quoteBrace > 1:
				pUpdate(r)
				quoteBrace--
			case quoteBrace == 1:
				if quoteBraceHide {
					quoteBraceHide = false
					pUpdate(r)
				}
				quoteBrace--
			case quoteBrace == 0:
				pErr = raiseErr(ErrClosingBraceQuoteNoOpen, i)
				return &nodes, pErr
			default:
				pErr = raiseErr(ErrUnexpectedParsingError, i) //+" No case found for `switch ')' { ... }`.", i)
				return &nodes, pErr
			}

		case ':':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
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
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName && len(*pop) == 0:
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
				ignoreWhitespace = false
				continue // skip `last=r`
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName && len(*pop) == 0:
				pUpdate(r)
				if i < len(block)-1 && block[i+1] == '>' {
					pErr = raiseErr(ErrUnexpectedPipeTokenEqGt, i)
					return &nodes, pErr
				} //else {
				startParameters()
				//}
			default:
				pUpdate(r)
			}

		case '[':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName:
				pUpdate(r)
				if i < len(block)-1 && block[i+1] != '[' {
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
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case scanFuncName:
				if len(*pop) == 0 {
					pErr = raiseErr(ErrUnexpectedOpenBraceFunc, i)
					return &nodes, pErr
				}
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
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case scanFuncName:
				pErr = raiseErr(ErrUnexpectedCloseBrace, i)
				return &nodes, pErr
			case braceCount == 0:
				pErr = raiseErr(ErrClosingBraceBlockNoOpen, i)
				return &nodes, pErr
			default:
				pUpdate(r)
				braceCount--
			}

		case ' ', '\t', '\r':
			switch {
			case escaped:
				for lookFwd := i; lookFwd < len(block)-1; lookFwd++ {
					if block[lookFwd] == '#' {
						escapedBfComment = true
						break
					}
					if block[lookFwd] != ' ' && block[lookFwd] != '\t' {
						break
					}
				}
				if !escapedBfComment {
					pUpdate(r)
					ignoreWhitespace = false
				}
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			case !scanFuncName && last != ' ':
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
				switch {
				case !scanFuncName && last != ' ' && last != ':':
					node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
					pCount++
					pToken = &node.ParamTokens[pCount][0]
					pop = &pToken.Key
				case scanFuncName && !ignoreWhitespace:
					startParameters()
				default:
					pUpdate(r)
				}
				escaped = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
			case braceCount > 0:
				pUpdate(r)
			case scanFuncName && ignoreWhitespace:
				// do nothing
			default:
				appendNode()
				node = AstNode{NewChain: true}
				pop = &node.Name
				scanFuncName = true
			}

		case '|':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeTokenPipe, i)
				return &nodes, pErr
			case last == '|':
				appendNode()
				node = AstNode{LogicOr: true, NewChain: true}
				pop = &node.Name
				scanFuncName = true
			case !next('|'):
				node.PipeOut = true
				appendNode()
				node = AstNode{Method: true}
				pop = &node.Name
				scanFuncName = true
			default:
				// do nothing
				//pErr = raiseErr(ErrUnknownParserErrorPipe, i)
				//return &nodes, pErr
			}

		case '?':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeTokenQm, i)
				return &nodes, pErr
			case last == ' ' || last == '\t':
				node.PipeErr = true
				appendNode()
				node = AstNode{Method: true}
				pop = &node.Name
				scanFuncName = true
			default:
				pUpdate(r)
			}

		case '&':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			case next('&'):
				if len(node.Name) == 0 {
					pErr = raiseErr(ErrUnexpectedLogicAnd, i)
					return &nodes, pErr
				}
				/**pop = (*pop)[:len(*pop)-1]
				if len(*pop) == 0 {
					pToken.Type = parameters.TokenTypeNil
				}*/
				appendNode()
				node = AstNode{LogicAnd: true, NewChain: true}
				pop = &node.Name
				scanFuncName = true
			case last == '&' && len(*pop) == 0:
				// do nothing
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}

		case '<':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
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
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			case last == '-':
				if len(node.ParamTokens) != 0 {
					l := len(node.ParamTokens[pCount])
					if l > 1 && node.ParamTokens[pCount][l-2].Type == parameters.TokenTypeTilde && len(node.ParamTokens[pCount][l-2].Key) > 0 {
						// work around '-' being an acceptable character for ~,
						// thus causing an index out of bounds panic.
						node.ParamTokens[pCount][l-2].Key = node.ParamTokens[pCount][l-2].Key[:len(node.ParamTokens[pCount][l-2].Key)-1]
					} else {
						*pop = (*pop)[:len(*pop)-1]
					}
					if len(*pop) == 0 {
						pToken.Type = parameters.TokenTypeNil
					}
				}
				node.PipeOut = true
				appendNode()
				node = AstNode{Method: true}
				pop = &node.Name
				scanFuncName = true
			case last == '=':
				// close last node
				*pop = (*pop)[:len(*pop)-1]
				if len(*pop) == 0 {
					pToken.Type = parameters.TokenTypeNil
				}
				node.PipeOut = true
				appendNode()

				// append -> format generic ->
				node = AstNode{
					Method:  true,
					PipeOut: true,
					Name:    "format",
					ParamTokens: [][]parameters.ParamToken{{{
						Type: parameters.TokenTypeValue,
						Key:  types.Generic,
					}}},
				}
				appendNode()

				// new node
				node = AstNode{Method: true}
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
				ignoreWhitespace = false
			case quoteSingle, quoteDouble, quoteBrace > 0:
				pUpdate(r)
				ignoreWhitespace = false
			case braceCount > 0:
				pUpdate(r)
			default:
				appendNode()
				node = AstNode{NewChain: true}
				pop = &node.Name
				scanFuncName = true
			}

		case '~':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
				ignoreWhitespace = false
			case scanFuncName || braceCount > 0 || quoteSingle:
				pUpdate(r)
				ignoreWhitespace = false
			case !scanFuncName && last == '=' && node.Name == "=":
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
				ignoreWhitespace = false
			case scanFuncName || braceCount > 0 || quoteSingle:
				pUpdate(r)
				ignoreWhitespace = false
			case next('{'):
				fallthrough
			case nextAlphaNumeric():
				node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{Type: parameters.TokenTypeString})
				pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				pop = &pToken.Key
			default:
				pUpdate(r)
			}

		case '@':
			switch {
			case escaped:
				pUpdate(r)
				escaped = false
				ignoreWhitespace = false
			case scanFuncName || braceCount > 0 || quoteSingle:
				pUpdate(r)
				ignoreWhitespace = false
			case last != ' ' && last != '\t':
				pUpdate(r)
			case next('{'):
				fallthrough
			case nextAlphaNumeric():
				node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
				pCount++
				pToken = &node.ParamTokens[pCount][0]
				pToken.Type = parameters.TokenTypeArray
				pop = &pToken.Key
			default:
				pUpdate(r)
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
			/*switch {
			case escaped:
				pUpdate(r)
				escaped = false
			default:
				ignoreWhitespace = false
				pUpdate(r)
			}*/
			ignoreWhitespace = false
			pUpdate(r)
			// skip last=r since the last char was escaped
			if escaped {
				escaped = false
				continue
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

	//debug.Json("params", nodes)

	return &nodes, pErr
}

package lang

import "github.com/lmorg/murex/debug"

func parseBlock(block []rune) (nodes Nodes, pErr ParserError) {
	defer debug.Json("Parser", nodes)
	var (
		// Current state
		commentLine              bool
		escaped                  bool
		quoteSingle, quoteDouble bool
		braceCount               int
		ignoreWhitespace         bool = true
		scanFuncName             bool = true
		//newLine                  bool

		// Parsed thus far
		node Node    = Node{NewChain: true}
		pop  *string = &node.Name
	)
	defer debug.Json("Last node", node)

	appendNode := func() {
		if len(node.Parameters) > 1 && len(node.Parameters[len(node.Parameters)-1]) == 0 {
			node.Parameters = node.Parameters[:len(node.Parameters)-1]
		}

		if node.Name != "" {
			nodes = append(nodes, node)
		}

		ignoreWhitespace = true
	}

	for i, b := range block {
		if commentLine {
			if b == '\n' {
				commentLine = false
			}
			continue
		}

		switch b {
		case '#':
			switch {
			case escaped, quoteSingle, quoteDouble, braceCount > 0:
				*pop += string(b)
			default:
				commentLine = true
			}

		case '\\':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case escaped:
				*pop += string(b)
				escaped = false
			default:
				escaped = true
			}

		case '\'':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle:
				quoteSingle = false
			case quoteDouble:
				*pop += string(b)
			default:
				quoteSingle = true
			}

		case '"':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle:
				*pop += string(b)
			case quoteDouble:
				quoteDouble = false
			default:
				quoteDouble = true
			}

		case ':':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case scanFuncName:
				scanFuncName = false
				node.Parameters = make([]string, 1)
				pop = &node.Parameters[0]
			default:
				*pop += string(b)
				//pErr = raiseErr(ErrUnexpectedColon, i)
				//return
			}

		case '{':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case scanFuncName:
				//pErr = raiseErr(ErrUnexpectedOpenBrace, i)
				//return
				// Update function name:
				scanFuncName = false
				node.Parameters = make([]string, 1)
				pop = &node.Parameters[0]
				// Update first parameter:
				*pop += string(b)
				braceCount++
			default:
				*pop += string(b)
				braceCount++
			}

		case '}':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case scanFuncName:
				pErr = raiseErr(ErrUnexpectedCloseBrace, i)
				return
			case braceCount == 0:
				pErr = raiseErr(ErrClosingBraceNoOpen, i)
				return
			default:
				*pop += string(b)
				braceCount--
			}

		case ' ', '\t', '\r':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case !scanFuncName:
				if len(*pop) > 0 {
					node.Parameters = append(node.Parameters, "")
					pop = &node.Parameters[len(node.Parameters)-1]
				}
			case scanFuncName && !ignoreWhitespace:
				scanFuncName = false
				node.Parameters = make([]string, 1)
				pop = &node.Parameters[0]
			default:
				// do nothing
			}

		case '\n':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case !scanFuncName:
				appendNode()
				node = Node{NewChain: true}
				pop = &node.Name
				scanFuncName = true
				//newLine = true
			case scanFuncName && !ignoreWhitespace:
				scanFuncName = false
				node.Parameters = make([]string, 1)
				pop = &node.Parameters[0]
			default:
				// do nothing
			}

		case '|':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
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
				node = Node{}
				pop = &node.Name
				scanFuncName = true
			}

		case '?':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeToken, i)
				return
			/*case newLine:
			newLine = false
			node.NewChain = false
			if len(nodes) > 0 {
				nodes.Last().PipeErr = true
			}*/
			default:
				node.PipeErr = true
				appendNode()
				node = Node{}
				pop = &node.Name
				scanFuncName = true
			}

		case '>':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case len(*pop) > 0 && (*pop)[len(*pop)-1] == '-':
				/*if len(node.Name) == 0 {
					pErr = raiseErr(ErrUnexpectedPipeToken, i)
					return
				}*/
				*pop = (*pop)[:len(*pop)-1]
				node.PipeOut = true
				appendNode()
				node = Node{Method: true}
				pop = &node.Name
				scanFuncName = true

				/*if newLine {
					node.NewChain = false
					node.Method = true
					nodes.Last().PipeOut = true
					newLine = false
				}*/
			default:
				*pop += string(b)
			}

		case ';':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
				//case !scanFuncName:
			default:
				appendNode()
				node = Node{NewChain: true}
				pop = &node.Name
				scanFuncName = true
				//default:
				// do nothing
			}

		case 's':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case quoteSingle:
				*pop += string(b)
			case escaped:
				*pop += " "
				escaped = false
			default:
				*pop += string(b)
			}

		case 't':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case quoteSingle:
				*pop += string(b)
			case escaped:
				*pop += "\t"
				escaped = false
			default:
				*pop += string(b)
			}

		case 'r':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case quoteSingle:
				*pop += string(b)
			case escaped:
				*pop += "\r"
				escaped = false
			default:
				*pop += string(b)
			}

		case 'n':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case quoteSingle:
				*pop += string(b)
			case escaped:
				*pop += "\n"
				escaped = false
			default:
				*pop += string(b)
			}

		default:
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			default:
				ignoreWhitespace = false
				*pop += string(b)
				/*if b != '-' {
					newLine = false
				}*/
			}
		}
	}

	switch {
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

package expressions

import "github.com/lmorg/murex/utils/lists"

var tokeniseCurlyBraceCommands = []string{
	"if", "!if",
	"foreach", "formap", "for",
	"switch",
}

func (tree *ParserT) tokeniseCurlyBrace() bool {
	if tree.statement == nil {
		return false
	}
	return lists.Match(tokeniseCurlyBraceCommands, string(tree.statement.command))
}

func (tree *ParserT) tokeniseScalar() bool {
	if tree.statement == nil || len(tree.statement.paramTemp) != 0 {
		return true
	}

	switch string(tree.statement.command) {

	case "set", "global", "export",
		"!set", "!global", "!export", "unset":
		if len(tree.statement.parameters) == 0 {
			return false
		}

	case "foreach":
		if len(tree.statement.parameters) == 0 {
			return false
		}
		if len(tree.statement.parameters) == 1 {
			s := string(tree.statement.parameters[0])
			if s == "--jmap" || s == "--step" {
				return false
			}
		}

	case "formap":
		if len(tree.statement.parameters) == 0 {
			return false
		}
		if len(tree.statement.parameters) == 1 {
			s := string(tree.statement.parameters[0])
			if s == "--jmap" {
				return false
			}
		}

	case "is-null":
		return false
	}

	return true
}

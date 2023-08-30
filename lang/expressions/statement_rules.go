package expressions

import "github.com/lmorg/murex/utils/lists"

var tokeniseCurlyBraceCommands = []string{
	"if", "!if",
	"foreach", "formap",
	"switch",
}

func (tree *ParserT) tokeniseCurlyBrace() bool {
	if tree.statement == nil {
		return false
	}
	return lists.Match(tokeniseCurlyBraceCommands, string(tree.statement.command))
}

var tokeniseScalarCommands = []string{
//	"set", "!set",
//	"global", "!global",
//	//"export", "!export", "unset",
//	"foreach", "formap",
}

func (tree *ParserT) tokeniseScalar() bool {
	if tree.statement == nil {
		return false
	}
	return !lists.Match(tokeniseScalarCommands, string(tree.statement.command))
}

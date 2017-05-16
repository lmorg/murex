package parameters

type Parameters struct {
	params []string
	Tokens [][]ParamToken
}

type ParamToken struct {
	Type int
	Key  string
}

const (
	TokenTypeNil = iota
	TokenTypeValue
	TokenTypeString
	TokenTypeBlockString
	TokenTypeArray
	TokenTypeBlockArray
)

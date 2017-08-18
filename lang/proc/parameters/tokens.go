package parameters

type Parameters struct {
	Params []string
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
	TokenTypeTilde
)

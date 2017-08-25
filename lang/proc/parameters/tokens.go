package parameters

// Parameters is the parameter object
type Parameters struct {
	Params []string
	Tokens [][]ParamToken
}

// ParamToken holds information on each parameter token before it is parsed into a string
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

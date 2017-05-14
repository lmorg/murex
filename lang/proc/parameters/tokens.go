package parameters

type ParamToken struct {
	Type   int
	StrLoc int
	Key    string
}

const (
	TokenTypeString = 1 + iota
	TokenTypeBlockString
	TokenTypeArray
	TokenTypeBlockArray
)

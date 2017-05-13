package parameters

type InStrToken struct {
	Type     int
	Location int
	Key      string
}

const (
	InStrTokenTypeString = 1 + iota
	InStrTokenTypeBlockString
	InStrTokenTypeArray
	InStrTokenTypeBlockArray
)

package parameters

type InStrToken struct {
	Type     int
	Location int
	Key      string
}

const (
	InStrTokenTypeString = iota + 1
	InStrTokenTypeBlockString
	InStrTokenTypeArray
	InStrTokenTypeBlockArray
)

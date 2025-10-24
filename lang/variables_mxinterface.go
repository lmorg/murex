package lang

type MxInterface interface {
	GetValue() any
	GetString() string
	Set(any, []string) error
	New(string) (MxInterface, error)
}

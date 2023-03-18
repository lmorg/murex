package lang

type MxInterface interface {
	GetValue() interface{}
	GetString() string
	Set(interface{}, []string) error
	New(string) (MxInterface, error)
}

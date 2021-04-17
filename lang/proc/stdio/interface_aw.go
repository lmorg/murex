package stdio

// ArrayWriter is a simple interface types can adopt for buffered writes of formatted arrays in structured types (eg JSON)
type ArrayWriter interface {
	Write([]byte) error
	WriteString(string) error
	Close() error
}

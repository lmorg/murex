package tabulate

type writer interface {
	Write([]string) error
	Flush()
	Error() error
}

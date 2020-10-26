package tabulate

type writer interface {
	Write([]string) error
	Merge(string, string) error
	Flush()
	Error() error
}

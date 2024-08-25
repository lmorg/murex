package process

import "os"

type SystemProcess interface {
	Signal(sig os.Signal) error
	Kill() error
	Pid() int
	ExitNum() int
}

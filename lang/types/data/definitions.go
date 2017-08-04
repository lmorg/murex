package data

import (
	"github.com/lmorg/murex/lang/proc"
	"strings"
)

var (
	ReadIndexes map[string]func(p *proc.Process, params []string) error         = make(map[string]func(*proc.Process, []string) error)
	Unmarshal   map[string]func(p *proc.Process) (interface{}, error)           = make(map[string]func(*proc.Process) (interface{}, error))
	Marshal     map[string]func(p *proc.Process, v interface{}) ([]byte, error) = make(map[string]func(*proc.Process, interface{}) ([]byte, error))
	Mimes       map[string]string                                               = make(map[string]string)
	fileExts    map[string]string                                               = make(map[string]string)
)

func SetMime(dt string, mime ...string) {
	for i := range mime {
		Mimes[mime[i]] = dt
	}
}

func SetFileExtensions(dt string, extension ...string) {
	for i := range extension {
		fileExts[extension[i]] = strings.ToLower(dt)
	}
}

func GetExtType(extension string) (dt string) {
	dt = fileExts[strings.ToLower(extension)]
	return
}

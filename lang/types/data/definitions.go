package data

import (
	"github.com/lmorg/murex/lang/proc"
)

var ReadIndexes map[string]func(p *proc.Process, params []string) error = make(map[string]func(*proc.Process, []string) error)
var Unmarshel map[string]func(p *proc.Process) (interface{}, error) = make(map[string]func(*proc.Process) (interface{}, error))
var Marshel map[string]func(p *proc.Process, t interface{}) ([]byte, error) = make(map[string]func(*proc.Process, interface{}) ([]byte, error))

/*func marshelCsv(p *proc.Process, t interface{}) ([]byte, error) {
	switch v:=t.(type) {
	case []string:
		//return csv.ArrayToCsv(v), nil
	case map[string]string:
	case map[string][]string:
	case []map[string]string{}:
	case []map[string]interface{}:
	case []interface{}:
	default:
	}
	return nil, nil
}*/

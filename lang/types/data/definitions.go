package data

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

var (
	ReadIndexes map[string]func(p *proc.Process, params []string) error         = make(map[string]func(*proc.Process, []string) error)
	Unmarshal   map[string]func(p *proc.Process) (interface{}, error)           = make(map[string]func(*proc.Process) (interface{}, error))
	Marshal     map[string]func(p *proc.Process, v interface{}) ([]byte, error) = make(map[string]func(*proc.Process, interface{}) ([]byte, error))
)

func init() {
	// Register builtin data types
	Marshal[types.String] = marshalString
	Unmarshal[types.String] = unmarshalString

	Marshal[types.Json] = marshalJson
	Unmarshal[types.Json] = unmarshalJson

	Marshal[types.Csv] = marshalCsv
	Unmarshal[types.Csv] = unmarshalCsv

	ReadIndexes[types.Json] = indexJson
	ReadIndexes[types.Csv] = indexCsv
	ReadIndexes[types.Generic] = indexGeneric
	ReadIndexes[types.String] = indexGeneric
}

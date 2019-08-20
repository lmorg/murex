package example

import (
	"encoding/json"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data-type
	define.Unmarshallers["example"] = unmarshal
}

// Describe unmarshaller
func unmarshal(p *lang.Process) (interface{}, error) {
	// Read data from stdio
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	var v interface{}
	err = json.Unmarshal(b, &v)

	// Return the Go data structure or error
	return v, err
}

package example

import (
	"encoding/json"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data-type
	define.Marshallers["json"] = marshal
}

// Describe marshaller
func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	if p.Stdout.IsTTY() {
		return json.MarshalIndent(v, "", "    ")
	} else {
		return json.Marshal(v)
	}
}

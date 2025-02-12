package example

import (
	"encoding/json"

	"github.com/lmorg/murex/lang"
)

func init() {
	// Register data-type
	lang.RegisterMarshaller("example", marshal)
}

// Describe marshaller
func marshal(p *lang.Process, v any) ([]byte, error) {
	if p.Stdout.IsTTY() {
		// If STDOUT is a TTY (ie not pipe, text file or other destination other
		// than a terminal) then output JSON in an indented, human readable,
		// format....
		return json.MarshalIndent(v, "", "    ")
	}

	// ....otherwise we might as well output it in a minified format
	return json.Marshal(v)
}

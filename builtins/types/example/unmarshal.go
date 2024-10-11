package example

import (
	"encoding/json"

	"github.com/lmorg/murex/lang"
)

func init() {
	// Register data-type
	lang.RegisterUnmarshaller("example", unmarshal)
}

// Describe unmarshaller
func unmarshal(p *lang.Process) (any, error) {
	// Read data from STDIN. Because JSON expects closing tokens, we should
	// read the entire stream before unmarshalling it. For formats like CSV or
	// jsonlines which are more line based, we might want to read STDIN line by
	// line. However given there is just one data return, you still effectively
	// head to read the entire file before returning the structure. There are
	// other APIs for iterative returns for streaming data - more akin to the
	// traditional way UNIX pipes would work.
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	var v any
	err = json.Unmarshal(b, &v)

	// Return the Go data structure or error
	return v, err
}

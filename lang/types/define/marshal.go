package define

import (
	"errors"

	"github.com/lmorg/murex/lang"
)

// MarshalData is a global marshaller which should be called from within murex
// builtin commands (etc).
// See docs/apis/marshaldata.md for more details
func MarshalData(p *lang.Process, dataType string, data interface{}) (b []byte, err error) {

	if Marshallers[dataType] == nil {
		return nil, errors.New("I don't know how to marshal `" + dataType + "`.")
	}

	b, err = Marshallers[dataType](p, data)
	if err != nil {
		return nil, errors.New("[" + dataType + " marshaller] " + err.Error())
	}

	return
}

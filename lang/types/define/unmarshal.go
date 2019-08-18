package define

import (
	"errors"

	"github.com/lmorg/murex/lang"
)

// UnmarshalData is a global unmarshaller which should be called from within
// murex builtin commands (etc).
// See docs/apis/marshaldata.md for more details
func UnmarshalData(p *lang.Process, dataType string) (v interface{}, err error) {

	if Unmarshallers[dataType] == nil {
		return nil, errors.New("I don't know how to unmarshal `" + dataType + "`.")
	}

	v, err = Unmarshallers[dataType](p)
	if err != nil {
		return nil, errors.New("[" + dataType + " unmarshaller] " + err.Error())
	}

	return v, nil
}

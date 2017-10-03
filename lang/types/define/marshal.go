package define

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
)

// Marshallers is a global marshaller
func MarshalData(p *proc.Process, dataType string, data interface{}) (b []byte, err error) {

	if Marshallers[dataType] == nil {
		return nil, errors.New("I don't know how to marshal `" + dataType + "`.")
	}

	b, err = Marshallers[dataType](p, data)
	if err != nil {
		return nil, errors.New("[" + dataType + " marshaller] " + err.Error())
	}

	return
}

// Unmarshallers is a global unmarshaller
func UnmarshalData(p *proc.Process, dataType string) (v interface{}, err error) {

	if Unmarshallers[dataType] == nil {
		return nil, errors.New("I don't know how to unmarshal `" + dataType + "`.")
	}

	v, err = Unmarshallers[dataType](p)
	if err != nil {
		return nil, errors.New("[" + dataType + " unmarshaller] " + err.Error())
	}

	return v, nil
}

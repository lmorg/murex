# _murex_ Development Guide

## API Reference: `define.MarshalData()

> Converts structured memory into a structured file format (eg for stdio)

### Description

This is a function you would write when programming a _murex_ data-type.
The marshal function takes in a Go (golang) `type` or `struct` and returns
a byte slice of a "string" representation of that format (eg in JSON) or an
error.

This marshaller is then registered to _murex_ inside an `init()` function
and _murex_ builtins can use that marshaller via the `MarshalData()` API.

### Usage

    b, err := define.MarshalData(p, dataType, data)

### Examples

    func exampleCommand(p *lang.Process) error) {
        data := map[string]string {
            "foo": "hello foo",
            "bar": "hello bar",
        }
    
        dataType := "json"
    
        b, err := define.MarshalData(p, dataType, data)
        if err != nil {
            return err
        }
    
        _, err := p.Stdout.Write(b)
        return err
    }

### Detail

Go source file:

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

### Parameters

* `*lang.Process`: Process's runtime state. Typically expressed as the variable `p
* `string`: _murex_ data type
* `interface{}`: data you wish to marshal

### See Also

* [`Marshal()](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [unmarshal](../apis/unmarshal.md):
  
* [unmarshaldata](../apis/unmarshaldata.md):
  
# _murex_ Development Guide

## API Reference: `Unmarshal()` 

> Converts a structured file format into structured memory

### Description

This is a function you would write when programming a _murex_ data-type.
The unmarshal function takes in a byte slice and returns a Go (golang)
`type` or `struct` or an error.

This unmarshaller is then registered to _murex_ inside an `init()` function
and _murex_ builtins can use that unmarshaller via the `UnmarshalData()`
API.

### Usage

Registering unmarshal (for writing builtin data-types)

    // To avoid data races, this should only happen inside func init()
    define.Unmarshallers["json"] = unmarshal
    
Using an existing unmarshaller (eg inside a builtin command)

    // See documentation on define.UnmarshalData for more details
    v, err := define.UnmarshalData(p *lang.Process, dataType string)

### Examples

Defining a marshaller for a murex data-type

    func init() {
        // Register data-type
        define.Unmarshallers["json"] = unmarshal
    }
    
    // Describe unmarshaller
    func unmarshal(p *lang.Process) (interface{}, error) {
      	b, err := p.Stdin.ReadAll()
        if err != nil {
            return nil, err
        }
        
        var v interface{}
        err = json.Unmarshal(b, &v)
        return v, err
    }

### Parameters

1. `*lang.Process`: Process's runtime state

### See Also

* [`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [`define.MarshalData()` ](../apis/marshaldata.md):
  Converts structured memory into a _murex_ data-type (eg for stdio)
* [unmashaldata](../apis/unmashaldata.md):
  
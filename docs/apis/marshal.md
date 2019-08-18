# _murex_ Development Guide

## API Reference: Marshal()

> Converts structured memory into a structured file format (eg for stdio)

### Description

This is a function you would write when programming a _murex_ data-type.
The marshal function takes in a Go (golang) `type` or `struct` and returns
a byte slice of a "string" representation of that format (eg in JSON) or an
error.

This marshaller is then registered to _murex_ inside an `init()` function
and _murex_ builtins can use that marshaller via the `MarshalData()` API.

### Usage

Registering marshaller (for writing builtin data-types)

    // To avoid data races, this should only happen inside func init()
    define.Marshallers[types.Json] = marshal
    
Using an existing marshaller (eg inside a builtin command)

    // See documentation on define.MarshalData for more details
    b, err := define.MarshalData(p, dataType, data)

### Examples

Defining a marshaller for a murex data-type

    func init() {
        // Register data-type
        define.Marshallers[types.Json] = marshal
    }
    
    // Describe marshaller
    func marshal(p *lang.Process, v interface{}) ([]byte, error) {
        if p.Stdout.IsTTY() {
            return json.MarshalIndent(v, "", "    ")
        } else {
            return json.Marshal(v)
        }
    }

### Parameters

* `*lang.Process`: Process's runtime state
* `interface{}`: data you wish to marshal


### See Also

* [define.MarshalData()](../apis/marshaldata.md):
  Converts structured memory into a structured file format (eg for stdio)
* [unmarshal](../apis/unmarshal.md):
  
* [unmashaldata](../apis/unmashaldata.md):
  
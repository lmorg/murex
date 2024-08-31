# `ReadNotIndex()` (type)

> Data type handler for the bang-prefixed index, `![`, builtin

## Description

This is a function you would write when programming a Murex data-type.

It's called by the index, `![`, builtin.

The purpose of this function is to allow builtins to support sequential reads
(where possible) and also create a standard interface for `![` (index), thus
allowing it to be data-type agnostic.

## Usage

Registering your `ReadNotIndex()`

```go
// To avoid data races, this should only happen inside func init()
lang.ReadNotIndexes[ /* your type name */ ] = /* your readIndex func */
```

## Examples

Example `ReadIndex()` function (the code structure is the same for `ReadIndex`
and `ReadNotIndex`):

```go
package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func index(p *lang.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	marshaller := func(iface interface{}) ([]byte, error) {
		return json.Marshal(iface, p.Stdout.IsTTY())
	}

	return lang.IndexTemplateObject(p, params, &jInterface, marshaller)
}
```

## Detail

While there is support for a dedicated `ReadNotIndex()` for instances of `![`,
the template APIs `lang.IndexTemplateObject` and `lang.IndexTemplateTable` are
both agnostic to the bang prefix.

## Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p` 
2. `[]string`: slice of parameters used in `![` 

## See Also

* [user-guide/Bang Prefix](../user-guide/bang-prefix.md):
  Bang prefixing to reverse default actions
* [parser/Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [apis/`ReadArrayWithType()` (type)](../apis/ReadArrayWithType.md):
  Read from a data type one array element at a time and return the elements contents and data type
* [apis/`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [apis/`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time
* [apis/`lang.IndexTemplateObject()` (template API)](../apis/lang.IndexTemplateObject.md):
  Returns element(s) from a data structure
* [apis/`lang.IndexTemplateTable()` (template API)](../apis/lang.IndexTemplateTable.md):
  Returns element(s) from a table
* [parser/index](../parser/item-index.md):
  Outputs an element from an array, map or table

<hr/>

This document was generated from [lang/stdio/interface_doc.yaml](https://github.com/lmorg/murex/blob/master/lang/stdio/interface_doc.yaml).
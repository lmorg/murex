# Define Method Relationships (`method`)

> Define a methods supported data-types

## Description

`method` defines what the typical data type would be for a function's stdin
and stdout.

## Usage

```
method: define name { json }
```

## Examples

```
method: define name {
    "Stdin":  "@Any",
    "Stdout": "json"
}
```

## Detail

### Type Groups

You can define a Murex data type or use a type group. The following type
groups are available to use:

```go
package types

// These are the different supported type groups
const (
	Any               = "@Any"
	Text              = "@Text"
	Math              = "@Math"
	Unmarshal         = "@Unmarshal"
	Marshal           = "@Marshal"
	ReadArray         = "@ReadArray"
	ReadArrayWithType = "@ReadArrayWithType"
	WriteArray        = "@WriteArray"
	ReadIndex         = "@ReadIndex"
	ReadNotIndex      = "@ReadNotIndex"
	ReadMap           = "@ReadMap"
)

// GroupText is an array of the data types that make up the `text` type
var GroupText = []string{
	Generic,
	String,
	`generic`,
	`string`,
}

// GroupMath is an array of the data types that make up the `math` type
var GroupMath = []string{
	Number,
	Integer,
	Float,
	Boolean,
}
```

## Synonyms

* `method`


## See Also

* [Alias Pointer (`alias`)](../commands/alias.md):
  Create an alias for a command
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [builtins/core/structs/function_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/function_doc.yaml).
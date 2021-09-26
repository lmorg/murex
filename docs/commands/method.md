# _murex_ Shell Docs

## Command Reference: `method`

> Define a methods supported data-types

## Description

`method` defines what the typical data type would be for a function's STDIN
and STDOUT.

## Usage

    method: define name { json }

## Examples

    method: define name {
        "Stdin":  "@Any",
        "Stdout": "json"
    }

## Detail

### Type Groups

You can define a _murex_ data type or use a type group. The following type
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
}

// GroupMath is an array of the data types that make up the `math` type
var GroupMath = []string{
	Number,
	Integer,
	Float,
	Boolean,
}
```

## See Also

* [commands/`alias`](../commands/alias.md):
  Create an alias for a command
* [commands/`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
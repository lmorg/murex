# `alter`

> Change a value within a structured data-type and pass that change along the pipeline without altering the original source input

## Description

`alter` a value within a structured data-type.

The path separator is defined by the first character in the path. For example
`/path/to/key`, `,path,to,key`, `|path|to|key` and `#path#to#key` are all valid
however you should remember to quote or escape any special characters (tokens)
used by the shell (such as pipe, `|`, and hash, `#`).

The **value** must always be supplied as JSON.

> When working with expressions, you may find the **Assign or Merge** operator
> more ergonomic ([read more](/docs/parser/assign-or-merge.md)) 

## Usage

```
<stdin> -> alter [ -m | --merge | -s | --sum ] /path value -> <stdout>
```

## Examples

```
» config -> [ shell ] -> [ prompt ] -> alter /Value moo
{
    "Data-Type": "block",
    "Default": "{ out 'murex » ' }",
    "Description": "Interactive shell prompt.",
    "Value": "moo"
}
```

`alter` also accepts JSON as a parameter for adding structured data:

```
config -> [ shell ] -> [ prompt ] -> alter /Example { "Foo": "Bar" }
{
    "Data-Type": "block",
    "Default": "{ out 'murex » ' }",
    "Description": "Interactive shell prompt.",
    "Example": {
        "Foo": "Bar"
    },
    "Value": "{ out 'murex » ' }"
}
```

However it is also data type aware so if they key you're updating holds a string
(for example) then the JSON data a will be stored as a string:

```
» config -> [ shell ] -> [ prompt ] -> alter /Value { "Foo": "Bar" }
{
    "Data-Type": "block",
    "Default": "{ out 'murex » ' }",
    "Description": "Interactive shell prompt.",
    "Value": "{ \"Foo\": \"Bar\" }"
}
```

Numbers will also follow the same transparent conversion treatment:

```
» tout json { "one": 1, "two": 2 } -> alter /two "3"
{
    "one": 1,
    "two": 3
}
```

> Please note: `alter` is not changing the value held inside `config` but
> instead took the STDOUT from `config`, altered a value and then passed that
> new complete structure through it's STDOUT.
>
> If you require modifying a structure inside Murex config (such as http
> headers) then you can use `config alter`. Read the config docs for reference.

### -m / --merge

Thus far all the examples have be changing existing keys. However you can also
alter a structure by appending to an array or a merging two maps together. You
do this with the `--merge` (or `-m`) flag.

```
» out a\nb\nc -> alter --merge / ([ "d", "e", "f" ])
a
b
c
d
e
f
```

### -s / --sum

This behaves similarly to `--merge` where structures are blended together.
However where a map exists with two keys the same and the values are numeric,
those values are added together.

```
» tout json { "a": 1, "b": 2 } -> alter --sum / { "b": 3, "c": 4 }
{
    "a": 1,
    "b": 5,
    "c": 4
}
```

## Flags

* `--merge`
    Merge data structures rather than overwrite
* `--sum`
    Sum values in a map, merge items in an array
* `-m`
    Alias for `--merge`
* `-s`
    Alias for `--sum`

## Detail

### Path

The path parameter can take any character as node separators. The separator is
assigned via the first character in the path. For example

```
config -> alter .shell.prompt.Value moo
config -> alter >shell>prompt>Value moo
```

Just make sure you quote or escape any characters used as shell tokens. eg

```
config -> alter '#shell#prompt#Value' moo
config -> alter ' shell prompt Value' moo
```

### Supported data-types

The *value* field must always be supplied as JSON however the *STDIN* struct
can be any data-type supported by murex.

You can check what data-types are available via the `runtime` command:

```
runtime --marshallers
```

Marshallers are enabled at compile time from the `builtins/data-types` directory.

## See Also

* [`<~` Assign or Merge](../parser/assign-or-merge.md):
  Merges the right hand value to a variable on the left hand side (expression)
* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`[[ Element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`prepend`](../commands/prepend.md):
  Add data to the start of an array
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex

<hr/>

This document was generated from [builtins/core/datatools/alter_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/datatools/alter_doc.yaml).
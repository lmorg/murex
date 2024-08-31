# Define Type (`cast`)

> Alters the data-type of the previous function without altering its output

## Description

`cast` alters the data-type annotation for a pipe.

The contents of the pipeline are preserved, only the reported data-type is
changed.

Additionally `cast` can be used to define the output type of a function.

## Usage

Changing type annotation for a pipe

```
<stdin> -> cast data-type -> <stdout>

<stdin> :data-type: command
```

Defining the output type of a function

```
cast data-type
```

## Examples

### As a command

```
» out {"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}} \
  -> cast json
{"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}}
```

### As a token

```
» out {"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}} \
  -> :json: cat
{"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}}
```

### Defining data-type

```
» function example {
    cast json
    out '{"foo": "bar"}'
}

» example -> debug -> [[ /Data-Type/Murex ]]
json
```

Please note you'd normally use the [Object Builder](/docs/parser/create-object.md) to create JSON objects.

## Detail

If you want to reformat the stdin into the new data type then use `format`
instead.

## See Also

* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [`%{}` Object Builder](../parser/create-object.md):
  Quickly generate objects (dictionaries / maps)

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).
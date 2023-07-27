# `cast` - Command Reference

> Alters the data type of the previous function without altering it's output

## Description

`cast` works a little like when you case variables in lower level languages
where the value of the variable is unchanged. In Murex the contents in
the pipeline are preserved however the reported data type is altered.

## Usage

    <stdin> -> cast data-type -> <stdout>

## Examples

    » out: {"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}} -> cast json
    {"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}}

## Detail

If you want to reformat the STDIN into the new data type then use `format`
instead.

## See Also

* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
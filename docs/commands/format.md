# _murex_ Shell Docs

## Command Reference: `format`

> Reformat one data-type into another data-type

## Description

`format` takes a data from STDIN and returns that data reformated in another
specified data-type

## Usage

    <stdin> -> format data-type -> <stdout>

## Examples

    » tout json { "One": 1, "Two": 2, "Three": 3 } -> format yaml
    One: 1
    Three: 3
    Two: 2

## See Also

* [apis/`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [apis/`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
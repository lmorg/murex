# Error String (`err`)

> Print a line to the stderr

## Description

Write parameters to stderr with a trailing new line character.

## Usage

```
err string to write -> <stderr>
```

## Examples

```
» err Hello, World!
Hello, World!
```

## Detail

`err` outputs as `string` data-type. This can be changed by casting

```
err { "Code": 404, "Message": "Page not found" } ? cast json
```

However passing structured data-types along the stderr stream is not recommended
as any other function within your code might also pass error messages along the
same stream and thus taint your structured data. This is why Murex does not
supply a `tout` function for stderr. The recommended solution for passing
messages like these which you want separate from your stdout stream is to create
a new Murex named pipe.

```
» pipe --create messages
» bg { <messages> -> pretty }
» tout <messages> json { "Code": 404, "Message": "Page not found" }
» pipe --close messages
{
    "Code": 404,
    "Message": "Page not found"
}
```

### ANSI Constants

`err` supports ANSI constants.

## Synonyms

* `err`


## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Create Named Pipe (`pipe`)](../commands/pipe.md):
  Manage Murex named pipes
* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Get Pipe Status (`pt`)](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Prettify JSON](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [Read / Write To A Named Pipe (`<pipe>`)](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [Truncate File (`>`)](../parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)
* [`>>` Append File](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists

<hr/>

This document was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml).
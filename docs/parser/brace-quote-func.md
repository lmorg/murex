# `(brace quote)`

> Write a string to the stdout without new line (deprecated)

## Description

Write parameters to stdout (does not include a new line)

## Usage

```
(string to write) -> <stdout>
```

## Examples

```
» (Hello, World!)
Hello, World!

» (Hello,\nWorld!)
Hello,
World!

» ((Hello,) (World!))
(Hello,) (World!)

# Print "Hello, World!" in red text
» {RED}Hello, World!{RESET}
Hello, World!
```

## Detail

The `(` function performs exactly like the `(` token for quoting so you do not
need to escape other tokens (eg single / double quotes, `'`/`"`, nor curly
braces, `{}`). However the braces are nestable so you will need to escape those
characters if you don't want them nested.

### ANSI Constants

`(` supports ANSI constants.

## Synonyms

* `(`


## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Get Pipe Status (`pt`)](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Truncate File (`>`)](../parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists
* [`>>` Append File](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists

<hr/>

This document was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml).
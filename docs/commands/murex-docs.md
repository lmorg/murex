# Murex's Offline Documentation (`murex-docs`)

> Displays the man pages for Murex builtins

## Description

Displays the man pages for Murex builtins.

## Usage

```
murex-docs [ flag ] command -> <stdout>
```

## Examples

```
# Output this man page
murex-docs murex-docs
```

## Flags

* `--docs`
    Returns a JSON object of every document available to read offline
* `--summary`
    Returns an abridged description of the command rather than the entire help page

## Detail

These man pages are compiled into the Murex executable.

## Synonyms

* `murex-docs`
* `help`


## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Parse Man-Page For Flags (`man-get-flags`)](../commands/man-get-flags.md):
  Parses man page files for command line flags 
* [Read With Type (`tread`) (removed 7.x)](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)
* [Truncate File (`>`)](../parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)
* [`>>` Append File](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists

<hr/>

This document was generated from [builtins/docs/docs_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/docs/docs_doc.yaml).
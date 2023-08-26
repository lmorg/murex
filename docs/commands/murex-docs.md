# `murex-docs`

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

* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`>>` (append file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`err`](../commands/err.md):
  Print a line to the STDERR
* [`man-get-flags` ](../commands/man-get-flags.md):
  Parses man page files for command line flags 
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [`tread`](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)

<hr/>

This document was generated from [builtins/docs/docs_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/docs/docs_doc.yaml).
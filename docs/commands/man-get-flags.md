# Parse Man-Page For Flags (`man-get-flags`)

> Parses man page files for command line flags 

## Description

Sometimes you might want to programmatically search `man` pages for any
supported flag. Particularly if you're writing a dynamic autocompletion.
`man-get-flags` does this and returns a JSON document.

You can either pipe a man page to `man-get-flags`, or pass the name of the
command as a parameter.

`man-get-flags` returns a JSON document. Either an array or an object,
depending on what flags (if any) are passed.

If no flags are passed, `man-get-flags` will default to just parsing the man
page for anything that looks like a flag (ie no descriptions or other detail).

## Usage

```
<stdin> -> man-get-flags [--descriptions] -> <stdout>

man-get-flags command [--descriptions] -> <stdout>
```

## Examples

```
» man-get-flags --descriptions find -> [{$.key =~ 'regex'}]
{
    "-iregex": "eg: pattern -- Like -regex, but the match is case insensitive.",
    "-regex": "eg: pattern -- True if the whole path of the file matches pattern using regular expression.  To match a file named “./foo/xyzzy”, you can use the regular expression “.*/[xyz]*” or “.*/foo/.*”, but not “xyzzy” or “/foo/”."
}
```

## Flags

* `--descriptions`
    return a map of flags with their described usage
* `-d`
    shorthand for `--descriptions`

## Detail

### Limitations

Due to the freeform nature of man pages - that they're intended to be human
readable rather than machine readable - and the flexibility that developers
have to parse command line parameters however they wish, there will always be
a margin for error with how reliably any parser can autodetect parameters. one
requirement is that flags are hyphen prefixed, eg `--flag`.

## Synonyms

* `man-get-flags`


## See Also

* [Man-Page Summary (`man-summary`)](../commands/man-summary.md):
  Outputs a man page summary of a command
* [Murex's Offline Documentation (`murex-docs`)](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [Set Command Summary Hint (`summary`)](../commands/summary.md):
  Defines a summary help text for a command

<hr/>

This document was generated from [builtins/core/management/shell_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/shell_doc.yaml).
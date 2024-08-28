# escape.cli: `esccli`

> Escapes an array so output is valid shell code

## Description

`esccli` takes an array and escapes any characters that might cause problems
when pasted back into the terminal. Typically you'd want to use this against
command parameters.

## Usage

```
<stdin> -> esccli -> <stdout>

esccli @array -> <stdout>
```

## Examples

### As a method

```
» alias foobar=out 'foo$b@r'
» alias -> [foobar]
[
    "out",
    "foo$b@r"
]
» alias -> [foobar] -> esccli
out foo\$b\@r
```

### As a function

```
» alias -> [foobar] -> set fb
» $fb
["out","foo$b@r"]
» esccli @fb
out foo\$b\@r
```

## Synonyms

* `esccli`
* `escape.cli`


## See Also

* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`alias`](../commands/alias.md):
  Create an alias for a command
* [escape.html: `eschtml`](../commands/eschtml.md):
  Encode or decodes text for HTML
* [escape.quote: `escape`](../commands/escape.md):
  Escape or unescape input
* [escape.url: `escurl`](../commands/escurl.md):
  Encode or decodes text for the URL
* [io.out: `out`](../commands/out.md):
  Print a string to the stdout with a trailing new line character

<hr/>

This document was generated from [builtins/core/escape/escape_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/escape/escape_doc.yaml).
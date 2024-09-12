# Quote String (`escape`)

> Escape or unescape input

## Description

`escape` takes input from either stdin or the parameters and returns the same
data, escaped.

`!escape` does the same process in reverse, where it takes escaped data and
returns its unescaped counterpart.

## Usage

### Escape

```
<stdin> -> escape -> <stdout>

escape string to escape -> <stdout>
```

### Unescape

```
<stdin> -> !escape -> <stdout>

!escape string to unescape -> <stdout>
```

## Examples

### Escape

```
» out (multi
» line
» string) -> escape
"multi\nline\nstring\n" 
```

## Synonyms

* `escape`
* `!escape`


## See Also

* [Escape Command Line String (`esccli`)](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [Escape HTML (`eschtml`)](../commands/eschtml.md):
  Encode or decodes text for HTML
* [Escape URL (`escurl`)](../commands/escurl.md):
  Encode or decodes text for the URL

<hr/>

This document was generated from [builtins/core/escape/escape_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/escape/escape_doc.yaml).
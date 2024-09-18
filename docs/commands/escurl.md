# Escape URL (`escurl`)

> Encode or decodes text for the URL

## Description

`escurl` takes input from either stdin or the parameters and returns the same
data, escaped for the URL.

`!eschtml` does the same process in reverse, where it takes URL escaped data
and returns its unescaped counterpart.

## Usage

### Escape

```
<stdin> -> escurl -> <stdout>

escurl string to escape -> <stdout>
```

### Unescape

```
<stdin> -> !escurl -> <stdout>

!escurl string to unescape -> <stdout>
```

## Examples

### Escape

```
» out "!? <>" -> escurl
%21%3F%20%3C%3E%0A 
```

### Unescape

```
out '%21%3F%20%3C%3E%0A' -> !escurl
!? <>
```

## Synonyms

* `escurl`
* `!escurl`


## See Also

* [Download File (`getfile`)](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [Escape Command Line String (`esccli`)](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [Escape HTML (`eschtml`)](../commands/eschtml.md):
  Encode or decodes text for HTML
* [Get Request (`get`)](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [Post Request (`post`)](../commands/post.md):
  HTTP POST request with a JSON-parsable return
* [Quote String (`escape`)](../commands/escape.md):
  Escape or unescape input

<hr/>

This document was generated from [builtins/core/escape/escape_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/escape/escape_doc.yaml).
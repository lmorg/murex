# Escape HTML (`eschtml`)

> Encode or decodes text for HTML

## Description

`eschtml` takes input from either stdin or the parameters and returns the same
data, HTML escaped.

`!eschtml` does the same process in reverse, where it takes HTML escaped data
and returns its unescaped counterpart.

## Usage

### Escape

```
<stdin> -> eschtml -> <stdout>

eschtml string to escape -> <stdout>
```

### Unescape

```
<stdin> -> !eschtml -> <stdout>

!eschtml string to unescape -> <stdout>
```

## Examples

### Escape

```
» out "<h1>foo & bar</h1>" -> eschtml
&lt;h1&gt;foo &amp; bar&lt;/h1&gt;
```

### Unescape

```
» out '&lt;h1&gt;foo &amp; bar&lt;/h1&gt;' -> !eschtml
<h1>foo & bar</h1>
```

## Synonyms

* `eschtml`
* `!eschtml`


## See Also

* [Download File (`getfile`)](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [Escape Command Line String (`esccli`)](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [Escape URL (`escurl`)](../commands/escurl.md):
  Encode or decodes text for the URL
* [Get Request (`get`)](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [Post Request (`post`)](../commands/post.md):
  HTTP POST request with a JSON-parsable return
* [Quote String (`escape`)](../commands/escape.md):
  Escape or unescape input

<hr/>

This document was generated from [builtins/core/escape/escape_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/escape/escape_doc.yaml).
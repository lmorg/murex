# `eschtml` - Command Reference

> Encode or decodes text for HTML

## Description

`eschtml` takes input from either STDIN or the parameters and returns the same
data, HTML escaped.

`!eschtml` does the same process in reverse, where it takes HTML escaped data
and returns its unescaped counterpart.

## Usage

Escape

```
<stdin> -> eschtml -> <stdout>

eschtml string to escape -> <stdout>
```

Unescape

```
<stdin> -> !eschtml -> <stdout>

!eschtml string to unescape -> <stdout>
```

## Examples

Escape

```
» out: "<h1>foo & bar</h1>" -> eschtml
&lt;h1&gt;foo &amp; bar&lt;/h1&gt;
```

Unescape

```
» out: '&lt;h1&gt;foo &amp; bar&lt;/h1&gt;' -> !eschtml
<h1>foo & bar</h1>
```

## Synonyms

* `eschtml`
* `!eschtml`


## See Also

* [`escape`](../commands/escape.md):
  Escape or unescape input 
* [`esccli`](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [`escurl`](../commands/escurl.md):
  Encode or decodes text for the URL
* [`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return
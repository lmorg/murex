# `!bz2`

> Decompress a bz2 file

## Description

`!bz2` is an optional builtin for decompressing a bz2 stream from stdin.

## Usage

```
<stdin> -> !bz2 -> <stdout>
```

## Detail

Currently there is no support for compressing a stream using bz2.

## Synonyms

* `!bz2`


## See Also

* [`base64` ](../optional/base64.md):
  Encode or decode a base64 string
* [`gz`](../optional/gz.md):
  Compress or decompress a gzip file
* [escape.cli: `esccli`](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [escape.html: `eschtml`](../commands/eschtml.md):
  Encode or decodes text for HTML
* [escape.quote: `escape`](../commands/escape.md):
  Escape or unescape input
* [escape.url: `escurl`](../commands/escurl.md):
  Encode or decodes text for the URL

<hr/>

This document was generated from [builtins/optional/encoders/bz2_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/optional/encoders/bz2_doc.yaml).
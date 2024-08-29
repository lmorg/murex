# `gz`

> Compress or decompress a gzip file

## Description

An optional builtin for compressing or decompressing a gzip stream from stdin.

## Usage

```
<stdin> -> gz -> <stdout>

<stdin> -> !gz -> <stdout>
```

## Synonyms

* `gz`
* `!gz`


## See Also

* [`!bz2`](../optional/bz2.md):
  Decompress a bz2 file
* [`base64` ](../optional/base64.md):
  Encode or decode a base64 string
* [escape.cli](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [escape.html](../commands/eschtml.md):
  Encode or decodes text for HTML
* [escape.quote](../commands/escape.md):
  Escape or unescape input
* [escape.url](../commands/escurl.md):
  Encode or decodes text for the URL

<hr/>

This document was generated from [builtins/optional/encoders/gz_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/optional/encoders/gz_doc.yaml).
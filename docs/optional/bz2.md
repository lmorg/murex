# _murex_ Shell Docs

## Optional Command Reference: `!bz2` 

> Decompress a bz2 file

## Description

`!bz2` is an optional builtin for decompressing a bz2 stream from STDIN.

## Usage

    <stdin> -> !bz2 -> <stdout>

## Detail

Currently there is no support for compressing a stream using bz2.

## Synonyms

* `!bz2`


## See Also

* [`base64` ](../optional/base64.md):
  Encode or decode a base64 string
* [`escape`](../commands/escape.md):
  Escape or unescape input 
* [`esccli`](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [`eschtml`](../commands/eschtml.md):
  Encode or decodes text for HTML
* [`escurl`](../commands/escurl.md):
  Encode or decodes text for the URL
* [`gz` ](../optional/gz.md):
  Compress or decompress a gzip file
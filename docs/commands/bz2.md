# _murex_ Language Guide

## Command Reference: `gz` (optional)

> Compress or decompress a bz2 file

### Description

An optional builtin for compressing or decompressing a gzip stream from STDIN.

### Usage

    <stdin> -> gz -> <stdout>
    
    <stdin> -> !gz -> <stdout>

### Synonyms

* `gz`
* `!gz`


### See Also

* [`!bz2` (optional)](../commands/bz2.md):
  Decompress a bz2 file
* [`base64` (optional)](../commands/base64.md):
  Encode or decode a base64 string
* [`esccli`](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [escape](../commands/escape.md):
  
* [eschtml](../commands/eschtml.md):
  
* [escurl](../commands/escurl.md):
  
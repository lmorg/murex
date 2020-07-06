# _murex_ Shell Docs

## Command Reference: `escape`

> Escape or unescapes input 

## Description

`escape` takes input from either STDIN or the parameters and returns the same
data, escaped.

`!escape` does the same process in reverse, where it takes escaped data and
returns its unescaped counterpart.

## Usage

Escape

    <stdin> -> escape -> <stdout>
    
    escape string to escape -> <stdout>
    
Unescape

    <stdin> -> !escape -> <stdout>
    
    !escape string to unescape -> <stdout>

## Examples

Escape

    » out (multi
    » line
    » string) -> escape
    "multi\nline\nstring\n" 

## Synonyms

* `escape`
* `!escape`


## See Also

* [commands/`esccli`](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [commands/`eschtml`](../commands/eschtml.md):
  Encode or decodes text for HTML
* [commands/`escurl`](../commands/escurl.md):
  Encode or decodes text for the URL
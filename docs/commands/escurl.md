# _murex_ Shell Docs

## Command Reference: `escurl`

> Encode or decodes text for the URL

## Description

`escurl` takes input from either STDIN or the parameters and returns the same
data, escaped for the URL.

`!eschtml` does the same process in reverse, where it takes URL escaped data
and returns its unescaped counterpart.

## Usage

Escape

    <stdin> -> escurl -> <stdout>
    
    escurl string to escape -> <stdout>
    
Unescape

    <stdin> -> !escurl -> <stdout>
    
    !escurl string to unescape -> <stdout>

## Examples

Escape

    » out: "!? <>" -> escurl
    %21%3F%20%3C%3E%0A 
    
Unescape

    out: '%21%3F%20%3C%3E%0A' -> !escurl
    !? <>

## Synonyms

* `escurl`
* `!escurl`


## See Also

* [`escape`](../commands/escape.md):
  Escape or unescape input 
* [`esccli`](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [`eschtml`](../commands/eschtml.md):
  Encode or decodes text for HTML
* [`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.
* [`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return
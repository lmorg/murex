# _murex_ Shell Docs

## Command Reference: `eschtml`

> Encode or decodes text for HTML

## Description

`eschtml` takes input from either STDIN or the parameters and returns the same
data, HTML escaped.

`!eschtml` does the same process in reverse, where it takes HTML escaped data
and returns its unescaped counterpart.

## Usage

Escape

    <stdin> -> eschtml -> <stdout>
    
    eschtml string to escape -> <stdout>
    
Unescape

    <stdin> -> !eschtml -> <stdout>
    
    !eschtml string to unescape -> <stdout>

## Examples

Escape

    » out: "<h1>foo & bar</h1>" -> eschtml
    &lt;h1&gt;foo &amp; bar&lt;/h1&gt;
    
Unescape

    » out: '&lt;h1&gt;foo &amp; bar&lt;/h1&gt;' -> !eschtml
    <h1>foo & bar</h1>

## Synonyms

* `eschtml`
* `!eschtml`


## See Also

* [commands/`escape`](../commands/escape.md):
  Escape or unescape input 
* [commands/`esccli`](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [commands/`escurl`](../commands/escurl.md):
  Encode or decodes text for the URL
* [commands/`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [commands/`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.
* [commands/`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return
# _murex_ Shell Docs

## Command Reference: `pt`

> Pipe telemetry. Writes data-types and bytes written

## Description

Pipe telemetry, `pt`, writes statistics about the pipeline. The telemetry is written
directly to the OS's STDERR so to preserved the pipeline.

## Usage

    <stdin> -> pt -> <stdout>

## Examples

    curl -s https://example.com/bigfile.bin -> pt -> > bigfile.bin
    (though _murex_ does also have it's own HTTP clients, `get`, `post` and
`getfile`)

## See Also

* [commands/`>>` (append file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [commands/`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [commands/`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [commands/`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.
* [commands/`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return
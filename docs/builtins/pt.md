# _murex_ Language Guide

## Command reference: pt

> Pipe telemetry. Writes data-types and bytes written

### Description

Pipe telemetry writes statistics about the pipeline. The telemetry is written
directly to the OS's STDERR the pipeline is preserved.

### Usage

    <stdin> -> pt -> <stdout>

### See also

* [`err`](err.md): Print a line to the STDERR
* `sprintf`
* [`tout`](tout.md): Print a string to the STDOUT and set it's data-type
* [`ttyfd`](ttyfd.md): Returns the TTY device of the parent.

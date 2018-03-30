# _murex_ Language Guide

## Command reference: event

> Event driven programming for shell scripts

### Description

Create or destroy an event interrupt

### Usage

    event: --event-type interrupt { code block }

    !event: --event-type interrupt

### Examples

Create an event:

    event: --timer autoquit=60 {
        out "You're 60 secound timeout has elapsed. Quitting murex"
        exit 1
    }

Destory an event:

    !event --timer autoquit

### Detail

### Synonyms

* !event

### See also

* `murex`
* `runtime`

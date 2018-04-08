# _murex_ Language Guide

## Command reference: event

> Event driven programming for shell scripts

### Description

Create or destroy an event interrupt

### Usage

    event: event-type name=interrupt { code block }

    !event: event-type name

### Examples

Create an event:

    event: afterSecondsElapsed autoquit=60 {
        out "You're 60 secound timeout has elapsed. Quitting murex"
        exit 1
    }

Destroy an event:

    !event afterSecondsElapsed autoquit

### details

The `interrupt` field in the CLI supports ANSI constants. eg

    event: onKeyPress f1={F1-VT100} {
        tout: qs HintText="Key F1 Pressed"
    }

### Synonyms

* !event

### See also

* `runtime`

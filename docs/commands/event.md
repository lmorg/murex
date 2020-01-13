# _murex_ Shell Docs

## Command Reference: `event`

> Event driven programming for shell scripts

## Description

Create or destroy an event interrupt

## Usage

    event: event-type name=interrupt { code block }
    
    !event: event-type name

## Examples

Create an event:

    event: onSecondsElapsed autoquit=60 {
        out "You're 60 second timeout has elapsed. Quitting murex"
        exit 1
    }
    
Destroy an event:

    !event onSecondsElapsed autoquit

## Detail

The `interrupt` field in the CLI supports ANSI constants. eg

    event: onKeyPress f1={F1-VT100} {
        tout: qs HintText="Key F1 Pressed"
    }
    
To list compiled event types:

    » runtime: --events -> formap k v { out $k }
    onFileSystemChange
    onKeyPress
    onSecondsElapsed

## Synonyms

* `event`
* `!event`


## See Also

* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/formap](../commands/formap.md):
  
* [commands/open](../commands/open.md):
  
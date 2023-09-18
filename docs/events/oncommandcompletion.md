# `onCommandCompletion`

> Trigger an event upon a command's completion

## Description

`onCommandCompletion` events are triggered after a command has finished
executing in the interactive terminal.

Background processes or commands ran from inside aliases, functions, nested
blocks or from shell scripts cannot trigger this event. This is to protect
against accidental race conditions, infinite loops and breaking expected
behaviour / the portability of Murex scripts. On those processes directly ran
from the prompt can trigger this event.

## Usage

```
event onCommandCompletion name=command { code block }

!event onCommandCompletion name
```
The following payload is passed to the function via STDIN:

```
{
    "Name": "",
    "Interrupt": {
        "Command": "",
        "Parameters": [],
        "Stdout": "",
        "Stderr": "",
        "ExitNum": 0
    }
}
```

## Valid Interrupts

* `<command>`
    Name of command that triggers this event

## Payload

### Name

This is the name you specified when defining the event.

### Command

Name of command executed prior to this event being triggered

### Operation

The commandline parameters of the aforementioned command

### Stdout

This is the name of the Murex named pipe which contains a copy of the STDOUT
from the command which executed prior to this event.

You can read this with `read-named-pipe`. eg

```
» <stdin> -> set: event
» read-named-pipe: $event.Interrupt.Stdout -> ...
```

### Stderr

This is the name of the Murex named pipe which contains a copy of the STDERR
from the command which executed prior to this event.

You can read this with `read-named-pipe`. eg

```
» <stdin> -> set: event
» read-named-pipe: $event.Interrupt.Stderr -> ...
```

### ExitNum

This is the exit number returned from the executed command.

## Examples

**Read STDERR:**

In this example we check the output from `pacman`, which is ArchLinux's package
management tool, to see if you have accidentally ran it as a non-root user. If
the STDERR contains a message saying you are no root, then this event function
will re-run `pacman` with `sudo`.

```
event onCommandCompletion sudo-pacman=pacman {
    <stdin> -> set event
    read-named-pipe $event.Interrupt.Stderr \
    -> regexp 'm/error: you cannot perform this operation unless you are root/' \
    -> if {
          sudo pacman @event.Interrupt.Parameters
       }
}
```

## Detail

### Stdout

Stdout is written to the terminal. So this can be used to provide multiple
additional lines to the prompt since readline only supports one line for the
prompt itself and three extra lines for the hint text.

## See Also

* [`onPrompt`](../events/onprompt.md):
  Events triggered by changes in state of the interactive shell
* [alias](../events/alias.md):
  
* [config](../events/config.md):
  
* [event](../events/event.md):
  
* [function](../events/function.md):
  
* [if](../events/if.md):
  
* [namedpipes](../events/namedpipes.md):
  
* [read-named-pipe](../events/read-named-pipe.md):
  
* [regexp](../events/regexp.md):
  
* [stdin](../events/stdin.md):
  

<hr/>

This document was generated from [builtins/events/onCommandCompletion/oncommandcompletion_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onCommandCompletion/oncommandcompletion_doc.yaml).
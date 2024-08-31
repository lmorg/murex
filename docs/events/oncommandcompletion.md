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

## Valid Interrupts

* `<command>`
    Name of command that triggers this event

## Payload

The following payload is passed to the function via stdin:

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

### Name

This is the name you specified when defining the event.

### Interrupt/Command

Name of command executed prior to this event being triggered.

### Interrupt/Parameters

The command line parameters of the aforementioned command.

This will be an array of strings, like `@ARGV`.

### Interrupt/Stdout

This is the name of the Murex named pipe which contains a copy of the stdout
from the command which executed prior to this event.

You can read this with `read-named-pipe`. eg

```
» <stdin> -> set: event
» read-named-pipe: $event.Interrupt.Stdout -> ...
```

### Interrupt/Stderr

This is the name of the Murex named pipe which contains a copy of the stderr
from the command which executed prior to this event.

You can read this with `read-named-pipe`. eg

```
» <stdin> -> set: event
» read-named-pipe: $event.Interrupt.Stderr -> ...
```

### Interrupt/ExitNum

This is the exit number returned from the executed command.

## Event Return

This event doesn't have any `$EVENT_RETURN` parameters.

## Examples

### Read stderr

In this example we check the output from `pacman`, which is ArchLinux's package
management tool, to see if you have accidentally ran it as a non-root user. If
the stderr contains a message saying you are no root, then this event function
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

### Standard out and error

Stdout and stderr are both written to the terminal's stderr.

## See Also

* [Alias Pointer (`alias`)](../commands/alias.md):
  Create an alias for a command
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Named Pipes](../user-guide/namedpipes.md):
  A detailed breakdown of named pipes in Murex
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Read From Stdin (`<stdin>`)](../parser/stdin.md):
  Read the stdin belonging to the parent code block
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`onPrompt`](../events/onprompt.md):
  Events triggered by changes in state of the interactive shell
* [read-named-pipe](../parser/namedpipe.md):
  Reads from a Murex named pipe

<hr/>

This document was generated from [builtins/events/onCommandCompletion/oncommandcompletion_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onCommandCompletion/oncommandcompletion_doc.yaml).
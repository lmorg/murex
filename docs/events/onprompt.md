# `onPrompt`

> Events triggered by changes in state of the interactive shell

## Description

`onPrompt` events are triggered by changes in state of the interactive shell
(often referred to as _readline_). Those states are defined in the interrupts
section below.

## Usage

```
event onPrompt name=(before|after|abort|eof) { code block }

!event onPrompt name[.before|.after|.abort|.eof]
```

## Valid Interrupts

* `abort`
    Triggered if `ctrl`+`c` pressed while in the interactive prompt
* `after`
    Triggered after user has written a command into the interactive prompt and then hit `enter`
* `before`
    Triggered before readline displays the interactive prompt
* `eof`
    Triggered if `ctrl`+`d` pressed while in the interactive prompt

## Payload

The following payload is passed to the function via stdin:

```
{
    "Name": "",
    "Interrupt": {
        "Name": "",
        "Operation": "",
        "CmdLine": ""
    }
}
```

### Name

This is the **namespaced** name -- ie the name and operation.

### Interrupt/Name

This is the name you specified when defining the event.

### Interrupt/Operation

This is the interrupt you specified when defining the event.

Valid interrupt operation values are specified below.

### Interrupt/CmdLine

This is the command line executed, ie what you typed into the _readline_ prompt.

Please note this is only populated if the interrupt is **after**.

## Event Return

This event doesn't have any `$EVENT_RETURN` parameters.

## Examples

### Interrupt 'before'

```
event onPrompt example=before {
    out "This will appear before your command prompt"
}
```

### Interrupt 'after'

```
event onPrompt example=after {
    out "This will appear after you've hit [enter] on your command prompt"
    out "...but before the command executes"
}
```

### Echo the command line

```
» event onPrompt echo=after { -> set event; out $event.Interrupt.CmdLine }
» echo hello world
echo hello world
hello world
```

## Detail

### Standard out and error

Stdout and stderr are both written to the terminal. So this event can be used
to provide multiple additional lines to the prompt since readline only supports
one line for the prompt itself and three extra lines for the hint text.

### Order of execution

Interrupts are run in alphabetical order. So an event named "alfa" would run
before an event named "zulu". If you are writing multiple events and the order
of execution matters, then you can prefix the names with a number, eg `10_jump`

### Namespacing

This event is namespaced as `$(NAME).$(OPERATION)`.

For example, if an event in `onPrompt` was defined as `example=eof` then its
namespace would be `example.eof` and thus a subsequent event with the same name
but different operation, eg `example=abort`, would not overwrite the former
event defined against the interrupt `eof`.

The reason for this namespacing is because you might legitimately want the same
name for different operations (eg a smart prompt that has elements triggered
from different interrupts).

## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`onCommandCompletion`](../events/oncommandcompletion.md):
  Trigger an event upon a command's completion
* [`onKeyPress`](../events/onkeypress.md):
  Custom definable key bindings and macros
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation

<hr/>

This document was generated from [builtins/events/onPrompt/onprompt_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onPrompt/onprompt_doc.yaml).
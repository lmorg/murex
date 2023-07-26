# `onPrompt` - events

> Events triggered by changes in state of the interactive shell

## Description

`onPrompt` events are triggered by changes in state of the interactive shell
(often referred to as _readline_). Those states are defined in the interrupts
section below.

## Usage

```
event: onPrompt name=[before|after|abort|eof] { code block }

!event: onPrompt [before_|after_|abort_|eof_]name
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

## Examples

**Interrupt 'before':**

```
event: onPrompt example=before {
    out: "This will appear before your command prompt"
}
```

**Interrupt 'after':**

```
event: onPrompt example=after {
    out: "This will appear after you've hit [enter] on your command prompt"
    out: "...but before the command executes"
}
```

**Echo the command line:**

```
» event: onPrompt echo=after { -> set event; out $event.Interrupt.CmdLine }
» echo hello world
echo hello world
hello world
```

## Detail

### Payload

The following payload is passed to the function via STDIN:

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

#### Name

This is the **namespaced** name -- ie the name and operation.

#### Interrupt/Name

This is the name you specified when defining the event.

#### Operation

This is the interrupt you specified when defining the event.

Valid interrupt operation values are specified below.

#### CmdLine

This is the commandline you typed in the prompt.

Please note this is only populated if the interrupt is **after**.

### Stdout

Stdout is written to the terminal. So this can be used to provide multiple
additional lines to the prompt since readline only supports one line for the
prompt itself and three extra lines for the hint text.

### Order of execution

Interrupts are run in alphabetical order. So an event named "alfa" would run
before an event named "zulu". If you are writing multiple events and the order
of execution matters, then you can prefix the names with a number, eg `10_jump`

### Namespacing

The `onPrompt` event differs a little from other events when it comes to the
namespacing of interrupts. Typically you cannot have multiple interrupts with
the same name for an event. However with `onPrompt` their names are further 
namespaced by the interrupt name. In layman's terms this means `example=before`
wouldn't overwrite `example=after`.

The reason for this namespacing is because, unlike other events, you might
legitimately want the same name for different interrupts (eg a smart prompt
that has elements triggered from different interrupts).

## See Also

* [Murex's Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`onCommandCompletion`](../events/oncommandcompletion.md):
  Trigger an event upon a command's completion
* [onkeypress](../events/onkeypress.md):
  
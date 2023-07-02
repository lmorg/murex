# `onPrompt` - events

> Changes in state of the interactive shell

## Description

`onPrompt` events triggered by changes in state of the interactive shell
(often referred to as _readline_).

## Usage

    event: onPrompt name=[before|after|abort|eof] { code block }
    
    !event: onPrompt [before_|after_|abort_|eof_]name

## Valid Interrupts

* `abort`
    Triggered if `ctrl`+`c` pressed while in the interactive prompt
* `after`
    Triggered after user has written a command into the interactive prompt and then hit `enter
* `before`
    Triggered before readline displays the interactive prompt
* `eof`
    Triggered if `ctrl`+`d` pressed while in the interactive prompt

## Detail

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
namespaced by the interrupt name. In laymans terms this means `example=before`
wouldn't overwrite `example=after`. The reason for this namespacing is because,
unlike other events, you might legitimately want the same name for different
interrupts (eg a smart prompt that has elements triggered from different 
interrupts).
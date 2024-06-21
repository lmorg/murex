# `onKeyPress`

> TODO

## Description

TODO

## Usage

TODO

## Payload

TODO

## Event Return

This event doesn't have any `$EVENT_RETURN` parameters.

## Detail

### Standard out and error

Stdout and stderr are both written to the terminal.

### Order of execution

Interrupts are run in alphabetical order. So an event named "alfa" would run
before an event named "zulu". If you are writing multiple events and the order
of execution matters, then you can prefix the names with a number, eg `10_jump`

### Namespacing

With `onPrompt`, an event is namespaced as `$(NAME).$(OPERATION)`. For example,
if an event in `onPrompt` was defined as `example=eof` then its namespace would
be `example.eof` and thus a subsequent event with the same name but different
operation, eg `example=abort`, would not overwrite the former event defined
against the interrupt `eof`.

The reason for this namespacing is because you might legitimately want the same
name for different operations (eg a smart prompt that has elements triggered
from different interrupts).

## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`onCommandCompletion`](../events/oncommandcompletion.md):
  Trigger an event upon a command's completion
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation
* [`onPrompt`](../events/onprompt.md):
  Events triggered by changes in state of the interactive shell

<hr/>

This document was generated from [builtins/events/onKeyPress/onkeypress_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onKeyPress/onkeypress_doc.yaml).
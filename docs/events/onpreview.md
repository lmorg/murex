# `onPreview`

> Full screen previews for files and command documentation

## Description

Murex's readline API supports {{bookmark "full screen previews" "interactive-shell" "autocomplete-preview"}}.
For example when autocompleting command line parameters, if that parameter is a
file then Murex can preview the contents if it is a text file or even an image.

This preview can also provide guides to command usage. Such as `man` pages or
AI generated cheatsheets.

## Usage

```
event onPreview name=(function|builtin|exec) { code block }

!event onPreview name[.function|.builtin|.exec]
```

## Valid Interrupts

* `builtin`
    Code to execute when previewing a builtin (for example, a `murex-docs` page)
* `exec`
    Code to execute when previewing an external executable (for example, a `man` page)
* `function`
    Code to execute when previewing a Murex function (for example, the function source code)

## Payload

The following payload is passed to the function via STDIN:

```
{
    "Name": "",
    "Interrupt": {
        "Name": "",
        "Operation": "",
        "PreviewItem": "",
        "CmdLine": "",
        "Width": 80
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

### Interrupt/PreviewItem

This will be the command name. For example if the command line is
`sudo apt-get update` then the **PreviewItem** value will be `sudo`.

### Interrupt/CmdLine

This is the full command line in the preview prompt (ie what you've typed).

### Interrupt/Width

Width of the preview pane. Please note that this will differ from the terminal
width due to borders surrounding the preview pane.

## Examples

### Creating an event

```
event onPreview example=exec {
    -> set event
    out "Preview event for $(event.Interrupt.PreviewItem)"
    
    $.CacheTTL = 0 # don't cache this response.
}
```

### ChatGPT

Murex's {{link "ChatGPT integration" "chatgpt"}} also uses this event.
The [source code can be found on Github](https://github.com/lmorg/murex/blob/master/integrations/chatgpt_any.mx),
of viewed from the terminal via:

```
runtime --events -> [[ /onPreview/chatgpt.exec/Block ]]
```

## Detail

### Meta values

Meta values are variables named `$.` that store a structure that is sometimes
writable to. ({{link "read more" "meta-values"}})

The `onPreview` event uses meta values as an API to write data back to the
event caller.

The meta values available to `onPreview` after:

```
{
    "CacheTTL": 0
}
```

#### $.CacheTTL

This just defines how long to cache the results for this `onPreview` event for
faster loading of `onPreview` events in the future.

### Standard

Stdout and stderr are both written to the preview pane. Output is stripped of
any ANSI escape sequences and stderr isn't written in red.

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

* [ChatGPT Integration](../user-guide/chatgpt.md):
  How to enable ChatGPT hints
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Meta Values (json)](../variables/meta-values.md):
  State information for iteration blocks
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`onCommandCompletion`](../events/oncommandcompletion.md):
  Trigger an event upon a command's completion
* [`onKeyPress`](../events/onkeypress.md):
  TODO
* [`onPrompt`](../events/onprompt.md):
  Events triggered by changes in state of the interactive shell

<hr/>

This document was generated from [builtins/events/onPreview/onpreview_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onPreview/onpreview_doc.yaml).
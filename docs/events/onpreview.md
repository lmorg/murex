# `onPreview`

> Full screen previews for files and command documentation

## Description

Murex's readline API supports [full screen previews](/docs/user-guide/interactive-shell.md#autocomplete-preview).
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

The following payload is passed to the function via stdin:

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

## Event Return

`$EVENT_RETURN`, is a special variable that stores a writable structure to
return back to the event caller.

The `$EVENT_RETURN` values available for this event are:

```
{
    "CacheCmdLine": false,
    "CacheTTL": 2592000,
    "Display": true,
}
```

### $EVENT_RETURN.CacheCmdLine

Should the cache be unique to the command or include the full command line? You
would generally only want **CacheCmdLine** to be `true` if the generated
preview is unique to the full command line (eg an AI generated page based on
the full command line) vs only specific to the command name (eg a `man` page).

### $EVENT_RETURN.CacheTTL

This just defines how long to cache the results for this `onPreview` event for
faster loading of `onPreview` events in the future.

**CacheTTL** takes an integer and is measured in seconds. It's default value is
30 days.

### $EVENT_RETURN.Display

Defines whenever to output this event invocation.

Defaults to `true`.

## Examples

### Creating a basic event

```
event onPreview example=exec {
    -> set event
    out "Preview event for $(event.Interrupt.PreviewItem)"
    
    $EVENT_RETURN.CacheTTL = 0 # don't cache this response.
}
```

### ChatGPT

Murex's [ChatGPT integration](/docs/integrations/chatgpt.md) also uses this event.
The [source code can be found on Github](https://github.com/lmorg/murex/blob/master/integrations/chatgpt_any.mx),
of viewed from the terminal via:

```
runtime --events -> [[ /onPreview/chatgpt.exec/Block ]]
```

## Detail

### Standard out and error

Stdout and stderr are both written to the preview pane. Output is stripped or
any ANSI escape sequences and stderr isn't written in red.

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

* [ChatGPT](../integrations/chatgpt.md):
  How to enable ChatGPT hints
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Man Pages (POSIX)](../integrations/man-pages.md):
  Linux/UNIX `man` page integrations
* [Murex's Offline Documentation (`murex-docs`)](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [Public Function (`function`)](../commands/function.md):
  Define a function block
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
* [`onPrompt`](../events/onprompt.md):
  Events triggered by changes in state of the interactive shell

<hr/>

This document was generated from [builtins/events/onPreview/onpreview_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onPreview/onpreview_doc.yaml).
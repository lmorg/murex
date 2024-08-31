# `event`

> Event driven programming for shell scripts

## Description

Create or destroy an event interrupt,

Each event will have subtly different behaviour depending on the event itself
due to the differing roles of each event system. Therefore it is recommended
that you read the docs on each event to understand its behaviour.

However while they might differ, the `event` API does try to retain a level of
external consistency. For example each event in defined via `name=interrupt`
where **name** is a user defined handle (like a variable or function would have
a name) and **interrupt** is a system state you wish the event to be fired on.

Each event function will have a payload sent via stdin which would look a
little like the following:

```
{
    "Name": "",
    "Interrupt": {}
}
```

**Name** will always refer to the name you passed when defining the event. And
**Interrupt** will carry any event specific metadata that might be useful to
the event function. Thus the value of **Interrupt** will vary from one event to
another.

## Usage

```
event event-type name=interrupt { code block }

!event event-type name
```

## Examples

### Create an event

```
event onSecondsElapsed autoquit=60 {
    out "You're 60 second timeout has elapsed. Quitting murex"
    exit 1
}
```

### Destroy an event

```
!event onSecondsElapsed autoquit
```

## Detail

### Supported events

* [`onCommandCompletion`](../events/oncommandcompletion.md):
  Trigger an event upon a command's completion
* [`onFileSystemChange`](../events/onfilesystemchange.md):
  Add a filesystem watch
* [`onKeyPress`](../events/onkeypress.md):
  Custom definable key bindings and macros
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation
* [`onPrompt`](../events/onprompt.md):
  Events triggered by changes in state of the interactive shell
* [`onSecondsElapsed`](../events/onsecondselapsed.md):
  Events triggered by time intervals
* [`onSignalReceived`](../events/onsignalreceived.md):
  Trap OS signals

### ANSI constants

The `interrupt` field in the CLI supports ANSI constants. eg

```
event onKeyPress f1={F1-VT100} {
    tout qs HintText="Key F1 Pressed"
}
```

### Compiled events

To list compiled event types:

```
» runtime --events -> formap event ! { out $event }
onCommandCompletion
onFileSystemChange
onKeyPress
onPrompt
onSecondsElapsed
```

## Synonyms

* `event`
* `!event`


## See Also

* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex

<hr/>

This document was generated from [builtins/events/events_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/events_doc.yaml).
# `onFileSystemChange`

> Add a filesystem watch

## Description

`onFileSystemChange` events are triggered whenever there is a change to a
watched path or file.

## Usage

```
event onFileSystemChange name=path { code block }

!event onFileSystemChange name
```

## Valid Interrupts

* `<path>`
    Path of directory or file to watch for filesystem events

## Payload

The following payload is passed to the function via stdin:

```
{
    "Name": "",
    "Interrupt": {
        "Path": "",
        "Operation": ""
    }
}
```

### Name

This is the name you specified when defining the event

### Interrupt/Path

The path of the file that has triggered the event

### Interrupt/Operation

This is the filesystem operation that triggered the event. The following
strings could be present in the **Operation** field:

* `create`: filesystem object created
* `remove`: filesystem object deleted
* `write`:  filesystem object has been written to
* `rename`: filesystem object has been renamed
* `chmod`:  filesystem object has had its POSIX permissions updated

Sometimes you might see more than one operation per interrupt. If that happens
the operation will be pipe delimited. For example `create|chmod`

## Event Return

This event doesn't have any `$EVENT_RETURN` parameters.

## Examples

This will automatically add any new files in your current working directory to
git upon file creation:

```
event onFileSystemChange example=. {
    -> set event
    if { $event.Interrupt.Operation =~ "create" } then {
        git add $event.Interrupt.Path
    }
}
```

## Detail

### Standard out and error

Stdout and stderr are both written to the terminal.

### POSIX only

At this stage, this event isn't available for Windows nor Plan 9. This is
chiefly down to a lack of testers on either platform so rather than release
untested and potentially broken code, the decision was made to restrict this
event to Linux, macOS and UNIX systems instead.

## See Also

* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [`event`](../commands/event.md):
  Event driven programming for shell scripts

<hr/>

This document was generated from [builtins/events/onFileSystemChange/onfilesystemchange_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onFileSystemChange/onfilesystemchange_doc.yaml).
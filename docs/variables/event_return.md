# `EVENT_RETURN` (json)

> Return values for events

## Description

Some events support return parameters outside of your typical stdout and stderr
streams. `$EVENT_RETURN` allows you to modify those parameters.

## Examples

```
event onPreview example=exec {
    -> set event
    out "Preview event for $(event.Interrupt.PreviewItem)"
    
    $EVENT_RETURN.CacheTTL = 0 # don't cache this response.
}
```

## Detail

`$EVENT_RETURN` will support different values for different events. Please read
the respective event document for details on using this variable.

## See Also

* [Murex Event Subsystem (`event`)](../commands/event.md):
  Event driven programming for shell scripts
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`onKeyPress`](../events/onkeypress.md):
  Custom definable key bindings and macros
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation

<hr/>

This document was generated from [gen/variables/EVENT_RETURN_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/EVENT_RETURN_doc.yaml).
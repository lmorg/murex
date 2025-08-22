```
{                                                                                                                 
    "Actions": [],
    "Continue": false,
    "SetCursorPos": 0,
    "SetHintText": "",
    "SetLine": ""
}
```

### $EVENT_RETURN.Actions

This is an array of strings, each defining a hotkey function to execute after
the event function has completed.

You can supply multiple functions in the array.

Supported values are:

```go
{{ file "builtins/events/onKeyPress/actions_generated.go" }}
```

### $EVENT_RETURN.Continue

This boolean value defines whether to execute other events for the same
keypress after this one.

On the surface of it, you might question why you'd wouldn't want multiple
actions bound to the same keypress. However since _readline_ supports different
operational modes, you might want different events bound to different states
of _readline_. In which case you'd need to set `$EVENT_RETURN.Continue` to
`true`.

### $EVENT_RETURN.SetCursorPos

This allows you to shift the text input cursor to an absolute location.

### $EVENT_RETURN.SetHintText

Forces a different message in the hint text.

### $EVENT_RETURN.SetLine

Change your input line in _readline_.
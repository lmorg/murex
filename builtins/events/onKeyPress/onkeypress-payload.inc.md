```
{
    "Name": "",
    "Interrupt": {
        "Line": "",
        "CursorPos": 0,
        "KeyPress": "",
        "IsMasked": false,
        "InputMode": "",
        "PreviewMode": ""
    }
}
```

### Name

This is the **namespaced** name -- ie the name and operation.

### Interrupt/Name

This is the name you specified when defining the event.

### Interrupt/Line

The current line as it appears in _readline_.

### Interrupt/CursorPos

Where the text input cursor is sat on the line

### Interrupt/KeyPress

The key which was pressed.

If the key stroke is represented by an ANSI escape sequence, then this field
will be multiple multiple characters long.

### Interrupt/IsMasked

This will be `true` if you have a password / input mask (eg `*`). Otherwise it
will be `false`.

### Interrupt/InputMode

This is the input mode of _readline_. Different input modes will utilise
keystrokes differently.

This field is a string and the following constants are supported:

* `Normal`:         regular input
* `VimKeys`:        where input behaves like `vim`
* `VimReplaceOnce`: `vim` mode, but next keystroke might normally overwrite
                    current character
* `VimReplaceMany`: `vim` mode, but where every keystroke overwrites characters
                    rather than inserts
* `VimDelete`:      `vim` mode, but where characters are deleted
* `Autocomplete`:   the autocomplete menu is shown
* `FuzzyFind`:      the autocomplete menu is shown with the Fuzzy Find input
                    enabled

More details about these modes can be found in the {{link "Terminal Hotkeys" "terminal-keys"}}
document.

### Interrupt/PreviewMode

Preview mode is independent to input mode.

* `Disabled`:     preview is not running
* `Autocomplete`: regular preview mode
* `CmdLine`:      command line preview

More details about these modes can be found in the {{bookmark "Interactive Shell" "interactive-shell" "preview"}}
document.
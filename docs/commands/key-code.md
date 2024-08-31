# ASCII And ANSI Escape Sequences (`key-code`)

> Returns character sequences for any key pressed (ie sent from the terminal)

## Description

`key-code` is a tool used for querying what byte sequence the terminal emulator

## Usage

```
key-code -> <stdout>

<stdin> -> key-code -> <stdout>
```

## Examples

### Typical use case

```
» key-code
Press any key to print its escape constants...
```

...then press [f9] and `key-code` returns...

```
ANSI Constants:   {F9}
Byte Sequence:    %[27 91 50 48 126]
Contains Unicode: false
```

### As a method

```
» tout str '{ESC}[20~' -> key-code
ANSI Constants:   {F9}
Byte Sequence:    %[27 91 50 48 126]
Contains Unicode: false
```

## Detail

### Redirection

If stdout is not a TTY then only the thing written is the ANSI Constant. This
is so that it can be used as a variable. eg

```
key-code -> set $key

event onKeyPress close=$key {
    exit
}
```

## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`onKeyPress`](../events/onkeypress.md):
  Custom definable key bindings and macros

<hr/>

This document was generated from [builtins/events/onKeyPress/keycodes_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onKeyPress/keycodes_doc.yaml).
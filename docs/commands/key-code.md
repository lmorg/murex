# `key-code`

> Returns character sequences for any key pressed (ie sent from the terminal)

## Description

`key-code` is a tool used for querying what byte sequence the terminal emulator

## Usage

```
key-code -> <stdout>
```

## Examples

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
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [onkeypress](../commands/onkeypress.md):
  

<hr/>

This document was generated from [builtins/events/onKeyPress/keycodes_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onKeyPress/keycodes_doc.yaml).
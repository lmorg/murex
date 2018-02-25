# _murex_ Language Guide

## Command reference: alter

> Change a value within a structured data-type and pass that change along then
pipeline without altering the original source input

### Description

`alter` a value within a structured data-type.

### Usage

    <stdin> -> alter: /path value -> <stdout>

### Examples

    » config: -> [ shell ] -> [ prompt ] -> alter: /Value moo
    {
        "Data-Type": "block",
        "Default": "{ out 'murex » ' }",
        "Description": "Interactive shell prompt.",
        "Value": "moo"
    }

> Please note: `alter` did not change the shell prompt value held inside `config`
  but instead took the STDOUT from `config`, altered a value and then passed that
  new complete structure through it's STDOUT.

The above example can also be performed without the index functions (`[`):

    config -> alter: /shell/prompt/Value moo

### Detail

#### Path

The path parameter can take any character as node separators. The separator is
assigned via the first character in the path. For example

    config -> alter: .shell.prompt.Value moo
    config -> alter: >shell>prompt>Value moo

Just make sure you quote or escape any characters used as shell tokens. eg

    config -> alter: '#shell#prompt#Value' moo
    config -> alter: ' shell prompt Value' moo

#### Supported data-types

You can check what data-types are available via the `runtime` command:

    runtime --marshallers

Marshallers are enabled at compile time from the `builtins/data-types` directory.

### See also

* [[]([.md)
* [append](append.md): Add data to the end of an array
* [cast](cast.md)
* [format](format.md)
* [prepend](prepend.md): Add data to the start of an array
* [runtime](runtime.md)

# Debugging Mode (`debug`)

> Debugging information

## Description

`debug` has two modes: as a function and as a method.

### Debug Method

This usage will return debug information about the previous function ran.

### Debug Function:

This will enable or disable debugging mode.

## Usage

### Enable or disable debug output

```
debug boolean -> <stdout>
```

### Output whether debug mode is enabled or disabled

```
debug -> <stdout>
```

### Run a pipeline with debug mode enabled

```
debug { code-block } -> <stdout>
```

### Print debugging information about the previous command

```
<stdin> -> debug -> <stdout>
```

## Examples

### Running a code block with debugging

```
» debug
false

» debug { out "debug mode is now ${debug}" }
debug mode is now true

» debug
false
```

### Debugging information on previous function

```
» out "hello, world!" -> debug 
{
    "DataType": {
        "Go": "[]string",
        "Murex": "str"
    },
    "Process": {
        "Context": {
            "Context": 0
        },
        "Stdin": {},
        "Stdout": {},
        "Stderr": {},
        "Parameters": {
            "Params": [
                "hello, world!"
            ],
            "Tokens": [
                [
                    {
                        "Type": 0,
                        "Key": ""
                    }
                ],
                [
                    {
                        "Type": 1,
                        "Key": "hello, world!"
                    }
                ],
                [
                    {
                        "Type": 0,
                        "Key": ""
                    }
                ]
            ]
        },
        "ExitNum": 0,
        "Name": "echo",
        "Id": 3750,
        "Exec": {
            "Pid": 0,
            "Cmd": null,
            "PipeR": null,
            "PipeW": null
        },
        "PromptGoProc": 1,
        "Path": "",
        "IsMethod": false,
        "IsNot": false,
        "NamedPipeOut": "out",
        "NamedPipeErr": "err",
        "NamedPipeTest": "",
        "State": 7,
        "IsBackground": false,
        "LineNumber": 1,
        "ColNumber": 1,
        "RunMode": 0,
        "Config": {},
        "Tests": {
            "Results": null
        },
        "Variables": {},
        "FidTree": [
            0,
            3750
        ],
        "CreationTime": "2019-01-20T00:00:52.167127131Z",
        "StartTime": "2019-01-20T00:00:52.167776212Z"
    }
}
```

### Enable or disable debug mode

```
» debug on
true

» debug off
false
```

### Print debug mode

Debug mode can be either enabled (`true`) or disabled (`false`):

```
» debug
false
```

## Detail

### Enable and Disable

When enabling or disabling debug mode, because the parameter is a murex
boolean type, it means you can use other boolean terms. eg

```
# enable debugging
» debug 1
» debug on
» debug yes
» debug true

# disable debugging
» debug 0
» debug off
» debug no
» debug false
```

It is also worth noting that the debugging information needs to be written
into the Go source code rather than in Murex's shell scripting language.
If you require debugging those processes then please use Murex's `test`
framework

### Generating a Panic

For testing purposes, you might want to force Murex to crash. You can do
this via:

```
debug panic
```

## Synonyms

* `debug`


## See Also

* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Shell Script Tests (`test`)](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts

<hr/>

This document was generated from [builtins/core/management/debug_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/debug_doc.yaml).
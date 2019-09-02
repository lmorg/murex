# _murex_ Shell Guide

## Command Reference: `debug`

> Debugging information

### Description

`debug` has two modes: as a function and as a method.

#### Debug Method

This usage will return debug information about the previous function ran.

#### Debug Function:

This will enable or disable debugging mode.

### Usage

    <stdin> -> debug -> <stdout>
    
    debug: boolean -> <stdout>
    
    debug -> <stdout>

### Examples

Return debugging information on the previous function:

    » echo: "hello, world!" -> debug 
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
    
Enable or disable debug mode:

    » debug: on
    true
    
    » debug: off
    false
    
Output whether debug mode is enabled or disabled:

    » debug
    false

### Detail

When enabling or disabling debug mode, because the parameter is a murex
boolean type, it means you can use other boolean terms. eg

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
    
It is also worth noting that the debugging information needs to be written
into the Go source code rather than in _murex_'s shell scripting language.
If you require debugging those processes then please use _murex_'s `test`
framework

### See Also

* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/test](../commands/test.md):
  
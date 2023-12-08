# `runtime`

> Returns runtime information on the internal state of Murex

## Description

`runtime` is a tool for querying the internal state of Murex. It's output
will be JSON dumps.

## Usage

```
runtime flags -> <stdout>
```

`builtins` is an alias for `runtime --builtins`:

```
builtins -> <stdout>
```

## Examples

List all the builtin data-types that support WriteArray()

```
» runtime --writearray
[
    "*",
    "commonlog",
    "csexp",
    "hcl",
    "json",
    "jsonl",
    "qs",
    "sexp",
    "str",
    "toml",
    "yaml"
]
```

List all the functions

```
» runtime --functions -> [ agent aliases ]
[
    {
        "Block": "\n    # Launch ssh-agent\n    ssh-agent -\u003e head -n2 -\u003e [ :0 ] -\u003e prefix \"export \" -\u003e source\n    ssh-add: @{g \u003c!null\u003e ~/.ssh/*.key} @{g \u003c!null\u003e ~/.ssh/*.pem}\n",
        "FileRef": {
            "Column": 1,
            "Line": 149,
            "Source": {
                "DateTime": "2019-07-07T14:06:11.05581+01:00",
                "Filename": "/home/lau/.murex_profile",
                "Module": "profile/.murex_profile"
            }
        },
        "Summary": "Launch ssh-agent"
    },
    {
        "Block": "\n\t# Output the aliases in human readable format\n\truntime --aliases -\u003e formap name alias {\n        $name -\u003e sprintf \"%10s =\u003e ${esccli @alias}\\n\"\n\t} -\u003e cast str\n",
        "FileRef": {
            "Column": 1,
            "Line": 6,
            "Source": {
                "DateTime": "2019-07-07T14:06:10.886706796+01:00",
                "Filename": "(builtin)",
                "Module": "source/builtin"
            }
        },
        "Summary": "Output the aliases in human readable format"
    }
]
```

To get a list of every flag supported by `runtime`

```
» runtime --help
[
    "--aliases",
    "--config",
    "--debug",
    "--events",
    "--fids",
    "--flags",
    "--functions",
    "--help",
    "--indexes",
    "--marshallers",
    "--memstats",
    "--modules",
    "--named-pipes",
    "--open-agents",
    "--pipes",
    "--privates",
    "--readarray",
    "--readmap",
    "--sources",
    "--test-results",
    "--tests",
    "--unmarshallers",
    "--variables",
    "--writearray"
]
```

Please also note that you can supply more than one flag. However when you
do use multiple flags the top level of the JSON output will be a map of the
flag names. eg

```
» runtime --pipes --tests
{
    "pipes": [
        "file",
        "std",
        "tcp-dial",
        "tcp-listen",
        "udp-dial",
        "udp-listen"
    ],
    "tests": {
        "state": {},
        "test": []
    }
}

» runtime --pipes
[
    "file",
    "std",
    "tcp-dial",
    "tcp-listen",
    "udp-dial",
    "udp-listen"
]

» runtime --tests
{
    "state": {},
    "test": []
}
```

## Flags

* `--aliases`
    Lists all aliases
* `--autocomplete`
    Lists all `autocomplete` schemas - both user defined and automatically generated one
* `--builtins`
    Lists all builtin commands, compiled into Murex
* `--config`
    Lists all properties available to `config`
* `--debug`
    Outputs the state of debug and inspect mode
* `--events`
    Lists all builtin event types and any defined events
* `--exports`
    Outputs environmental variables. For Murex variables (`global` and `set`/`let`) use `--variables`
* `--fids`
    Lists all running processes / functions
* `--functions`
    Lists all Murex global functions
* `--globals`
    Lists all global variables
* `--help`
    Outputs a list of `runtimes`'s flags
* `--indexes`
    Lists all builtin data-types which are supported by index (`[`)
* `--marshallers`
    Lists all builtin data-types with marshallers (eg required for `format`)
* `--memstats`
    Outputs the running state of Go's runtime
* `--methods`
    Lists all commands with a defined STDOUT and STDIN data type. This is used to generate smarter autocompletion suggestions with `->`
* `--modules`
    Lists all installed modules
* `--named-pipes`
    Lists all named pipes defined
* `--not-indexes`
    Lists all builtin data-types which are supported by index (`![`)
* `--open-agents`
    Lists all registered `open` handlers 
* `--pipes`
    Lists builtin pipes compiled into Murex. These can be then be defined as named-pipes
* `--privates`
    Lists all Murex private functions
* `--readarray`
    Lists all builtin data-types which support ReadArray()
* `--readarraywithtype`
    Lists all builtin data-types which support ReadArrayWithType()
* `--readmap`
    Lists all builtin data-types which support ReadMap()
* `--sources`
    Lists all loaded murex sources
* `--summaries`
    Outputs all the override summaries 
* `--test-results`
    A dump of any unreported test results
* `--tests`
    Lists defined tests
* `--unmarshallers`
    Lists all builtin data-types with unmarshallers (eg required for `format`)
* `--variables`
    Lists all local Murex variables which doesn't include environmental nor global variables
* `--writearray`
    Lists all builtin data-types which support WriteArray()

## Detail

### Usage in scripts

`runtime` should not be used in scripts because the output of `runtime` may
be subject to change as and when the internal mechanics of Murex change.
The purpose behind `runtime` is not to provide an API but rather to provide
a verbose "dump" of the internal running state of Murex.

If you require a stable API to script against then please use the respective
command line tool. For example `fid-list` instead of `runtime --fids`. Some
tools will provide a human readable output when STDOUT is a TTY but output
a script parsable version when STDOUT is not a terminal.

```
» fid-list
    FID   Parent    Scope  State         Run Mode  BG   Out Pipe    Err Pipe    Command     Parameters
      0        0        0  Executing     Shell     no                           -murex
 265499        0        0  Executing     Normal    no   out         err         fid-list

» fid-list -> pretty
[
    {
        "FID": 0,
        "Parent": 0,
        "Scope": 0,
        "State": "Executing",
        "Run Mode": "Shell",
        "BG": false,
        "Out Pipe": "",
        "Err Pipe": "",
        "Command": "-murex",
        "Parameters": ""
    },
    {
        "FID": 265540,
        "Parent": 0,
        "Scope": 0,
        "State": "Executing",
        "Run Mode": "Normal",
        "BG": false,
        "Out Pipe": "out",
        "Err Pipe": "err",
        "Command": "fid-list",
        "Parameters": ""
    },
    {
        "FID": 265541,
        "Parent": 0,
        "Scope": 0,
        "State": "Executing",
        "Run Mode": "Normal",
        "BG": false,
        "Out Pipe": "out",
        "Err Pipe": "err",
        "Command": "pretty",
        "Parameters": ""
    }
]
```

### File reference

Some of the JSON dumps produced from `runtime` will include a map called
`FileRef`. This is a trace of the source file that defined it. It is used
by Murex to help provide meaningful errors (eg with line and character
positions) however it is also useful for manually debugging user-defined
properties such as which module or script defined an `autocomplete` schema.

### Debug mode

When `debug` is enabled garbage collection is disabled for variables and
FIDs. This means the output of `runtime --variables` and `runtime --fids`
will contain more than just the currently defined variables and running
functions.

## Synonyms

* `runtime`
* `builtins`


## See Also

* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`debug`](../commands/debug.md):
  Debugging information
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`export`](../commands/export.md):
  Define an environmental variable and set it's value
* [`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [`foreach`](../commands/foreach.md):
  Iterate through an array
* [`formap`](../commands/formap.md):
  Iterate through a map or other collection of data
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`function`](../commands/function.md):
  Define a function block
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`method`](../commands/method.md):
  Define a methods supported data-types
* [`open`](../commands/open.md):
  Open a file with a preferred handler
* [`openagent`](../commands/openagent.md):
  Creates a handler function for `open`
* [`pipe`](../commands/pipe.md):
  Manage Murex named pipes
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`private`](../commands/private.md):
  Define a private function block
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`source`](../commands/source.md):
  Import Murex code from another file of code block
* [`test`](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts

<hr/>

This document was generated from [builtins/core/runtime/runtime_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/runtime/runtime_doc.yaml).
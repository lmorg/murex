# Shell Runtime (`runtime`)

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
runtime --help
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
    Lists all [aliases](/docs/commands/alias.md)
* `--autocomplete`
    Lists all `autocomplete` schemas - both [user defined](/docs/commands/autocomplete.md) and automatically generated ones
* `--builtins`
    Lists all builtin commands, compiled into Murex
* `--cache`
    Returns a complete dump of everything in the local cache as well as cache DB (if compiled with persistent sqlite3 cache support)
* `--cache-db-enabled`
    Returns boolean value stating if the cache DB is enabled
* `--cache-db-path`
    Returns a string representation of the cache DB path (or a zero length string if cache DB support is not compiled)
* `--cache-namespaces`
    Lists namespaces in cache
* `--clear-cache`
    Clears all items from both the local cache and cache DB. Returns all items removed
* `--config`
    Lists all properties available to [`config`](/docs/commands/config.md)
* `--debug`
    Outputs the state of [debug mode](/docs/commands/debug.md)
* `--event-types`
    Lists all builtin event types
* `--events`
    Lists all builtin event types and any [defined events](/docs/commands/event.md)
* `--exports`
    Outputs [environmental variables](/docs/commands/export.md)
* `--fids`
    Lists all running processes / functions
* `--functions`
    Lists all Murex [global functions](/docs/commands/function.md)
* `--globals`
    Lists all [global variables](/docs/commands/global.md)
* `--go-gc`
    Forces the Go runtime to run its garbage collection
* `--help`
    Outputs a list of `runtimes`'s flags
* `--indexes`
    Lists all builtin data-types which are supported by [index](/docs/parser/item-index.md) (`[`)
* `--integrations`
    Lists all compiled [integrations](/docs/user-guide/integrations.md)
* `--marshallers`
    Lists all builtin data-types with marshallers (eg required for [`format`](/docs/commands/format.md))
* `--memstats`
    Outputs the running state of Go's runtime
* `--methods`
    Lists all commands with a defined stdout and stdin data type. This is used to generate smarter autocompletion suggestions with `->`
* `--modules`
    Lists all installed [modules](/docs/user-guide/modules.md)
* `--named-pipes`
    Lists all [named pipes defined](/docs/commands/pipe.md)
* `--not-indexes`
    Lists all builtin data-types which are supported by [not-index](/docs/parser/item-index.md) (`![`)
* `--open-agents`
    Lists all registered [`open`](/docs/commands/open.md) handlers (defined with [`openagent`](/docs/commands/openagent.md))
* `--pipes`
    Lists builtin pipes compiled into Murex. These can be then be defined as [named-pipes](/docs/user-guide/namedpipes.md)
* `--privates`
    Lists all Murex [private functions](/docs/commands/private.md)
* `--readarray`
    Lists all builtin data-types which support ReadArray()
* `--readarraywithtype`
    Lists all builtin data-types which support ReadArrayWithType()
* `--readmap`
    Lists all builtin data-types which support ReadMap()
* `--sources`
    Lists all [loaded murex sources](/docs/user-guide/fileref.md)
* `--summaries`
    Outputs all the [override summaries](/docs/commands/summary.md)
* `--test-results`
    A dump of any unreported [test](/docs/commands/test.md) results
* `--tests`
    Lists [defined tests](/docs/commands/test.md)
* `--trim-cache`
    Clears out-of-date items from both the local cache and cache DB. Returns all items removed
* `--unmarshallers`
    Lists all builtin data-types with unmarshallers (eg required for [`format`](/docs/commands/format.md))
* `--variables`
    Lists all [local variables](/docs/commands/set.md) (excludes [environmental](/docs/commands/export.md) and [global variables](/docs/commands/global.md))
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
tools will provide a human readable output when stdout is a TTY but output
a script parsable version when stdout is not a terminal.

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
* `shell.runtime`


## See Also

* [Create Named Pipe (`pipe`)](../commands/pipe.md):
  Manage Murex named pipes
* [Debugging Mode (`debug`)](../commands/debug.md):
  Debugging information
* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Define Global (`global`)](../commands/global.md):
  Define a global variable and set it's value
* [Define Handlers For "`open`" (`openagent`)](../commands/openagent.md):
  Creates a handler function for `open`
* [Define Method Relationships (`method`)](../commands/method.md):
  Define a methods supported data-types
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Integrations](../user-guide/integrations.md):
  Default integrations shipped with Murex
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Prettify JSON](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Shell Script Tests (`test`)](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)

<hr/>

This document was generated from [builtins/core/runtime/runtime_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/runtime/runtime_doc.yaml).
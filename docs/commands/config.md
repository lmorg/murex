# Shell Configuration And Settings (`config`)

> Query or define Murex runtime settings

## Description

Rather than Murex runtime settings being definable via obscure environmental
variables, Murex instead supports a registry of config defined via the
`config` command. This means any preferences and/or runtime config becomes
centralised and discoverable.

## Terminology

The following terms have special means with regards to `config`:

### app

_app_ refers to a grouped category of settings. For example the name of an
integration or builtin.
    
Other _app_ names include

* `shell`: for "global" (system wide) Murex settings
* `proc`: for scoped Murex settings
* `http`: for settings that are applied to any processes which use the builtin
    HTTP user agent (eg `open`, `get`, `getfile`, `post`)
* `test`: settings for Murex's test frameworks
* `index`: settings for `[` (index)

### key

_key_ refers to the config setting itself. For example the _app_ might be `http`
but the _key_ might be `timeout` - where the _key_, in this instance, holds the
value for how long any HTTP user agents might wait before timing out.

### value

_value_ is the actual value of a setting. So the value for _app_: `http`, _key_:
`timeout` might be `10`. eg

```
» config get http timeout
10
```

## Usage

### Get value

```
config get app key -> <stdout>
```

### Set value

```
           config set  app key value
<stdin> -> config set  app key
           config eval app key { code-block }
```

### Reset to default

```
!config app key
config default app key
```

### Define custom configs

```
config define app key { json }
```

## Examples

### eval

Using `eval` to append to an array (in this instance, adding a function
name to the list of "safe" commands):

```
» config eval shell safe-commands { -> append function-name }
```

You could also use the `~>` operator too:

```
» config eval shell safe-commands { ~> %[function-name] }
```

## Flags

* `default`
    Reset a the value of _app_'s _key_ to its default _value_ (the default _value_ is defined by the same process that defines the config field)
* `define`
    Allows you to create custom config options. See [Custom Config Directives](/docs/commands/config.md#custom-config-directives) to learn how to use `define`
* `get`
    Output the currently held config _value_ without changing it
* `set`
    Change the _value_ of an _app_'s _key_. `set` does not print any output

## Detail

### scope

Settings in `config`, by default, are scoped per function and module. Any
functions called will inherit the settings of it's caller parent. However any
child functions that then change the settings will only change settings for it's
own function and not the parent caller.

### global

Global settings defined inside a function will affect settings queried inside
another executing function (same concept as global variables).

## Custom Config Directives

> This section relates to creating custom configs via `config define`.
> You do not need to refer to this for any regular usage of `config`.

<div id="toc">

- [DataType](#datatype)
- [Description](#description)
- [Global](#global)
- [Default](#default)
- [Options](#options)
- [Dynamic](#dynamic)
  - [Dynamic Read](#dynamic-read)
  - [Dynamic Write](#dynamic-write)

</div>


Where "default value" is what will be auto-populated if you don't include that
directive (or "required" if the directive must be included).

### DataType

> Value: `str` (required)

This is the Murex data-type for the value.

### Description

> Value: `str` (required)

Description is a required field to force developers into writing meaning hints
enabling the discoverability of settings within Murex.

### Global

> Value: `bool` (default: `false`)

This defines whether this setting is global or scoped.

All **Dynamic** config must also be **Global**. This is because **Dynamic**
config rely on a state that likely isn't scoped (eg the contents of a file on
disk or environmental variable).

### Default

> Value: any (required)

This is the initialized and default value.

### Options

> Value: array (default: `null`)

Some suggested options (if known) to provide as autocompletion suggestions in
the interactive command line.

### Dynamic

> Value: map of strings (default: `null`)

Only use this if config options need to be more than just static values stored
inside Murex's runtime. Using **Dynamic** means `autocomplete get app key`
and `autocomplete set app key value` will spawn off a subshell running a code
block defined from the `Read` and `Write` mapped values. eg

```
# Create the example config file
out "this is the default value" |> example.conf

config define example test5 %{
    Description: This is only an example
    DataType: str
    Global: true
    Dynamic: {
        Read: '{
            open example.conf
        }'
        Write: '{
            |> example.conf
        }'
    },
    
    # read the config file to get the default value
    Default: ${open example.conf}
}
```

It's also worth noting the different syntax between **Read** and **Default**.
The **Read** code block is being executed when the **Read** directive is being
requested, whereas the **Default** code block is being executed when the JSON
is being read.

In technical terms, the **Default** code block is being executed by Murex 
when `config define` is getting executed where as the **Read** and **Write**
code blocks are getting stored as a JSON string and then executed only when
those hooks are getting triggered.

#### Dynamic Read

> Value: `str` (default: empty)

This is executed when `autocomplete get app key` is ran. The stdout of the code
block is the setting's value.

#### Dynamic Write

> Value: `str` (default: empty)

This is executed when `autocomplete` is setting a value (eg `set`, `default`,
`eval`). is ran. The stdin of the code block is the new value.

## Synonyms

* `config`
* `!config`


## See Also

* [Alter Data Structure (`alter`)](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Download File (`getfile`)](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Get Request (`get`)](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [Match String (`match`)](../commands/match.md):
  Match an exact value in an array
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Post Request (`post`)](../commands/post.md):
  HTTP POST request with a JSON-parsable return
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [`%{}` Object Builder](../parser/create-object.md):
  Quickly generate objects (dictionaries / maps)
* [`event`](../commands/event.md):
  Event driven programming for shell scripts

<hr/>

This document was generated from [builtins/core/config/config_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/config/config_doc.yaml).
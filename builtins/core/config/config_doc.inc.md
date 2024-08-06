> This section relates to creating custom configs via `config define`.
> You do not need to refer to this for any regular usage of `config`.

{{ if env "DOCGEN_TARGET=" }}<div id="toc">

- [DataType](#datatype)
- [Description](#description)
- [Global](#global)
- [Default](#default)
- [Options](#options)
- [Dynamic](#dynamic)
  - [Dynamic Read](#dynamic-read)
  - [Dynamic Write](#dynamic-write)

</div>
{{ end }}

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

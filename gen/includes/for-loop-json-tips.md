#### Tips when writing JSON inside for loops

One of the drawbacks (or maybe advantages, depending on your perspective) of
JSON is that parsers generally expect a complete file for processing in that
the JSON specification requires closing tags for every opening tag. This means
it's not always suitable for streaming. For example

```
» ja [1..3] -> foreach i { out ({ "$i": $i }) }
{ "1": 1 }
{ "2": 2 }
{ "3": 3 }
```

**What does this even mean and how can you build a JSON file up sequentially?**

One answer if to write the output in a streaming file format and convert back
to JSON

```
» ja [1..3] -> foreach i { out (- "$i": $i) }
- "1": 1
- "2": 2
- "3": 3

» ja [1..3] -> foreach i { out (- "$i": $i) } -> cast yaml -> format json
[
    {
        "1": 1
    },
    {
        "2": 2
    },
    {
        "3": 3
    }
]
```

**What if I'm returning an object rather than writing one?**

The problem with building JSON structures from existing structures is that you
can quickly end up with invalid JSON due to the specifications strict use of
commas.
```
» config -> [ shell ] -> formap k v { $v -> alter /Foo Bar }
{
    "Data-Type": "bool",
    "Default": true,
    "Description": "Display the interactive shell's hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions",
    "Dynamic": false,
    "Foo": "Bar",
    "Global": true,
    "Value": true
}
{
    "Data-Type": "block",
    "Default": "{ progress $PID }",
    "Description": "Murex function to execute when an `exec` process is stopped",
    "Dynamic": false,
    "Foo": "Bar",
    "Global": true,
    "Value": "{ progress $PID }"
}
{
    "Data-Type": "bool",
    "Default": true,
    "Description": "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc",
    "Dynamic": false,
    "Foo": "Bar",
    "Global": true,
    "Value": true
}
...
```

Luckily JSON also has it's own streaming format: jsonlines (`jsonl`)

```
» config -> [ shell ] -> formap k v { $v -> alter /Foo Bar } -> cast jsonl -> format json
[
    {
        "Data-Type": "bool",
        "Default": true,
        "Description": "Write shell history (interactive shell) to disk",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": true
    },
    {
        "Data-Type": "int",
        "Default": 4,
        "Description": "Maximum number of lines with auto-completion suggestions to display",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": "6"
    },
    {
        "Data-Type": "bool",
        "Default": true,
        "Description": "Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": true
    },
...
```
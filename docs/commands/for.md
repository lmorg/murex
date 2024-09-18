# For Loop (`for`)

> A more familiar iteration loop to existing developers

## Description

This `for` loop is fills a small niche where `foreach` or `formap` are
inappropriate in your script. It's generally not recommended to use `for`
because it performs slower and doesn't adhere to Murex's design
philosophy. However it does offer additional flexibility around recursion.

## Usage

```
for { variable; conditional; incrementor } { code-block } -> <stdout>
```

## Examples

```
» for {$i=1; $i<6; $i++} { out "iteration $i" }
iteration 1
iteration 2
iteration 3
iteration 4
iteration 5
```

## Detail

### Syntax

`for` is a little naughty in terms of breaking Murex's style guidelines due
to the first parameter being entered as one string treated as 3 separate code
blocks. The syntax is like this for two reasons:
  
1. readability, having multiple `{ blocks }` would make scripts unsightly
2. familiarity for those using to `for` loops in other languages

Take the following example:

```
for {$i=1; $i<6; $i++} { out "iteration $i" }
```

The first parameter is: `{$i=1; $i<6; $i++}`, this is then converted into the
following code:

1. `$i=1` - declare the loop iteration variable
2. `$i<6` - if the condition is true then proceed to run the code in
the second parameter - `{ echo $i }`
3. `$i++` - increment the loop iterator variable

The second parameter is the code to execute upon each iteration

### Tips when writing JSON inside for loops

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

For example in the code below, each item block is it's own object and there are
no `[ ... ]` encapsulating them to denote it is an array of objects, nor are
the objects terminated by a comma.

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

Luckily JSON also has it's own streaming format: JSON lines (`jsonl`). We can
`cast` this output as `jsonl` then `format` it back into valid JSON:

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

#### `foreach` will automatically cast it's output as `jsonl` _if_ it's stdin type is `json`

```
» ja [Tom,Dick,Sally] -> foreach name { out Hello $name }
Hello Tom
Hello Dick
Hello Sally

» ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> debug -> [[ /Data-Type/Murex ]]
jsonl

» ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> format json
[
    "Hello Tom",
    "Hello Dick",
    "Hello Sally"
]
```

## See Also

* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Loop While (`while`)](../commands/while.md):
  Loop until condition false
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)

<hr/>

This document was generated from [builtins/core/structs/for_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/for_doc.yaml).
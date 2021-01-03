# _murex_ Shell Docs

## Command Reference: `for`

> A more familiar iteration loop to existing developers

## Description

This `for` loop is fills a small niche where `foreach` or `formap` idioms will
fail within your scripts. It's generally not recommended to use `for` because
it performs slower and doesn't adhere to _murex_'s design philosiphy.

## Usage

    for ( variable; conditional; incrementation ) { code-block } -> <stdout>

## Examples

    » for ( i=1; i<6; i++ ) { echo $i }
    1
    2
    3
    4
    5

## Detail

### Syntax

`for` is a little naughty in terms of breaking _murex_'s style guidelines due
to the first parameter being entered as one string treated as 3 separate code
blocks. The syntax is like this for two reasons:
  
1. readability (having multiple `{ blocks }` would make scripts unsightly
2. familiarity (for those using to `for` loops in other languages

The first parameter is: `( i=1; i<6; i++ )`, but it is then converted into the
following code:

1. `let i=0` - declare the loop iteration variable
2. `= i<0` - if the condition is true then proceed to run the code in
the second parameter - `{ echo $i }`
3. `let i++` - increment the loop iteration variable

The second parameter is the code to execute upon each iteration

### Better `for` loops

Because each iteration of a `for` loop reruns the 2nd 2 parts in the first
parameter (the conditional and incrementation), `for` is very slow. Plus the
weird, non-idiomatic, way of writing the 3 parts, it's fair to say `for` is
not the recommended method of iteration and in fact there are better functions
to achieve the same thing...most of the time at least.

For example:

    a: [1..5] -> foreach: i { echo $i }
    1
    2
    3
    4
    5
    
The different in performance can be measured. eg:

    » time { a: [1..9999] -> foreach: i { out: <null> $i } }
    0.097643108
    
    » time { for ( i=1; i<10000; i=i+1 ) { out: <null> $i } }
    0.663812496
    
You can also do step ranges with `foreach`:

    » time { for ( i=10; i<10001; i=i+2 ) { out: <null> $i } }
    0.346254973
    
    » time { a: [1..999][0,2,4,6,8],10000 -> foreach i { out: <null> $i } }
    0.053924326
    
...though granted the latter is a little less readable.

### Tips when writing JSON inside for loops

One of the drawbacks (or maybe advantages, depending on your perspective) of
JSON is that parsers generally expect a complete file for processing in that
the JSON specification requires closing tags for every opening tag. This means
it's not always suitable for streaming. For example

    » ja [1..3] -> foreach i { out ({ "$i": $i }) }
    { "1": 1 }
    { "2": 2 }
    { "3": 3 }
    
**What does this even mean and how can you build a JSON file up sequentially?**

One answer if to write the output in a streaming file format and convert back
to JSON

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
    
**What if I'm returning an object rather than writing one?**

The problem with building JSON structures from existing structures is that you
can quickly end up with invalid JSON due to the specifications strict use of
commas.

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
    
Luckily JSON also has it's own streaming format: JSON lines (`jsonl`)

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
    
#### `foreach` will automatically cast it's output as `jsonl` _if_ it's STDIN type is `json`

    » ja: [Tom,Dick,Sally] -> foreach: name { out Hello $name }
    Hello Tom
    Hello Dick
    Hello Sally
    
    » ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> debug -> [[ /Data-Type/Murex ]]
    jsonl
    
    » ja: [Tom,Dick,Sally] -> foreach: name { out Hello $name } -> format: json
    [
        "Hello Tom",
        "Hello Dick",
        "Hello Sally"
    ]

## See Also

* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`foreach`](../commands/foreach.md):
  Iterate through an array
* [commands/`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [commands/`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value
* [commands/`while`](../commands/while.md):
  Loop until condition false
* [commands/formap](../commands/formap.md):
  
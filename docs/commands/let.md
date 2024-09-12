# `let`

> Evaluate a mathematical function and assign to variable (deprecated)

## Description

`let` evaluates a mathematical function and then assigns it to a locally
scoped variable (like `set`)

**This is a deprecated feature. Please refer to [`expr`](expr.md) instead.**

## Usage

```
let var_name=evaluation

let var_name++

let var_name--
```

## Examples

```
» let age=18
» $age
18

» let age++
» $age
19

» let under18=age<18
» $under18
false

» let under21 = age < 21
» $under21
true
```

## Detail

### Other Operators

`let` also supports the following operators (substitute **VAR** with your
variable name, and **NUM** with a number):

* `VAR--`, subtract 1 from VAR
* `VAR++`, add 1 to VAR
* `VAR -= NUM`, subtract NUM from VAR
* `VAR += NUM`, add NUM to VAR
* `VAR /= NUM`, divide VAR by NUM
* `VAR *= NUM`, multiply VAR by NUM

eg

```
» let i=0
» let i++
» $i
1

» let i+=8
» $i
9

» let i/=3
3
```

Please note these operators are not supported by `=`.

### Variables

There are two ways you can use variables with the math functions. Either by
string interpolation like you would normally with any other function, or
directly by name.

String interpolation:

```
» set abc=123
» = $abc==123
true
```

Directly by name:

```
» set abc=123
» = abc==123
false
```

To understand the difference between the two, you must first understand how
string interpolation works; which is where the parser tokenised the parameters
like so

```
command line: = $abc==123
token 1: command (name: "=")
token 2: parameter 1, string (content: "")
token 3: parameter 1, variable (name: "abc")
token 4: parameter 1, string (content: "==123")
```

Then when the command line gets executed, the parameters are compiled on demand
similarly to this crude pseudo-code

```
command: "="
parameters 1: concatenate("", GetValue(abc), "==123")
output: "=" "123==123"
```

Thus the actual command getting run is literally `123==123` due to the variable
being replace **before** the command executes.

Whereas when you call the variable by name it's up to `=` or `let` to do the
variable substitution.

```
command line: = abc==123
token 1: command (name: "=")
token 2: parameter 1, string (content: "abc==123")
```

```
command: "="
parameters 1: concatenate("abc==123")
output: "=" "abc==123"
```

The main advantage (or disadvantage, depending on your perspective) of using
variables this way is that their data-type is preserved.

```
» set str abc=123
» = abc==123
false

» set int abc=123
» = abc==123
true
```

Unfortunately is one of the biggest areas in Murex where you'd need to be
careful. The simple addition or omission of the dollar prefix, `$`, can change
the behavior of `=` and `let`.

### Strings

Because the usual Murex tools for encapsulating a string (`"`, `'` and `()`)
are interpreted by the shell language parser, it means we need a new token for
handling strings inside `=` and `let`. This is where backtick comes to our
rescue.

```
» set str abc=123
» = abc==`123`
true
```

Please be mindful that if you use string interpolation then you will need to
instruct `=` and `let` that your field is a string

```
» set str abc=123
» = `$abc`==`123`
true
```

### Best practice recommendation

As you can see from the sections above, string interpolation offers us some
conveniences when comparing variables of differing data-types, such as a `str`
type with a number (eg `num` or `int`). However it makes for less readable code
when just comparing strings. Thus the recommendation is to avoid using string
interpolation except only where it really makes sense (ie use it sparingly).

### Non-boolean logic

Thus far the examples given have been focused on comparisons however `=` and
`let` supports all the usual arithmetic operators:

```
» = 10+10
20

» = 10/10
1

» = (4 * (3 + 2))
20

» = `foo`+`bar`
foobar
```

### Read more

Murex uses the [govaluate package](https://github.com/Knetic/govaluate). More information can be found in it's manual.

### Type Annotations

When `set` or `global` are used as a function, the parameters are passed as a
string which means the variables are defined as a `str`. If you wish to define
them as an alternate data type then you should add type annotations:

```
» set int age = 30
```

(`$age` is an integer, `int`)

```
» global bool dark_theme = true
```

(`$dark_theme` is a boolean, `bool`)

When using `set` or `global` as a method, by default they will define the
variable as the data type of the pipe:

```
» open example.json -> set: file
```

(`$file` is defined a `json` type because `open` wrote to `set`'s pipe with a
`json` type)

You can also annotate `set` and `global` when used as a method too:

```
out 30 -> set: int age
```

(`$age` is an integer, `int`, despite `out` writing a string, `str, to the pipe)

> `export` does not support type annotations because environmental variables
> must always be strings. This is a limitation of the current operating systems.

### Scoping

Variable scoping is simplified to three layers:

1. Local variables (`set`, `!set`, `let`)
2. Global variables (`global`, `!global`)
3. Environmental variables (`export`, `!export`, `unset`)

Variables are looked up in that order of too. For example a the following
code where `set` overrides both the global and environmental variable:

```
» set    foobar=1
» global foobar=2
» export foobar=3
» out $foobar
1
```

#### Local variables

These are defined via `set` and `let`. They're variables that are persistent
across any blocks within a function. Functions will typically be blocks
encapsulated like so:

```
function example {
    # variables scoped inside here
}
```

...or...

```
private example {
    # variables scoped inside here
}

```

...however dynamic autocompletes, events, unit tests and any blocks defined in
`config` will also be triggered as functions.

Code running inside any control flow or error handing structures will be
treated as part of the same part of the same scope as the parent function:

```
» function example {
»     try {
»         # set 'foobar' inside a `try` block
»         set foobar=example
»     }
»     # 'foobar' exists outside of `try` because it is scoped to `function`
»     out $foobar
» }
example
```

Where this behavior might catch you out is with iteration blocks which create
variables, eg `for`, `foreach` and `formap`. Any variables created inside them
are still shared with any code outside of those structures but still inside the
function block.

Any local variables are only available to that function. If a variable is
defined in a parent function that goes on to call child functions, then those
local variables are not inherited but the child functions:

```
» function parent {
»     # set a local variable
»     set foobar=example
»     child
» }
»
» function child {
»     # returns the `global` value, "not set", because the local `set` isn't inherited
»     out $foobar
» }
»
» global $foobar="not set"
» parent
not set
```

It's also worth remembering that any variable defined using `set` in the shells
FID (ie in the interactive shell) is localised to structures running in the
interactive, REPL, shell and are not inherited by any called functions.

#### Global variables

Where `global` differs from `set` is that the variables defined with `global`
will be scoped at the global shell level (please note this is not the same as
environmental variables!) so will cascade down through all scoped code-blocks
including those running in other threads.

#### Environmental variables

Exported variables (defined via `export`) are system environmental variables.
Inside Murex environmental variables behave much like `global` variables
however their real purpose is passing data to external processes. For example
`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):

```
» export foo=bar
» env -> grep foo
foo=bar
```

### Function Names

As a security feature function names cannot include variables. This is done to
reduce the risk of code executing by mistake due to executables being hidden
behind variable names.

Instead Murex will assume you want the output of the variable printed:

```
» out "Hello, world!" -> set hw
» $hw
Hello, world!
```

On the rare occasions you want to force variables to be expanded inside a
function name, then call that function via `exec`:

```
» set cmd=grep
» ls -> exec $cmd main.go
main.go
```

This only works for external executables. There is currently no way to call
aliases, functions nor builtins from a variable and even the above `exec` trick
is considered bad form because it reduces the readability of your shell scripts.

### Usage Inside Quotation Marks

Like with Bash, Perl and PHP: Murex will expand the variable when it is used
inside a double quotes but will escape the variable name when used inside single
quotes:

```
» out "$foo"
bar

» out '$foo'
$foo

» out %($foo)
bar
```

## See Also

* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Define Global (`global`)](../commands/global.md):
  Define a global variable and set it's value
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`=` (arithmetic evaluation)](../parser/equ.md):
  Evaluate a mathematical function (deprecated)

<hr/>

This document was generated from [builtins/core/typemgmt/math_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/math_doc.yaml).
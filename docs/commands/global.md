# Define Global (`global`)

> Define a global variable and set it's value

## Description

Defines, updates or deallocates a global variable.

## Usage

```
# Assume data type and value from stdin
<stdin> -> global var_name

# Assume value from stdin, define the data type manually
<stdin> -> global datatype var_name

# Define value manually (data type defaults to string; `str`)
global var_name=data

# Define value and data type manually
global datatype var_name=data

# Define a variable but don't set any value
global var_name
global datatype var_name
```

## Examples

As a method:

```
» out "Hello, world!" -> global hw
» out "$hw"
Hello, World!
```

As a function:

```
» global hw="Hello, world!"
» out "$hw"
Hello, World!
```

## Detail

### Deallocation

You can unset variable names with the bang prefix:

```
!global var_name
```

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

## Synonyms

* `global`
* `!global`


## See Also

* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Is Value Null (`is-null`)](../commands/is-null.md):
  Checks if a variable is null or undefined
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`??` Null Coalescing Operator](../parser/null-coalescing.md):
  Returns the right operand if the left operand is empty / undefined (expression)

<hr/>

This document was generated from [builtins/core/typemgmt/variables_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/variables_doc.yaml).
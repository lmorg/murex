# Public Function (`function`)

> Define a function block

## Description

`function` defines a block of code as a function

## Usage

Define a function:

```
function name { code-block }
```

Define a function with variable names defined (**default value** and
**description** are optional parameters):

```
function name (
    variable1: data-type [default-value] "description",
    variable2: data-type [default-value] "description"
) {
    code-block
}
```

Undefine a function:

```
!function command
```

## Examples

```
» function hw { out "Hello, World!" }
» hw
Hello, World!

» !function hw
» hw
exec "hw": executable file not found in $PATH
```

## Detail

### Allowed characters

Function names can only include any characters apart from dollar (`$`).
This is to prevent functions from overwriting variables (see the order of
preference below).

### Undefining a function

Like all other definable states in Murex, you can delete a function with
the bang prefix `!function` (see the example above).

### Using parameterized variable names

By default, if you wanted to query the parameters passed to a Murex function
you would have to use either:

* the Bash syntax where of `$2` style numbered reserved variables,

* and/or the Murex convention of `$PARAM` / `$ARGS` arrays (see **reserved-vars**
  document below),
  
* and/or the older Murex convention of the `args` builtin for any flags.

Starting from Murex `2.7.x` it's been possible to declare parameters from
within the function declaration:

```
function name (
    variable1: data-type [default-value] "description",
    variable2: data-type [default-value] "description"
) {
    code-block
}
```

#### Syntax

First off, the syntax doesn't have to follow exactly as above:

* **Variables** shouldn't be prefixed with a dollar (`$`). This is a little like
  declaring variables via `set`, etc. However it should be followed by a colon
  (`:`) or comma (`,`). Normal rules apply with regards to allowed characters
  in variable names: limited to ASCII letters (upper and lower case), numbers,
  underscore (`_`), and hyphen (`-`). Unicode characters as variable names are
  not currently supported.

* **data-type** is the Murex data type. This is an optional field in version
  `2.8.x` (defaults to `str`) but is required in `2.7.x`.

* The **default value** must be inside square brackets (`[...]`). Any value is
  allowed (including Unicode) _except_ for carriage returns / new lines (`\r`,
  `\n`) and a closing square bracket (`]`) as the latter would indicate the end
  of this field. You cannot escape these characters either.

  This field is optional.

* The **description** must sit inside double quotes (`"..."`). Any value is allowed
  (including Unicode) _except_ for carriage returns / new lines (`\r`, `\n`)
  and double quotes (`"`) as the latter would indicate the end of this field.
  You cannot escape these characters either.

  This field is optional.

* You do not need a new line between each parameter, however you do need to
  separate them with a comma (like with JSON, there should not be a trailing
  comma at the end of the parameters). Thus the following is valid:
  `variable1: data-type, variable2: data-type`.

#### Variables

Any variable name you declare in your function declaration will be exposed in
your function body as a local variable. For example:

```
function hello (name: str) {
    out "Hello $name, pleased to meet you."
}
```

If the function isn't called with the complete list of parameters and it is
running in the foreground (ie not part of `autocomplete`, `event`, `bg`, etc)
then you will be prompted for it's value. That could look something like this:

```
» function hello (name: str) {
»     out "Hello $name, pleased to meet you."
» }

» hello
Please enter a value for 'name': Bob
Hello Bob, pleased to meet you.
```

(in this example you typed `Bob` when prompted)

#### Data-Types

This is the Murex data type of the variable. From version `2.8.x` this field
is optional and will default to `str` when omitted.

The advantage of setting this field is that values are type checked and the
function will fail early if an incorrect value is presented. For example:

```
» function age (age: int) { out "$age is a great age." }

» age
Please enter a value for 'age': ten
Error in `age` ( 2,1): cannot convert parameter 1 'ten' to data type 'int'

» age ten
Error in `age` ( 2,1): cannot convert parameter 1 'ten' to data type 'int'
```

However it will try to automatically convert values if it can:

```
» age 1.2
1 is a great age.
```

#### Default values

Default values are only relevant when functions are run interactively. It
allows the user to press enter without inputting a value:

```
» function hello (name: str [John]) { out "Hello $name, pleased to meet you." }

» hello
Please enter a value for 'name' [John]: 
Hello John, pleased to meet you.
```

Here no value was entered so `$name` defaulted to `John`.

Default values will not auto-populate when the function is run in the
background. For example:

```
» bg {hello}
Error in `hello` ( 2,2): cannot prompt for parameters when a function is run in the background: too few parameters
```

#### Description

Descriptions are only relevant when functions are run interactively. It allows
you to define a more useful prompt should that function be called without
sufficient parameters. For example:

```
» function hello (name: str "What is your name?") { out "Hello $name" }

» hello
What is your name?: Sally
Hello Sally
```

### Order of precedence

There is an order of precedence for which commands are looked up:

1. `runmode`: this is executed before the rest of the script. It is invoked by
   the pre-compiler forking process and is required to sit at the top of any
   scripts.

1. `test` and `pipe` functions also alter the behavior of the compiler and thus
   are executed ahead of any scripts.

4. private functions - defined via `private`. Private's cannot be global and
   are scoped only to the module or source that defined them. For example, You
   cannot call a private function directly from the interactive command line
   (however you can force an indirect call via `fexec`).

2. Aliases - defined via `alias`. All aliases are global.

3. Murex functions - defined via `function`. All functions are global.

5. Variables (dollar prefixed) which are declared via `global`, `set` or `let`.
   Also environmental variables too, declared via `export`.

6. globbing: however this only applies for commands executed in the interactive
   shell.

7. Murex builtins.

8. External executable files

You can override this order of precedence via the `fexec` and `exec` builtins.

## Synonyms

* `function`
* `!function`


## See Also

* [Alias Pointer (`alias`)](../commands/alias.md):
  Create an alias for a command
* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Define Function Arguments (`args`)](../commands/args.md):
  Command line flag parser for Murex shell scripting
* [Define Global (`global`)](../commands/global.md):
  Define a global variable and set it's value
* [Define Method Relationships (`method`)](../commands/method.md):
  Define a methods supported data-types
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [Globbing (`g`)](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Murex Version (`version`)](../commands/version.md):
  Get Murex version
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)

<hr/>

This document was generated from [builtins/core/structs/function_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/function_doc.yaml).
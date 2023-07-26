# Reserved Variables - User Guide

> Special variables reserved by Murex

Murex reserves a few special variables which cannot be assigned via `set` nor
`let`.

The following is a list of reserved variables, their data type, and its usage:

## `SELF` (json)

This returns meta information about the running scope.

A 'scope' in Murex is a collection of code blocks to which variables and
config are persistent within. In Murex, a variable declared inside an `if` or
`foreach` block will be persistent outside of their blocks as long as you're
still inside the same function.

Please see scoping document (link below) for more information on scoping.

```
» function example { out $SELF }
» example
{
    "Parent": 11357,
    "Scope": 11357,
    "TTY": true,
    "Method": false,
    "Not": false,
    "Background": false,
    "Module": "murex"
}
```

#### Parent (num)

This is the function ID of the parent function that created the scope. In
some instances this will be the same value as scope FID. However if in doubt
then please using **Scope** instead.

#### Scope (num)

The scope value here returns the function ID of the top level function in the
scope.

#### TTY (bool)

A boolean value as to whether STDOUT is a TTY (ie are we printing to the
terminal (TTY) or a pipe?)

#### Method (bool)

A boolean value to describe whether the current scope is a method (ie being
called mid-way or at the end of a pipeline). 

#### Not (bool)

A boolean value which represents whether the function was called with a bang-
prefix or not.

#### Background (bool)

A boolean value to identify whether the current scope is running in the
background for foreground.

#### Module (str)

This will be the module string for the current scope.

### `ARGS` (json)

This returns a JSON array of the command name and parameters within a given
scope.

Unlike `$PARAMS`, `$ARGS` includes the function name.

```
» function example { out $ARGS }
» example abc 1 2 3
[
    "example",
    "abc",
    "1",
    "2",
    "3"
]
```

### `PARAMS` (json)

This returns a JSON array of the parameters within a given scope.

Unlike `$ARGS`, `$PARAMS` does not include the function name.

```
» function example { out $PARAMS }
» example abc 1 2 3
[
    "abc",
    "1",
    "2",
    "3"
]
```

### `MUREX_EXE` (str)

This is very similar to the `$SHELL` environmental variable in that it holds
the full path to the running shell. The reason for defining a reserved variable
is so that the shell path cannot be overridden.

### `MUREX_ARGS` (json)

This is TODO: [https://github.com/lmorg/murex/issues/304](Github issue 304)

### `HOSTNAME` (str)

This returns the hostname of the machine Murex is running from.

### `0` (str)

This returns the name of the executable (like `$ARGS[0]`)

### `1`, `2`, `3`... (str)

This returns parameter _n_ (like `$ARGS[n]`). If there is no parameter _n_
then the variable will not be set thus the upper limit variable is determined
by how many parameters are set. For example if you have 19 parameters passed
then variables `$1` through to `$19` (inclusive) will all be set.

## See Also

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [Bang Prefix](../user-guide/bang-prefix.md):
  Bang prefixing to reverse default actions
* [Modules and Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [String (`$`) Token](../parser/string.md):
  Expand values as a string
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`foreach`](../commands/foreach.md):
  Iterate through an array
* [`function`](../commands/function.md):
  Define a function block
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`private`](../commands/private.md):
  Define a private function block
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`switch`](../commands/switch.md):
  Blocks of cascading conditionals
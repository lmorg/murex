# `SELF` (json)

> Meta information about the running scope.

## Description

`SELF` returns information about the functional scope that the code is running
inside. Such as whether that functions STDOUT is a TTY, running in the
background or a method.

A 'scope' in Murex is a collection of code blocks to which variables and
config are persistent within. In Murex, a variable declared inside an `if` or
`foreach` block will be persistent outside of their blocks as long as you're
still inside the same function.

Please see scoping document (link below) for more information on scoping.

This is a reserved variable so it cannot be changed.



## Examples

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

## See Also

* [Modules and Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [String (`$`) Token](../parser/string.md):
  Expand values as a string
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`expr`](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [`foreach`](../commands/foreach.md):
  Iterate through an array
* [`function`](../commands/function.md):
  Define a function block
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`json` ](../types/json.md):
  JavaScript Object Notation (JSON)
* [`private`](../commands/private.md):
  Define a private function block
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`switch`](../commands/switch.md):
  Blocks of cascading conditionals
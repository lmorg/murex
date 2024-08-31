# Variable and Config Scoping

> How scoping works within Murex

## Description

A 'scope' in Murex is a collection of code blocks to which variables and
config are persistent within. In Murex, a variable declared inside an `if` or
`foreach` block will be persistent outside of their blocks as long as you're
still inside the same function.

For example lets start with the following function that sets a variable called
**foo**

```
function example {
    if { true } then { set foo=bar }
    out $foo
}
```

In here the value is getting set inside an `if` block but its value is is
retrieved outside of that block. `out` and `set` have different parents but
the same scoping.

Then lets set **foo** outside of that function and see what happens:

```
» set foo=oof
» $foo
oof

» example
bar

» $foo
oof
```

Despite setting a variable named **foo**, the value inside **example** does not
overwrite the value outside **example** because they occupy different scoping.

## What Instantiates A New Scope?

A new scope is instantiated by anything which resembles a function. This would
be code inside events, dynamic autocompletes, open agents, any code blocks
defined in `config`, as well as public and private functions too.

Code inside an `if`, `switch`, `foreach` and `source` do not create a new
scope. Subshells also do not create a new scope either.

## See Also

* [Define Handlers For "`open`" (`openagent`)](../commands/openagent.md):
  Creates a handler function for `open`
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Switch Conditional (`switch`)](../commands/switch.md):
  Blocks of cascading conditionals
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)

<hr/>

This document was generated from [gen/user-guide/scoping_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/scoping_doc.yaml).
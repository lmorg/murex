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
scope. Subshells also do not create a new scoping either.

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [autocomplete](../user-guide/autocomplete.md):
  
* [config](../user-guide/config.md):
  
* [event](../user-guide/event.md):
  
* [foreach](../user-guide/foreach.md):
  
* [function](../user-guide/function.md):
  
* [if](../user-guide/if.md):
  
* [let](../user-guide/let.md):
  
* [openagent](../user-guide/openagent.md):
  
* [private](../user-guide/private.md):
  
* [set](../user-guide/set.md):
  
* [source](../user-guide/source.md):
  
* [switch](../user-guide/switch.md):
  

<hr/>

This document was generated from [gen/user-guide/scoping_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/scoping_doc.yaml).
# _murex_ Shell Docs

## Command Reference: `while`

> Loop until condition false

## Description

`while` loops until loops until **condition** is false.

Normally the **conditional** and executed code block are 2 separate parameters
however you can call `while` with just 1 parameter where the code block acts
as both the conditional and the code to be ran.

## Usage

Until true

    while { condition } { code-block } -> <stdout>
    
    while { code-block } -> <stdout>
    
Until false

    !while { condition } { code-block } -> <stdout>
    
``
!while { code-block } -> <std

## Examples

`while` **$i** is less then **5**

    » let i=0; while { =i<5 } { let i=i+1; out $i }
    1
    2
    3
    4
    5
    
    » let i=0; while { let i=i+1; = i<5; out }
    true
    true
    true
    true
    false
    
`while` **$i** is _NOT_ greater than or equal to **5**

    » let i=0; !while { =i>=5 } { let i=i+1; out $i }
    1
    2
    3
    4
    5
    
    » let i=0; while { let i=i+1; = i>=5; out }
    true
    true
    true
    true
    false

## Synonyms

* `while`
* `!while`


## See Also

* [commands/`err`](../commands/err.md):
  Print a line to the STDERR
* [commands/`for`](../commands/for.md):
  A more familiar iteration loop to existing developers
* [commands/`foreach`](../commands/foreach.md):
  Iterate through an array
* [commands/`global`](../commands/global.md):
  Define a global variable and set it's value
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value
* [commands/formap](../commands/formap.md):
  
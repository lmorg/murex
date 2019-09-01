# _murex_ Shell Guide

## Command Reference: `exec`

> Runs an executable

### Description

With _murex_, like most other shells, you launch a process by calling the
name of that executable directly. While this is suitable 99% of the time,
occasionally you might run into an edge case where that wouldn't work. The
primary reason being if you needed to launch a process from a variable, eg

    » set exe=uname
    » $exe
    uname
    
As you can see here, _murex_'s behavior here is to output the contents of
the variable rather then executing the contents of the variable. This is
done for safety reasons, however if you wanted to override that behavior
then you could prefix the variable with exec:

    » set exe=uname
    » exec $exe
    Linux

### Usage

    <stdin> -> exec
    <stdin> -> exec -> <stdout>
               exec -> <stdout>

### Examples

    » exec printf "Hello, world!"
    Hello, world!

### See Also

* commands/[`=` (artithmetic evaluation)](../commands/equ.md):
  Evaluate a mathmatical function
* commands/[`let`](../commands/let.md):
  Evaluate a mathmatical function and assign to variable
* commands/[`set`](../commands/set.md):
  Define a local variable and set it's value
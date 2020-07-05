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
* [commands/formap](../commands/formap.md):
  
* [commands/while](../commands/while.md):
  
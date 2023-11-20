# Parser Reference

This section is a glossary of Murex tokens and parser behavior.

## Other Reference Material

### Language Guides

1. [Language Tour](/docs/tour.md), which is an introduction into
    the Murex language.

2. [Rosetta Stone](/docs/user-guide/rosetta-stone.md), which is a reference
    table comparing Bash syntax to Murex's.

3. [Builtins](/docs/commands/), for docs on the core builtins.

### Murex's Source Code

The parser is located Murex's source under the `lang/` path of the project
files.

## Pages

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [Tilde (`~`) Token](../parser/tilde.md):
  Home directory path variable
* [`!` (not)](../parser/not-func.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [`"Double Quote"`](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`$Variable`](../parser/scalar.md):
  Expand values as a scalar
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%[]` Create Array](../parser/create-array.md):
  Quickly generate arrays
* [`%{}` Create Map](../parser/create-object.md):
  Quickly generate objects and maps
* [`&&` And Logical Operator](../parser/logical-and.md):
  Continues next operation if previous operation passes
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the STDOUT without new line (deprecated)
* [`*=` Multiply By Operator](../parser/multiply-by.md):
  Multiplies a variable by the right hand value (expression)
* [`*` Multiplication Operator](../parser/multiplication.md):
  Multiplies one numeric value with another (expression)
* [`+=` Add With Operator](../parser/add-with.md):
  Adds the right hand value to a variable (expression)
* [`+` Addition Operator](../parser/addition.md):
  Adds two numeric values together (expression)
* [`-=` Subtract By Operator](../parser/subtract-by.md):
  Subtracts a variable by the right hand value (expression)
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`-` Subtraction Operator](../parser/subtraction.md):
  Subtracts one numeric value from another (expression)
* [`/=` Divide By Operator](../parser/divide-by.md):
  Divides a variable by the right hand value (expression)
* [`/` Division Operator](../parser/division.md):
  Divides one numeric value from another (expression)
* [`<read-named-pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [`=>` Generic Pipe](../parser/pipe-generic.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [`=` (arithmetic evaluation)](../parser/equ.md):
  Evaluate a mathematical function (deprecated)
* [`>>` Append File](../parser/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>>` Append Pipe](../parser/pipe-append.md):
  Redirects STDOUT to a file and append its contents
* [`?:` Elvis Operator](../parser/elvis.md):
  Returns the right operand if the left operand is falsy (expression)
* [`??` Null Coalescing Operator](../parser/null-coalescing.md):
  Returns the right operand if the left operand is empty / undefined (expression)
* [`?` STDERR Pipe](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command (DEPRECATED)
* [`[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from STDIN
* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`[[ Element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [`[{ Lambda }]`](../parser/lambda.md):
  Iterate through structured data
* [`{ Curly Brace }`](../parser/curly-brace.md):
  Initiates or terminates a code block
* [`|>` Truncate File](../parser/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`||` Or Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails
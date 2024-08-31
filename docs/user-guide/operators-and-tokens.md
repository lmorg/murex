# Operators And Tokens

> A table of all supported operators and tokens

<h2>Table of Contents</h2>

<div id="toc">

- [Syntax](#syntax)
  - [Expressions](#expressions)
  - [Statements](#statements)
- [Order Of Operations](#order-of-operations)
  - [Expression Or Statement Discovery](#expression-or-statement-discovery)
- [Operators And Tokens](#operators-and-tokens)
  - [Terminology](#terminology)
  - [Modifiers](#modifiers)
  - [Immutable Merge](#immutable-merge)
  - [Comparators](#comparators)
  - [Assignment](#assignment)
  - [Conditionals](#conditionals)
  - [Sigils](#sigils)
  - [Constants](#constants)
  - [Sub-shells](#sub-shells)
  - [Boolean Operations](#boolean-operations)
  - [Pipes](#pipes)
  - [Terminators](#terminators)
  - [Escape Codes](#escape-codes)

</div>


## Syntax

Murex supports both expressions and statements. You can use the interchangeably
with your code and the Murex parser will decide whether to run that code as an
expression or statement.

### Expressions

Expressions are patterns formed like equations (eg `$val = 1 + 2`).

Strings must be quoted in expressions.

### Statements

Statements are traditional shell command calls (eg `command parameters...`).

Quoting strings is optional in statements.

Not all operators are supported in statements.

## Order Of Operations

Expressions and statements are split by pipes and terminators. Each statement
and expression is executed from left to right, with the statement or expression
parsed by the following rules of operation

Order of operations:
1. expression or statement discovery
2. sub-shells / sub-expressions
3. multiplication / division (expressions only)
4. addition / subtraction (expressions only)
5. immutable merge
6. comparisons, eg greater than (expressions only)
7. logical and (sub-expressions only)
8. logical or (sub-expressions only)
9. elvis (expressions only)
10. assign (expressions only)
11. _left_ to _right_

### Expression Or Statement Discovery

First a command is read as an expression. Because the rules of parsing
expressions are stricter than statements, everything is assumed to be an
expression unless the expression parser fails, which then it is assumed to be a
statement.

## Operators And Tokens

### Terminology

* _left_: this is the value to the left hand side of the operator
* _right_: this is the value to the right hand side of the operator

Example: _left_ `operator` _right_

### Modifiers

All modifiers replace the _left_, operator and _right_ with the returned value
of the modifier.

All returns will be `num` data type (or their original type if strict types is
enabled).

Modifiers are only supported in expressions.

| Operator | Name           | Operation                  |
|----------|----------------|----------------------------|
| `*`      | Multiplication | Multiply _left_ by _right_ |
| `/`      | Divide         | Divide _left_ by _right_   |
| `+`      | Addition       | Add _left_ with _right_    |
| `-`      | Subtraction    | Subtract _left_ by _right_ |

Read more:
* Data types: [num](/docs/types/num.md), [int](/docs/types/int.md), [float](/docs/types/float.md)
* Strict types config: [strict types](/docs/user-guide/strict-types.md)
* Operators: [+](/docs/parser/addition.md), [-](/docs/parser/subtraction.md), [*](/docs/parser/multiplication.md), [/](/docs/parser/division.md)

### Immutable Merge

Returns the result of merging _right_ into _left_.

_immutable merge_ does not modify the contents of either _left_ nor _right_.

The direction of the arrow indicates that the result returned is a new value
rather than an updated assignment.

_Left_ can be a statement or expression, whereas _right_ can only be an
expression. However you can still use a sub-shell as part of, or the entirety,
of, that expression.

| Operator | Name            | Operation                                   |
|----------|-----------------|---------------------------------------------|
| `~>`     | Immutable Merge | Returns merged value of _right_ into _left_ |

### Comparators

All comparators replace the _left_, operator and _right_ with the returned
value of the comparator.

All returns will be `bool` data type, either `true` or `false`.

Comparators are only supported in expressions.

| Operator | Name                  | Operation                                           |
|----------|-----------------------|-----------------------------------------------------|
| `>`      | Greater Than          | `true` if _left_ is greater than _right_            |
| `>=`     | Greater Or Equal To   | `true` if _left_ is greater or equal to _right_     |
| `<`      | Less Than             | `true` if _left_ is less than _right_               |
| `<=`     | Less Or Equal To      | `true` if _left_ is less or equal to _right_        |
| `==`     | Equal To              | `true` if _left_ is equal to _right_                |
| `!=`     | Not Equal To          | `false` if _left_ is equal to _right_               |
| `~~`     | Like                  | `true` if _left_ string is like _right_ string      |
| `!!`     | Not Like              | `false` if _left_ string is like _right_ string     |
| `=~`     | Matches Regexp        | `true` if _left_ matches regexp pattern on _right_  |
| `!~`     | Does Not Match Regexp | `false` if _left_ matches regexp pattern on _right_ |

Read more:
* Data types: [bool](/docs/types/bool.md)

### Assignment

Assignment returns `null` if successful.

Assignment is only supported in expressions.

| Operator | Name                  | Operation                                         |
|----------|-----------------------|---------------------------------------------------|
| `=`      | Assign (overwrite)    | Assign _right_ to _left_                          |
| `:=`     | Assign (retain)       | **EXPERIMENTAL**                                  |
| `<~`     | Assign Or Merge       | Merge _right_ (array / object) into _left_        |
| `+=`     | Assign And Add        | Add _right_ to _left_ and assign to _left_        |
| `-=`     | Assign And Subtract   | Subtract _right_ from _left_ and assign to _left_ |
| `*=`     | Assign And Multiply   | Multiply _right_ with _left_ and assign to _left_ |
| `/=`     | Assign And Divide     | Divide _right_ with _left_ and assign to _left_   |
| `++`     | Add one to variable   | Adds one to _right_ and reassigns                 |
| `--`     | Subtract one from var | Subtracts one from _right_ and reassigns          |

Read more:
* Data types: [bool](/docs/types/bool.md)
* Operators: =, [<~](/docs/parser/assign-or-merge.md), [+=](/docs/parser/add-with.md),  [-=](/docs/parser/subtract-by.md), [*=](/docs/parser/multiply-by.md), [/=](/docs/parser/divide-by.md)

### Conditionals

Conditionals replace _left_, operator and _right_ with the value defined in
_operation_.

These conditionals are only supported in expressions.

| Operator | Name               | Operation                                       |
|----------|--------------------|-------------------------------------------------|
| `??`     | Null Coalescence   | Returns _left_ if not `null`, otherwise _right_ |
| `?:`     | Elvis              | Returns _left_ if truthy, otherwise _right_     |

Read more:
* Operators: [??](/docs/parser/null-coalescing.md), [?:](/docs/parser/elvis.md)

### Sigils

Sigils are special prefixes that provide hints to the parser.

Sigils are supported in both expressions and statements.

| Token    | Name           | Operation                                  |
|----------|----------------|--------------------------------------------|
| `$`      | Scalar         | Expand value as a string                   |
| `@`      | Array          | Expand value as an array                   |
| `~`      | Home           | Expand value as the persons home directory |
| `%`      | Builder        | Create an array, map or nestable string    |

### Constants

Constants are supported in both expressions and statements. However `null`,
`true`, `false` and _number_ will all be interpreted as strings in statements.

| Token         | Name           | Operation                                          |
|---------------|----------------|----------------------------------------------------|
| `null`        | Null           | `null` (null / nil / void) type                    |
| `true`        | True           | `bool` (boolean) true                              |
| `false`       | False          | `bool` (boolean) false                             |
| number        | Number         | `num` (numeric) value                              |
| `'`string`'`  | String Literal | `str` (string) literal value                       |
| `"`string`"`  | Infix String   | `str` (string) value, supports escaping & infixing |
| `%(`string`)` | String Builder | Creates a nestable `str` (string)                  |
| `%[`array`]`  | Array Builder  | Creates a `json` (JSON) array (list)               |
| `%{`map`}`    | Object Builder | Creates a `json` (JSON) object (map / dictionary)  |

Read more:
* Operators: ['string'](/docs/parser/single-quote.md), ["string"](/docs/parser/double-quote.md), [%(string)](/docs/parser/brace-quote.md), [%[array]](/docs/parser/create-array.md), [%{map}](/docs/parser/create-object.md)

### Sub-shells

Sub-shells are a way of inlining expressions or statements into an existing
expression or statement. Because of this they are supported in both.

| Syntax                       | Name               | Operation                          |
|------------------------------|--------------------|------------------------------------|
| command`(` parameters... `)` | C-Style Functions  | Inline a command as a function     |
| `${`command parameters...`}` | Sub-shell (scalar) | Inline a command line as a string  |
| `@{`command parameters...`}` | Sub-shell (array)  | expand a command line as an array  |
| `(`expression`)`             | Sub-expression     | Inline an expression (_statement_) |
| `(`expression`)`             | Sub-expression     | Order of evaluation (_expression_) |

Read more:
* [C-style functions](/docs/parser/c-style-fun.md), [sub-shells](/docs/tour.md#sub-shells), [sub-expressions](/docs/parser/expr-inlined.md)

### Boolean Operations

Boolean operators behave like pipes.

They are supported in both expressions and statements.

| Operator | Name           | Operation                                 |
|----------|----------------|-------------------------------------------|
| `&&`     | And            | Evaluates _right_ if _left_ is truthy     |
| `\|\|`   | Or             | Evaluates _right_ if _left_ is falsy      |

### Pipes

Pipes always flow from _left_ to _right_.

They are supported in both expressions and statements.

| Operator | Name           | Operation                                  |
|----------|----------------|--------------------------------------------|
| `\|`     | POSIX Pipe     | POSIX compatibility                        |
| `->`     | Arrow Pipe     | Context aware pipe                         |
| `=>`     | Generic Pipe   | Convert stdout to `*` (generic) then pipe  |
| `\|>`    | Truncate File  | Write stdout to file, overwriting contents |
| `>>`     | Append File    | Write stdout to file, appending contents   |

### Terminators

"LF" refers to the life feed character, which is a new line.

| Token    | Name              | Operation                                 |
|----------|-------------------|-------------------------------------------|
| `;`      | Semi-Colon        | End of statement or expression (optional) |
| LF       | Line Feed         | End of statement or expression (new line) |

### Escape Codes

Any character can be escaped via `\` to signal it isn't a token. However some
characters have special meanings when escaped.

"LF" refers to the life feed character, which is a new line.

| Token    | Name              | Operation                                  |
|----------|-------------------|--------------------------------------------|
| `\s`     | Space             | Same as a space character                  |
| `\t`     | Tab               | Same as a tab character                    |
| `\r`     | Carriage Return   | Carriage Return (CR) sometimes precedes LF |
| `\n`     | Line Feed         | Line Feed (LF), typically a new line       |
| `\`LF    | Escaped Line Feed | Statement continues on next line           |

## See Also

* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Language Tour](../Murex/tour.md):
  Getting started with Murex

<hr/>

This document was generated from [gen/user-guide/operators-tokens_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/operators-tokens_doc.yaml).
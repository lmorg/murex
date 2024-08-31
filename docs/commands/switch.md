# Switch Conditional (`switch`)

> Blocks of cascading conditionals

## Description

`switch` is a large block for simplifying cascades of conditional statements.

## Usage

```
switch [value] {
  case | if { conditional } [then] { code-block }
  case | if { conditional } [then] { code-block }
  ...
  [ default { code-block } ]
} -> <stdout>
```

The first parameter should be either **case** or **if** -- the statements are
subtly different and thus alter the behavior of `switch`.

**then** is optional ('then' is assumed even if not explicitly present).

## Examples

Output an array of editors installed:

```
switch {
    if { which vi    } { out vi    }
    if { which vim   } { out vim   }
    if { which nano  } { out nano  }
    if { which emacs } { out emacs }
} -> format: json
```

A higher/lower game written using `switch`:

```
function higherlower {
  try {
    rand int 100 -> set rand
    while { $rand } {
      read guess "Guess a number between 1 and 100: "

      switch {
        case: { = $guess < $rand } then {
          out "Too low"
        }

        case: { = $guess > $rand } then {
          out "Too high"
        }

        default: {
          out "Correct"
          let rand=0
        }
      }
    }
  }
}
```

String matching with `switch`:

```
read name "What is your name? "
switch $name {
    case "Tom"   { out "I have a brother called Tom" }
    case "Dick"  { out "I have an uncle called Dick" }
    case "Sally" { out "I have a sister called Sally" }
    default      { err "That is an odd name" }
}
```

## Detail

### Comparing Values vs Boolean State

#### By Values

If you supply a value with `switch`...

```
switch value { ... }
```

...then all the conditionals are compared against that value. For example:

```
switch foo {
    case bar {
        # not executed because foo != bar
    }
    case foo {
        # executed because foo != foo
    }
}
```

You can use code blocks to return strings too

```
switch foo {
    case {out bar} then {
        # not executed because foo != bar
    }
    case {out foo} then {
        # executed because foo != foo
    }
}
```

#### By Boolean State

This style of syntax could be argued as a prettier counterpart to if/else if.
Only code blocks are support and each block is checked for its boolean state
rather than string matching.

This is simply written as:

```
switch { ... }
```

### When To Use `case`, `if` and `default`?

A `switch` command may contain multiple **case** and **if** blocks. These
statements subtly alter the behavior of `switch`. You can mix and match **if**
and **case** statements within the same `switch` block.

#### case

A **case** statement will only move on to the next statement if the result of
the **case** statement is **false**. If a **case** statement is **true** then
`switch` will exit with an exit number of `0`.

```
switch {
    case { false } then {
        # ignored because case == false
    }
    case { true } then {
        # executed because case == true
    }
    case { true } then {
        # ignored because a previous case was true
    }
}
```

### if

An **if** statement will proceed to the next statement _even_ if the result of
the **if** statement is **true**.

```
switch {
    if { false } then {
        # ignored because if == false
    }
    if { true } then {
        # executed because if == true
    }
    if { true } then {
        # executed because if == true
    }
}
```

### default

**default** statements are only run if _all_ **case** _and_ **if** statements are
false.

```
switch {
    if { false } then {
        # ignored because if == false
    }
    if { true } then {
        # executed because if == true
    }
    if { true } then {
        # executed because if == true
    }
    if { false } then {
        # ignored because if == false
    }
    default {
        # ignored because one or more previous if's were true
    }
}
```

> **default** was added in Murex version 3.1

### catch

**catch** has been deprecated in version 3.1 and replaced with **default**.

## See Also

* [Caught Error Block (`catch`)](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [False (`false`)](../commands/false.md):
  Returns a `false` value
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Logic And Statements (`and`)](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [Logic Or Statements (`or`)](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [Loop While (`while`)](../commands/while.md):
  Loop until condition false
* [Not (`!`)](../commands/not-func.md):
  Reads the stdin and exit number from previous process and not's it's condition
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [True (`true`)](../commands/true.md):
  Returns a `true` value
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)

<hr/>

This document was generated from [builtins/core/structs/switch_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/switch_doc.yaml).
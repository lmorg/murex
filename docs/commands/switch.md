# _murex_ Shell Docs

## Command Reference: `switch`

> Blocks of cascading conditionals

## Description

`switch` is a large block for simplifying cascades of conditional statements.

## Usage

    switch {
      case | if { conditional } then { code-block }
      case | if { conditional } then { code-block }
      ...
      [ catch { code-block } ]
    } -> <stdout>

## Examples

Output an array of editors installed

    switch {
      if { which: vim   } { out: vim   }
      if { which: vi    } { out: vi    }
      if { which: nano  } { out: nano  }
      if { which: emacs } { out: emacs }
    } -> format: json
    
    function higherlower {
      try {
        rand: int 100 -> set rand
        while { $rand } {
          read: guess "Guess a number between 1 and 100: "
    
          switch {
            case: { = $guess < $rand } then {
              out: "Too low"
            }
    
            case: { = $guess > $rand } then {
              out: "Too high"
            }
    
            catch: {
              out: "Correct"
              let: rand=0
            }
          }
        }
      }
    }

## See Also

* [commands/`!` (not)](../commands/not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [commands/`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [commands/`catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe` 
* [commands/`false`](../commands/false.md):
  Returns a `false` value
* [commands/`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable
* [commands/`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value
* [commands/`true`](../commands/true.md):
  Returns a `true` value
* [commands/`try`](../commands/try.md):
  Handles errors inside a block of code
* [commands/`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error
* [commands/`while`](../commands/while.md):
  Loop until condition false